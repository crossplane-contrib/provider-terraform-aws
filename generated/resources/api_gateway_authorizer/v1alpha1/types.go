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

// ApiGatewayAuthorizer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayAuthorizer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayAuthorizerSpec   `json:"spec"`
	Status ApiGatewayAuthorizerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayAuthorizer contains a list of ApiGatewayAuthorizerList
type ApiGatewayAuthorizerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayAuthorizer `json:"items"`
}

// A ApiGatewayAuthorizerSpec defines the desired state of a ApiGatewayAuthorizer
type ApiGatewayAuthorizerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayAuthorizerParameters `json:"forProvider"`
}

// A ApiGatewayAuthorizerParameters defines the desired state of a ApiGatewayAuthorizer
type ApiGatewayAuthorizerParameters struct {
	AuthorizerResultTtlInSeconds int64    `json:"authorizer_result_ttl_in_seconds"`
	IdentityValidationExpression string   `json:"identity_validation_expression"`
	ProviderArns                 []string `json:"provider_arns"`
	RestApiId                    string   `json:"rest_api_id"`
	AuthorizerCredentials        string   `json:"authorizer_credentials"`
	AuthorizerUri                string   `json:"authorizer_uri"`
	IdentitySource               string   `json:"identity_source"`
	Name                         string   `json:"name"`
	Type                         string   `json:"type"`
}

// A ApiGatewayAuthorizerStatus defines the observed state of a ApiGatewayAuthorizer
type ApiGatewayAuthorizerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayAuthorizerObservation `json:"atProvider"`
}

// A ApiGatewayAuthorizerObservation records the observed state of a ApiGatewayAuthorizer
type ApiGatewayAuthorizerObservation struct{}