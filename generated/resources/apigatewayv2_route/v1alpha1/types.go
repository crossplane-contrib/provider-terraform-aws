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

// Apigatewayv2Route is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Apigatewayv2Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Apigatewayv2RouteSpec   `json:"spec"`
	Status Apigatewayv2RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2Route contains a list of Apigatewayv2RouteList
type Apigatewayv2RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2Route `json:"items"`
}

// A Apigatewayv2RouteSpec defines the desired state of a Apigatewayv2Route
type Apigatewayv2RouteSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Apigatewayv2RouteParameters `json:",inline"`
}

// A Apigatewayv2RouteParameters defines the desired state of a Apigatewayv2Route
type Apigatewayv2RouteParameters struct {
	AuthorizationScopes              []string          `json:"authorization_scopes"`
	AuthorizationType                string            `json:"authorization_type"`
	Id                               string            `json:"id"`
	OperationName                    string            `json:"operation_name"`
	RouteResponseSelectionExpression string            `json:"route_response_selection_expression"`
	Target                           string            `json:"target"`
	ApiId                            string            `json:"api_id"`
	ApiKeyRequired                   bool              `json:"api_key_required"`
	AuthorizerId                     string            `json:"authorizer_id"`
	ModelSelectionExpression         string            `json:"model_selection_expression"`
	RequestModels                    map[string]string `json:"request_models"`
	RouteKey                         string            `json:"route_key"`
}

// A Apigatewayv2RouteStatus defines the observed state of a Apigatewayv2Route
type Apigatewayv2RouteStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Apigatewayv2RouteObservation `json:",inline"`
}

// A Apigatewayv2RouteObservation records the observed state of a Apigatewayv2Route
type Apigatewayv2RouteObservation struct{}