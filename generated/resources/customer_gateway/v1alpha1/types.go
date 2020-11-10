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

// CustomerGateway is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CustomerGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomerGatewaySpec   `json:"spec"`
	Status CustomerGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomerGateway contains a list of CustomerGatewayList
type CustomerGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomerGateway `json:"items"`
}

// A CustomerGatewaySpec defines the desired state of a CustomerGateway
type CustomerGatewaySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CustomerGatewayParameters `json:",inline"`
}

// A CustomerGatewayParameters defines the desired state of a CustomerGateway
type CustomerGatewayParameters struct {
	IpAddress string            `json:"ip_address"`
	Tags      map[string]string `json:"tags"`
	Type      string            `json:"type"`
	BgpAsn    string            `json:"bgp_asn"`
}

// A CustomerGatewayStatus defines the observed state of a CustomerGateway
type CustomerGatewayStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CustomerGatewayObservation `json:",inline"`
}

// A CustomerGatewayObservation records the observed state of a CustomerGateway
type CustomerGatewayObservation struct {
	Id  string `json:"id"`
	Arn string `json:"arn"`
}