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

// RouteTableAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouteTableAssociationSpec   `json:"spec"`
	Status RouteTableAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableAssociation contains a list of RouteTableAssociationList
type RouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTableAssociation `json:"items"`
}

// A RouteTableAssociationSpec defines the desired state of a RouteTableAssociation
type RouteTableAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RouteTableAssociationParameters `json:"forProvider"`
}

// A RouteTableAssociationParameters defines the desired state of a RouteTableAssociation
type RouteTableAssociationParameters struct {
	GatewayId    string `json:"gateway_id"`
	RouteTableId string `json:"route_table_id"`
	SubnetId     string `json:"subnet_id"`
}

// A RouteTableAssociationStatus defines the observed state of a RouteTableAssociation
type RouteTableAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RouteTableAssociationObservation `json:"atProvider"`
}

// A RouteTableAssociationObservation records the observed state of a RouteTableAssociation
type RouteTableAssociationObservation struct{}