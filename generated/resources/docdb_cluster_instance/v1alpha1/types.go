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

// DocdbClusterInstance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DocdbClusterInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DocdbClusterInstanceSpec   `json:"spec"`
	Status DocdbClusterInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocdbClusterInstance contains a list of DocdbClusterInstanceList
type DocdbClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocdbClusterInstance `json:"items"`
}

// A DocdbClusterInstanceSpec defines the desired state of a DocdbClusterInstance
type DocdbClusterInstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DocdbClusterInstanceParameters `json:",inline"`
}

// A DocdbClusterInstanceParameters defines the desired state of a DocdbClusterInstance
type DocdbClusterInstanceParameters struct {
	IdentifierPrefix           string            `json:"identifier_prefix"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	AvailabilityZone           string            `json:"availability_zone"`
	CaCertIdentifier           string            `json:"ca_cert_identifier"`
	PromotionTier              int               `json:"promotion_tier"`
	Engine                     string            `json:"engine"`
	Identifier                 string            `json:"identifier"`
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	InstanceClass              string            `json:"instance_class"`
	Tags                       map[string]string `json:"tags"`
	ClusterIdentifier          string            `json:"cluster_identifier"`
	Id                         string            `json:"id"`
	PreferredMaintenanceWindow string            `json:"preferred_maintenance_window"`
	Timeouts                   []Timeouts        `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A DocdbClusterInstanceStatus defines the observed state of a DocdbClusterInstance
type DocdbClusterInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DocdbClusterInstanceObservation `json:",inline"`
}

// A DocdbClusterInstanceObservation records the observed state of a DocdbClusterInstance
type DocdbClusterInstanceObservation struct {
	DbSubnetGroupName     string `json:"db_subnet_group_name"`
	Endpoint              string `json:"endpoint"`
	PubliclyAccessible    bool   `json:"publicly_accessible"`
	EngineVersion         string `json:"engine_version"`
	KmsKeyId              string `json:"kms_key_id"`
	Port                  int    `json:"port"`
	DbiResourceId         string `json:"dbi_resource_id"`
	StorageEncrypted      bool   `json:"storage_encrypted"`
	Writer                bool   `json:"writer"`
	Arn                   string `json:"arn"`
	PreferredBackupWindow string `json:"preferred_backup_window"`
}