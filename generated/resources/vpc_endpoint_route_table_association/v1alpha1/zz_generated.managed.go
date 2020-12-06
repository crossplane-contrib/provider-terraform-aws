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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"

// GetCondition of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) GetDeletionPolicy() runtimev1alpha1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) GetProviderConfigReference() *runtimev1alpha1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VpcEndpointRouteTableAssociation.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VpcEndpointRouteTableAssociation) GetProviderReference() *runtimev1alpha1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) SetDeletionPolicy(r runtimev1alpha1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) SetProviderConfigReference(r *runtimev1alpha1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VpcEndpointRouteTableAssociation.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VpcEndpointRouteTableAssociation) SetProviderReference(r *runtimev1alpha1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
