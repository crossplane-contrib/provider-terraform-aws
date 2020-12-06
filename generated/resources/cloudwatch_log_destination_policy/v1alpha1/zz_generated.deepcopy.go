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
func (in *CloudwatchLogDestinationPolicy) DeepCopyInto(out *CloudwatchLogDestinationPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogDestinationPolicy.
func (in *CloudwatchLogDestinationPolicy) DeepCopy() *CloudwatchLogDestinationPolicy {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogDestinationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudwatchLogDestinationPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchLogDestinationPolicyList) DeepCopyInto(out *CloudwatchLogDestinationPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudwatchLogDestinationPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogDestinationPolicyList.
func (in *CloudwatchLogDestinationPolicyList) DeepCopy() *CloudwatchLogDestinationPolicyList {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogDestinationPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudwatchLogDestinationPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchLogDestinationPolicyObservation) DeepCopyInto(out *CloudwatchLogDestinationPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogDestinationPolicyObservation.
func (in *CloudwatchLogDestinationPolicyObservation) DeepCopy() *CloudwatchLogDestinationPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogDestinationPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchLogDestinationPolicyParameters) DeepCopyInto(out *CloudwatchLogDestinationPolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogDestinationPolicyParameters.
func (in *CloudwatchLogDestinationPolicyParameters) DeepCopy() *CloudwatchLogDestinationPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogDestinationPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchLogDestinationPolicySpec) DeepCopyInto(out *CloudwatchLogDestinationPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogDestinationPolicySpec.
func (in *CloudwatchLogDestinationPolicySpec) DeepCopy() *CloudwatchLogDestinationPolicySpec {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogDestinationPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchLogDestinationPolicyStatus) DeepCopyInto(out *CloudwatchLogDestinationPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogDestinationPolicyStatus.
func (in *CloudwatchLogDestinationPolicyStatus) DeepCopy() *CloudwatchLogDestinationPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogDestinationPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
