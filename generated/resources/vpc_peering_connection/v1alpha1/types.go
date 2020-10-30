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

// VpcPeeringConnection is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpcPeeringConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpcPeeringConnectionSpec   `json:"spec"`
	Status VpcPeeringConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcPeeringConnection contains a list of VpcPeeringConnectionList
type VpcPeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcPeeringConnection `json:"items"`
}

// A VpcPeeringConnectionSpec defines the desired state of a VpcPeeringConnection
type VpcPeeringConnectionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpcPeeringConnectionParameters `json:",inline"`
}

// A VpcPeeringConnectionParameters defines the desired state of a VpcPeeringConnection
type VpcPeeringConnectionParameters struct {
	VpcId      string `json:"vpc_id"`
	AutoAccept bool   `json:"auto_accept"`
	PeerVpcId  string `json:"peer_vpc_id"`
}

// A VpcPeeringConnectionStatus defines the observed state of a VpcPeeringConnection
type VpcPeeringConnectionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpcPeeringConnectionObservation `json:",inline"`
}

// A VpcPeeringConnectionObservation records the observed state of a VpcPeeringConnection
type VpcPeeringConnectionObservation struct {
	AcceptStatus string `json:"accept_status"`
	Id           string `json:"id"`
	PeerOwnerId  string `json:"peer_owner_id"`
	PeerRegion   string `json:"peer_region"`
}