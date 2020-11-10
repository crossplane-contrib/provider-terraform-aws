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

// RdsGlobalCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RdsGlobalCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RdsGlobalClusterSpec   `json:"spec"`
	Status RdsGlobalClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdsGlobalCluster contains a list of RdsGlobalClusterList
type RdsGlobalClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RdsGlobalCluster `json:"items"`
}

// A RdsGlobalClusterSpec defines the desired state of a RdsGlobalCluster
type RdsGlobalClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RdsGlobalClusterParameters `json:",inline"`
}

// A RdsGlobalClusterParameters defines the desired state of a RdsGlobalCluster
type RdsGlobalClusterParameters struct {
	ForceDestroy            bool   `json:"force_destroy"`
	GlobalClusterIdentifier string `json:"global_cluster_identifier"`
	DatabaseName            string `json:"database_name"`
	StorageEncrypted        bool   `json:"storage_encrypted"`
	DeletionProtection      bool   `json:"deletion_protection"`
}

// A RdsGlobalClusterStatus defines the observed state of a RdsGlobalCluster
type RdsGlobalClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RdsGlobalClusterObservation `json:",inline"`
}

// A RdsGlobalClusterObservation records the observed state of a RdsGlobalCluster
type RdsGlobalClusterObservation struct {
	GlobalClusterResourceId   string `json:"global_cluster_resource_id"`
	SourceDbClusterIdentifier string `json:"source_db_cluster_identifier"`
	Engine                    string `json:"engine"`
	EngineVersion             string `json:"engine_version"`
	Id                        string `json:"id"`
	Arn                       string `json:"arn"`
}