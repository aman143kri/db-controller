//go:build !ignore_autogenerated

/*
Copyright 2024.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XNetworkRecord) DeepCopyInto(out *XNetworkRecord) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XNetworkRecord.
func (in *XNetworkRecord) DeepCopy() *XNetworkRecord {
	if in == nil {
		return nil
	}
	out := new(XNetworkRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XNetworkRecord) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XNetworkRecordList) DeepCopyInto(out *XNetworkRecordList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]XNetworkRecord, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XNetworkRecordList.
func (in *XNetworkRecordList) DeepCopy() *XNetworkRecordList {
	if in == nil {
		return nil
	}
	out := new(XNetworkRecordList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XNetworkRecordList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XNetworkRecordParameters) DeepCopyInto(out *XNetworkRecordParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XNetworkRecordParameters.
func (in *XNetworkRecordParameters) DeepCopy() *XNetworkRecordParameters {
	if in == nil {
		return nil
	}
	out := new(XNetworkRecordParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XNetworkRecordSpec) DeepCopyInto(out *XNetworkRecordSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.Parameters = in.Parameters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XNetworkRecordSpec.
func (in *XNetworkRecordSpec) DeepCopy() *XNetworkRecordSpec {
	if in == nil {
		return nil
	}
	out := new(XNetworkRecordSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XNetworkRecordStatus) DeepCopyInto(out *XNetworkRecordStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XNetworkRecordStatus.
func (in *XNetworkRecordStatus) DeepCopy() *XNetworkRecordStatus {
	if in == nil {
		return nil
	}
	out := new(XNetworkRecordStatus)
	in.DeepCopyInto(out)
	return out
}