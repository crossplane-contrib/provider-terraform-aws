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
func (in *AutoscalingAttachment) DeepCopyInto(out *AutoscalingAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingAttachment.
func (in *AutoscalingAttachment) DeepCopy() *AutoscalingAttachment {
	if in == nil {
		return nil
	}
	out := new(AutoscalingAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoscalingAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingAttachmentList) DeepCopyInto(out *AutoscalingAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutoscalingAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingAttachmentList.
func (in *AutoscalingAttachmentList) DeepCopy() *AutoscalingAttachmentList {
	if in == nil {
		return nil
	}
	out := new(AutoscalingAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoscalingAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingAttachmentObservation) DeepCopyInto(out *AutoscalingAttachmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingAttachmentObservation.
func (in *AutoscalingAttachmentObservation) DeepCopy() *AutoscalingAttachmentObservation {
	if in == nil {
		return nil
	}
	out := new(AutoscalingAttachmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingAttachmentParameters) DeepCopyInto(out *AutoscalingAttachmentParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingAttachmentParameters.
func (in *AutoscalingAttachmentParameters) DeepCopy() *AutoscalingAttachmentParameters {
	if in == nil {
		return nil
	}
	out := new(AutoscalingAttachmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingAttachmentSpec) DeepCopyInto(out *AutoscalingAttachmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingAttachmentSpec.
func (in *AutoscalingAttachmentSpec) DeepCopy() *AutoscalingAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(AutoscalingAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingAttachmentStatus) DeepCopyInto(out *AutoscalingAttachmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingAttachmentStatus.
func (in *AutoscalingAttachmentStatus) DeepCopy() *AutoscalingAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(AutoscalingAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}
