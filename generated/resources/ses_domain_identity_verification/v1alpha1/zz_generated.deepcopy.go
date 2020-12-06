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
func (in *SesDomainIdentityVerification) DeepCopyInto(out *SesDomainIdentityVerification) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesDomainIdentityVerification.
func (in *SesDomainIdentityVerification) DeepCopy() *SesDomainIdentityVerification {
	if in == nil {
		return nil
	}
	out := new(SesDomainIdentityVerification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SesDomainIdentityVerification) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesDomainIdentityVerificationList) DeepCopyInto(out *SesDomainIdentityVerificationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SesDomainIdentityVerification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesDomainIdentityVerificationList.
func (in *SesDomainIdentityVerificationList) DeepCopy() *SesDomainIdentityVerificationList {
	if in == nil {
		return nil
	}
	out := new(SesDomainIdentityVerificationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SesDomainIdentityVerificationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesDomainIdentityVerificationObservation) DeepCopyInto(out *SesDomainIdentityVerificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesDomainIdentityVerificationObservation.
func (in *SesDomainIdentityVerificationObservation) DeepCopy() *SesDomainIdentityVerificationObservation {
	if in == nil {
		return nil
	}
	out := new(SesDomainIdentityVerificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesDomainIdentityVerificationParameters) DeepCopyInto(out *SesDomainIdentityVerificationParameters) {
	*out = *in
	out.Timeouts = in.Timeouts
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesDomainIdentityVerificationParameters.
func (in *SesDomainIdentityVerificationParameters) DeepCopy() *SesDomainIdentityVerificationParameters {
	if in == nil {
		return nil
	}
	out := new(SesDomainIdentityVerificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesDomainIdentityVerificationSpec) DeepCopyInto(out *SesDomainIdentityVerificationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesDomainIdentityVerificationSpec.
func (in *SesDomainIdentityVerificationSpec) DeepCopy() *SesDomainIdentityVerificationSpec {
	if in == nil {
		return nil
	}
	out := new(SesDomainIdentityVerificationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesDomainIdentityVerificationStatus) DeepCopyInto(out *SesDomainIdentityVerificationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesDomainIdentityVerificationStatus.
func (in *SesDomainIdentityVerificationStatus) DeepCopy() *SesDomainIdentityVerificationStatus {
	if in == nil {
		return nil
	}
	out := new(SesDomainIdentityVerificationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timeouts) DeepCopyInto(out *Timeouts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timeouts.
func (in *Timeouts) DeepCopy() *Timeouts {
	if in == nil {
		return nil
	}
	out := new(Timeouts)
	in.DeepCopyInto(out)
	return out
}