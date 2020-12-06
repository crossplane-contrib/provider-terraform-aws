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
func (in *DatasyncLocationEfs) DeepCopyInto(out *DatasyncLocationEfs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationEfs.
func (in *DatasyncLocationEfs) DeepCopy() *DatasyncLocationEfs {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationEfs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasyncLocationEfs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationEfsList) DeepCopyInto(out *DatasyncLocationEfsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatasyncLocationEfs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationEfsList.
func (in *DatasyncLocationEfsList) DeepCopy() *DatasyncLocationEfsList {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationEfsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasyncLocationEfsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationEfsObservation) DeepCopyInto(out *DatasyncLocationEfsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationEfsObservation.
func (in *DatasyncLocationEfsObservation) DeepCopy() *DatasyncLocationEfsObservation {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationEfsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationEfsParameters) DeepCopyInto(out *DatasyncLocationEfsParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Ec2Config.DeepCopyInto(&out.Ec2Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationEfsParameters.
func (in *DatasyncLocationEfsParameters) DeepCopy() *DatasyncLocationEfsParameters {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationEfsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationEfsSpec) DeepCopyInto(out *DatasyncLocationEfsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationEfsSpec.
func (in *DatasyncLocationEfsSpec) DeepCopy() *DatasyncLocationEfsSpec {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationEfsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasyncLocationEfsStatus) DeepCopyInto(out *DatasyncLocationEfsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasyncLocationEfsStatus.
func (in *DatasyncLocationEfsStatus) DeepCopy() *DatasyncLocationEfsStatus {
	if in == nil {
		return nil
	}
	out := new(DatasyncLocationEfsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2Config) DeepCopyInto(out *Ec2Config) {
	*out = *in
	if in.SecurityGroupArns != nil {
		in, out := &in.SecurityGroupArns, &out.SecurityGroupArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2Config.
func (in *Ec2Config) DeepCopy() *Ec2Config {
	if in == nil {
		return nil
	}
	out := new(Ec2Config)
	in.DeepCopyInto(out)
	return out
}