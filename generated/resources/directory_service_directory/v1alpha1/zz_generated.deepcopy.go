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
func (in *ConnectSettings) DeepCopyInto(out *ConnectSettings) {
	*out = *in
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConnectIps != nil {
		in, out := &in.ConnectIps, &out.ConnectIps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomerDnsIps != nil {
		in, out := &in.CustomerDnsIps, &out.CustomerDnsIps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectSettings.
func (in *ConnectSettings) DeepCopy() *ConnectSettings {
	if in == nil {
		return nil
	}
	out := new(ConnectSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryServiceDirectory) DeepCopyInto(out *DirectoryServiceDirectory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryServiceDirectory.
func (in *DirectoryServiceDirectory) DeepCopy() *DirectoryServiceDirectory {
	if in == nil {
		return nil
	}
	out := new(DirectoryServiceDirectory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DirectoryServiceDirectory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryServiceDirectoryList) DeepCopyInto(out *DirectoryServiceDirectoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DirectoryServiceDirectory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryServiceDirectoryList.
func (in *DirectoryServiceDirectoryList) DeepCopy() *DirectoryServiceDirectoryList {
	if in == nil {
		return nil
	}
	out := new(DirectoryServiceDirectoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DirectoryServiceDirectoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryServiceDirectoryObservation) DeepCopyInto(out *DirectoryServiceDirectoryObservation) {
	*out = *in
	if in.DnsIpAddresses != nil {
		in, out := &in.DnsIpAddresses, &out.DnsIpAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryServiceDirectoryObservation.
func (in *DirectoryServiceDirectoryObservation) DeepCopy() *DirectoryServiceDirectoryObservation {
	if in == nil {
		return nil
	}
	out := new(DirectoryServiceDirectoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryServiceDirectoryParameters) DeepCopyInto(out *DirectoryServiceDirectoryParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.ConnectSettings.DeepCopyInto(&out.ConnectSettings)
	in.VpcSettings.DeepCopyInto(&out.VpcSettings)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryServiceDirectoryParameters.
func (in *DirectoryServiceDirectoryParameters) DeepCopy() *DirectoryServiceDirectoryParameters {
	if in == nil {
		return nil
	}
	out := new(DirectoryServiceDirectoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryServiceDirectorySpec) DeepCopyInto(out *DirectoryServiceDirectorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryServiceDirectorySpec.
func (in *DirectoryServiceDirectorySpec) DeepCopy() *DirectoryServiceDirectorySpec {
	if in == nil {
		return nil
	}
	out := new(DirectoryServiceDirectorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryServiceDirectoryStatus) DeepCopyInto(out *DirectoryServiceDirectoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryServiceDirectoryStatus.
func (in *DirectoryServiceDirectoryStatus) DeepCopy() *DirectoryServiceDirectoryStatus {
	if in == nil {
		return nil
	}
	out := new(DirectoryServiceDirectoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcSettings) DeepCopyInto(out *VpcSettings) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcSettings.
func (in *VpcSettings) DeepCopy() *VpcSettings {
	if in == nil {
		return nil
	}
	out := new(VpcSettings)
	in.DeepCopyInto(out)
	return out
}
