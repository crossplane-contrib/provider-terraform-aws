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

// NeptuneClusterSnapshot is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type NeptuneClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NeptuneClusterSnapshotSpec   `json:"spec"`
	Status NeptuneClusterSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneClusterSnapshot contains a list of NeptuneClusterSnapshotList
type NeptuneClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NeptuneClusterSnapshot `json:"items"`
}

// A NeptuneClusterSnapshotSpec defines the desired state of a NeptuneClusterSnapshot
type NeptuneClusterSnapshotSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  NeptuneClusterSnapshotParameters `json:",inline"`
}

// A NeptuneClusterSnapshotParameters defines the desired state of a NeptuneClusterSnapshot
type NeptuneClusterSnapshotParameters struct {
	Id                          string   `json:"id"`
	DbClusterIdentifier         string   `json:"db_cluster_identifier"`
	DbClusterSnapshotIdentifier string   `json:"db_cluster_snapshot_identifier"`
	Timeouts                    Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
}

// A NeptuneClusterSnapshotStatus defines the observed state of a NeptuneClusterSnapshot
type NeptuneClusterSnapshotStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     NeptuneClusterSnapshotObservation `json:",inline"`
}

// A NeptuneClusterSnapshotObservation records the observed state of a NeptuneClusterSnapshot
type NeptuneClusterSnapshotObservation struct {
	Engine                     string   `json:"engine"`
	EngineVersion              string   `json:"engine_version"`
	LicenseModel               string   `json:"license_model"`
	Status                     string   `json:"status"`
	VpcId                      string   `json:"vpc_id"`
	AllocatedStorage           int64    `json:"allocated_storage"`
	AvailabilityZones          []string `json:"availability_zones"`
	DbClusterSnapshotArn       string   `json:"db_cluster_snapshot_arn"`
	SnapshotType               string   `json:"snapshot_type"`
	StorageEncrypted           bool     `json:"storage_encrypted"`
	KmsKeyId                   string   `json:"kms_key_id"`
	Port                       int64    `json:"port"`
	SourceDbClusterSnapshotArn string   `json:"source_db_cluster_snapshot_arn"`
}