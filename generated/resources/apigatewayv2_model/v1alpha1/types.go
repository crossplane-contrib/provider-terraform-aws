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

// Apigatewayv2Model is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Apigatewayv2Model struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Apigatewayv2ModelSpec   `json:"spec"`
	Status Apigatewayv2ModelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2Model contains a list of Apigatewayv2ModelList
type Apigatewayv2ModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2Model `json:"items"`
}

// A Apigatewayv2ModelSpec defines the desired state of a Apigatewayv2Model
type Apigatewayv2ModelSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Apigatewayv2ModelParameters `json:",inline"`
}

// A Apigatewayv2ModelParameters defines the desired state of a Apigatewayv2Model
type Apigatewayv2ModelParameters struct {
	ApiId       string `json:"api_id"`
	ContentType string `json:"content_type"`
	Description string `json:"description"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Schema      string `json:"schema"`
}

// A Apigatewayv2ModelStatus defines the observed state of a Apigatewayv2Model
type Apigatewayv2ModelStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Apigatewayv2ModelObservation `json:",inline"`
}

// A Apigatewayv2ModelObservation records the observed state of a Apigatewayv2Model
type Apigatewayv2ModelObservation struct{}