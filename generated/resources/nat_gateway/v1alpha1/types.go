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

// NatGateway is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type NatGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NatGatewaySpec   `json:"spec"`
	Status NatGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NatGateway contains a list of NatGatewayList
type NatGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NatGateway `json:"items"`
}

// A NatGatewaySpec defines the desired state of a NatGateway
type NatGatewaySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  NatGatewayParameters `json:"forProvider"`
}

// A NatGatewayParameters defines the desired state of a NatGateway
type NatGatewayParameters struct {
	SubnetId     string            `json:"subnet_id"`
	Tags         map[string]string `json:"tags"`
	AllocationId string            `json:"allocation_id"`
}

// A NatGatewayStatus defines the observed state of a NatGateway
type NatGatewayStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     NatGatewayObservation `json:"atProvider"`
}

// A NatGatewayObservation records the observed state of a NatGateway
type NatGatewayObservation struct {
	NetworkInterfaceId string `json:"network_interface_id"`
	PrivateIp          string `json:"private_ip"`
	PublicIp           string `json:"public_ip"`
}