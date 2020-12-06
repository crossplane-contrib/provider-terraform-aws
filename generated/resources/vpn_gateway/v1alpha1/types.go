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

// VpnGateway is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpnGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpnGatewaySpec   `json:"spec"`
	Status VpnGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpnGateway contains a list of VpnGatewayList
type VpnGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpnGateway `json:"items"`
}

// A VpnGatewaySpec defines the desired state of a VpnGateway
type VpnGatewaySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpnGatewayParameters `json:",inline"`
}

// A VpnGatewayParameters defines the desired state of a VpnGateway
type VpnGatewayParameters struct {
	VpcId            string            `json:"vpc_id"`
	AmazonSideAsn    string            `json:"amazon_side_asn"`
	AvailabilityZone string            `json:"availability_zone"`
	Id               string            `json:"id"`
	Tags             map[string]string `json:"tags"`
}

// A VpnGatewayStatus defines the observed state of a VpnGateway
type VpnGatewayStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpnGatewayObservation `json:",inline"`
}

// A VpnGatewayObservation records the observed state of a VpnGateway
type VpnGatewayObservation struct {
	Arn string `json:"arn"`
}