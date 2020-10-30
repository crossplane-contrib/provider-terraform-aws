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

// Apigatewayv2DomainName is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Apigatewayv2DomainName struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Apigatewayv2DomainNameSpec   `json:"spec"`
	Status Apigatewayv2DomainNameStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2DomainName contains a list of Apigatewayv2DomainNameList
type Apigatewayv2DomainNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2DomainName `json:"items"`
}

// A Apigatewayv2DomainNameSpec defines the desired state of a Apigatewayv2DomainName
type Apigatewayv2DomainNameSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Apigatewayv2DomainNameParameters `json:",inline"`
}

// A Apigatewayv2DomainNameParameters defines the desired state of a Apigatewayv2DomainName
type Apigatewayv2DomainNameParameters struct {
	DomainName string `json:"domain_name"`
}

// A Apigatewayv2DomainNameStatus defines the observed state of a Apigatewayv2DomainName
type Apigatewayv2DomainNameStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Apigatewayv2DomainNameObservation `json:",inline"`
}

// A Apigatewayv2DomainNameObservation records the observed state of a Apigatewayv2DomainName
type Apigatewayv2DomainNameObservation struct {
	ApiMappingSelectionExpression string `json:"api_mapping_selection_expression"`
	Arn                           string `json:"arn"`
	Id                            string `json:"id"`
}