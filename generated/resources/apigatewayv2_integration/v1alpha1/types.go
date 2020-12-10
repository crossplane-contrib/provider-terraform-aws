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

// Apigatewayv2Integration is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Apigatewayv2Integration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Apigatewayv2IntegrationSpec   `json:"spec"`
	Status Apigatewayv2IntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2Integration contains a list of Apigatewayv2IntegrationList
type Apigatewayv2IntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2Integration `json:"items"`
}

// A Apigatewayv2IntegrationSpec defines the desired state of a Apigatewayv2Integration
type Apigatewayv2IntegrationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Apigatewayv2IntegrationParameters `json:",inline"`
}

// A Apigatewayv2IntegrationParameters defines the desired state of a Apigatewayv2Integration
type Apigatewayv2IntegrationParameters struct {
	IntegrationUri              string            `json:"integration_uri"`
	Id                          string            `json:"id"`
	ConnectionId                string            `json:"connection_id"`
	ConnectionType              string            `json:"connection_type"`
	CredentialsArn              string            `json:"credentials_arn"`
	IntegrationMethod           string            `json:"integration_method"`
	IntegrationType             string            `json:"integration_type"`
	RequestTemplates            map[string]string `json:"request_templates"`
	ApiId                       string            `json:"api_id"`
	TemplateSelectionExpression string            `json:"template_selection_expression"`
	PayloadFormatVersion        string            `json:"payload_format_version"`
	RequestParameters           map[string]string `json:"request_parameters"`
	IntegrationSubtype          string            `json:"integration_subtype"`
	Description                 string            `json:"description"`
	PassthroughBehavior         string            `json:"passthrough_behavior"`
	TimeoutMilliseconds         int64             `json:"timeout_milliseconds"`
	ContentHandlingStrategy     string            `json:"content_handling_strategy"`
	TlsConfig                   TlsConfig         `json:"tls_config"`
}

type TlsConfig struct {
	ServerNameToVerify string `json:"server_name_to_verify"`
}

// A Apigatewayv2IntegrationStatus defines the observed state of a Apigatewayv2Integration
type Apigatewayv2IntegrationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Apigatewayv2IntegrationObservation `json:",inline"`
}

// A Apigatewayv2IntegrationObservation records the observed state of a Apigatewayv2Integration
type Apigatewayv2IntegrationObservation struct {
	IntegrationResponseSelectionExpression string `json:"integration_response_selection_expression"`
}