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
func (in *DatasyncLocationNfs) DeepCopyInto(out *DatasyncLocationNfs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationNfs.
func (in *DatasyncLocationNfs) DeepCopy() *DatasyncLocationNfs {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationNfs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasyncLocationNfs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationNfsList) DeepCopyInto(out *DatasyncLocationNfsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatasyncLocationNfs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationNfsList.
func (in *DatasyncLocationNfsList) DeepCopy() *DatasyncLocationNfsList {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationNfsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasyncLocationNfsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationNfsObservation) DeepCopyInto(out *DatasyncLocationNfsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationNfsObservation.
func (in *DatasyncLocationNfsObservation) DeepCopy() *DatasyncLocationNfsObservation {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationNfsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationNfsParameters) DeepCopyInto(out *DatasyncLocationNfsParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.OnPremConfig.DeepCopyInto(&out.OnPremConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationNfsParameters.
func (in *DatasyncLocationNfsParameters) DeepCopy() *DatasyncLocationNfsParameters {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationNfsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationNfsSpec) DeepCopyInto(out *DatasyncLocationNfsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationNfsSpec.
func (in *DatasyncLocationNfsSpec) DeepCopy() *DatasyncLocationNfsSpec {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationNfsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationNfsStatus) DeepCopyInto(out *DatasyncLocationNfsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationNfsStatus.
func (in *DatasyncLocationNfsStatus) DeepCopy() *DatasyncLocationNfsStatus {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationNfsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnPremConfig) DeepCopyInto(out *OnPremConfig) {
	*out = *in
	if in.AgentArns != nil {
		in, out := &in.AgentArns, &out.AgentArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnPremConfig.
func (in *OnPremConfig) DeepCopy() *OnPremConfig {
	if in == nil {
		return nil
	}
	out := new(OnPremConfig)
	in.DeepCopyInto(out)
	return out
}