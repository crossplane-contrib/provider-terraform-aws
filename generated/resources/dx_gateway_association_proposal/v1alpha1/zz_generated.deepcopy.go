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
func (in *DxGatewayAssociationProposal) DeepCopyInto(out *DxGatewayAssociationProposal) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxGatewayAssociationProposal.
func (in *DxGatewayAssociationProposal) DeepCopy() *DxGatewayAssociationProposal {
	if in == nil {
		return nil
	}
	out := new(DxGatewayAssociationProposal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DxGatewayAssociationProposal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxGatewayAssociationProposalList) DeepCopyInto(out *DxGatewayAssociationProposalList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DxGatewayAssociationProposal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxGatewayAssociationProposalList.
func (in *DxGatewayAssociationProposalList) DeepCopy() *DxGatewayAssociationProposalList {
	if in == nil {
		return nil
	}
	out := new(DxGatewayAssociationProposalList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DxGatewayAssociationProposalList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxGatewayAssociationProposalObservation) DeepCopyInto(out *DxGatewayAssociationProposalObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxGatewayAssociationProposalObservation.
func (in *DxGatewayAssociationProposalObservation) DeepCopy() *DxGatewayAssociationProposalObservation {
	if in == nil {
		return nil
	}
	out := new(DxGatewayAssociationProposalObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxGatewayAssociationProposalParameters) DeepCopyInto(out *DxGatewayAssociationProposalParameters) {
	*out = *in
	if in.AllowedPrefixes != nil {
		in, out := &in.AllowedPrefixes, &out.AllowedPrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxGatewayAssociationProposalParameters.
func (in *DxGatewayAssociationProposalParameters) DeepCopy() *DxGatewayAssociationProposalParameters {
	if in == nil {
		return nil
	}
	out := new(DxGatewayAssociationProposalParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxGatewayAssociationProposalSpec) DeepCopyInto(out *DxGatewayAssociationProposalSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxGatewayAssociationProposalSpec.
func (in *DxGatewayAssociationProposalSpec) DeepCopy() *DxGatewayAssociationProposalSpec {
	if in == nil {
		return nil
	}
	out := new(DxGatewayAssociationProposalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DxGatewayAssociationProposalStatus) DeepCopyInto(out *DxGatewayAssociationProposalStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DxGatewayAssociationProposalStatus.
func (in *DxGatewayAssociationProposalStatus) DeepCopy() *DxGatewayAssociationProposalStatus {
	if in == nil {
		return nil
	}
	out := new(DxGatewayAssociationProposalStatus)
	in.DeepCopyInto(out)
	return out
}
