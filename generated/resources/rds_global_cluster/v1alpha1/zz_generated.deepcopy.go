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
func (in *GlobalClusterMembers) DeepCopyInto(out *GlobalClusterMembers) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalClusterMembers.
func (in *GlobalClusterMembers) DeepCopy() *GlobalClusterMembers {
	if in == nil {
		return nil
	}
	out := new(GlobalClusterMembers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsGlobalCluster) DeepCopyInto(out *RdsGlobalCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsGlobalCluster.
func (in *RdsGlobalCluster) DeepCopy() *RdsGlobalCluster {
	if in == nil {
		return nil
	}
	out := new(RdsGlobalCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RdsGlobalCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsGlobalClusterList) DeepCopyInto(out *RdsGlobalClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RdsGlobalCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsGlobalClusterList.
func (in *RdsGlobalClusterList) DeepCopy() *RdsGlobalClusterList {
	if in == nil {
		return nil
	}
	out := new(RdsGlobalClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RdsGlobalClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsGlobalClusterObservation) DeepCopyInto(out *RdsGlobalClusterObservation) {
	*out = *in
	if in.GlobalClusterMembers != nil {
		in, out := &in.GlobalClusterMembers, &out.GlobalClusterMembers
		*out = make([]GlobalClusterMembers, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsGlobalClusterObservation.
func (in *RdsGlobalClusterObservation) DeepCopy() *RdsGlobalClusterObservation {
	if in == nil {
		return nil
	}
	out := new(RdsGlobalClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsGlobalClusterParameters) DeepCopyInto(out *RdsGlobalClusterParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsGlobalClusterParameters.
func (in *RdsGlobalClusterParameters) DeepCopy() *RdsGlobalClusterParameters {
	if in == nil {
		return nil
	}
	out := new(RdsGlobalClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsGlobalClusterSpec) DeepCopyInto(out *RdsGlobalClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsGlobalClusterSpec.
func (in *RdsGlobalClusterSpec) DeepCopy() *RdsGlobalClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RdsGlobalClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RdsGlobalClusterStatus) DeepCopyInto(out *RdsGlobalClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RdsGlobalClusterStatus.
func (in *RdsGlobalClusterStatus) DeepCopy() *RdsGlobalClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RdsGlobalClusterStatus)
	in.DeepCopyInto(out)
	return out
}