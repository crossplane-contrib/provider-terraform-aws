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
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivatedRule) DeepCopyInto(out *ActivatedRule) {
	*out = *in
	out.Action = in.Action
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivatedRule.
func (in *ActivatedRule) DeepCopy() *ActivatedRule {
	if in == nil {
		return nil
	}
	out := new(ActivatedRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRuleGroup) DeepCopyInto(out *WafregionalRuleGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRuleGroup.
func (in *WafregionalRuleGroup) DeepCopy() *WafregionalRuleGroup {
	if in == nil {
		return nil
	}
	out := new(WafregionalRuleGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafregionalRuleGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRuleGroupList) DeepCopyInto(out *WafregionalRuleGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WafregionalRuleGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRuleGroupList.
func (in *WafregionalRuleGroupList) DeepCopy() *WafregionalRuleGroupList {
	if in == nil {
		return nil
	}
	out := new(WafregionalRuleGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafregionalRuleGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRuleGroupObservation) DeepCopyInto(out *WafregionalRuleGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRuleGroupObservation.
func (in *WafregionalRuleGroupObservation) DeepCopy() *WafregionalRuleGroupObservation {
	if in == nil {
		return nil
	}
	out := new(WafregionalRuleGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRuleGroupParameters) DeepCopyInto(out *WafregionalRuleGroupParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.ActivatedRule = in.ActivatedRule
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRuleGroupParameters.
func (in *WafregionalRuleGroupParameters) DeepCopy() *WafregionalRuleGroupParameters {
	if in == nil {
		return nil
	}
	out := new(WafregionalRuleGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRuleGroupSpec) DeepCopyInto(out *WafregionalRuleGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRuleGroupSpec.
func (in *WafregionalRuleGroupSpec) DeepCopy() *WafregionalRuleGroupSpec {
	if in == nil {
		return nil
	}
	out := new(WafregionalRuleGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRuleGroupStatus) DeepCopyInto(out *WafregionalRuleGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRuleGroupStatus.
func (in *WafregionalRuleGroupStatus) DeepCopy() *WafregionalRuleGroupStatus {
	if in == nil {
		return nil
	}
	out := new(WafregionalRuleGroupStatus)
	in.DeepCopyInto(out)
	return out
}
