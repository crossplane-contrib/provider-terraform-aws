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
func (in *DbOptionGroup) DeepCopyInto(out *DbOptionGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbOptionGroup.
func (in *DbOptionGroup) DeepCopy() *DbOptionGroup {
	if in == nil {
		return nil
	}
	out := new(DbOptionGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DbOptionGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbOptionGroupList) DeepCopyInto(out *DbOptionGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DbOptionGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbOptionGroupList.
func (in *DbOptionGroupList) DeepCopy() *DbOptionGroupList {
	if in == nil {
		return nil
	}
	out := new(DbOptionGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DbOptionGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbOptionGroupObservation) DeepCopyInto(out *DbOptionGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbOptionGroupObservation.
func (in *DbOptionGroupObservation) DeepCopy() *DbOptionGroupObservation {
	if in == nil {
		return nil
	}
	out := new(DbOptionGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbOptionGroupParameters) DeepCopyInto(out *DbOptionGroupParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Option.DeepCopyInto(&out.Option)
	out.Timeouts = in.Timeouts
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbOptionGroupParameters.
func (in *DbOptionGroupParameters) DeepCopy() *DbOptionGroupParameters {
	if in == nil {
		return nil
	}
	out := new(DbOptionGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbOptionGroupSpec) DeepCopyInto(out *DbOptionGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbOptionGroupSpec.
func (in *DbOptionGroupSpec) DeepCopy() *DbOptionGroupSpec {
	if in == nil {
		return nil
	}
	out := new(DbOptionGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbOptionGroupStatus) DeepCopyInto(out *DbOptionGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbOptionGroupStatus.
func (in *DbOptionGroupStatus) DeepCopy() *DbOptionGroupStatus {
	if in == nil {
		return nil
	}
	out := new(DbOptionGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Option) DeepCopyInto(out *Option) {
	*out = *in
	if in.VpcSecurityGroupMemberships != nil {
		in, out := &in.VpcSecurityGroupMemberships, &out.VpcSecurityGroupMemberships
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DbSecurityGroupMemberships != nil {
		in, out := &in.DbSecurityGroupMemberships, &out.DbSecurityGroupMemberships
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.OptionSettings = in.OptionSettings
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Option.
func (in *Option) DeepCopy() *Option {
	if in == nil {
		return nil
	}
	out := new(Option)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionSettings) DeepCopyInto(out *OptionSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionSettings.
func (in *OptionSettings) DeepCopy() *OptionSettings {
	if in == nil {
		return nil
	}
	out := new(OptionSettings)
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
