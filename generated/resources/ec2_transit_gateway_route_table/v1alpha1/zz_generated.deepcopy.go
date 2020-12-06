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
func (in *Ec2TransitGatewayRouteTable) DeepCopyInto(out *Ec2TransitGatewayRouteTable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayRouteTable.
func (in *Ec2TransitGatewayRouteTable) DeepCopy() *Ec2TransitGatewayRouteTable {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayRouteTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TransitGatewayRouteTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayRouteTableList) DeepCopyInto(out *Ec2TransitGatewayRouteTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ec2TransitGatewayRouteTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayRouteTableList.
func (in *Ec2TransitGatewayRouteTableList) DeepCopy() *Ec2TransitGatewayRouteTableList {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayRouteTableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TransitGatewayRouteTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayRouteTableObservation) DeepCopyInto(out *Ec2TransitGatewayRouteTableObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayRouteTableObservation.
func (in *Ec2TransitGatewayRouteTableObservation) DeepCopy() *Ec2TransitGatewayRouteTableObservation {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayRouteTableObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayRouteTableParameters) DeepCopyInto(out *Ec2TransitGatewayRouteTableParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayRouteTableParameters.
func (in *Ec2TransitGatewayRouteTableParameters) DeepCopy() *Ec2TransitGatewayRouteTableParameters {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayRouteTableParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayRouteTableSpec) DeepCopyInto(out *Ec2TransitGatewayRouteTableSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayRouteTableSpec.
func (in *Ec2TransitGatewayRouteTableSpec) DeepCopy() *Ec2TransitGatewayRouteTableSpec {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayRouteTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TransitGatewayRouteTableStatus) DeepCopyInto(out *Ec2TransitGatewayRouteTableStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TransitGatewayRouteTableStatus.
func (in *Ec2TransitGatewayRouteTableStatus) DeepCopy() *Ec2TransitGatewayRouteTableStatus {
	if in == nil {
		return nil
	}
	out := new(Ec2TransitGatewayRouteTableStatus)
	in.DeepCopyInto(out)
	return out
}
