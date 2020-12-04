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

// RdsClusterEndpoint is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RdsClusterEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RdsClusterEndpointSpec   `json:"spec"`
	Status RdsClusterEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdsClusterEndpoint contains a list of RdsClusterEndpointList
type RdsClusterEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RdsClusterEndpoint `json:"items"`
}

// A RdsClusterEndpointSpec defines the desired state of a RdsClusterEndpoint
type RdsClusterEndpointSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RdsClusterEndpointParameters `json:",inline"`
}

// A RdsClusterEndpointParameters defines the desired state of a RdsClusterEndpoint
type RdsClusterEndpointParameters struct {
	CustomEndpointType        string            `json:"custom_endpoint_type"`
	ExcludedMembers           []string          `json:"excluded_members"`
	Id                        string            `json:"id"`
	StaticMembers             []string          `json:"static_members"`
	ClusterEndpointIdentifier string            `json:"cluster_endpoint_identifier"`
	ClusterIdentifier         string            `json:"cluster_identifier"`
	Tags                      map[string]string `json:"tags"`
}

// A RdsClusterEndpointStatus defines the observed state of a RdsClusterEndpoint
type RdsClusterEndpointStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RdsClusterEndpointObservation `json:",inline"`
}

// A RdsClusterEndpointObservation records the observed state of a RdsClusterEndpoint
type RdsClusterEndpointObservation struct {
	Arn      string `json:"arn"`
	Endpoint string `json:"endpoint"`
}