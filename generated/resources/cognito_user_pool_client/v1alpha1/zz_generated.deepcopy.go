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
func (in *AnalyticsConfiguration) DeepCopyInto(out *AnalyticsConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsConfiguration.
func (in *AnalyticsConfiguration) DeepCopy() *AnalyticsConfiguration {
	if in == nil {
		return nil
	}
	out := new(AnalyticsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolClient) DeepCopyInto(out *CognitoUserPoolClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolClient.
func (in *CognitoUserPoolClient) DeepCopy() *CognitoUserPoolClient {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CognitoUserPoolClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolClientList) DeepCopyInto(out *CognitoUserPoolClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CognitoUserPoolClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolClientList.
func (in *CognitoUserPoolClientList) DeepCopy() *CognitoUserPoolClientList {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CognitoUserPoolClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolClientObservation) DeepCopyInto(out *CognitoUserPoolClientObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolClientObservation.
func (in *CognitoUserPoolClientObservation) DeepCopy() *CognitoUserPoolClientObservation {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolClientObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolClientParameters) DeepCopyInto(out *CognitoUserPoolClientParameters) {
	*out = *in
	if in.WriteAttributes != nil {
		in, out := &in.WriteAttributes, &out.WriteAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedOauthScopes != nil {
		in, out := &in.AllowedOauthScopes, &out.AllowedOauthScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReadAttributes != nil {
		in, out := &in.ReadAttributes, &out.ReadAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExplicitAuthFlows != nil {
		in, out := &in.ExplicitAuthFlows, &out.ExplicitAuthFlows
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedOauthFlows != nil {
		in, out := &in.AllowedOauthFlows, &out.AllowedOauthFlows
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CallbackUrls != nil {
		in, out := &in.CallbackUrls, &out.CallbackUrls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SupportedIdentityProviders != nil {
		in, out := &in.SupportedIdentityProviders, &out.SupportedIdentityProviders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LogoutUrls != nil {
		in, out := &in.LogoutUrls, &out.LogoutUrls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.AnalyticsConfiguration = in.AnalyticsConfiguration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolClientParameters.
func (in *CognitoUserPoolClientParameters) DeepCopy() *CognitoUserPoolClientParameters {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolClientParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolClientSpec) DeepCopyInto(out *CognitoUserPoolClientSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolClientSpec.
func (in *CognitoUserPoolClientSpec) DeepCopy() *CognitoUserPoolClientSpec {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolClientStatus) DeepCopyInto(out *CognitoUserPoolClientStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolClientStatus.
func (in *CognitoUserPoolClientStatus) DeepCopy() *CognitoUserPoolClientStatus {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolClientStatus)
	in.DeepCopyInto(out)
	return out
}