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
func (in *DatasyncLocationFsxWindowsFileSystem) DeepCopyInto(out *DatasyncLocationFsxWindowsFileSystem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationFsxWindowsFileSystem.
func (in *DatasyncLocationFsxWindowsFileSystem) DeepCopy() *DatasyncLocationFsxWindowsFileSystem {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationFsxWindowsFileSystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasyncLocationFsxWindowsFileSystem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationFsxWindowsFileSystemList) DeepCopyInto(out *DatasyncLocationFsxWindowsFileSystemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatasyncLocationFsxWindowsFileSystem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationFsxWindowsFileSystemList.
func (in *DatasyncLocationFsxWindowsFileSystemList) DeepCopy() *DatasyncLocationFsxWindowsFileSystemList {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationFsxWindowsFileSystemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasyncLocationFsxWindowsFileSystemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationFsxWindowsFileSystemObservation) DeepCopyInto(out *DatasyncLocationFsxWindowsFileSystemObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationFsxWindowsFileSystemObservation.
func (in *DatasyncLocationFsxWindowsFileSystemObservation) DeepCopy() *DatasyncLocationFsxWindowsFileSystemObservation {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationFsxWindowsFileSystemObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationFsxWindowsFileSystemParameters) DeepCopyInto(out *DatasyncLocationFsxWindowsFileSystemParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityGroupArns != nil {
		in, out := &in.SecurityGroupArns, &out.SecurityGroupArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationFsxWindowsFileSystemParameters.
func (in *DatasyncLocationFsxWindowsFileSystemParameters) DeepCopy() *DatasyncLocationFsxWindowsFileSystemParameters {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationFsxWindowsFileSystemParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationFsxWindowsFileSystemSpec) DeepCopyInto(out *DatasyncLocationFsxWindowsFileSystemSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationFsxWindowsFileSystemSpec.
func (in *DatasyncLocationFsxWindowsFileSystemSpec) DeepCopy() *DatasyncLocationFsxWindowsFileSystemSpec {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationFsxWindowsFileSystemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationFsxWindowsFileSystemStatus) DeepCopyInto(out *DatasyncLocationFsxWindowsFileSystemStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationFsxWindowsFileSystemStatus.
func (in *DatasyncLocationFsxWindowsFileSystemStatus) DeepCopy() *DatasyncLocationFsxWindowsFileSystemStatus {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationFsxWindowsFileSystemStatus)
	in.DeepCopyInto(out)
	return out
}
