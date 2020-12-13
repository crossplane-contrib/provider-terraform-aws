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

// Apigatewayv2IntegrationResponse is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Apigatewayv2IntegrationResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Apigatewayv2IntegrationResponseSpec   `json:"spec"`
	Status Apigatewayv2IntegrationResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2IntegrationResponse contains a list of Apigatewayv2IntegrationResponseList
type Apigatewayv2IntegrationResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2IntegrationResponse `json:"items"`
}

// A Apigatewayv2IntegrationResponseSpec defines the desired state of a Apigatewayv2IntegrationResponse
type Apigatewayv2IntegrationResponseSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Apigatewayv2IntegrationResponseParameters `json:"forProvider"`
}

// A Apigatewayv2IntegrationResponseParameters defines the desired state of a Apigatewayv2IntegrationResponse
type Apigatewayv2IntegrationResponseParameters struct {
	IntegrationId               string            `json:"integration_id"`
	IntegrationResponseKey      string            `json:"integration_response_key"`
	ResponseTemplates           map[string]string `json:"response_templates"`
	TemplateSelectionExpression string            `json:"template_selection_expression"`
	ApiId                       string            `json:"api_id"`
	ContentHandlingStrategy     string            `json:"content_handling_strategy"`
	Id                          string            `json:"id"`
}

// A Apigatewayv2IntegrationResponseStatus defines the observed state of a Apigatewayv2IntegrationResponse
type Apigatewayv2IntegrationResponseStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Apigatewayv2IntegrationResponseObservation `json:"atProvider"`
}

// A Apigatewayv2IntegrationResponseObservation records the observed state of a Apigatewayv2IntegrationResponse
type Apigatewayv2IntegrationResponseObservation struct{}