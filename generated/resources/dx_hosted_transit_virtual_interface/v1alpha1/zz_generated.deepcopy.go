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
func (in *DxHostedTransitVirtualInterface) DeepCopyInto(out *DxHostedTransitVirtualInterface) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxHostedTransitVirtualInterface.
func (in *DxHostedTransitVirtualInterface) DeepCopy() *DxHostedTransitVirtualInterface {
	if in == nil {
		return nil
	}
	out := new(DxHostedTransitVirtualInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DxHostedTransitVirtualInterface) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxHostedTransitVirtualInterfaceList) DeepCopyInto(out *DxHostedTransitVirtualInterfaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DxHostedTransitVirtualInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxHostedTransitVirtualInterfaceList.
func (in *DxHostedTransitVirtualInterfaceList) DeepCopy() *DxHostedTransitVirtualInterfaceList {
	if in == nil {
		return nil
	}
	out := new(DxHostedTransitVirtualInterfaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DxHostedTransitVirtualInterfaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxHostedTransitVirtualInterfaceObservation) DeepCopyInto(out *DxHostedTransitVirtualInterfaceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxHostedTransitVirtualInterfaceObservation.
func (in *DxHostedTransitVirtualInterfaceObservation) DeepCopy() *DxHostedTransitVirtualInterfaceObservation {
	if in == nil {
		return nil
	}
	out := new(DxHostedTransitVirtualInterfaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxHostedTransitVirtualInterfaceParameters) DeepCopyInto(out *DxHostedTransitVirtualInterfaceParameters) {
	*out = *in
	out.Timeouts = in.Timeouts
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxHostedTransitVirtualInterfaceParameters.
func (in *DxHostedTransitVirtualInterfaceParameters) DeepCopy() *DxHostedTransitVirtualInterfaceParameters {
	if in == nil {
		return nil
	}
	out := new(DxHostedTransitVirtualInterfaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxHostedTransitVirtualInterfaceSpec) DeepCopyInto(out *DxHostedTransitVirtualInterfaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxHostedTransitVirtualInterfaceSpec.
func (in *DxHostedTransitVirtualInterfaceSpec) DeepCopy() *DxHostedTransitVirtualInterfaceSpec {
	if in == nil {
		return nil
	}
	out := new(DxHostedTransitVirtualInterfaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxHostedTransitVirtualInterfaceStatus) DeepCopyInto(out *DxHostedTransitVirtualInterfaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxHostedTransitVirtualInterfaceStatus.
func (in *DxHostedTransitVirtualInterfaceStatus) DeepCopy() *DxHostedTransitVirtualInterfaceStatus {
	if in == nil {
		return nil
	}
	out := new(DxHostedTransitVirtualInterfaceStatus)
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
