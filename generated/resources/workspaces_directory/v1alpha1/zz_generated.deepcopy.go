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
func (in *SelfServicePermissions) DeepCopyInto(out *SelfServicePermissions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfServicePermissions.
func (in *SelfServicePermissions) DeepCopy() *SelfServicePermissions {
	if in == nil {
		return nil
	}
	out := new(SelfServicePermissions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacesDirectory) DeepCopyInto(out *WorkspacesDirectory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacesDirectory.
func (in *WorkspacesDirectory) DeepCopy() *WorkspacesDirectory {
	if in == nil {
		return nil
	}
	out := new(WorkspacesDirectory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspacesDirectory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacesDirectoryList) DeepCopyInto(out *WorkspacesDirectoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkspacesDirectory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacesDirectoryList.
func (in *WorkspacesDirectoryList) DeepCopy() *WorkspacesDirectoryList {
	if in == nil {
		return nil
	}
	out := new(WorkspacesDirectoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspacesDirectoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacesDirectoryObservation) DeepCopyInto(out *WorkspacesDirectoryObservation) {
	*out = *in
	if in.DnsIpAddresses != nil {
		in, out := &in.DnsIpAddresses, &out.DnsIpAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IpGroupIds != nil {
		in, out := &in.IpGroupIds, &out.IpGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacesDirectoryObservation.
func (in *WorkspacesDirectoryObservation) DeepCopy() *WorkspacesDirectoryObservation {
	if in == nil {
		return nil
	}
	out := new(WorkspacesDirectoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacesDirectoryParameters) DeepCopyInto(out *WorkspacesDirectoryParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.SelfServicePermissions = in.SelfServicePermissions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacesDirectoryParameters.
func (in *WorkspacesDirectoryParameters) DeepCopy() *WorkspacesDirectoryParameters {
	if in == nil {
		return nil
	}
	out := new(WorkspacesDirectoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacesDirectorySpec) DeepCopyInto(out *WorkspacesDirectorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacesDirectorySpec.
func (in *WorkspacesDirectorySpec) DeepCopy() *WorkspacesDirectorySpec {
	if in == nil {
		return nil
	}
	out := new(WorkspacesDirectorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacesDirectoryStatus) DeepCopyInto(out *WorkspacesDirectoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacesDirectoryStatus.
func (in *WorkspacesDirectoryStatus) DeepCopy() *WorkspacesDirectoryStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspacesDirectoryStatus)
	in.DeepCopyInto(out)
	return out
}
