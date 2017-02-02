// Copyright 2016 Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Handlers for namespace updates.  Keeps an index of namespace
// annotations

package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	v1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/controller"
)

func (cont *aciController) initNamespaceInformer() {
	cont.namespaceInformer = cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return cont.kubeClient.Core().Namespaces().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return cont.kubeClient.Core().Namespaces().Watch(options)
			},
		},
		&v1.Namespace{},
		controller.NoResyncPeriodFunc(),
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	cont.namespaceInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			cont.namespaceChanged(obj)
		},
		UpdateFunc: func(_ interface{}, obj interface{}) {
			cont.namespaceChanged(obj)
		},
		DeleteFunc: func(obj interface{}) {
			cont.namespaceChanged(obj)
		},
	})

	go cont.namespaceInformer.GetController().Run(wait.NeverStop)
	go cont.namespaceInformer.Run(wait.NeverStop)
}

func (cont *aciController) namespaceChanged(obj interface{}) {
	cont.indexMutex.Lock()
	defer cont.indexMutex.Unlock()

	ns := obj.(*v1.Namespace)

	pods := cont.podInformer.GetStore().List()
	for _, podobj := range pods {
		pod := podobj.(*v1.Pod)

		if ns.Name == pod.ObjectMeta.Namespace {
			cont.podChangedLocked(pod)
		}
	}
}
