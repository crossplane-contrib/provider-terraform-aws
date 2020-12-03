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
	EngineVersion              string            `json:"engine_version"`
	PreferredMaintenanceWindow string            `json:"preferred_maintenance_window"`
	PubliclyAccessible         bool              `json:"publicly_accessible"`
	Tags                       map[string]string `json:"tags"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	IdentifierPrefix           string            `json:"identifier_prefix"`
	InstanceClass              string            `json:"instance_class"`
	Port                       int               `json:"port"`
	PromotionTier              int               `json:"promotion_tier"`
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	ClusterIdentifier          string            `json:"cluster_identifier"`
	Engine                     string            `json:"engine"`
	Id                         string            `json:"id"`
	NeptuneParameterGroupName  string            `json:"neptune_parameter_group_name"`
	NeptuneSubnetGroupName     string            `json:"neptune_subnet_group_name"`
	AvailabilityZone           string            `json:"availability_zone"`
	Identifier                 string            `json:"identifier"`
	PreferredBackupWindow      string            `json:"preferred_backup_window"`
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
	Arn              string `json:"arn"`
	Endpoint         string `json:"endpoint"`
	StorageEncrypted bool   `json:"storage_encrypted"`
	KmsKeyArn        string `json:"kms_key_arn"`
	Address          string `json:"address"`
	DbiResourceId    string `json:"dbi_resource_id"`
	Writer           bool   `json:"writer"`
}