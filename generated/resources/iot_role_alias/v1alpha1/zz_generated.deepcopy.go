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
func (in *IotRoleAlias) DeepCopyInto(out *IotRoleAlias) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IotRoleAlias.
func (in *IotRoleAlias) DeepCopy() *IotRoleAlias {
	if in == nil {
		return nil
	}
	out := new(IotRoleAlias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IotRoleAlias) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IotRoleAliasList) DeepCopyInto(out *IotRoleAliasList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IotRoleAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IotRoleAliasList.
func (in *IotRoleAliasList) DeepCopy() *IotRoleAliasList {
	if in == nil {
		return nil
	}
	out := new(IotRoleAliasList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IotRoleAliasList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IotRoleAliasObservation) DeepCopyInto(out *IotRoleAliasObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IotRoleAliasObservation.
func (in *IotRoleAliasObservation) DeepCopy() *IotRoleAliasObservation {
	if in == nil {
		return nil
	}
	out := new(IotRoleAliasObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IotRoleAliasParameters) DeepCopyInto(out *IotRoleAliasParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IotRoleAliasParameters.
func (in *IotRoleAliasParameters) DeepCopy() *IotRoleAliasParameters {
	if in == nil {
		return nil
	}
	out := new(IotRoleAliasParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IotRoleAliasSpec) DeepCopyInto(out *IotRoleAliasSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IotRoleAliasSpec.
func (in *IotRoleAliasSpec) DeepCopy() *IotRoleAliasSpec {
	if in == nil {
		return nil
	}
	out := new(IotRoleAliasSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IotRoleAliasStatus) DeepCopyInto(out *IotRoleAliasStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IotRoleAliasStatus.
func (in *IotRoleAliasStatus) DeepCopy() *IotRoleAliasStatus {
	if in == nil {
		return nil
	}
	out := new(IotRoleAliasStatus)
	in.DeepCopyInto(out)
	return out
}
