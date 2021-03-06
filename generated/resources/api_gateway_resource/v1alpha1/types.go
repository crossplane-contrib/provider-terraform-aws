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

// ApiGatewayResource is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayResourceSpec   `json:"spec"`
	Status ApiGatewayResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayResource contains a list of ApiGatewayResourceList
type ApiGatewayResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayResource `json:"items"`
}

// A ApiGatewayResourceSpec defines the desired state of a ApiGatewayResource
type ApiGatewayResourceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayResourceParameters `json:"forProvider"`
}

// A ApiGatewayResourceParameters defines the desired state of a ApiGatewayResource
type ApiGatewayResourceParameters struct {
	RestApiId string `json:"rest_api_id"`
	ParentId  string `json:"parent_id"`
	PathPart  string `json:"path_part"`
}

// A ApiGatewayResourceStatus defines the observed state of a ApiGatewayResource
type ApiGatewayResourceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayResourceObservation `json:"atProvider"`
}

// A ApiGatewayResourceObservation records the observed state of a ApiGatewayResource
type ApiGatewayResourceObservation struct {
	Path string `json:"path"`
}