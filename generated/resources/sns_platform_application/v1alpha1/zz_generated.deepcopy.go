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
func (in *SnsPlatformApplication) DeepCopyInto(out *SnsPlatformApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsPlatformApplication.
func (in *SnsPlatformApplication) DeepCopy() *SnsPlatformApplication {
	if in == nil {
		return nil
	}
	out := new(SnsPlatformApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsPlatformApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsPlatformApplicationList) DeepCopyInto(out *SnsPlatformApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnsPlatformApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsPlatformApplicationList.
func (in *SnsPlatformApplicationList) DeepCopy() *SnsPlatformApplicationList {
	if in == nil {
		return nil
	}
	out := new(SnsPlatformApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsPlatformApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsPlatformApplicationObservation) DeepCopyInto(out *SnsPlatformApplicationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsPlatformApplicationObservation.
func (in *SnsPlatformApplicationObservation) DeepCopy() *SnsPlatformApplicationObservation {
	if in == nil {
		return nil
	}
	out := new(SnsPlatformApplicationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsPlatformApplicationParameters) DeepCopyInto(out *SnsPlatformApplicationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsPlatformApplicationParameters.
func (in *SnsPlatformApplicationParameters) DeepCopy() *SnsPlatformApplicationParameters {
	if in == nil {
		return nil
	}
	out := new(SnsPlatformApplicationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsPlatformApplicationSpec) DeepCopyInto(out *SnsPlatformApplicationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsPlatformApplicationSpec.
func (in *SnsPlatformApplicationSpec) DeepCopy() *SnsPlatformApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(SnsPlatformApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsPlatformApplicationStatus) DeepCopyInto(out *SnsPlatformApplicationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsPlatformApplicationStatus.
func (in *SnsPlatformApplicationStatus) DeepCopy() *SnsPlatformApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(SnsPlatformApplicationStatus)
	in.DeepCopyInto(out)
	return out
}
