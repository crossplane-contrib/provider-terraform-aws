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

// SagemakerEndpoint is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SagemakerEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SagemakerEndpointSpec   `json:"spec"`
	Status SagemakerEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerEndpoint contains a list of SagemakerEndpointList
type SagemakerEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerEndpoint `json:"items"`
}

// A SagemakerEndpointSpec defines the desired state of a SagemakerEndpoint
type SagemakerEndpointSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SagemakerEndpointParameters `json:"forProvider"`
}

// A SagemakerEndpointParameters defines the desired state of a SagemakerEndpoint
type SagemakerEndpointParameters struct {
	EndpointConfigName string            `json:"endpoint_config_name"`
	Id                 string            `json:"id"`
	Name               string            `json:"name"`
	Tags               map[string]string `json:"tags"`
}

// A SagemakerEndpointStatus defines the observed state of a SagemakerEndpoint
type SagemakerEndpointStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SagemakerEndpointObservation `json:"atProvider"`
}

// A SagemakerEndpointObservation records the observed state of a SagemakerEndpoint
type SagemakerEndpointObservation struct {
	Arn string `json:"arn"`
}