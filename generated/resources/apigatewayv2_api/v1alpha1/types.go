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

// Apigatewayv2Api is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Apigatewayv2Api struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Apigatewayv2ApiSpec   `json:"spec"`
	Status Apigatewayv2ApiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2Api contains a list of Apigatewayv2ApiList
type Apigatewayv2ApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2Api `json:"items"`
}

// A Apigatewayv2ApiSpec defines the desired state of a Apigatewayv2Api
type Apigatewayv2ApiSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Apigatewayv2ApiParameters `json:",inline"`
}

// A Apigatewayv2ApiParameters defines the desired state of a Apigatewayv2Api
type Apigatewayv2ApiParameters struct {
	Description               string            `json:"description"`
	RouteSelectionExpression  string            `json:"route_selection_expression"`
	Body                      string            `json:"body"`
	CredentialsArn            string            `json:"credentials_arn"`
	Version                   string            `json:"version"`
	ApiKeySelectionExpression string            `json:"api_key_selection_expression"`
	Name                      string            `json:"name"`
	Target                    string            `json:"target"`
	Tags                      map[string]string `json:"tags"`
	DisableExecuteApiEndpoint bool              `json:"disable_execute_api_endpoint"`
	Id                        string            `json:"id"`
	ProtocolType              string            `json:"protocol_type"`
	RouteKey                  string            `json:"route_key"`
	CorsConfiguration         CorsConfiguration `json:"cors_configuration"`
}

type CorsConfiguration struct {
	AllowCredentials bool     `json:"allow_credentials"`
	AllowHeaders     []string `json:"allow_headers"`
	AllowMethods     []string `json:"allow_methods"`
	AllowOrigins     []string `json:"allow_origins"`
	ExposeHeaders    []string `json:"expose_headers"`
	MaxAge           int64    `json:"max_age"`
}

// A Apigatewayv2ApiStatus defines the observed state of a Apigatewayv2Api
type Apigatewayv2ApiStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Apigatewayv2ApiObservation `json:",inline"`
}

// A Apigatewayv2ApiObservation records the observed state of a Apigatewayv2Api
type Apigatewayv2ApiObservation struct {
	ExecutionArn string `json:"execution_arn"`
	ApiEndpoint  string `json:"api_endpoint"`
	Arn          string `json:"arn"`
}