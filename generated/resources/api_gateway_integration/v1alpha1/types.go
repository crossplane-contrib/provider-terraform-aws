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

// ApiGatewayIntegration is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayIntegration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayIntegrationSpec   `json:"spec"`
	Status ApiGatewayIntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayIntegration contains a list of ApiGatewayIntegrationList
type ApiGatewayIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayIntegration `json:"items"`
}

// A ApiGatewayIntegrationSpec defines the desired state of a ApiGatewayIntegration
type ApiGatewayIntegrationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayIntegrationParameters `json:",inline"`
}

// A ApiGatewayIntegrationParameters defines the desired state of a ApiGatewayIntegration
type ApiGatewayIntegrationParameters struct {
	RestApiId             string            `json:"rest_api_id"`
	TimeoutMilliseconds   int               `json:"timeout_milliseconds"`
	Uri                   string            `json:"uri"`
	ConnectionId          string            `json:"connection_id"`
	ContentHandling       string            `json:"content_handling"`
	Credentials           string            `json:"credentials"`
	HttpMethod            string            `json:"http_method"`
	IntegrationHttpMethod string            `json:"integration_http_method"`
	CacheKeyParameters    []string          `json:"cache_key_parameters"`
	ConnectionType        string            `json:"connection_type"`
	RequestParameters     map[string]string `json:"request_parameters"`
	RequestTemplates      map[string]string `json:"request_templates"`
	Type                  string            `json:"type"`
	ResourceId            string            `json:"resource_id"`
}

// A ApiGatewayIntegrationStatus defines the observed state of a ApiGatewayIntegration
type ApiGatewayIntegrationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayIntegrationObservation `json:",inline"`
}

// A ApiGatewayIntegrationObservation records the observed state of a ApiGatewayIntegration
type ApiGatewayIntegrationObservation struct {
	CacheNamespace      string `json:"cache_namespace"`
	PassthroughBehavior string `json:"passthrough_behavior"`
	Id                  string `json:"id"`
}