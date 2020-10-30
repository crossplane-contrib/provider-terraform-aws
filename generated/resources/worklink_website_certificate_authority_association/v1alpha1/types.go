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

// WorklinkWebsiteCertificateAuthorityAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WorklinkWebsiteCertificateAuthorityAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorklinkWebsiteCertificateAuthorityAssociationSpec   `json:"spec"`
	Status WorklinkWebsiteCertificateAuthorityAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorklinkWebsiteCertificateAuthorityAssociation contains a list of WorklinkWebsiteCertificateAuthorityAssociationList
type WorklinkWebsiteCertificateAuthorityAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorklinkWebsiteCertificateAuthorityAssociation `json:"items"`
}

// A WorklinkWebsiteCertificateAuthorityAssociationSpec defines the desired state of a WorklinkWebsiteCertificateAuthorityAssociation
type WorklinkWebsiteCertificateAuthorityAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WorklinkWebsiteCertificateAuthorityAssociationParameters `json:",inline"`
}

// A WorklinkWebsiteCertificateAuthorityAssociationParameters defines the desired state of a WorklinkWebsiteCertificateAuthorityAssociation
type WorklinkWebsiteCertificateAuthorityAssociationParameters struct {
	Certificate string `json:"certificate"`
	DisplayName string `json:"display_name"`
	FleetArn    string `json:"fleet_arn"`
}

// A WorklinkWebsiteCertificateAuthorityAssociationStatus defines the observed state of a WorklinkWebsiteCertificateAuthorityAssociation
type WorklinkWebsiteCertificateAuthorityAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WorklinkWebsiteCertificateAuthorityAssociationObservation `json:",inline"`
}

// A WorklinkWebsiteCertificateAuthorityAssociationObservation records the observed state of a WorklinkWebsiteCertificateAuthorityAssociation
type WorklinkWebsiteCertificateAuthorityAssociationObservation struct {
	Id          string `json:"id"`
	WebsiteCaId string `json:"website_ca_id"`
}