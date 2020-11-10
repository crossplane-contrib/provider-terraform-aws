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

// ApiGatewayMethodResponse is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayMethodResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayMethodResponseSpec   `json:"spec"`
	Status ApiGatewayMethodResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayMethodResponse contains a list of ApiGatewayMethodResponseList
type ApiGatewayMethodResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayMethodResponse `json:"items"`
}

// A ApiGatewayMethodResponseSpec defines the desired state of a ApiGatewayMethodResponse
type ApiGatewayMethodResponseSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayMethodResponseParameters `json:",inline"`
}

// A ApiGatewayMethodResponseParameters defines the desired state of a ApiGatewayMethodResponse
type ApiGatewayMethodResponseParameters struct {
	ResourceId         string            `json:"resource_id"`
	ResponseModels     map[string]string `json:"response_models"`
	ResponseParameters map[string]bool   `json:"response_parameters"`
	RestApiId          string            `json:"rest_api_id"`
	StatusCode         string            `json:"status_code"`
	HttpMethod         string            `json:"http_method"`
}

// A ApiGatewayMethodResponseStatus defines the observed state of a ApiGatewayMethodResponse
type ApiGatewayMethodResponseStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayMethodResponseObservation `json:",inline"`
}

// A ApiGatewayMethodResponseObservation records the observed state of a ApiGatewayMethodResponse
type ApiGatewayMethodResponseObservation struct {
	Id string `json:"id"`
}