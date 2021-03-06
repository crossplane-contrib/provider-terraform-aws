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
func (in *PinpointEmailChannel) DeepCopyInto(out *PinpointEmailChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointEmailChannel.
func (in *PinpointEmailChannel) DeepCopy() *PinpointEmailChannel {
	if in == nil {
		return nil
	}
	out := new(PinpointEmailChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PinpointEmailChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointEmailChannelList) DeepCopyInto(out *PinpointEmailChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PinpointEmailChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointEmailChannelList.
func (in *PinpointEmailChannelList) DeepCopy() *PinpointEmailChannelList {
	if in == nil {
		return nil
	}
	out := new(PinpointEmailChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PinpointEmailChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointEmailChannelObservation) DeepCopyInto(out *PinpointEmailChannelObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointEmailChannelObservation.
func (in *PinpointEmailChannelObservation) DeepCopy() *PinpointEmailChannelObservation {
	if in == nil {
		return nil
	}
	out := new(PinpointEmailChannelObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointEmailChannelParameters) DeepCopyInto(out *PinpointEmailChannelParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointEmailChannelParameters.
func (in *PinpointEmailChannelParameters) DeepCopy() *PinpointEmailChannelParameters {
	if in == nil {
		return nil
	}
	out := new(PinpointEmailChannelParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointEmailChannelSpec) DeepCopyInto(out *PinpointEmailChannelSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointEmailChannelSpec.
func (in *PinpointEmailChannelSpec) DeepCopy() *PinpointEmailChannelSpec {
	if in == nil {
		return nil
	}
	out := new(PinpointEmailChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointEmailChannelStatus) DeepCopyInto(out *PinpointEmailChannelStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointEmailChannelStatus.
func (in *PinpointEmailChannelStatus) DeepCopy() *PinpointEmailChannelStatus {
	if in == nil {
		return nil
	}
	out := new(PinpointEmailChannelStatus)
	in.DeepCopyInto(out)
	return out
}
