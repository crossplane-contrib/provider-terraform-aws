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

// MacieS3BucketAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type MacieS3BucketAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MacieS3BucketAssociationSpec   `json:"spec"`
	Status MacieS3BucketAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MacieS3BucketAssociation contains a list of MacieS3BucketAssociationList
type MacieS3BucketAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MacieS3BucketAssociation `json:"items"`
}

// A MacieS3BucketAssociationSpec defines the desired state of a MacieS3BucketAssociation
type MacieS3BucketAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  MacieS3BucketAssociationParameters `json:",inline"`
}

// A MacieS3BucketAssociationParameters defines the desired state of a MacieS3BucketAssociation
type MacieS3BucketAssociationParameters struct {
	MemberAccountId    string             `json:"member_account_id"`
	Prefix             string             `json:"prefix"`
	BucketName         string             `json:"bucket_name"`
	Id                 string             `json:"id"`
	ClassificationType ClassificationType `json:"classification_type"`
}

type ClassificationType struct {
	Continuous string `json:"continuous"`
	OneTime    string `json:"one_time"`
}

// A MacieS3BucketAssociationStatus defines the observed state of a MacieS3BucketAssociation
type MacieS3BucketAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     MacieS3BucketAssociationObservation `json:",inline"`
}

// A MacieS3BucketAssociationObservation records the observed state of a MacieS3BucketAssociation
type MacieS3BucketAssociationObservation struct{}