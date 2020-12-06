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
func (in *EfsFileSystemPolicy) DeepCopyInto(out *EfsFileSystemPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EfsFileSystemPolicy.
func (in *EfsFileSystemPolicy) DeepCopy() *EfsFileSystemPolicy {
	if in == nil {
		return nil
	}
	out := new(EfsFileSystemPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EfsFileSystemPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EfsFileSystemPolicyList) DeepCopyInto(out *EfsFileSystemPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EfsFileSystemPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EfsFileSystemPolicyList.
func (in *EfsFileSystemPolicyList) DeepCopy() *EfsFileSystemPolicyList {
	if in == nil {
		return nil
	}
	out := new(EfsFileSystemPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EfsFileSystemPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EfsFileSystemPolicyObservation) DeepCopyInto(out *EfsFileSystemPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EfsFileSystemPolicyObservation.
func (in *EfsFileSystemPolicyObservation) DeepCopy() *EfsFileSystemPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(EfsFileSystemPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EfsFileSystemPolicyParameters) DeepCopyInto(out *EfsFileSystemPolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EfsFileSystemPolicyParameters.
func (in *EfsFileSystemPolicyParameters) DeepCopy() *EfsFileSystemPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(EfsFileSystemPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EfsFileSystemPolicySpec) DeepCopyInto(out *EfsFileSystemPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EfsFileSystemPolicySpec.
func (in *EfsFileSystemPolicySpec) DeepCopy() *EfsFileSystemPolicySpec {
	if in == nil {
		return nil
	}
	out := new(EfsFileSystemPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EfsFileSystemPolicyStatus) DeepCopyInto(out *EfsFileSystemPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EfsFileSystemPolicyStatus.
func (in *EfsFileSystemPolicyStatus) DeepCopy() *EfsFileSystemPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(EfsFileSystemPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
