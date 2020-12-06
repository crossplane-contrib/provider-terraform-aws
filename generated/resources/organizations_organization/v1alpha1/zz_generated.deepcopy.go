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
func (in *Accounts) DeepCopyInto(out *Accounts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Accounts.
func (in *Accounts) DeepCopy() *Accounts {
	if in == nil {
		return nil
	}
	out := new(Accounts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonMasterAccounts) DeepCopyInto(out *NonMasterAccounts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonMasterAccounts.
func (in *NonMasterAccounts) DeepCopy() *NonMasterAccounts {
	if in == nil {
		return nil
	}
	out := new(NonMasterAccounts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationsOrganization) DeepCopyInto(out *OrganizationsOrganization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationsOrganization.
func (in *OrganizationsOrganization) DeepCopy() *OrganizationsOrganization {
	if in == nil {
		return nil
	}
	out := new(OrganizationsOrganization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationsOrganization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationsOrganizationList) DeepCopyInto(out *OrganizationsOrganizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OrganizationsOrganization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationsOrganizationList.
func (in *OrganizationsOrganizationList) DeepCopy() *OrganizationsOrganizationList {
	if in == nil {
		return nil
	}
	out := new(OrganizationsOrganizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationsOrganizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationsOrganizationObservation) DeepCopyInto(out *OrganizationsOrganizationObservation) {
	*out = *in
	if in.NonMasterAccounts != nil {
		in, out := &in.NonMasterAccounts, &out.NonMasterAccounts
		*out = make([]NonMasterAccounts, len(*in))
		copy(*out, *in)
	}
	if in.Accounts != nil {
		in, out := &in.Accounts, &out.Accounts
		*out = make([]Accounts, len(*in))
		copy(*out, *in)
	}
	if in.Roots != nil {
		in, out := &in.Roots, &out.Roots
		*out = make([]Roots, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationsOrganizationObservation.
func (in *OrganizationsOrganizationObservation) DeepCopy() *OrganizationsOrganizationObservation {
	if in == nil {
		return nil
	}
	out := new(OrganizationsOrganizationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationsOrganizationParameters) DeepCopyInto(out *OrganizationsOrganizationParameters) {
	*out = *in
	if in.EnabledPolicyTypes != nil {
		in, out := &in.EnabledPolicyTypes, &out.EnabledPolicyTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AwsServiceAccessPrincipals != nil {
		in, out := &in.AwsServiceAccessPrincipals, &out.AwsServiceAccessPrincipals
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationsOrganizationParameters.
func (in *OrganizationsOrganizationParameters) DeepCopy() *OrganizationsOrganizationParameters {
	if in == nil {
		return nil
	}
	out := new(OrganizationsOrganizationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationsOrganizationSpec) DeepCopyInto(out *OrganizationsOrganizationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationsOrganizationSpec.
func (in *OrganizationsOrganizationSpec) DeepCopy() *OrganizationsOrganizationSpec {
	if in == nil {
		return nil
	}
	out := new(OrganizationsOrganizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationsOrganizationStatus) DeepCopyInto(out *OrganizationsOrganizationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationsOrganizationStatus.
func (in *OrganizationsOrganizationStatus) DeepCopy() *OrganizationsOrganizationStatus {
	if in == nil {
		return nil
	}
	out := new(OrganizationsOrganizationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyTypes) DeepCopyInto(out *PolicyTypes) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyTypes.
func (in *PolicyTypes) DeepCopy() *PolicyTypes {
	if in == nil {
		return nil
	}
	out := new(PolicyTypes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Roots) DeepCopyInto(out *Roots) {
	*out = *in
	if in.PolicyTypes != nil {
		in, out := &in.PolicyTypes, &out.PolicyTypes
		*out = make([]PolicyTypes, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Roots.
func (in *Roots) DeepCopy() *Roots {
	if in == nil {
		return nil
	}
	out := new(Roots)
	in.DeepCopyInto(out)
	return out
}
