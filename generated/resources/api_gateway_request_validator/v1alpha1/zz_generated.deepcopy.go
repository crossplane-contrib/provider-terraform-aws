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
func (in *ApiGatewayRequestValidator) DeepCopyInto(out *ApiGatewayRequestValidator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayRequestValidator.
func (in *ApiGatewayRequestValidator) DeepCopy() *ApiGatewayRequestValidator {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayRequestValidator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiGatewayRequestValidator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayRequestValidatorList) DeepCopyInto(out *ApiGatewayRequestValidatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApiGatewayRequestValidator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayRequestValidatorList.
func (in *ApiGatewayRequestValidatorList) DeepCopy() *ApiGatewayRequestValidatorList {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayRequestValidatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiGatewayRequestValidatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayRequestValidatorObservation) DeepCopyInto(out *ApiGatewayRequestValidatorObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayRequestValidatorObservation.
func (in *ApiGatewayRequestValidatorObservation) DeepCopy() *ApiGatewayRequestValidatorObservation {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayRequestValidatorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayRequestValidatorParameters) DeepCopyInto(out *ApiGatewayRequestValidatorParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayRequestValidatorParameters.
func (in *ApiGatewayRequestValidatorParameters) DeepCopy() *ApiGatewayRequestValidatorParameters {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayRequestValidatorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayRequestValidatorSpec) DeepCopyInto(out *ApiGatewayRequestValidatorSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayRequestValidatorSpec.
func (in *ApiGatewayRequestValidatorSpec) DeepCopy() *ApiGatewayRequestValidatorSpec {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayRequestValidatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayRequestValidatorStatus) DeepCopyInto(out *ApiGatewayRequestValidatorStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayRequestValidatorStatus.
func (in *ApiGatewayRequestValidatorStatus) DeepCopy() *ApiGatewayRequestValidatorStatus {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayRequestValidatorStatus)
	in.DeepCopyInto(out)
	return out
}
