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
func (in *NeptuneClusterParameterGroup) DeepCopyInto(out *NeptuneClusterParameterGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeptuneClusterParameterGroup.
func (in *NeptuneClusterParameterGroup) DeepCopy() *NeptuneClusterParameterGroup {
	if in == nil {
		return nil
	}
	out := new(NeptuneClusterParameterGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NeptuneClusterParameterGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeptuneClusterParameterGroupList) DeepCopyInto(out *NeptuneClusterParameterGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NeptuneClusterParameterGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeptuneClusterParameterGroupList.
func (in *NeptuneClusterParameterGroupList) DeepCopy() *NeptuneClusterParameterGroupList {
	if in == nil {
		return nil
	}
	out := new(NeptuneClusterParameterGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NeptuneClusterParameterGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeptuneClusterParameterGroupObservation) DeepCopyInto(out *NeptuneClusterParameterGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeptuneClusterParameterGroupObservation.
func (in *NeptuneClusterParameterGroupObservation) DeepCopy() *NeptuneClusterParameterGroupObservation {
	if in == nil {
		return nil
	}
	out := new(NeptuneClusterParameterGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeptuneClusterParameterGroupParameters) DeepCopyInto(out *NeptuneClusterParameterGroupParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Parameter = in.Parameter
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeptuneClusterParameterGroupParameters.
func (in *NeptuneClusterParameterGroupParameters) DeepCopy() *NeptuneClusterParameterGroupParameters {
	if in == nil {
		return nil
	}
	out := new(NeptuneClusterParameterGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeptuneClusterParameterGroupSpec) DeepCopyInto(out *NeptuneClusterParameterGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeptuneClusterParameterGroupSpec.
func (in *NeptuneClusterParameterGroupSpec) DeepCopy() *NeptuneClusterParameterGroupSpec {
	if in == nil {
		return nil
	}
	out := new(NeptuneClusterParameterGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeptuneClusterParameterGroupStatus) DeepCopyInto(out *NeptuneClusterParameterGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeptuneClusterParameterGroupStatus.
func (in *NeptuneClusterParameterGroupStatus) DeepCopy() *NeptuneClusterParameterGroupStatus {
	if in == nil {
		return nil
	}
	out := new(NeptuneClusterParameterGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameter) DeepCopyInto(out *Parameter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameter.
func (in *Parameter) DeepCopy() *Parameter {
	if in == nil {
		return nil
	}
	out := new(Parameter)
	in.DeepCopyInto(out)
	return out
}
