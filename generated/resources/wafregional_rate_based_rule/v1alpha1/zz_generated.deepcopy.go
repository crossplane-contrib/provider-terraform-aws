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
func (in *Predicate) DeepCopyInto(out *Predicate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Predicate.
func (in *Predicate) DeepCopy() *Predicate {
	if in == nil {
		return nil
	}
	out := new(Predicate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRateBasedRule) DeepCopyInto(out *WafregionalRateBasedRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRateBasedRule.
func (in *WafregionalRateBasedRule) DeepCopy() *WafregionalRateBasedRule {
	if in == nil {
		return nil
	}
	out := new(WafregionalRateBasedRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafregionalRateBasedRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRateBasedRuleList) DeepCopyInto(out *WafregionalRateBasedRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WafregionalRateBasedRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRateBasedRuleList.
func (in *WafregionalRateBasedRuleList) DeepCopy() *WafregionalRateBasedRuleList {
	if in == nil {
		return nil
	}
	out := new(WafregionalRateBasedRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafregionalRateBasedRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRateBasedRuleObservation) DeepCopyInto(out *WafregionalRateBasedRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRateBasedRuleObservation.
func (in *WafregionalRateBasedRuleObservation) DeepCopy() *WafregionalRateBasedRuleObservation {
	if in == nil {
		return nil
	}
	out := new(WafregionalRateBasedRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRateBasedRuleParameters) DeepCopyInto(out *WafregionalRateBasedRuleParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Predicate = in.Predicate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRateBasedRuleParameters.
func (in *WafregionalRateBasedRuleParameters) DeepCopy() *WafregionalRateBasedRuleParameters {
	if in == nil {
		return nil
	}
	out := new(WafregionalRateBasedRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRateBasedRuleSpec) DeepCopyInto(out *WafregionalRateBasedRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRateBasedRuleSpec.
func (in *WafregionalRateBasedRuleSpec) DeepCopy() *WafregionalRateBasedRuleSpec {
	if in == nil {
		return nil
	}
	out := new(WafregionalRateBasedRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafregionalRateBasedRuleStatus) DeepCopyInto(out *WafregionalRateBasedRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafregionalRateBasedRuleStatus.
func (in *WafregionalRateBasedRuleStatus) DeepCopy() *WafregionalRateBasedRuleStatus {
	if in == nil {
		return nil
	}
	out := new(WafregionalRateBasedRuleStatus)
	in.DeepCopyInto(out)
	return out
}
