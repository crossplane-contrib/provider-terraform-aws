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

// Apigatewayv2ApiMapping is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Apigatewayv2ApiMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Apigatewayv2ApiMappingSpec   `json:"spec"`
	Status Apigatewayv2ApiMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2ApiMapping contains a list of Apigatewayv2ApiMappingList
type Apigatewayv2ApiMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2ApiMapping `json:"items"`
}

// A Apigatewayv2ApiMappingSpec defines the desired state of a Apigatewayv2ApiMapping
type Apigatewayv2ApiMappingSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Apigatewayv2ApiMappingParameters `json:"forProvider"`
}

// A Apigatewayv2ApiMappingParameters defines the desired state of a Apigatewayv2ApiMapping
type Apigatewayv2ApiMappingParameters struct {
	ApiId         string `json:"api_id"`
	ApiMappingKey string `json:"api_mapping_key"`
	DomainName    string `json:"domain_name"`
	Id            string `json:"id"`
	Stage         string `json:"stage"`
}

// A Apigatewayv2ApiMappingStatus defines the observed state of a Apigatewayv2ApiMapping
type Apigatewayv2ApiMappingStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Apigatewayv2ApiMappingObservation `json:"atProvider"`
}

// A Apigatewayv2ApiMappingObservation records the observed state of a Apigatewayv2ApiMapping
type Apigatewayv2ApiMappingObservation struct{}