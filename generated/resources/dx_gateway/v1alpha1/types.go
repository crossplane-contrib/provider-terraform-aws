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

// DxGateway is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxGatewaySpec   `json:"spec"`
	Status DxGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxGateway contains a list of DxGatewayList
type DxGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxGateway `json:"items"`
}

// A DxGatewaySpec defines the desired state of a DxGateway
type DxGatewaySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxGatewayParameters `json:"forProvider"`
}

// A DxGatewayParameters defines the desired state of a DxGateway
type DxGatewayParameters struct {
	AmazonSideAsn string   `json:"amazon_side_asn"`
	Name          string   `json:"name"`
	Timeouts      Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A DxGatewayStatus defines the observed state of a DxGateway
type DxGatewayStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxGatewayObservation `json:"atProvider"`
}

// A DxGatewayObservation records the observed state of a DxGateway
type DxGatewayObservation struct {
	OwnerAccountId string `json:"owner_account_id"`
}