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

// ApiGatewayApiKey is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayApiKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayApiKeySpec   `json:"spec"`
	Status ApiGatewayApiKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayApiKey contains a list of ApiGatewayApiKeyList
type ApiGatewayApiKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayApiKey `json:"items"`
}

// A ApiGatewayApiKeySpec defines the desired state of a ApiGatewayApiKey
type ApiGatewayApiKeySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayApiKeyParameters `json:"forProvider"`
}

// A ApiGatewayApiKeyParameters defines the desired state of a ApiGatewayApiKey
type ApiGatewayApiKeyParameters struct {
	Id          string            `json:"id"`
	Tags        map[string]string `json:"tags"`
	Value       string            `json:"value"`
	Description string            `json:"description"`
	Enabled     bool              `json:"enabled"`
	Name        string            `json:"name"`
}

// A ApiGatewayApiKeyStatus defines the observed state of a ApiGatewayApiKey
type ApiGatewayApiKeyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayApiKeyObservation `json:"atProvider"`
}

// A ApiGatewayApiKeyObservation records the observed state of a ApiGatewayApiKey
type ApiGatewayApiKeyObservation struct {
	Arn             string `json:"arn"`
	CreatedDate     string `json:"created_date"`
	LastUpdatedDate string `json:"last_updated_date"`
}