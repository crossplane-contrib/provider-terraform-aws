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
func (in *WafRegexPatternSet) DeepCopyInto(out *WafRegexPatternSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafRegexPatternSet.
func (in *WafRegexPatternSet) DeepCopy() *WafRegexPatternSet {
	if in == nil {
		return nil
	}
	out := new(WafRegexPatternSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafRegexPatternSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafRegexPatternSetList) DeepCopyInto(out *WafRegexPatternSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WafRegexPatternSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafRegexPatternSetList.
func (in *WafRegexPatternSetList) DeepCopy() *WafRegexPatternSetList {
	if in == nil {
		return nil
	}
	out := new(WafRegexPatternSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafRegexPatternSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafRegexPatternSetObservation) DeepCopyInto(out *WafRegexPatternSetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafRegexPatternSetObservation.
func (in *WafRegexPatternSetObservation) DeepCopy() *WafRegexPatternSetObservation {
	if in == nil {
		return nil
	}
	out := new(WafRegexPatternSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafRegexPatternSetParameters) DeepCopyInto(out *WafRegexPatternSetParameters) {
	*out = *in
	if in.RegexPatternStrings != nil {
		in, out := &in.RegexPatternStrings, &out.RegexPatternStrings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafRegexPatternSetParameters.
func (in *WafRegexPatternSetParameters) DeepCopy() *WafRegexPatternSetParameters {
	if in == nil {
		return nil
	}
	out := new(WafRegexPatternSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafRegexPatternSetSpec) DeepCopyInto(out *WafRegexPatternSetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafRegexPatternSetSpec.
func (in *WafRegexPatternSetSpec) DeepCopy() *WafRegexPatternSetSpec {
	if in == nil {
		return nil
	}
	out := new(WafRegexPatternSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafRegexPatternSetStatus) DeepCopyInto(out *WafRegexPatternSetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafRegexPatternSetStatus.
func (in *WafRegexPatternSetStatus) DeepCopy() *WafRegexPatternSetStatus {
	if in == nil {
		return nil
	}
	out := new(WafRegexPatternSetStatus)
	in.DeepCopyInto(out)
	return out
}
