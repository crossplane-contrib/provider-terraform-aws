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
func (in *GlueConnection) DeepCopyInto(out *GlueConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueConnection.
func (in *GlueConnection) DeepCopy() *GlueConnection {
	if in == nil {
		return nil
	}
	out := new(GlueConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueConnectionList) DeepCopyInto(out *GlueConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlueConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueConnectionList.
func (in *GlueConnectionList) DeepCopy() *GlueConnectionList {
	if in == nil {
		return nil
	}
	out := new(GlueConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueConnectionObservation) DeepCopyInto(out *GlueConnectionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueConnectionObservation.
func (in *GlueConnectionObservation) DeepCopy() *GlueConnectionObservation {
	if in == nil {
		return nil
	}
	out := new(GlueConnectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueConnectionParameters) DeepCopyInto(out *GlueConnectionParameters) {
	*out = *in
	if in.MatchCriteria != nil {
		in, out := &in.MatchCriteria, &out.MatchCriteria
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConnectionProperties != nil {
		in, out := &in.ConnectionProperties, &out.ConnectionProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.PhysicalConnectionRequirements.DeepCopyInto(&out.PhysicalConnectionRequirements)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueConnectionParameters.
func (in *GlueConnectionParameters) DeepCopy() *GlueConnectionParameters {
	if in == nil {
		return nil
	}
	out := new(GlueConnectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueConnectionSpec) DeepCopyInto(out *GlueConnectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueConnectionSpec.
func (in *GlueConnectionSpec) DeepCopy() *GlueConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(GlueConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueConnectionStatus) DeepCopyInto(out *GlueConnectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueConnectionStatus.
func (in *GlueConnectionStatus) DeepCopy() *GlueConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(GlueConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhysicalConnectionRequirements) DeepCopyInto(out *PhysicalConnectionRequirements) {
	*out = *in
	if in.SecurityGroupIdList != nil {
		in, out := &in.SecurityGroupIdList, &out.SecurityGroupIdList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhysicalConnectionRequirements.
func (in *PhysicalConnectionRequirements) DeepCopy() *PhysicalConnectionRequirements {
	if in == nil {
		return nil
	}
	out := new(PhysicalConnectionRequirements)
	in.DeepCopyInto(out)
	return out
}
