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

// Ec2TransitGatewayPeeringAttachment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2TransitGatewayPeeringAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2TransitGatewayPeeringAttachmentSpec   `json:"spec"`
	Status Ec2TransitGatewayPeeringAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayPeeringAttachment contains a list of Ec2TransitGatewayPeeringAttachmentList
type Ec2TransitGatewayPeeringAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayPeeringAttachment `json:"items"`
}

// A Ec2TransitGatewayPeeringAttachmentSpec defines the desired state of a Ec2TransitGatewayPeeringAttachment
type Ec2TransitGatewayPeeringAttachmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2TransitGatewayPeeringAttachmentParameters `json:",inline"`
}

// A Ec2TransitGatewayPeeringAttachmentParameters defines the desired state of a Ec2TransitGatewayPeeringAttachment
type Ec2TransitGatewayPeeringAttachmentParameters struct {
	Tags                 map[string]string `json:"tags"`
	TransitGatewayId     string            `json:"transit_gateway_id"`
	Id                   string            `json:"id"`
	PeerAccountId        string            `json:"peer_account_id"`
	PeerRegion           string            `json:"peer_region"`
	PeerTransitGatewayId string            `json:"peer_transit_gateway_id"`
}

// A Ec2TransitGatewayPeeringAttachmentStatus defines the observed state of a Ec2TransitGatewayPeeringAttachment
type Ec2TransitGatewayPeeringAttachmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2TransitGatewayPeeringAttachmentObservation `json:",inline"`
}

// A Ec2TransitGatewayPeeringAttachmentObservation records the observed state of a Ec2TransitGatewayPeeringAttachment
type Ec2TransitGatewayPeeringAttachmentObservation struct{}