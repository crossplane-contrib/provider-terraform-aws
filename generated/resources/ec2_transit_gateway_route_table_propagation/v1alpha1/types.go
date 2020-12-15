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

// Ec2TransitGatewayRouteTablePropagation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2TransitGatewayRouteTablePropagation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2TransitGatewayRouteTablePropagationSpec   `json:"spec"`
	Status Ec2TransitGatewayRouteTablePropagationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayRouteTablePropagation contains a list of Ec2TransitGatewayRouteTablePropagationList
type Ec2TransitGatewayRouteTablePropagationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayRouteTablePropagation `json:"items"`
}

// A Ec2TransitGatewayRouteTablePropagationSpec defines the desired state of a Ec2TransitGatewayRouteTablePropagation
type Ec2TransitGatewayRouteTablePropagationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2TransitGatewayRouteTablePropagationParameters `json:"forProvider"`
}

// A Ec2TransitGatewayRouteTablePropagationParameters defines the desired state of a Ec2TransitGatewayRouteTablePropagation
type Ec2TransitGatewayRouteTablePropagationParameters struct {
	TransitGatewayAttachmentId string `json:"transit_gateway_attachment_id"`
	TransitGatewayRouteTableId string `json:"transit_gateway_route_table_id"`
}

// A Ec2TransitGatewayRouteTablePropagationStatus defines the observed state of a Ec2TransitGatewayRouteTablePropagation
type Ec2TransitGatewayRouteTablePropagationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2TransitGatewayRouteTablePropagationObservation `json:"atProvider"`
}

// A Ec2TransitGatewayRouteTablePropagationObservation records the observed state of a Ec2TransitGatewayRouteTablePropagation
type Ec2TransitGatewayRouteTablePropagationObservation struct {
	ResourceId   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}