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

// Ec2ClientVpnRoute is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2ClientVpnRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2ClientVpnRouteSpec   `json:"spec"`
	Status Ec2ClientVpnRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2ClientVpnRoute contains a list of Ec2ClientVpnRouteList
type Ec2ClientVpnRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2ClientVpnRoute `json:"items"`
}

// A Ec2ClientVpnRouteSpec defines the desired state of a Ec2ClientVpnRoute
type Ec2ClientVpnRouteSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2ClientVpnRouteParameters `json:",inline"`
}

// A Ec2ClientVpnRouteParameters defines the desired state of a Ec2ClientVpnRoute
type Ec2ClientVpnRouteParameters struct {
	ClientVpnEndpointId  string `json:"client_vpn_endpoint_id"`
	Description          string `json:"description"`
	DestinationCidrBlock string `json:"destination_cidr_block"`
	TargetVpcSubnetId    string `json:"target_vpc_subnet_id"`
}

// A Ec2ClientVpnRouteStatus defines the observed state of a Ec2ClientVpnRoute
type Ec2ClientVpnRouteStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2ClientVpnRouteObservation `json:",inline"`
}

// A Ec2ClientVpnRouteObservation records the observed state of a Ec2ClientVpnRoute
type Ec2ClientVpnRouteObservation struct {
	Id     string `json:"id"`
	Origin string `json:"origin"`
	Type   string `json:"type"`
}