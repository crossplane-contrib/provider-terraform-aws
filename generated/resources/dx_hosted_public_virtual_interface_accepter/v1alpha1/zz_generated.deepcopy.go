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
func (in *DxHostedPublicVirtualInterfaceAccepter) DeepCopyInto(out *DxHostedPublicVirtualInterfaceAccepter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxHostedPublicVirtualInterfaceAccepter.
func (in *DxHostedPublicVirtualInterfaceAccepter) DeepCopy() *DxHostedPublicVirtualInterfaceAccepter {
	if in == nil {
		return nil
	}
	out := new(DxHostedPublicVirtualInterfaceAccepter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DxHostedPublicVirtualInterfaceAccepter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxHostedPublicVirtualInterfaceAccepterList) DeepCopyInto(out *DxHostedPublicVirtualInterfaceAccepterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DxHostedPublicVirtualInterfaceAccepter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxHostedPublicVirtualInterfaceAccepterList.
func (in *DxHostedPublicVirtualInterfaceAccepterList) DeepCopy() *DxHostedPublicVirtualInterfaceAccepterList {
	if in == nil {
		return nil
	}
	out := new(DxHostedPublicVirtualInterfaceAccepterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DxHostedPublicVirtualInterfaceAccepterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxHostedPublicVirtualInterfaceAccepterObservation) DeepCopyInto(out *DxHostedPublicVirtualInterfaceAccepterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxHostedPublicVirtualInterfaceAccepterObservation.
func (in *DxHostedPublicVirtualInterfaceAccepterObservation) DeepCopy() *DxHostedPublicVirtualInterfaceAccepterObservation {
	if in == nil {
		return nil
	}
	out := new(DxHostedPublicVirtualInterfaceAccepterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxHostedPublicVirtualInterfaceAccepterParameters) DeepCopyInto(out *DxHostedPublicVirtualInterfaceAccepterParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Timeouts = in.Timeouts
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxHostedPublicVirtualInterfaceAccepterParameters.
func (in *DxHostedPublicVirtualInterfaceAccepterParameters) DeepCopy() *DxHostedPublicVirtualInterfaceAccepterParameters {
	if in == nil {
		return nil
	}
	out := new(DxHostedPublicVirtualInterfaceAccepterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxHostedPublicVirtualInterfaceAccepterSpec) DeepCopyInto(out *DxHostedPublicVirtualInterfaceAccepterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxHostedPublicVirtualInterfaceAccepterSpec.
func (in *DxHostedPublicVirtualInterfaceAccepterSpec) DeepCopy() *DxHostedPublicVirtualInterfaceAccepterSpec {
	if in == nil {
		return nil
	}
	out := new(DxHostedPublicVirtualInterfaceAccepterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxHostedPublicVirtualInterfaceAccepterStatus) DeepCopyInto(out *DxHostedPublicVirtualInterfaceAccepterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxHostedPublicVirtualInterfaceAccepterStatus.
func (in *DxHostedPublicVirtualInterfaceAccepterStatus) DeepCopy() *DxHostedPublicVirtualInterfaceAccepterStatus {
	if in == nil {
		return nil
	}
	out := new(DxHostedPublicVirtualInterfaceAccepterStatus)
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
