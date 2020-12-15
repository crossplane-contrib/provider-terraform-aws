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

// ApiGatewayBasePathMapping is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayBasePathMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayBasePathMappingSpec   `json:"spec"`
	Status ApiGatewayBasePathMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayBasePathMapping contains a list of ApiGatewayBasePathMappingList
type ApiGatewayBasePathMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayBasePathMapping `json:"items"`
}

// A ApiGatewayBasePathMappingSpec defines the desired state of a ApiGatewayBasePathMapping
type ApiGatewayBasePathMappingSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayBasePathMappingParameters `json:"forProvider"`
}

// A ApiGatewayBasePathMappingParameters defines the desired state of a ApiGatewayBasePathMapping
type ApiGatewayBasePathMappingParameters struct {
	StageName  string `json:"stage_name"`
	ApiId      string `json:"api_id"`
	BasePath   string `json:"base_path"`
	DomainName string `json:"domain_name"`
}

// A ApiGatewayBasePathMappingStatus defines the observed state of a ApiGatewayBasePathMapping
type ApiGatewayBasePathMappingStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayBasePathMappingObservation `json:"atProvider"`
}

// A ApiGatewayBasePathMappingObservation records the observed state of a ApiGatewayBasePathMapping
type ApiGatewayBasePathMappingObservation struct{}