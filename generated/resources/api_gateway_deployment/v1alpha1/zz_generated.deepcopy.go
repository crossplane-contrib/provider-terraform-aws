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
func (in *ApiGatewayDeployment) DeepCopyInto(out *ApiGatewayDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDeployment.
func (in *ApiGatewayDeployment) DeepCopy() *ApiGatewayDeployment {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiGatewayDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDeploymentList) DeepCopyInto(out *ApiGatewayDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApiGatewayDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDeploymentList.
func (in *ApiGatewayDeploymentList) DeepCopy() *ApiGatewayDeploymentList {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiGatewayDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDeploymentObservation) DeepCopyInto(out *ApiGatewayDeploymentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDeploymentObservation.
func (in *ApiGatewayDeploymentObservation) DeepCopy() *ApiGatewayDeploymentObservation {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDeploymentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDeploymentParameters) DeepCopyInto(out *ApiGatewayDeploymentParameters) {
	*out = *in
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDeploymentParameters.
func (in *ApiGatewayDeploymentParameters) DeepCopy() *ApiGatewayDeploymentParameters {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDeploymentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDeploymentSpec) DeepCopyInto(out *ApiGatewayDeploymentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDeploymentSpec.
func (in *ApiGatewayDeploymentSpec) DeepCopy() *ApiGatewayDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayDeploymentStatus) DeepCopyInto(out *ApiGatewayDeploymentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayDeploymentStatus.
func (in *ApiGatewayDeploymentStatus) DeepCopy() *ApiGatewayDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}
