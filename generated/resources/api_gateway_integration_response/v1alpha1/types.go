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

// ApiGatewayIntegrationResponse is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayIntegrationResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayIntegrationResponseSpec   `json:"spec"`
	Status ApiGatewayIntegrationResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayIntegrationResponse contains a list of ApiGatewayIntegrationResponseList
type ApiGatewayIntegrationResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayIntegrationResponse `json:"items"`
}

// A ApiGatewayIntegrationResponseSpec defines the desired state of a ApiGatewayIntegrationResponse
type ApiGatewayIntegrationResponseSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayIntegrationResponseParameters `json:"forProvider"`
}

// A ApiGatewayIntegrationResponseParameters defines the desired state of a ApiGatewayIntegrationResponse
type ApiGatewayIntegrationResponseParameters struct {
	StatusCode         string            `json:"status_code"`
	ResponseParameters map[string]string `json:"response_parameters"`
	HttpMethod         string            `json:"http_method"`
	ResourceId         string            `json:"resource_id"`
	ResponseTemplates  map[string]string `json:"response_templates"`
	RestApiId          string            `json:"rest_api_id"`
	SelectionPattern   string            `json:"selection_pattern"`
	ContentHandling    string            `json:"content_handling"`
}

// A ApiGatewayIntegrationResponseStatus defines the observed state of a ApiGatewayIntegrationResponse
type ApiGatewayIntegrationResponseStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayIntegrationResponseObservation `json:"atProvider"`
}

// A ApiGatewayIntegrationResponseObservation records the observed state of a ApiGatewayIntegrationResponse
type ApiGatewayIntegrationResponseObservation struct{}