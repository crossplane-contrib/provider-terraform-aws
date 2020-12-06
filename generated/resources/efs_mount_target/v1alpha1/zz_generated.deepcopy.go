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
func (in *EfsMountTarget) DeepCopyInto(out *EfsMountTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EfsMountTarget.
func (in *EfsMountTarget) DeepCopy() *EfsMountTarget {
	if in == nil {
		return nil
	}
	out := new(EfsMountTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EfsMountTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EfsMountTargetList) DeepCopyInto(out *EfsMountTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EfsMountTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EfsMountTargetList.
func (in *EfsMountTargetList) DeepCopy() *EfsMountTargetList {
	if in == nil {
		return nil
	}
	out := new(EfsMountTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EfsMountTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EfsMountTargetObservation) DeepCopyInto(out *EfsMountTargetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EfsMountTargetObservation.
func (in *EfsMountTargetObservation) DeepCopy() *EfsMountTargetObservation {
	if in == nil {
		return nil
	}
	out := new(EfsMountTargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EfsMountTargetParameters) DeepCopyInto(out *EfsMountTargetParameters) {
	*out = *in
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EfsMountTargetParameters.
func (in *EfsMountTargetParameters) DeepCopy() *EfsMountTargetParameters {
	if in == nil {
		return nil
	}
	out := new(EfsMountTargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EfsMountTargetSpec) DeepCopyInto(out *EfsMountTargetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EfsMountTargetSpec.
func (in *EfsMountTargetSpec) DeepCopy() *EfsMountTargetSpec {
	if in == nil {
		return nil
	}
	out := new(EfsMountTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EfsMountTargetStatus) DeepCopyInto(out *EfsMountTargetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EfsMountTargetStatus.
func (in *EfsMountTargetStatus) DeepCopy() *EfsMountTargetStatus {
	if in == nil {
		return nil
	}
	out := new(EfsMountTargetStatus)
	in.DeepCopyInto(out)
	return out
}
