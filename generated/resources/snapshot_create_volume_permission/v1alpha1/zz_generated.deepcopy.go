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
func (in *SnapshotCreateVolumePermission) DeepCopyInto(out *SnapshotCreateVolumePermission) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotCreateVolumePermission.
func (in *SnapshotCreateVolumePermission) DeepCopy() *SnapshotCreateVolumePermission {
	if in == nil {
		return nil
	}
	out := new(SnapshotCreateVolumePermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotCreateVolumePermission) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotCreateVolumePermissionList) DeepCopyInto(out *SnapshotCreateVolumePermissionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnapshotCreateVolumePermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotCreateVolumePermissionList.
func (in *SnapshotCreateVolumePermissionList) DeepCopy() *SnapshotCreateVolumePermissionList {
	if in == nil {
		return nil
	}
	out := new(SnapshotCreateVolumePermissionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotCreateVolumePermissionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotCreateVolumePermissionObservation) DeepCopyInto(out *SnapshotCreateVolumePermissionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotCreateVolumePermissionObservation.
func (in *SnapshotCreateVolumePermissionObservation) DeepCopy() *SnapshotCreateVolumePermissionObservation {
	if in == nil {
		return nil
	}
	out := new(SnapshotCreateVolumePermissionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotCreateVolumePermissionParameters) DeepCopyInto(out *SnapshotCreateVolumePermissionParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotCreateVolumePermissionParameters.
func (in *SnapshotCreateVolumePermissionParameters) DeepCopy() *SnapshotCreateVolumePermissionParameters {
	if in == nil {
		return nil
	}
	out := new(SnapshotCreateVolumePermissionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotCreateVolumePermissionSpec) DeepCopyInto(out *SnapshotCreateVolumePermissionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotCreateVolumePermissionSpec.
func (in *SnapshotCreateVolumePermissionSpec) DeepCopy() *SnapshotCreateVolumePermissionSpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotCreateVolumePermissionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotCreateVolumePermissionStatus) DeepCopyInto(out *SnapshotCreateVolumePermissionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotCreateVolumePermissionStatus.
func (in *SnapshotCreateVolumePermissionStatus) DeepCopy() *SnapshotCreateVolumePermissionStatus {
	if in == nil {
		return nil
	}
	out := new(SnapshotCreateVolumePermissionStatus)
	in.DeepCopyInto(out)
	return out
}