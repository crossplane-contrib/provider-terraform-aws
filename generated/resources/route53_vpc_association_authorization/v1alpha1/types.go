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

// Route53VpcAssociationAuthorization is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Route53VpcAssociationAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Route53VpcAssociationAuthorizationSpec   `json:"spec"`
	Status Route53VpcAssociationAuthorizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53VpcAssociationAuthorization contains a list of Route53VpcAssociationAuthorizationList
type Route53VpcAssociationAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53VpcAssociationAuthorization `json:"items"`
}

// A Route53VpcAssociationAuthorizationSpec defines the desired state of a Route53VpcAssociationAuthorization
type Route53VpcAssociationAuthorizationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Route53VpcAssociationAuthorizationParameters `json:"forProvider"`
}

// A Route53VpcAssociationAuthorizationParameters defines the desired state of a Route53VpcAssociationAuthorization
type Route53VpcAssociationAuthorizationParameters struct {
	ZoneId    string `json:"zone_id"`
	VpcId     string `json:"vpc_id"`
	VpcRegion string `json:"vpc_region"`
}

// A Route53VpcAssociationAuthorizationStatus defines the observed state of a Route53VpcAssociationAuthorization
type Route53VpcAssociationAuthorizationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Route53VpcAssociationAuthorizationObservation `json:"atProvider"`
}

// A Route53VpcAssociationAuthorizationObservation records the observed state of a Route53VpcAssociationAuthorization
type Route53VpcAssociationAuthorizationObservation struct{}