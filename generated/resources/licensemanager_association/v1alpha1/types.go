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

// LicensemanagerAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LicensemanagerAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LicensemanagerAssociationSpec   `json:"spec"`
	Status LicensemanagerAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LicensemanagerAssociation contains a list of LicensemanagerAssociationList
type LicensemanagerAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LicensemanagerAssociation `json:"items"`
}

// A LicensemanagerAssociationSpec defines the desired state of a LicensemanagerAssociation
type LicensemanagerAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LicensemanagerAssociationParameters `json:"forProvider"`
}

// A LicensemanagerAssociationParameters defines the desired state of a LicensemanagerAssociation
type LicensemanagerAssociationParameters struct {
	LicenseConfigurationArn string `json:"license_configuration_arn"`
	ResourceArn             string `json:"resource_arn"`
	Id                      string `json:"id"`
}

// A LicensemanagerAssociationStatus defines the observed state of a LicensemanagerAssociation
type LicensemanagerAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LicensemanagerAssociationObservation `json:"atProvider"`
}

// A LicensemanagerAssociationObservation records the observed state of a LicensemanagerAssociation
type LicensemanagerAssociationObservation struct{}