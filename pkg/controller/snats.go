// Copyright 2019 Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"github.com/Sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
//	"k8s.io/apimachinery/pkg/runtime"
//	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
//	"k8s.io/kubernetes/pkg/controller"
	snatclientset "github.com/noironetworks/aci-containers/pkg/snatpolicy/clientset/versioned"
	snatpolicy "github.com/noironetworks/aci-containers/pkg/snatpolicy/apis/aci.snat/v1"
)

type MyLabel struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type MyPodSelector struct {
	Labels     []MyLabel
	Deployment string
	Namespace  string
}

type MyPortRange struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

type MySnatPolicy struct {
	SnatIp    []string
	Selector  MyPodSelector
	PortRange []MyPortRange
	Protocols []string
}

func SnatPolicyLogger(log *logrus.Logger, snat *snatpolicy.SnatPolicy) *logrus.Entry {
	return log.WithFields(logrus.Fields{
		"namespace": snat.ObjectMeta.Namespace,
		"name":      snat.ObjectMeta.Name,
		"spec":      snat.Spec,
	})
}

func (cont *AciController) initSnatInformerFromClient(
	snatClient *snatclientset.Clientset) {
	cont.log.Debug("NEW INFROMER CLIENT")
	cont.initSnatInformerBase(
		cache.NewListWatchFromClient(
			snatClient.AciV1().RESTClient(), "snatpolicies",
			metav1.NamespaceAll, fields.Everything()))
}

func (cont *AciController) initSnatInformerBase(listWatch *cache.ListWatch) {
	cont.snatIndexer, cont.snatInformer = cache.NewIndexerInformer(
		listWatch,
		&snatpolicy.SnatPolicy{}, 0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				cont.log.Debug("POLICY ADDED")
			//	cont.snatPolicyUpdate(obj)
				cont.snatAdded(obj)
			},
			UpdateFunc: func(_ interface{}, obj interface{}) {
				//cont.snatPolicyUpdate(obj)
				cont.log.Debug("POLICY UPDATED")
				cont.snatAdded(obj)
			},
			DeleteFunc: func(obj interface{}) {
				cont.snatPolicyDelete(obj)
			},
		},
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	cont.log.Debug("Initializing Snat Policy Informers")

}

func(cont *AciController) snatAdded(obj interface{}) {
	cont.log.Debug("NEW POLICY ADDER")
	snat := obj.(*snatpolicy.SnatPolicy)
	key, err := cache.MetaNamespaceKeyFunc(snat)
	if err != nil {
		SnatPolicyLogger(cont.log, snat).
			Error("Could not create key:" + err.Error())
		return
	}
	cont.queueSnatUpdateByKey(key)
}

func (cont *AciController) queueSnatUpdateByKey(key string) {
	cont.snatQ.Add(key)
}

func (cont *AciController) handleSnatUpdate(snatpolicy *snatpolicy.SnatPolicy) bool {

	_, err := cache.MetaNamespaceKeyFunc(snatpolicy)
	if err != nil {
		SnatPolicyLogger(cont.log, snatpolicy).
			Error("Could not create key:" + err.Error())
		return false
	}

	policyName := snatpolicy.ObjectMeta.Name
	var requeue bool
	cont.indexMutex.Lock()
	cont.updateSnatPolicyCache(policyName, snatpolicy)
	cont.log.Debug("map issssssssss ", cont.snatPolicyCache)
	cont.indexMutex.Unlock()
	err = cont.updateServiceDeviceInstanceSnat("MYSNAT")
	if err != nil {
		requeue = false
	}
	cont.log.Debug("map issssssssss ", cont.snatPolicyCache)
	return requeue
}

//this should take snatpolicy, not obj
func (cont *AciController) updateSnatPolicyCache(key string, snatobj interface{}) {
	snatpolicy := snatobj.(*snatpolicy.SnatPolicy)
	var mypolicy MySnatPolicy
	mypolicy.SnatIp = snatpolicy.Spec.SnatIp
	snatLabels := snatpolicy.Spec.Selector.Labels
	snatDeploy := snatpolicy.Spec.Selector.Deployment
	snatNS := snatpolicy.Spec.Selector.Namespace
	var myLabels []MyLabel
	for _, val := range snatLabels {
		lab := MyLabel{Key: val.Key, Value: val.Value}
		myLabels = append(myLabels, lab)
	}
	mypolicy.Selector = MyPodSelector{Labels: myLabels, Deployment: snatDeploy, Namespace: snatNS}
	cont.snatPolicyCache[key] = mypolicy
//	return mypolicy
}

func (cont *AciController) snatPolicyDelete(snatobj interface{}) {
	cont.indexMutex.Lock()
        snatpolicy := snatobj.(*snatpolicy.SnatPolicy)
	delete(cont.snatPolicyCache, snatpolicy.ObjectMeta.Name)

	cont.log.Debug("cache after deleting is ", cont.snatPolicyCache)
        if len(cont.snatPolicyCache) == 0 {
                cont.log.Debug("Cache is empty now....")
		graphName := cont.aciNameForKey("snat", "MYSNAT")
		go cont.apicConn.ClearApicObjects(graphName)
        } else {
		go cont.updateServiceDeviceInstanceSnat("MYSNAT")
	}
	cont.indexMutex.Unlock()
}
