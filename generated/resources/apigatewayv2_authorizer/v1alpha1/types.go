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

// Apigatewayv2Authorizer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Apigatewayv2Authorizer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Apigatewayv2AuthorizerSpec   `json:"spec"`
	Status Apigatewayv2AuthorizerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2Authorizer contains a list of Apigatewayv2AuthorizerList
type Apigatewayv2AuthorizerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2Authorizer `json:"items"`
}

// A Apigatewayv2AuthorizerSpec defines the desired state of a Apigatewayv2Authorizer
type Apigatewayv2AuthorizerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Apigatewayv2AuthorizerParameters `json:",inline"`
}

// A Apigatewayv2AuthorizerParameters defines the desired state of a Apigatewayv2Authorizer
type Apigatewayv2AuthorizerParameters struct {
	AuthorizerUri                  string           `json:"authorizer_uri"`
	EnableSimpleResponses          bool             `json:"enable_simple_responses"`
	Id                             string           `json:"id"`
	IdentitySources                []string         `json:"identity_sources"`
	ApiId                          string           `json:"api_id"`
	AuthorizerCredentialsArn       string           `json:"authorizer_credentials_arn"`
	AuthorizerType                 string           `json:"authorizer_type"`
	Name                           string           `json:"name"`
	AuthorizerPayloadFormatVersion string           `json:"authorizer_payload_format_version"`
	AuthorizerResultTtlInSeconds   int              `json:"authorizer_result_ttl_in_seconds"`
	JwtConfiguration               JwtConfiguration `json:"jwt_configuration"`
}

type JwtConfiguration struct {
	Audience []string `json:"audience"`
	Issuer   string   `json:"issuer"`
}

// A Apigatewayv2AuthorizerStatus defines the observed state of a Apigatewayv2Authorizer
type Apigatewayv2AuthorizerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Apigatewayv2AuthorizerObservation `json:",inline"`
}

// A Apigatewayv2AuthorizerObservation records the observed state of a Apigatewayv2Authorizer
type Apigatewayv2AuthorizerObservation struct{}