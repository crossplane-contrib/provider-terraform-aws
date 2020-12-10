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

// ApiGatewayMethod is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayMethod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayMethodSpec   `json:"spec"`
	Status ApiGatewayMethodStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayMethod contains a list of ApiGatewayMethodList
type ApiGatewayMethodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayMethod `json:"items"`
}

// A ApiGatewayMethodSpec defines the desired state of a ApiGatewayMethod
type ApiGatewayMethodSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayMethodParameters `json:",inline"`
}

// A ApiGatewayMethodParameters defines the desired state of a ApiGatewayMethod
type ApiGatewayMethodParameters struct {
	ResourceId          string            `json:"resource_id"`
	ApiKeyRequired      bool              `json:"api_key_required"`
	AuthorizationScopes []string          `json:"authorization_scopes"`
	AuthorizerId        string            `json:"authorizer_id"`
	Id                  string            `json:"id"`
	RequestModels       map[string]string `json:"request_models"`
	Authorization       string            `json:"authorization"`
	HttpMethod          string            `json:"http_method"`
	RequestParameters   map[string]bool   `json:"request_parameters"`
	RequestValidatorId  string            `json:"request_validator_id"`
	RestApiId           string            `json:"rest_api_id"`
}

// A ApiGatewayMethodStatus defines the observed state of a ApiGatewayMethod
type ApiGatewayMethodStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayMethodObservation `json:",inline"`
}

// A ApiGatewayMethodObservation records the observed state of a ApiGatewayMethod
type ApiGatewayMethodObservation struct{}