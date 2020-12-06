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
func (in *DocdbClusterParameterGroup) DeepCopyInto(out *DocdbClusterParameterGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DocdbClusterParameterGroup.
func (in *DocdbClusterParameterGroup) DeepCopy() *DocdbClusterParameterGroup {
	if in == nil {
		return nil
	}
	out := new(DocdbClusterParameterGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DocdbClusterParameterGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DocdbClusterParameterGroupList) DeepCopyInto(out *DocdbClusterParameterGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DocdbClusterParameterGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DocdbClusterParameterGroupList.
func (in *DocdbClusterParameterGroupList) DeepCopy() *DocdbClusterParameterGroupList {
	if in == nil {
		return nil
	}
	out := new(DocdbClusterParameterGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DocdbClusterParameterGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DocdbClusterParameterGroupObservation) DeepCopyInto(out *DocdbClusterParameterGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DocdbClusterParameterGroupObservation.
func (in *DocdbClusterParameterGroupObservation) DeepCopy() *DocdbClusterParameterGroupObservation {
	if in == nil {
		return nil
	}
	out := new(DocdbClusterParameterGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DocdbClusterParameterGroupParameters) DeepCopyInto(out *DocdbClusterParameterGroupParameters) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DocdbClusterParameterGroupParameters.
func (in *DocdbClusterParameterGroupParameters) DeepCopy() *DocdbClusterParameterGroupParameters {
	if in == nil {
		return nil
	}
	out := new(DocdbClusterParameterGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DocdbClusterParameterGroupSpec) DeepCopyInto(out *DocdbClusterParameterGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DocdbClusterParameterGroupSpec.
func (in *DocdbClusterParameterGroupSpec) DeepCopy() *DocdbClusterParameterGroupSpec {
	if in == nil {
		return nil
	}
	out := new(DocdbClusterParameterGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DocdbClusterParameterGroupStatus) DeepCopyInto(out *DocdbClusterParameterGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DocdbClusterParameterGroupStatus.
func (in *DocdbClusterParameterGroupStatus) DeepCopy() *DocdbClusterParameterGroupStatus {
	if in == nil {
		return nil
	}
	out := new(DocdbClusterParameterGroupStatus)
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