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

// VpnGatewayRoutePropagation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpnGatewayRoutePropagation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpnGatewayRoutePropagationSpec   `json:"spec"`
	Status VpnGatewayRoutePropagationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpnGatewayRoutePropagation contains a list of VpnGatewayRoutePropagationList
type VpnGatewayRoutePropagationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpnGatewayRoutePropagation `json:"items"`
}

// A VpnGatewayRoutePropagationSpec defines the desired state of a VpnGatewayRoutePropagation
type VpnGatewayRoutePropagationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpnGatewayRoutePropagationParameters `json:"forProvider"`
}

// A VpnGatewayRoutePropagationParameters defines the desired state of a VpnGatewayRoutePropagation
type VpnGatewayRoutePropagationParameters struct {
	RouteTableId string `json:"route_table_id"`
	VpnGatewayId string `json:"vpn_gateway_id"`
}

// A VpnGatewayRoutePropagationStatus defines the observed state of a VpnGatewayRoutePropagation
type VpnGatewayRoutePropagationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpnGatewayRoutePropagationObservation `json:"atProvider"`
}

// A VpnGatewayRoutePropagationObservation records the observed state of a VpnGatewayRoutePropagation
type VpnGatewayRoutePropagationObservation struct{}