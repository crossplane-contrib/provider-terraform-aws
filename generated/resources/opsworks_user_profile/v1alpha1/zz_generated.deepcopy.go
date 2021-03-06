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
func (in *OpsworksUserProfile) DeepCopyInto(out *OpsworksUserProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksUserProfile.
func (in *OpsworksUserProfile) DeepCopy() *OpsworksUserProfile {
	if in == nil {
		return nil
	}
	out := new(OpsworksUserProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsworksUserProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksUserProfileList) DeepCopyInto(out *OpsworksUserProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpsworksUserProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksUserProfileList.
func (in *OpsworksUserProfileList) DeepCopy() *OpsworksUserProfileList {
	if in == nil {
		return nil
	}
	out := new(OpsworksUserProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsworksUserProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksUserProfileObservation) DeepCopyInto(out *OpsworksUserProfileObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksUserProfileObservation.
func (in *OpsworksUserProfileObservation) DeepCopy() *OpsworksUserProfileObservation {
	if in == nil {
		return nil
	}
	out := new(OpsworksUserProfileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksUserProfileParameters) DeepCopyInto(out *OpsworksUserProfileParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksUserProfileParameters.
func (in *OpsworksUserProfileParameters) DeepCopy() *OpsworksUserProfileParameters {
	if in == nil {
		return nil
	}
	out := new(OpsworksUserProfileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksUserProfileSpec) DeepCopyInto(out *OpsworksUserProfileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksUserProfileSpec.
func (in *OpsworksUserProfileSpec) DeepCopy() *OpsworksUserProfileSpec {
	if in == nil {
		return nil
	}
	out := new(OpsworksUserProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksUserProfileStatus) DeepCopyInto(out *OpsworksUserProfileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksUserProfileStatus.
func (in *OpsworksUserProfileStatus) DeepCopy() *OpsworksUserProfileStatus {
	if in == nil {
		return nil
	}
	out := new(OpsworksUserProfileStatus)
	in.DeepCopyInto(out)
	return out
}
