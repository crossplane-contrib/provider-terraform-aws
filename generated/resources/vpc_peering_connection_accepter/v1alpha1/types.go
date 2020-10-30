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

// VpcPeeringConnectionAccepter is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpcPeeringConnectionAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpcPeeringConnectionAccepterSpec   `json:"spec"`
	Status VpcPeeringConnectionAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcPeeringConnectionAccepter contains a list of VpcPeeringConnectionAccepterList
type VpcPeeringConnectionAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcPeeringConnectionAccepter `json:"items"`
}

// A VpcPeeringConnectionAccepterSpec defines the desired state of a VpcPeeringConnectionAccepter
type VpcPeeringConnectionAccepterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpcPeeringConnectionAccepterParameters `json:",inline"`
}

// A VpcPeeringConnectionAccepterParameters defines the desired state of a VpcPeeringConnectionAccepter
type VpcPeeringConnectionAccepterParameters struct {
	VpcPeeringConnectionId string `json:"vpc_peering_connection_id"`
	AutoAccept             bool   `json:"auto_accept"`
}

// A VpcPeeringConnectionAccepterStatus defines the observed state of a VpcPeeringConnectionAccepter
type VpcPeeringConnectionAccepterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpcPeeringConnectionAccepterObservation `json:",inline"`
}

// A VpcPeeringConnectionAccepterObservation records the observed state of a VpcPeeringConnectionAccepter
type VpcPeeringConnectionAccepterObservation struct {
	PeerVpcId    string `json:"peer_vpc_id"`
	VpcId        string `json:"vpc_id"`
	PeerRegion   string `json:"peer_region"`
	AcceptStatus string `json:"accept_status"`
	Id           string `json:"id"`
	PeerOwnerId  string `json:"peer_owner_id"`
}