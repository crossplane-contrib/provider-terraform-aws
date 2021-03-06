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
func (in *AutoscalingSchedule) DeepCopyInto(out *AutoscalingSchedule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingSchedule.
func (in *AutoscalingSchedule) DeepCopy() *AutoscalingSchedule {
	if in == nil {
		return nil
	}
	out := new(AutoscalingSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoscalingSchedule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingScheduleList) DeepCopyInto(out *AutoscalingScheduleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutoscalingSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingScheduleList.
func (in *AutoscalingScheduleList) DeepCopy() *AutoscalingScheduleList {
	if in == nil {
		return nil
	}
	out := new(AutoscalingScheduleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoscalingScheduleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingScheduleObservation) DeepCopyInto(out *AutoscalingScheduleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingScheduleObservation.
func (in *AutoscalingScheduleObservation) DeepCopy() *AutoscalingScheduleObservation {
	if in == nil {
		return nil
	}
	out := new(AutoscalingScheduleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingScheduleParameters) DeepCopyInto(out *AutoscalingScheduleParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingScheduleParameters.
func (in *AutoscalingScheduleParameters) DeepCopy() *AutoscalingScheduleParameters {
	if in == nil {
		return nil
	}
	out := new(AutoscalingScheduleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingScheduleSpec) DeepCopyInto(out *AutoscalingScheduleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingScheduleSpec.
func (in *AutoscalingScheduleSpec) DeepCopy() *AutoscalingScheduleSpec {
	if in == nil {
		return nil
	}
	out := new(AutoscalingScheduleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingScheduleStatus) DeepCopyInto(out *AutoscalingScheduleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingScheduleStatus.
func (in *AutoscalingScheduleStatus) DeepCopy() *AutoscalingScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(AutoscalingScheduleStatus)
	in.DeepCopyInto(out)
	return out
}
