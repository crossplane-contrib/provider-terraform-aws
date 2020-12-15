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

// Ec2TransitGatewayPeeringAttachmentAccepter is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2TransitGatewayPeeringAttachmentAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2TransitGatewayPeeringAttachmentAccepterSpec   `json:"spec"`
	Status Ec2TransitGatewayPeeringAttachmentAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayPeeringAttachmentAccepter contains a list of Ec2TransitGatewayPeeringAttachmentAccepterList
type Ec2TransitGatewayPeeringAttachmentAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayPeeringAttachmentAccepter `json:"items"`
}

// A Ec2TransitGatewayPeeringAttachmentAccepterSpec defines the desired state of a Ec2TransitGatewayPeeringAttachmentAccepter
type Ec2TransitGatewayPeeringAttachmentAccepterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2TransitGatewayPeeringAttachmentAccepterParameters `json:"forProvider"`
}

// A Ec2TransitGatewayPeeringAttachmentAccepterParameters defines the desired state of a Ec2TransitGatewayPeeringAttachmentAccepter
type Ec2TransitGatewayPeeringAttachmentAccepterParameters struct {
	TransitGatewayAttachmentId string            `json:"transit_gateway_attachment_id"`
	Tags                       map[string]string `json:"tags"`
}

// A Ec2TransitGatewayPeeringAttachmentAccepterStatus defines the observed state of a Ec2TransitGatewayPeeringAttachmentAccepter
type Ec2TransitGatewayPeeringAttachmentAccepterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2TransitGatewayPeeringAttachmentAccepterObservation `json:"atProvider"`
}

// A Ec2TransitGatewayPeeringAttachmentAccepterObservation records the observed state of a Ec2TransitGatewayPeeringAttachmentAccepter
type Ec2TransitGatewayPeeringAttachmentAccepterObservation struct {
	TransitGatewayId     string `json:"transit_gateway_id"`
	PeerAccountId        string `json:"peer_account_id"`
	PeerRegion           string `json:"peer_region"`
	PeerTransitGatewayId string `json:"peer_transit_gateway_id"`
}