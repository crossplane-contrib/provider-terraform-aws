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
func (in *FsxWindowsFileSystem) DeepCopyInto(out *FsxWindowsFileSystem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsxWindowsFileSystem.
func (in *FsxWindowsFileSystem) DeepCopy() *FsxWindowsFileSystem {
	if in == nil {
		return nil
	}
	out := new(FsxWindowsFileSystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FsxWindowsFileSystem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FsxWindowsFileSystemList) DeepCopyInto(out *FsxWindowsFileSystemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FsxWindowsFileSystem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsxWindowsFileSystemList.
func (in *FsxWindowsFileSystemList) DeepCopy() *FsxWindowsFileSystemList {
	if in == nil {
		return nil
	}
	out := new(FsxWindowsFileSystemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FsxWindowsFileSystemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FsxWindowsFileSystemObservation) DeepCopyInto(out *FsxWindowsFileSystemObservation) {
	*out = *in
	if in.NetworkInterfaceIds != nil {
		in, out := &in.NetworkInterfaceIds, &out.NetworkInterfaceIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsxWindowsFileSystemObservation.
func (in *FsxWindowsFileSystemObservation) DeepCopy() *FsxWindowsFileSystemObservation {
	if in == nil {
		return nil
	}
	out := new(FsxWindowsFileSystemObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FsxWindowsFileSystemParameters) DeepCopyInto(out *FsxWindowsFileSystemParameters) {
	*out = *in
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.SelfManagedActiveDirectory.DeepCopyInto(&out.SelfManagedActiveDirectory)
	out.Timeouts = in.Timeouts
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsxWindowsFileSystemParameters.
func (in *FsxWindowsFileSystemParameters) DeepCopy() *FsxWindowsFileSystemParameters {
	if in == nil {
		return nil
	}
	out := new(FsxWindowsFileSystemParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FsxWindowsFileSystemSpec) DeepCopyInto(out *FsxWindowsFileSystemSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsxWindowsFileSystemSpec.
func (in *FsxWindowsFileSystemSpec) DeepCopy() *FsxWindowsFileSystemSpec {
	if in == nil {
		return nil
	}
	out := new(FsxWindowsFileSystemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FsxWindowsFileSystemStatus) DeepCopyInto(out *FsxWindowsFileSystemStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsxWindowsFileSystemStatus.
func (in *FsxWindowsFileSystemStatus) DeepCopy() *FsxWindowsFileSystemStatus {
	if in == nil {
		return nil
	}
	out := new(FsxWindowsFileSystemStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfManagedActiveDirectory) DeepCopyInto(out *SelfManagedActiveDirectory) {
	*out = *in
	if in.DnsIps != nil {
		in, out := &in.DnsIps, &out.DnsIps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfManagedActiveDirectory.
func (in *SelfManagedActiveDirectory) DeepCopy() *SelfManagedActiveDirectory {
	if in == nil {
		return nil
	}
	out := new(SelfManagedActiveDirectory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timeouts) DeepCopyInto(out *Timeouts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timeouts.
func (in *Timeouts) DeepCopy() *Timeouts {
	if in == nil {
		return nil
	}
	out := new(Timeouts)
	in.DeepCopyInto(out)
	return out
}