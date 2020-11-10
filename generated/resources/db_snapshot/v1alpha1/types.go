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

// DbSnapshot is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DbSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbSnapshotSpec   `json:"spec"`
	Status DbSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbSnapshot contains a list of DbSnapshotList
type DbSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbSnapshot `json:"items"`
}

// A DbSnapshotSpec defines the desired state of a DbSnapshot
type DbSnapshotSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DbSnapshotParameters `json:",inline"`
}

// A DbSnapshotParameters defines the desired state of a DbSnapshot
type DbSnapshotParameters struct {
	Tags                 map[string]string `json:"tags"`
	DbSnapshotIdentifier string            `json:"db_snapshot_identifier"`
	DbInstanceIdentifier string            `json:"db_instance_identifier"`
	Timeouts             []Timeouts        `json:"timeouts"`
}

type Timeouts struct {
	Read string `json:"read"`
}

// A DbSnapshotStatus defines the observed state of a DbSnapshot
type DbSnapshotStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DbSnapshotObservation `json:",inline"`
}

// A DbSnapshotObservation records the observed state of a DbSnapshot
type DbSnapshotObservation struct {
	Engine                     string `json:"engine"`
	KmsKeyId                   string `json:"kms_key_id"`
	OptionGroupName            string `json:"option_group_name"`
	Port                       int    `json:"port"`
	SnapshotType               string `json:"snapshot_type"`
	StorageType                string `json:"storage_type"`
	AvailabilityZone           string `json:"availability_zone"`
	EngineVersion              string `json:"engine_version"`
	LicenseModel               string `json:"license_model"`
	AllocatedStorage           int    `json:"allocated_storage"`
	Encrypted                  bool   `json:"encrypted"`
	Id                         string `json:"id"`
	SourceDbSnapshotIdentifier string `json:"source_db_snapshot_identifier"`
	SourceRegion               string `json:"source_region"`
	Status                     string `json:"status"`
	VpcId                      string `json:"vpc_id"`
	DbSnapshotArn              string `json:"db_snapshot_arn"`
	Iops                       int    `json:"iops"`
}