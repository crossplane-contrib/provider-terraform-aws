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
func (in *DefaultVpcDhcpOptions) DeepCopyInto(out *DefaultVpcDhcpOptions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcDhcpOptions.
func (in *DefaultVpcDhcpOptions) DeepCopy() *DefaultVpcDhcpOptions {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcDhcpOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultVpcDhcpOptions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcDhcpOptionsList) DeepCopyInto(out *DefaultVpcDhcpOptionsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DefaultVpcDhcpOptions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcDhcpOptionsList.
func (in *DefaultVpcDhcpOptionsList) DeepCopy() *DefaultVpcDhcpOptionsList {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcDhcpOptionsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultVpcDhcpOptionsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcDhcpOptionsObservation) DeepCopyInto(out *DefaultVpcDhcpOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcDhcpOptionsObservation.
func (in *DefaultVpcDhcpOptionsObservation) DeepCopy() *DefaultVpcDhcpOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcDhcpOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcDhcpOptionsParameters) DeepCopyInto(out *DefaultVpcDhcpOptionsParameters) {
	*out = *in
	if in.NetbiosNameServers != nil {
		in, out := &in.NetbiosNameServers, &out.NetbiosNameServers
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcDhcpOptionsParameters.
func (in *DefaultVpcDhcpOptionsParameters) DeepCopy() *DefaultVpcDhcpOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcDhcpOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcDhcpOptionsSpec) DeepCopyInto(out *DefaultVpcDhcpOptionsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcDhcpOptionsSpec.
func (in *DefaultVpcDhcpOptionsSpec) DeepCopy() *DefaultVpcDhcpOptionsSpec {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcDhcpOptionsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultVpcDhcpOptionsStatus) DeepCopyInto(out *DefaultVpcDhcpOptionsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultVpcDhcpOptionsStatus.
func (in *DefaultVpcDhcpOptionsStatus) DeepCopy() *DefaultVpcDhcpOptionsStatus {
	if in == nil {
		return nil
	}
	out := new(DefaultVpcDhcpOptionsStatus)
	in.DeepCopyInto(out)
	return out
}
