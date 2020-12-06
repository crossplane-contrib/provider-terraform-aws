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
func (in *PublicAccessBlockConfiguration) DeepCopyInto(out *PublicAccessBlockConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicAccessBlockConfiguration.
func (in *PublicAccessBlockConfiguration) DeepCopy() *PublicAccessBlockConfiguration {
	if in == nil {
		return nil
	}
	out := new(PublicAccessBlockConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3AccessPoint) DeepCopyInto(out *S3AccessPoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3AccessPoint.
func (in *S3AccessPoint) DeepCopy() *S3AccessPoint {
	if in == nil {
		return nil
	}
	out := new(S3AccessPoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3AccessPoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3AccessPointList) DeepCopyInto(out *S3AccessPointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3AccessPoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3AccessPointList.
func (in *S3AccessPointList) DeepCopy() *S3AccessPointList {
	if in == nil {
		return nil
	}
	out := new(S3AccessPointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3AccessPointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3AccessPointObservation) DeepCopyInto(out *S3AccessPointObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3AccessPointObservation.
func (in *S3AccessPointObservation) DeepCopy() *S3AccessPointObservation {
	if in == nil {
		return nil
	}
	out := new(S3AccessPointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3AccessPointParameters) DeepCopyInto(out *S3AccessPointParameters) {
	*out = *in
	out.PublicAccessBlockConfiguration = in.PublicAccessBlockConfiguration
	out.VpcConfiguration = in.VpcConfiguration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3AccessPointParameters.
func (in *S3AccessPointParameters) DeepCopy() *S3AccessPointParameters {
	if in == nil {
		return nil
	}
	out := new(S3AccessPointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3AccessPointSpec) DeepCopyInto(out *S3AccessPointSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3AccessPointSpec.
func (in *S3AccessPointSpec) DeepCopy() *S3AccessPointSpec {
	if in == nil {
		return nil
	}
	out := new(S3AccessPointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3AccessPointStatus) DeepCopyInto(out *S3AccessPointStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3AccessPointStatus.
func (in *S3AccessPointStatus) DeepCopy() *S3AccessPointStatus {
	if in == nil {
		return nil
	}
	out := new(S3AccessPointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcConfiguration) DeepCopyInto(out *VpcConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcConfiguration.
func (in *VpcConfiguration) DeepCopy() *VpcConfiguration {
	if in == nil {
		return nil
	}
	out := new(VpcConfiguration)
	in.DeepCopyInto(out)
	return out
}
