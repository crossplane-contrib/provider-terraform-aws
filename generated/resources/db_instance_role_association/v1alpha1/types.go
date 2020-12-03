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

// DbInstanceRoleAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DbInstanceRoleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbInstanceRoleAssociationSpec   `json:"spec"`
	Status DbInstanceRoleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbInstanceRoleAssociation contains a list of DbInstanceRoleAssociationList
type DbInstanceRoleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbInstanceRoleAssociation `json:"items"`
}

// A DbInstanceRoleAssociationSpec defines the desired state of a DbInstanceRoleAssociation
type DbInstanceRoleAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DbInstanceRoleAssociationParameters `json:",inline"`
}

// A DbInstanceRoleAssociationParameters defines the desired state of a DbInstanceRoleAssociation
type DbInstanceRoleAssociationParameters struct {
	DbInstanceIdentifier string `json:"db_instance_identifier"`
	FeatureName          string `json:"feature_name"`
	Id                   string `json:"id"`
	RoleArn              string `json:"role_arn"`
}

// A DbInstanceRoleAssociationStatus defines the observed state of a DbInstanceRoleAssociation
type DbInstanceRoleAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DbInstanceRoleAssociationObservation `json:",inline"`
}

// A DbInstanceRoleAssociationObservation records the observed state of a DbInstanceRoleAssociation
type DbInstanceRoleAssociationObservation struct{}