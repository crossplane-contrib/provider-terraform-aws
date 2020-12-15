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

// Apigatewayv2RouteResponse is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Apigatewayv2RouteResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Apigatewayv2RouteResponseSpec   `json:"spec"`
	Status Apigatewayv2RouteResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2RouteResponse contains a list of Apigatewayv2RouteResponseList
type Apigatewayv2RouteResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2RouteResponse `json:"items"`
}

// A Apigatewayv2RouteResponseSpec defines the desired state of a Apigatewayv2RouteResponse
type Apigatewayv2RouteResponseSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Apigatewayv2RouteResponseParameters `json:"forProvider"`
}

// A Apigatewayv2RouteResponseParameters defines the desired state of a Apigatewayv2RouteResponse
type Apigatewayv2RouteResponseParameters struct {
	ModelSelectionExpression string            `json:"model_selection_expression"`
	ResponseModels           map[string]string `json:"response_models"`
	RouteId                  string            `json:"route_id"`
	RouteResponseKey         string            `json:"route_response_key"`
	ApiId                    string            `json:"api_id"`
}

// A Apigatewayv2RouteResponseStatus defines the observed state of a Apigatewayv2RouteResponse
type Apigatewayv2RouteResponseStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Apigatewayv2RouteResponseObservation `json:"atProvider"`
}

// A Apigatewayv2RouteResponseObservation records the observed state of a Apigatewayv2RouteResponse
type Apigatewayv2RouteResponseObservation struct{}