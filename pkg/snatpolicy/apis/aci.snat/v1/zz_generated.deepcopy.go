// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Label) DeepCopyInto(out *Label) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Label.
func (in *Label) DeepCopy() *Label {
	if in == nil {
		return nil
	}
	out := new(Label)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSelector) DeepCopyInto(out *PodSelector) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSelector.
func (in *PodSelector) DeepCopy() *PodSelector {
	if in == nil {
		return nil
	}
	out := new(PodSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRange) DeepCopyInto(out *PortRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRange.
func (in *PortRange) DeepCopy() *PortRange {
	if in == nil {
		return nil
	}
	out := new(PortRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatPolicy) DeepCopyInto(out *SnatPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatPolicy.
func (in *SnatPolicy) DeepCopy() *SnatPolicy {
	if in == nil {
		return nil
	}
	out := new(SnatPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnatPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatPolicyList) DeepCopyInto(out *SnatPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnatPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatPolicyList.
func (in *SnatPolicyList) DeepCopy() *SnatPolicyList {
	if in == nil {
		return nil
	}
	out := new(SnatPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnatPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatPolicySpec) DeepCopyInto(out *SnatPolicySpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.PortRange != nil {
		in, out := &in.PortRange, &out.PortRange
		*out = make([]PortRange, len(*in))
		copy(*out, *in)
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatPolicySpec.
func (in *SnatPolicySpec) DeepCopy() *SnatPolicySpec {
	if in == nil {
		return nil
	}
	out := new(SnatPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatPolicyStatus) DeepCopyInto(out *SnatPolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatPolicyStatus.
func (in *SnatPolicyStatus) DeepCopy() *SnatPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(SnatPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
