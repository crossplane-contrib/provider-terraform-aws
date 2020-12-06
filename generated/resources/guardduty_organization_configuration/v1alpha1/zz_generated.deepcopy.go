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
func (in *GuarddutyOrganizationConfiguration) DeepCopyInto(out *GuarddutyOrganizationConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuarddutyOrganizationConfiguration.
func (in *GuarddutyOrganizationConfiguration) DeepCopy() *GuarddutyOrganizationConfiguration {
	if in == nil {
		return nil
	}
	out := new(GuarddutyOrganizationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GuarddutyOrganizationConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuarddutyOrganizationConfigurationList) DeepCopyInto(out *GuarddutyOrganizationConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GuarddutyOrganizationConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuarddutyOrganizationConfigurationList.
func (in *GuarddutyOrganizationConfigurationList) DeepCopy() *GuarddutyOrganizationConfigurationList {
	if in == nil {
		return nil
	}
	out := new(GuarddutyOrganizationConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GuarddutyOrganizationConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuarddutyOrganizationConfigurationObservation) DeepCopyInto(out *GuarddutyOrganizationConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuarddutyOrganizationConfigurationObservation.
func (in *GuarddutyOrganizationConfigurationObservation) DeepCopy() *GuarddutyOrganizationConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(GuarddutyOrganizationConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuarddutyOrganizationConfigurationParameters) DeepCopyInto(out *GuarddutyOrganizationConfigurationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuarddutyOrganizationConfigurationParameters.
func (in *GuarddutyOrganizationConfigurationParameters) DeepCopy() *GuarddutyOrganizationConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(GuarddutyOrganizationConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuarddutyOrganizationConfigurationSpec) DeepCopyInto(out *GuarddutyOrganizationConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuarddutyOrganizationConfigurationSpec.
func (in *GuarddutyOrganizationConfigurationSpec) DeepCopy() *GuarddutyOrganizationConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(GuarddutyOrganizationConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuarddutyOrganizationConfigurationStatus) DeepCopyInto(out *GuarddutyOrganizationConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuarddutyOrganizationConfigurationStatus.
func (in *GuarddutyOrganizationConfigurationStatus) DeepCopy() *GuarddutyOrganizationConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(GuarddutyOrganizationConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}
