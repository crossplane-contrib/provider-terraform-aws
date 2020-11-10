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

// DocdbClusterSnapshot is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DocdbClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DocdbClusterSnapshotSpec   `json:"spec"`
	Status DocdbClusterSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocdbClusterSnapshot contains a list of DocdbClusterSnapshotList
type DocdbClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocdbClusterSnapshot `json:"items"`
}

// A DocdbClusterSnapshotSpec defines the desired state of a DocdbClusterSnapshot
type DocdbClusterSnapshotSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DocdbClusterSnapshotParameters `json:",inline"`
}

// A DocdbClusterSnapshotParameters defines the desired state of a DocdbClusterSnapshot
type DocdbClusterSnapshotParameters struct {
	DbClusterIdentifier         string     `json:"db_cluster_identifier"`
	DbClusterSnapshotIdentifier string     `json:"db_cluster_snapshot_identifier"`
	Timeouts                    []Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
}

// A DocdbClusterSnapshotStatus defines the observed state of a DocdbClusterSnapshot
type DocdbClusterSnapshotStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DocdbClusterSnapshotObservation `json:",inline"`
}

// A DocdbClusterSnapshotObservation records the observed state of a DocdbClusterSnapshot
type DocdbClusterSnapshotObservation struct {
	AvailabilityZones          []string `json:"availability_zones"`
	Status                     string   `json:"status"`
	SourceDbClusterSnapshotArn string   `json:"source_db_cluster_snapshot_arn"`
	StorageEncrypted           bool     `json:"storage_encrypted"`
	KmsKeyId                   string   `json:"kms_key_id"`
	SnapshotType               string   `json:"snapshot_type"`
	Port                       int      `json:"port"`
	VpcId                      string   `json:"vpc_id"`
	DbClusterSnapshotArn       string   `json:"db_cluster_snapshot_arn"`
	Engine                     string   `json:"engine"`
	EngineVersion              string   `json:"engine_version"`
	Id                         string   `json:"id"`
}