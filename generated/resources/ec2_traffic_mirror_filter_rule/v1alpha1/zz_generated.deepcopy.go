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
func (in *DestinationPortRange) DeepCopyInto(out *DestinationPortRange) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationPortRange.
func (in *DestinationPortRange) DeepCopy() *DestinationPortRange {
	if in == nil {
		return nil
	}
	out := new(DestinationPortRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorFilterRule) DeepCopyInto(out *Ec2TrafficMirrorFilterRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorFilterRule.
func (in *Ec2TrafficMirrorFilterRule) DeepCopy() *Ec2TrafficMirrorFilterRule {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorFilterRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TrafficMirrorFilterRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorFilterRuleList) DeepCopyInto(out *Ec2TrafficMirrorFilterRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ec2TrafficMirrorFilterRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorFilterRuleList.
func (in *Ec2TrafficMirrorFilterRuleList) DeepCopy() *Ec2TrafficMirrorFilterRuleList {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorFilterRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TrafficMirrorFilterRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorFilterRuleObservation) DeepCopyInto(out *Ec2TrafficMirrorFilterRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorFilterRuleObservation.
func (in *Ec2TrafficMirrorFilterRuleObservation) DeepCopy() *Ec2TrafficMirrorFilterRuleObservation {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorFilterRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorFilterRuleParameters) DeepCopyInto(out *Ec2TrafficMirrorFilterRuleParameters) {
	*out = *in
	out.DestinationPortRange = in.DestinationPortRange
	out.SourcePortRange = in.SourcePortRange
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorFilterRuleParameters.
func (in *Ec2TrafficMirrorFilterRuleParameters) DeepCopy() *Ec2TrafficMirrorFilterRuleParameters {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorFilterRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorFilterRuleSpec) DeepCopyInto(out *Ec2TrafficMirrorFilterRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorFilterRuleSpec.
func (in *Ec2TrafficMirrorFilterRuleSpec) DeepCopy() *Ec2TrafficMirrorFilterRuleSpec {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorFilterRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorFilterRuleStatus) DeepCopyInto(out *Ec2TrafficMirrorFilterRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorFilterRuleStatus.
func (in *Ec2TrafficMirrorFilterRuleStatus) DeepCopy() *Ec2TrafficMirrorFilterRuleStatus {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorFilterRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourcePortRange) DeepCopyInto(out *SourcePortRange) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourcePortRange.
func (in *SourcePortRange) DeepCopy() *SourcePortRange {
	if in == nil {
		return nil
	}
	out := new(SourcePortRange)
	in.DeepCopyInto(out)
	return out
}