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

// ApiGatewayDocumentationPart is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayDocumentationPart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayDocumentationPartSpec   `json:"spec"`
	Status ApiGatewayDocumentationPartStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayDocumentationPart contains a list of ApiGatewayDocumentationPartList
type ApiGatewayDocumentationPartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayDocumentationPart `json:"items"`
}

// A ApiGatewayDocumentationPartSpec defines the desired state of a ApiGatewayDocumentationPart
type ApiGatewayDocumentationPartSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayDocumentationPartParameters `json:"forProvider"`
}

// A ApiGatewayDocumentationPartParameters defines the desired state of a ApiGatewayDocumentationPart
type ApiGatewayDocumentationPartParameters struct {
	RestApiId  string   `json:"rest_api_id"`
	Id         string   `json:"id"`
	Properties string   `json:"properties"`
	Location   Location `json:"location"`
}

type Location struct {
	Type       string `json:"type"`
	Method     string `json:"method"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	StatusCode string `json:"status_code"`
}

// A ApiGatewayDocumentationPartStatus defines the observed state of a ApiGatewayDocumentationPart
type ApiGatewayDocumentationPartStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayDocumentationPartObservation `json:"atProvider"`
}

// A ApiGatewayDocumentationPartObservation records the observed state of a ApiGatewayDocumentationPart
type ApiGatewayDocumentationPartObservation struct{}