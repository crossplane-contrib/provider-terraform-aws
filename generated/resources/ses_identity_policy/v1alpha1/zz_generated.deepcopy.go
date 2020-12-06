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
func (in *SesIdentityPolicy) DeepCopyInto(out *SesIdentityPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesIdentityPolicy.
func (in *SesIdentityPolicy) DeepCopy() *SesIdentityPolicy {
	if in == nil {
		return nil
	}
	out := new(SesIdentityPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SesIdentityPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesIdentityPolicyList) DeepCopyInto(out *SesIdentityPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SesIdentityPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesIdentityPolicyList.
func (in *SesIdentityPolicyList) DeepCopy() *SesIdentityPolicyList {
	if in == nil {
		return nil
	}
	out := new(SesIdentityPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SesIdentityPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesIdentityPolicyObservation) DeepCopyInto(out *SesIdentityPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesIdentityPolicyObservation.
func (in *SesIdentityPolicyObservation) DeepCopy() *SesIdentityPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(SesIdentityPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesIdentityPolicyParameters) DeepCopyInto(out *SesIdentityPolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesIdentityPolicyParameters.
func (in *SesIdentityPolicyParameters) DeepCopy() *SesIdentityPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(SesIdentityPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesIdentityPolicySpec) DeepCopyInto(out *SesIdentityPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesIdentityPolicySpec.
func (in *SesIdentityPolicySpec) DeepCopy() *SesIdentityPolicySpec {
	if in == nil {
		return nil
	}
	out := new(SesIdentityPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesIdentityPolicyStatus) DeepCopyInto(out *SesIdentityPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesIdentityPolicyStatus.
func (in *SesIdentityPolicyStatus) DeepCopy() *SesIdentityPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(SesIdentityPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
