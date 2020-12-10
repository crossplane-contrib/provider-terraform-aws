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
func (in *ConfigOrganizationCustomRule) DeepCopyInto(out *ConfigOrganizationCustomRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigOrganizationCustomRule.
func (in *ConfigOrganizationCustomRule) DeepCopy() *ConfigOrganizationCustomRule {
	if in == nil {
		return nil
	}
	out := new(ConfigOrganizationCustomRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigOrganizationCustomRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigOrganizationCustomRuleList) DeepCopyInto(out *ConfigOrganizationCustomRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigOrganizationCustomRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigOrganizationCustomRuleList.
func (in *ConfigOrganizationCustomRuleList) DeepCopy() *ConfigOrganizationCustomRuleList {
	if in == nil {
		return nil
	}
	out := new(ConfigOrganizationCustomRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigOrganizationCustomRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigOrganizationCustomRuleObservation) DeepCopyInto(out *ConfigOrganizationCustomRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigOrganizationCustomRuleObservation.
func (in *ConfigOrganizationCustomRuleObservation) DeepCopy() *ConfigOrganizationCustomRuleObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigOrganizationCustomRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigOrganizationCustomRuleParameters) DeepCopyInto(out *ConfigOrganizationCustomRuleParameters) {
	*out = *in
	if in.ExcludedAccounts != nil {
		in, out := &in.ExcludedAccounts, &out.ExcludedAccounts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TriggerTypes != nil {
		in, out := &in.TriggerTypes, &out.TriggerTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceTypesScope != nil {
		in, out := &in.ResourceTypesScope, &out.ResourceTypesScope
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Timeouts = in.Timeouts
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigOrganizationCustomRuleParameters.
func (in *ConfigOrganizationCustomRuleParameters) DeepCopy() *ConfigOrganizationCustomRuleParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigOrganizationCustomRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigOrganizationCustomRuleSpec) DeepCopyInto(out *ConfigOrganizationCustomRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigOrganizationCustomRuleSpec.
func (in *ConfigOrganizationCustomRuleSpec) DeepCopy() *ConfigOrganizationCustomRuleSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigOrganizationCustomRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigOrganizationCustomRuleStatus) DeepCopyInto(out *ConfigOrganizationCustomRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigOrganizationCustomRuleStatus.
func (in *ConfigOrganizationCustomRuleStatus) DeepCopy() *ConfigOrganizationCustomRuleStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigOrganizationCustomRuleStatus)
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
