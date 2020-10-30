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

// ApiGatewayModel is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayModelSpec   `json:"spec"`
	Status ApiGatewayModelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayModel contains a list of ApiGatewayModelList
type ApiGatewayModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayModel `json:"items"`
}

// A ApiGatewayModelSpec defines the desired state of a ApiGatewayModel
type ApiGatewayModelSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayModelParameters `json:",inline"`
}

// A ApiGatewayModelParameters defines the desired state of a ApiGatewayModel
type ApiGatewayModelParameters struct {
	RestApiId   string `json:"rest_api_id"`
	Schema      string `json:"schema"`
	ContentType string `json:"content_type"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// A ApiGatewayModelStatus defines the observed state of a ApiGatewayModel
type ApiGatewayModelStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayModelObservation `json:",inline"`
}

// A ApiGatewayModelObservation records the observed state of a ApiGatewayModel
type ApiGatewayModelObservation struct {
	Id string `json:"id"`
}