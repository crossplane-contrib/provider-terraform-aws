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
func (in *Ec2ClientVpnRoute) DeepCopyInto(out *Ec2ClientVpnRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2ClientVpnRoute.
func (in *Ec2ClientVpnRoute) DeepCopy() *Ec2ClientVpnRoute {
	if in == nil {
		return nil
	}
	out := new(Ec2ClientVpnRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2ClientVpnRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2ClientVpnRouteList) DeepCopyInto(out *Ec2ClientVpnRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ec2ClientVpnRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2ClientVpnRouteList.
func (in *Ec2ClientVpnRouteList) DeepCopy() *Ec2ClientVpnRouteList {
	if in == nil {
		return nil
	}
	out := new(Ec2ClientVpnRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2ClientVpnRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2ClientVpnRouteObservation) DeepCopyInto(out *Ec2ClientVpnRouteObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2ClientVpnRouteObservation.
func (in *Ec2ClientVpnRouteObservation) DeepCopy() *Ec2ClientVpnRouteObservation {
	if in == nil {
		return nil
	}
	out := new(Ec2ClientVpnRouteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2ClientVpnRouteParameters) DeepCopyInto(out *Ec2ClientVpnRouteParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2ClientVpnRouteParameters.
func (in *Ec2ClientVpnRouteParameters) DeepCopy() *Ec2ClientVpnRouteParameters {
	if in == nil {
		return nil
	}
	out := new(Ec2ClientVpnRouteParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2ClientVpnRouteSpec) DeepCopyInto(out *Ec2ClientVpnRouteSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2ClientVpnRouteSpec.
func (in *Ec2ClientVpnRouteSpec) DeepCopy() *Ec2ClientVpnRouteSpec {
	if in == nil {
		return nil
	}
	out := new(Ec2ClientVpnRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2ClientVpnRouteStatus) DeepCopyInto(out *Ec2ClientVpnRouteStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2ClientVpnRouteStatus.
func (in *Ec2ClientVpnRouteStatus) DeepCopy() *Ec2ClientVpnRouteStatus {
	if in == nil {
		return nil
	}
	out := new(Ec2ClientVpnRouteStatus)
	in.DeepCopyInto(out)
	return out
}