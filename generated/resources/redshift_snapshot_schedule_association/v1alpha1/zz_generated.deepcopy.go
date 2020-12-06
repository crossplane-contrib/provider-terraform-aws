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
func (in *RedshiftSnapshotScheduleAssociation) DeepCopyInto(out *RedshiftSnapshotScheduleAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedshiftSnapshotScheduleAssociation.
func (in *RedshiftSnapshotScheduleAssociation) DeepCopy() *RedshiftSnapshotScheduleAssociation {
	if in == nil {
		return nil
	}
	out := new(RedshiftSnapshotScheduleAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedshiftSnapshotScheduleAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedshiftSnapshotScheduleAssociationList) DeepCopyInto(out *RedshiftSnapshotScheduleAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedshiftSnapshotScheduleAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedshiftSnapshotScheduleAssociationList.
func (in *RedshiftSnapshotScheduleAssociationList) DeepCopy() *RedshiftSnapshotScheduleAssociationList {
	if in == nil {
		return nil
	}
	out := new(RedshiftSnapshotScheduleAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedshiftSnapshotScheduleAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedshiftSnapshotScheduleAssociationObservation) DeepCopyInto(out *RedshiftSnapshotScheduleAssociationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedshiftSnapshotScheduleAssociationObservation.
func (in *RedshiftSnapshotScheduleAssociationObservation) DeepCopy() *RedshiftSnapshotScheduleAssociationObservation {
	if in == nil {
		return nil
	}
	out := new(RedshiftSnapshotScheduleAssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedshiftSnapshotScheduleAssociationParameters) DeepCopyInto(out *RedshiftSnapshotScheduleAssociationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedshiftSnapshotScheduleAssociationParameters.
func (in *RedshiftSnapshotScheduleAssociationParameters) DeepCopy() *RedshiftSnapshotScheduleAssociationParameters {
	if in == nil {
		return nil
	}
	out := new(RedshiftSnapshotScheduleAssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedshiftSnapshotScheduleAssociationSpec) DeepCopyInto(out *RedshiftSnapshotScheduleAssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedshiftSnapshotScheduleAssociationSpec.
func (in *RedshiftSnapshotScheduleAssociationSpec) DeepCopy() *RedshiftSnapshotScheduleAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(RedshiftSnapshotScheduleAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedshiftSnapshotScheduleAssociationStatus) DeepCopyInto(out *RedshiftSnapshotScheduleAssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedshiftSnapshotScheduleAssociationStatus.
func (in *RedshiftSnapshotScheduleAssociationStatus) DeepCopy() *RedshiftSnapshotScheduleAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(RedshiftSnapshotScheduleAssociationStatus)
	in.DeepCopyInto(out)
	return out
}
