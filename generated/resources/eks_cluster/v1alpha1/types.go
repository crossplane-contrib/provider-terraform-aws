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

// EksCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EksCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EksClusterSpec   `json:"spec"`
	Status EksClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EksCluster contains a list of EksClusterList
type EksClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EksCluster `json:"items"`
}

// A EksClusterSpec defines the desired state of a EksCluster
type EksClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EksClusterParameters `json:",inline"`
}

// A EksClusterParameters defines the desired state of a EksCluster
type EksClusterParameters struct {
	Name                   string   `json:"name"`
	RoleArn                string   `json:"role_arn"`
	EnabledClusterLogTypes []string `json:"enabled_cluster_log_types"`
}

// A EksClusterStatus defines the observed state of a EksCluster
type EksClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EksClusterObservation `json:",inline"`
}

// A EksClusterObservation records the observed state of a EksCluster
type EksClusterObservation struct {
	CreatedAt       string `json:"created_at"`
	Endpoint        string `json:"endpoint"`
	Version         string `json:"version"`
	Arn             string `json:"arn"`
	Id              string `json:"id"`
	PlatformVersion string `json:"platform_version"`
	Status          string `json:"status"`
}