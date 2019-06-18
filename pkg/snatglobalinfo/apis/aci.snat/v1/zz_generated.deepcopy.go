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
func (in *GlobalInfo) DeepCopyInto(out *GlobalInfo) {
	*out = *in
	if in.PortRanges != nil {
		in, out := &in.PortRanges, &out.PortRanges
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalInfo.
func (in *GlobalInfo) DeepCopy() *GlobalInfo {
	if in == nil {
		return nil
	}
	out := new(GlobalInfo)
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
func (in *SnatGlobalInfo) DeepCopyInto(out *SnatGlobalInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatGlobalInfo.
func (in *SnatGlobalInfo) DeepCopy() *SnatGlobalInfo {
	if in == nil {
		return nil
	}
	out := new(SnatGlobalInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnatGlobalInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatGlobalInfoList) DeepCopyInto(out *SnatGlobalInfoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnatGlobalInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatGlobalInfoList.
func (in *SnatGlobalInfoList) DeepCopy() *SnatGlobalInfoList {
	if in == nil {
		return nil
	}
	out := new(SnatGlobalInfoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnatGlobalInfoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatGlobalInfoSpec) DeepCopyInto(out *SnatGlobalInfoSpec) {
	*out = *in
	if in.GlobalInfos != nil {
		in, out := &in.GlobalInfos, &out.GlobalInfos
		*out = make(map[string][]GlobalInfo, len(*in))
		for key, val := range *in {
			var outVal []GlobalInfo
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]GlobalInfo, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatGlobalInfoSpec.
func (in *SnatGlobalInfoSpec) DeepCopy() *SnatGlobalInfoSpec {
	if in == nil {
		return nil
	}
	out := new(SnatGlobalInfoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatGlobalInfoStatus) DeepCopyInto(out *SnatGlobalInfoStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatGlobalInfoStatus.
func (in *SnatGlobalInfoStatus) DeepCopy() *SnatGlobalInfoStatus {
	if in == nil {
		return nil
	}
	out := new(SnatGlobalInfoStatus)
	in.DeepCopyInto(out)
	return out
}
