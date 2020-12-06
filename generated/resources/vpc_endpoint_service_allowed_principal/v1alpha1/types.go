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

// VpcEndpointServiceAllowedPrincipal is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpcEndpointServiceAllowedPrincipal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpcEndpointServiceAllowedPrincipalSpec   `json:"spec"`
	Status VpcEndpointServiceAllowedPrincipalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcEndpointServiceAllowedPrincipal contains a list of VpcEndpointServiceAllowedPrincipalList
type VpcEndpointServiceAllowedPrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcEndpointServiceAllowedPrincipal `json:"items"`
}

// A VpcEndpointServiceAllowedPrincipalSpec defines the desired state of a VpcEndpointServiceAllowedPrincipal
type VpcEndpointServiceAllowedPrincipalSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpcEndpointServiceAllowedPrincipalParameters `json:",inline"`
}

// A VpcEndpointServiceAllowedPrincipalParameters defines the desired state of a VpcEndpointServiceAllowedPrincipal
type VpcEndpointServiceAllowedPrincipalParameters struct {
	Id                   string `json:"id"`
	PrincipalArn         string `json:"principal_arn"`
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`
}

// A VpcEndpointServiceAllowedPrincipalStatus defines the observed state of a VpcEndpointServiceAllowedPrincipal
type VpcEndpointServiceAllowedPrincipalStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpcEndpointServiceAllowedPrincipalObservation `json:",inline"`
}

// A VpcEndpointServiceAllowedPrincipalObservation records the observed state of a VpcEndpointServiceAllowedPrincipal
type VpcEndpointServiceAllowedPrincipalObservation struct{}