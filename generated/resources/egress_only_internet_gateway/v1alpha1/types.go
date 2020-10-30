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

// EgressOnlyInternetGateway is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EgressOnlyInternetGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EgressOnlyInternetGatewaySpec   `json:"spec"`
	Status EgressOnlyInternetGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EgressOnlyInternetGateway contains a list of EgressOnlyInternetGatewayList
type EgressOnlyInternetGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EgressOnlyInternetGateway `json:"items"`
}

// A EgressOnlyInternetGatewaySpec defines the desired state of a EgressOnlyInternetGateway
type EgressOnlyInternetGatewaySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EgressOnlyInternetGatewayParameters `json:",inline"`
}

// A EgressOnlyInternetGatewayParameters defines the desired state of a EgressOnlyInternetGateway
type EgressOnlyInternetGatewayParameters struct {
	VpcId string `json:"vpc_id"`
}

// A EgressOnlyInternetGatewayStatus defines the observed state of a EgressOnlyInternetGateway
type EgressOnlyInternetGatewayStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EgressOnlyInternetGatewayObservation `json:",inline"`
}

// A EgressOnlyInternetGatewayObservation records the observed state of a EgressOnlyInternetGateway
type EgressOnlyInternetGatewayObservation struct {
	Id string `json:"id"`
}