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
func (in *Apigatewayv2Route) DeepCopyInto(out *Apigatewayv2Route) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2Route.
func (in *Apigatewayv2Route) DeepCopy() *Apigatewayv2Route {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Apigatewayv2Route) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2RouteList) DeepCopyInto(out *Apigatewayv2RouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Apigatewayv2Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2RouteList.
func (in *Apigatewayv2RouteList) DeepCopy() *Apigatewayv2RouteList {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2RouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Apigatewayv2RouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2RouteObservation) DeepCopyInto(out *Apigatewayv2RouteObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2RouteObservation.
func (in *Apigatewayv2RouteObservation) DeepCopy() *Apigatewayv2RouteObservation {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2RouteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2RouteParameters) DeepCopyInto(out *Apigatewayv2RouteParameters) {
	*out = *in
	if in.RequestModels != nil {
		in, out := &in.RequestModels, &out.RequestModels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AuthorizationScopes != nil {
		in, out := &in.AuthorizationScopes, &out.AuthorizationScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2RouteParameters.
func (in *Apigatewayv2RouteParameters) DeepCopy() *Apigatewayv2RouteParameters {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2RouteParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2RouteSpec) DeepCopyInto(out *Apigatewayv2RouteSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2RouteSpec.
func (in *Apigatewayv2RouteSpec) DeepCopy() *Apigatewayv2RouteSpec {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2RouteStatus) DeepCopyInto(out *Apigatewayv2RouteStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2RouteStatus.
func (in *Apigatewayv2RouteStatus) DeepCopy() *Apigatewayv2RouteStatus {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2RouteStatus)
	in.DeepCopyInto(out)
	return out
}
