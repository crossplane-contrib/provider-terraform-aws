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
func (in *SnsTopicSubscription) DeepCopyInto(out *SnsTopicSubscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSubscription.
func (in *SnsTopicSubscription) DeepCopy() *SnsTopicSubscription {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsTopicSubscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicSubscriptionList) DeepCopyInto(out *SnsTopicSubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnsTopicSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSubscriptionList.
func (in *SnsTopicSubscriptionList) DeepCopy() *SnsTopicSubscriptionList {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsTopicSubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicSubscriptionObservation) DeepCopyInto(out *SnsTopicSubscriptionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSubscriptionObservation.
func (in *SnsTopicSubscriptionObservation) DeepCopy() *SnsTopicSubscriptionObservation {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSubscriptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicSubscriptionParameters) DeepCopyInto(out *SnsTopicSubscriptionParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSubscriptionParameters.
func (in *SnsTopicSubscriptionParameters) DeepCopy() *SnsTopicSubscriptionParameters {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSubscriptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicSubscriptionSpec) DeepCopyInto(out *SnsTopicSubscriptionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSubscriptionSpec.
func (in *SnsTopicSubscriptionSpec) DeepCopy() *SnsTopicSubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicSubscriptionStatus) DeepCopyInto(out *SnsTopicSubscriptionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSubscriptionStatus.
func (in *SnsTopicSubscriptionStatus) DeepCopy() *SnsTopicSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}
