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

// NeptuneClusterInstance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type NeptuneClusterInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NeptuneClusterInstanceSpec   `json:"spec"`
	Status NeptuneClusterInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneClusterInstance contains a list of NeptuneClusterInstanceList
type NeptuneClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NeptuneClusterInstance `json:"items"`
}

// A NeptuneClusterInstanceSpec defines the desired state of a NeptuneClusterInstance
type NeptuneClusterInstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  NeptuneClusterInstanceParameters `json:",inline"`
}

// A NeptuneClusterInstanceParameters defines the desired state of a NeptuneClusterInstance
type NeptuneClusterInstanceParameters struct {
	AvailabilityZone           string            `json:"availability_zone"`
	ClusterIdentifier          string            `json:"cluster_identifier"`
	NeptuneParameterGroupName  string            `json:"neptune_parameter_group_name"`
	PromotionTier              int               `json:"promotion_tier"`
	NeptuneSubnetGroupName     string            `json:"neptune_subnet_group_name"`
	PreferredBackupWindow      string            `json:"preferred_backup_window"`
	PreferredMaintenanceWindow string            `json:"preferred_maintenance_window"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	EngineVersion              string            `json:"engine_version"`
	Identifier                 string            `json:"identifier"`
	InstanceClass              string            `json:"instance_class"`
	Tags                       map[string]string `json:"tags"`
	Id                         string            `json:"id"`
	PubliclyAccessible         bool              `json:"publicly_accessible"`
	Engine                     string            `json:"engine"`
	IdentifierPrefix           string            `json:"identifier_prefix"`
	Port                       int               `json:"port"`
	Timeouts                   []Timeouts        `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A NeptuneClusterInstanceStatus defines the observed state of a NeptuneClusterInstance
type NeptuneClusterInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     NeptuneClusterInstanceObservation `json:",inline"`
}

// A NeptuneClusterInstanceObservation records the observed state of a NeptuneClusterInstance
type NeptuneClusterInstanceObservation struct {
	StorageEncrypted bool   `json:"storage_encrypted"`
	Writer           bool   `json:"writer"`
	Arn              string `json:"arn"`
	DbiResourceId    string `json:"dbi_resource_id"`
	Endpoint         string `json:"endpoint"`
	Address          string `json:"address"`
	KmsKeyArn        string `json:"kms_key_arn"`
}