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
func (in *ConfigConfigRule) DeepCopyInto(out *ConfigConfigRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConfigRule.
func (in *ConfigConfigRule) DeepCopy() *ConfigConfigRule {
	if in == nil {
		return nil
	}
	out := new(ConfigConfigRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigConfigRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConfigRuleList) DeepCopyInto(out *ConfigConfigRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigConfigRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConfigRuleList.
func (in *ConfigConfigRuleList) DeepCopy() *ConfigConfigRuleList {
	if in == nil {
		return nil
	}
	out := new(ConfigConfigRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigConfigRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConfigRuleObservation) DeepCopyInto(out *ConfigConfigRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConfigRuleObservation.
func (in *ConfigConfigRuleObservation) DeepCopy() *ConfigConfigRuleObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigConfigRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConfigRuleParameters) DeepCopyInto(out *ConfigConfigRuleParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Scope.DeepCopyInto(&out.Scope)
	in.Source.DeepCopyInto(&out.Source)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConfigRuleParameters.
func (in *ConfigConfigRuleParameters) DeepCopy() *ConfigConfigRuleParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigConfigRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConfigRuleSpec) DeepCopyInto(out *ConfigConfigRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConfigRuleSpec.
func (in *ConfigConfigRuleSpec) DeepCopy() *ConfigConfigRuleSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigConfigRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConfigRuleStatus) DeepCopyInto(out *ConfigConfigRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConfigRuleStatus.
func (in *ConfigConfigRuleStatus) DeepCopy() *ConfigConfigRuleStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigConfigRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scope) DeepCopyInto(out *Scope) {
	*out = *in
	if in.ComplianceResourceTypes != nil {
		in, out := &in.ComplianceResourceTypes, &out.ComplianceResourceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scope.
func (in *Scope) DeepCopy() *Scope {
	if in == nil {
		return nil
	}
	out := new(Scope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	if in.SourceDetail != nil {
		in, out := &in.SourceDetail, &out.SourceDetail
		*out = make([]SourceDetail, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceDetail) DeepCopyInto(out *SourceDetail) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceDetail.
func (in *SourceDetail) DeepCopy() *SourceDetail {
	if in == nil {
		return nil
	}
	out := new(SourceDetail)
	in.DeepCopyInto(out)
	return out
}