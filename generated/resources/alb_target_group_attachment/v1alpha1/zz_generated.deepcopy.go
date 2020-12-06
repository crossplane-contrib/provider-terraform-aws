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
func (in *AlbTargetGroupAttachment) DeepCopyInto(out *AlbTargetGroupAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlbTargetGroupAttachment.
func (in *AlbTargetGroupAttachment) DeepCopy() *AlbTargetGroupAttachment {
	if in == nil {
		return nil
	}
	out := new(AlbTargetGroupAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlbTargetGroupAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlbTargetGroupAttachmentList) DeepCopyInto(out *AlbTargetGroupAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlbTargetGroupAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlbTargetGroupAttachmentList.
func (in *AlbTargetGroupAttachmentList) DeepCopy() *AlbTargetGroupAttachmentList {
	if in == nil {
		return nil
	}
	out := new(AlbTargetGroupAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlbTargetGroupAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlbTargetGroupAttachmentObservation) DeepCopyInto(out *AlbTargetGroupAttachmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlbTargetGroupAttachmentObservation.
func (in *AlbTargetGroupAttachmentObservation) DeepCopy() *AlbTargetGroupAttachmentObservation {
	if in == nil {
		return nil
	}
	out := new(AlbTargetGroupAttachmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlbTargetGroupAttachmentParameters) DeepCopyInto(out *AlbTargetGroupAttachmentParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlbTargetGroupAttachmentParameters.
func (in *AlbTargetGroupAttachmentParameters) DeepCopy() *AlbTargetGroupAttachmentParameters {
	if in == nil {
		return nil
	}
	out := new(AlbTargetGroupAttachmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlbTargetGroupAttachmentSpec) DeepCopyInto(out *AlbTargetGroupAttachmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlbTargetGroupAttachmentSpec.
func (in *AlbTargetGroupAttachmentSpec) DeepCopy() *AlbTargetGroupAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(AlbTargetGroupAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlbTargetGroupAttachmentStatus) DeepCopyInto(out *AlbTargetGroupAttachmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlbTargetGroupAttachmentStatus.
func (in *AlbTargetGroupAttachmentStatus) DeepCopy() *AlbTargetGroupAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(AlbTargetGroupAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}
