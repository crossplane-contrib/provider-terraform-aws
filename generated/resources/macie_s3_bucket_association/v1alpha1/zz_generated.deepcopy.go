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
func (in *ClassificationType) DeepCopyInto(out *ClassificationType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassificationType.
func (in *ClassificationType) DeepCopy() *ClassificationType {
	if in == nil {
		return nil
	}
	out := new(ClassificationType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MacieS3BucketAssociation) DeepCopyInto(out *MacieS3BucketAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MacieS3BucketAssociation.
func (in *MacieS3BucketAssociation) DeepCopy() *MacieS3BucketAssociation {
	if in == nil {
		return nil
	}
	out := new(MacieS3BucketAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MacieS3BucketAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MacieS3BucketAssociationList) DeepCopyInto(out *MacieS3BucketAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MacieS3BucketAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MacieS3BucketAssociationList.
func (in *MacieS3BucketAssociationList) DeepCopy() *MacieS3BucketAssociationList {
	if in == nil {
		return nil
	}
	out := new(MacieS3BucketAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MacieS3BucketAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MacieS3BucketAssociationObservation) DeepCopyInto(out *MacieS3BucketAssociationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MacieS3BucketAssociationObservation.
func (in *MacieS3BucketAssociationObservation) DeepCopy() *MacieS3BucketAssociationObservation {
	if in == nil {
		return nil
	}
	out := new(MacieS3BucketAssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MacieS3BucketAssociationParameters) DeepCopyInto(out *MacieS3BucketAssociationParameters) {
	*out = *in
	out.ClassificationType = in.ClassificationType
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MacieS3BucketAssociationParameters.
func (in *MacieS3BucketAssociationParameters) DeepCopy() *MacieS3BucketAssociationParameters {
	if in == nil {
		return nil
	}
	out := new(MacieS3BucketAssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MacieS3BucketAssociationSpec) DeepCopyInto(out *MacieS3BucketAssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MacieS3BucketAssociationSpec.
func (in *MacieS3BucketAssociationSpec) DeepCopy() *MacieS3BucketAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(MacieS3BucketAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MacieS3BucketAssociationStatus) DeepCopyInto(out *MacieS3BucketAssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MacieS3BucketAssociationStatus.
func (in *MacieS3BucketAssociationStatus) DeepCopy() *MacieS3BucketAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(MacieS3BucketAssociationStatus)
	in.DeepCopyInto(out)
	return out
}
