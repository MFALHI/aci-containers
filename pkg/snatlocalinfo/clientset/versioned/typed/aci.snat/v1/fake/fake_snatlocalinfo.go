/***
Copyright 2019 Cisco Systems Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	acisnatv1 "github.com/noironetworks/aci-containers/pkg/snatlocalinfo/apis/aci.snat/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSnatLocalInfos implements SnatLocalInfoInterface
type FakeSnatLocalInfos struct {
	Fake *FakeAciV1
	ns   string
}

var snatlocalinfosResource = schema.GroupVersionResource{Group: "aci.snat", Version: "v1", Resource: "snatlocalinfos"}

var snatlocalinfosKind = schema.GroupVersionKind{Group: "aci.snat", Version: "v1", Kind: "SnatLocalInfo"}

// Get takes name of the snatLocalInfo, and returns the corresponding snatLocalInfo object, and an error if there is any.
func (c *FakeSnatLocalInfos) Get(name string, options v1.GetOptions) (result *acisnatv1.SnatLocalInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(snatlocalinfosResource, c.ns, name), &acisnatv1.SnatLocalInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acisnatv1.SnatLocalInfo), err
}

// List takes label and field selectors, and returns the list of SnatLocalInfos that match those selectors.
func (c *FakeSnatLocalInfos) List(opts v1.ListOptions) (result *acisnatv1.SnatLocalInfoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(snatlocalinfosResource, snatlocalinfosKind, c.ns, opts), &acisnatv1.SnatLocalInfoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &acisnatv1.SnatLocalInfoList{ListMeta: obj.(*acisnatv1.SnatLocalInfoList).ListMeta}
	for _, item := range obj.(*acisnatv1.SnatLocalInfoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested snatLocalInfos.
func (c *FakeSnatLocalInfos) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(snatlocalinfosResource, c.ns, opts))

}

// Create takes the representation of a snatLocalInfo and creates it.  Returns the server's representation of the snatLocalInfo, and an error, if there is any.
func (c *FakeSnatLocalInfos) Create(snatLocalInfo *acisnatv1.SnatLocalInfo) (result *acisnatv1.SnatLocalInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(snatlocalinfosResource, c.ns, snatLocalInfo), &acisnatv1.SnatLocalInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acisnatv1.SnatLocalInfo), err
}

// Update takes the representation of a snatLocalInfo and updates it. Returns the server's representation of the snatLocalInfo, and an error, if there is any.
func (c *FakeSnatLocalInfos) Update(snatLocalInfo *acisnatv1.SnatLocalInfo) (result *acisnatv1.SnatLocalInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(snatlocalinfosResource, c.ns, snatLocalInfo), &acisnatv1.SnatLocalInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acisnatv1.SnatLocalInfo), err
}

// Delete takes name of the snatLocalInfo and deletes it. Returns an error if one occurs.
func (c *FakeSnatLocalInfos) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(snatlocalinfosResource, c.ns, name), &acisnatv1.SnatLocalInfo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSnatLocalInfos) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(snatlocalinfosResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &acisnatv1.SnatLocalInfoList{})
	return err
}

// Patch applies the patch and returns the patched snatLocalInfo.
func (c *FakeSnatLocalInfos) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *acisnatv1.SnatLocalInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(snatlocalinfosResource, c.ns, name, pt, data, subresources...), &acisnatv1.SnatLocalInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acisnatv1.SnatLocalInfo), err
}
