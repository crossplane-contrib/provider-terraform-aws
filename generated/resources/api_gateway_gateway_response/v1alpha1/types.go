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

// ApiGatewayGatewayResponse is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayGatewayResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayGatewayResponseSpec   `json:"spec"`
	Status ApiGatewayGatewayResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayGatewayResponse contains a list of ApiGatewayGatewayResponseList
type ApiGatewayGatewayResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayGatewayResponse `json:"items"`
}

// A ApiGatewayGatewayResponseSpec defines the desired state of a ApiGatewayGatewayResponse
type ApiGatewayGatewayResponseSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayGatewayResponseParameters `json:",inline"`
}

// A ApiGatewayGatewayResponseParameters defines the desired state of a ApiGatewayGatewayResponse
type ApiGatewayGatewayResponseParameters struct {
	ResponseParameters map[string]string `json:"response_parameters"`
	ResponseTemplates  map[string]string `json:"response_templates"`
	ResponseType       string            `json:"response_type"`
	RestApiId          string            `json:"rest_api_id"`
	StatusCode         string            `json:"status_code"`
}

// A ApiGatewayGatewayResponseStatus defines the observed state of a ApiGatewayGatewayResponse
type ApiGatewayGatewayResponseStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayGatewayResponseObservation `json:",inline"`
}

// A ApiGatewayGatewayResponseObservation records the observed state of a ApiGatewayGatewayResponse
type ApiGatewayGatewayResponseObservation struct {
	Id string `json:"id"`
}