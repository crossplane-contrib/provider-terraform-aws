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
func (in *VpnGatewayRoutePropagation) DeepCopyInto(out *VpnGatewayRoutePropagation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayRoutePropagation.
func (in *VpnGatewayRoutePropagation) DeepCopy() *VpnGatewayRoutePropagation {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayRoutePropagation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnGatewayRoutePropagation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayRoutePropagationList) DeepCopyInto(out *VpnGatewayRoutePropagationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpnGatewayRoutePropagation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayRoutePropagationList.
func (in *VpnGatewayRoutePropagationList) DeepCopy() *VpnGatewayRoutePropagationList {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayRoutePropagationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnGatewayRoutePropagationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayRoutePropagationObservation) DeepCopyInto(out *VpnGatewayRoutePropagationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayRoutePropagationObservation.
func (in *VpnGatewayRoutePropagationObservation) DeepCopy() *VpnGatewayRoutePropagationObservation {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayRoutePropagationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayRoutePropagationParameters) DeepCopyInto(out *VpnGatewayRoutePropagationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayRoutePropagationParameters.
func (in *VpnGatewayRoutePropagationParameters) DeepCopy() *VpnGatewayRoutePropagationParameters {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayRoutePropagationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayRoutePropagationSpec) DeepCopyInto(out *VpnGatewayRoutePropagationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayRoutePropagationSpec.
func (in *VpnGatewayRoutePropagationSpec) DeepCopy() *VpnGatewayRoutePropagationSpec {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayRoutePropagationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnGatewayRoutePropagationStatus) DeepCopyInto(out *VpnGatewayRoutePropagationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnGatewayRoutePropagationStatus.
func (in *VpnGatewayRoutePropagationStatus) DeepCopy() *VpnGatewayRoutePropagationStatus {
	if in == nil {
		return nil
	}
	out := new(VpnGatewayRoutePropagationStatus)
	in.DeepCopyInto(out)
	return out
}
