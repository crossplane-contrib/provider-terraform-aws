/*
	Copyright 2019 The Crossplane Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true

// CodeartifactDomainPermissionsPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodeartifactDomainPermissionsPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodeartifactDomainPermissionsPolicySpec   `json:"spec"`
	Status CodeartifactDomainPermissionsPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodeartifactDomainPermissionsPolicy contains a list of CodeartifactDomainPermissionsPolicyList
type CodeartifactDomainPermissionsPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodeartifactDomainPermissionsPolicy `json:"items"`
}

// A CodeartifactDomainPermissionsPolicySpec defines the desired state of a CodeartifactDomainPermissionsPolicy
type CodeartifactDomainPermissionsPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodeartifactDomainPermissionsPolicyParameters `json:",inline"`
}

// A CodeartifactDomainPermissionsPolicyParameters defines the desired state of a CodeartifactDomainPermissionsPolicy
type CodeartifactDomainPermissionsPolicyParameters struct {
	Domain         string `json:"domain"`
	DomainOwner    string `json:"domain_owner"`
	Id             string `json:"id"`
	PolicyDocument string `json:"policy_document"`
	PolicyRevision string `json:"policy_revision"`
}

// A CodeartifactDomainPermissionsPolicyStatus defines the observed state of a CodeartifactDomainPermissionsPolicy
type CodeartifactDomainPermissionsPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodeartifactDomainPermissionsPolicyObservation `json:",inline"`
}

// A CodeartifactDomainPermissionsPolicyObservation records the observed state of a CodeartifactDomainPermissionsPolicy
type CodeartifactDomainPermissionsPolicyObservation struct {
	ResourceArn string `json:"resource_arn"`
}