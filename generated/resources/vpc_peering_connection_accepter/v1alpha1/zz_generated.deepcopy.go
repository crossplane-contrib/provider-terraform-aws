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
func (in *Accepter) DeepCopyInto(out *Accepter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Accepter.
func (in *Accepter) DeepCopy() *Accepter {
	if in == nil {
		return nil
	}
	out := new(Accepter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Requester) DeepCopyInto(out *Requester) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Requester.
func (in *Requester) DeepCopy() *Requester {
	if in == nil {
		return nil
	}
	out := new(Requester)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcPeeringConnectionAccepter) DeepCopyInto(out *VpcPeeringConnectionAccepter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcPeeringConnectionAccepter.
func (in *VpcPeeringConnectionAccepter) DeepCopy() *VpcPeeringConnectionAccepter {
	if in == nil {
		return nil
	}
	out := new(VpcPeeringConnectionAccepter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcPeeringConnectionAccepter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcPeeringConnectionAccepterList) DeepCopyInto(out *VpcPeeringConnectionAccepterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpcPeeringConnectionAccepter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcPeeringConnectionAccepterList.
func (in *VpcPeeringConnectionAccepterList) DeepCopy() *VpcPeeringConnectionAccepterList {
	if in == nil {
		return nil
	}
	out := new(VpcPeeringConnectionAccepterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcPeeringConnectionAccepterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcPeeringConnectionAccepterObservation) DeepCopyInto(out *VpcPeeringConnectionAccepterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcPeeringConnectionAccepterObservation.
func (in *VpcPeeringConnectionAccepterObservation) DeepCopy() *VpcPeeringConnectionAccepterObservation {
	if in == nil {
		return nil
	}
	out := new(VpcPeeringConnectionAccepterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcPeeringConnectionAccepterParameters) DeepCopyInto(out *VpcPeeringConnectionAccepterParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Accepter = in.Accepter
	out.Requester = in.Requester
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcPeeringConnectionAccepterParameters.
func (in *VpcPeeringConnectionAccepterParameters) DeepCopy() *VpcPeeringConnectionAccepterParameters {
	if in == nil {
		return nil
	}
	out := new(VpcPeeringConnectionAccepterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcPeeringConnectionAccepterSpec) DeepCopyInto(out *VpcPeeringConnectionAccepterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcPeeringConnectionAccepterSpec.
func (in *VpcPeeringConnectionAccepterSpec) DeepCopy() *VpcPeeringConnectionAccepterSpec {
	if in == nil {
		return nil
	}
	out := new(VpcPeeringConnectionAccepterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcPeeringConnectionAccepterStatus) DeepCopyInto(out *VpcPeeringConnectionAccepterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcPeeringConnectionAccepterStatus.
func (in *VpcPeeringConnectionAccepterStatus) DeepCopy() *VpcPeeringConnectionAccepterStatus {
	if in == nil {
		return nil
	}
	out := new(VpcPeeringConnectionAccepterStatus)
	in.DeepCopyInto(out)
	return out
}
