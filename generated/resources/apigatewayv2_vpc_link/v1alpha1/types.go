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

// Apigatewayv2VpcLink is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Apigatewayv2VpcLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Apigatewayv2VpcLinkSpec   `json:"spec"`
	Status Apigatewayv2VpcLinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2VpcLink contains a list of Apigatewayv2VpcLinkList
type Apigatewayv2VpcLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2VpcLink `json:"items"`
}

// A Apigatewayv2VpcLinkSpec defines the desired state of a Apigatewayv2VpcLink
type Apigatewayv2VpcLinkSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Apigatewayv2VpcLinkParameters `json:"forProvider"`
}

// A Apigatewayv2VpcLinkParameters defines the desired state of a Apigatewayv2VpcLink
type Apigatewayv2VpcLinkParameters struct {
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	SecurityGroupIds []string          `json:"security_group_ids"`
	SubnetIds        []string          `json:"subnet_ids"`
	Tags             map[string]string `json:"tags"`
}

// A Apigatewayv2VpcLinkStatus defines the observed state of a Apigatewayv2VpcLink
type Apigatewayv2VpcLinkStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Apigatewayv2VpcLinkObservation `json:"atProvider"`
}

// A Apigatewayv2VpcLinkObservation records the observed state of a Apigatewayv2VpcLink
type Apigatewayv2VpcLinkObservation struct {
	Arn string `json:"arn"`
}