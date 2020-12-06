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
func (in *ApiGatewayDocumentationVersion) DeepCopyInto(out *ApiGatewayDocumentationVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDocumentationVersion.
func (in *ApiGatewayDocumentationVersion) DeepCopy() *ApiGatewayDocumentationVersion {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDocumentationVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiGatewayDocumentationVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDocumentationVersionList) DeepCopyInto(out *ApiGatewayDocumentationVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApiGatewayDocumentationVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDocumentationVersionList.
func (in *ApiGatewayDocumentationVersionList) DeepCopy() *ApiGatewayDocumentationVersionList {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDocumentationVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiGatewayDocumentationVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDocumentationVersionObservation) DeepCopyInto(out *ApiGatewayDocumentationVersionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDocumentationVersionObservation.
func (in *ApiGatewayDocumentationVersionObservation) DeepCopy() *ApiGatewayDocumentationVersionObservation {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDocumentationVersionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDocumentationVersionParameters) DeepCopyInto(out *ApiGatewayDocumentationVersionParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDocumentationVersionParameters.
func (in *ApiGatewayDocumentationVersionParameters) DeepCopy() *ApiGatewayDocumentationVersionParameters {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDocumentationVersionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDocumentationVersionSpec) DeepCopyInto(out *ApiGatewayDocumentationVersionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDocumentationVersionSpec.
func (in *ApiGatewayDocumentationVersionSpec) DeepCopy() *ApiGatewayDocumentationVersionSpec {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDocumentationVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDocumentationVersionStatus) DeepCopyInto(out *ApiGatewayDocumentationVersionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDocumentationVersionStatus.
func (in *ApiGatewayDocumentationVersionStatus) DeepCopy() *ApiGatewayDocumentationVersionStatus {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDocumentationVersionStatus)
	in.DeepCopyInto(out)
	return out
}
