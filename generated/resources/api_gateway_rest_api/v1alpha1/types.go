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

// ApiGatewayRestApi is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayRestApi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayRestApiSpec   `json:"spec"`
	Status ApiGatewayRestApiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayRestApi contains a list of ApiGatewayRestApiList
type ApiGatewayRestApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayRestApi `json:"items"`
}

// A ApiGatewayRestApiSpec defines the desired state of a ApiGatewayRestApi
type ApiGatewayRestApiSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayRestApiParameters `json:",inline"`
}

// A ApiGatewayRestApiParameters defines the desired state of a ApiGatewayRestApi
type ApiGatewayRestApiParameters struct {
	MinimumCompressionSize int                   `json:"minimum_compression_size"`
	Policy                 string                `json:"policy"`
	Tags                   map[string]string     `json:"tags"`
	ApiKeySource           string                `json:"api_key_source"`
	BinaryMediaTypes       []string              `json:"binary_media_types"`
	Id                     string                `json:"id"`
	Name                   string                `json:"name"`
	Body                   string                `json:"body"`
	Description            string                `json:"description"`
	EndpointConfiguration  EndpointConfiguration `json:"endpoint_configuration"`
}

type EndpointConfiguration struct {
	Types          []string `json:"types"`
	VpcEndpointIds []string `json:"vpc_endpoint_ids"`
}

// A ApiGatewayRestApiStatus defines the observed state of a ApiGatewayRestApi
type ApiGatewayRestApiStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayRestApiObservation `json:",inline"`
}

// A ApiGatewayRestApiObservation records the observed state of a ApiGatewayRestApi
type ApiGatewayRestApiObservation struct {
	Arn            string `json:"arn"`
	ExecutionArn   string `json:"execution_arn"`
	RootResourceId string `json:"root_resource_id"`
	CreatedDate    string `json:"created_date"`
}