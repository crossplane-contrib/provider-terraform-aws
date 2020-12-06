// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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
func (in *TransferSshKey) DeepCopyInto(out *TransferSshKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferSshKey.
func (in *TransferSshKey) DeepCopy() *TransferSshKey {
	if in == nil {
		return nil
	}
	out := new(TransferSshKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TransferSshKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferSshKeyList) DeepCopyInto(out *TransferSshKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TransferSshKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferSshKeyList.
func (in *TransferSshKeyList) DeepCopy() *TransferSshKeyList {
	if in == nil {
		return nil
	}
	out := new(TransferSshKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TransferSshKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferSshKeyObservation) DeepCopyInto(out *TransferSshKeyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferSshKeyObservation.
func (in *TransferSshKeyObservation) DeepCopy() *TransferSshKeyObservation {
	if in == nil {
		return nil
	}
	out := new(TransferSshKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferSshKeyParameters) DeepCopyInto(out *TransferSshKeyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferSshKeyParameters.
func (in *TransferSshKeyParameters) DeepCopy() *TransferSshKeyParameters {
	if in == nil {
		return nil
	}
	out := new(TransferSshKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferSshKeySpec) DeepCopyInto(out *TransferSshKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferSshKeySpec.
func (in *TransferSshKeySpec) DeepCopy() *TransferSshKeySpec {
	if in == nil {
		return nil
	}
	out := new(TransferSshKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferSshKeyStatus) DeepCopyInto(out *TransferSshKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferSshKeyStatus.
func (in *TransferSshKeyStatus) DeepCopy() *TransferSshKeyStatus {
	if in == nil {
		return nil
	}
	out := new(TransferSshKeyStatus)
	in.DeepCopyInto(out)
	return out
}