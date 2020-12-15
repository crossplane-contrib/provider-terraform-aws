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

// DmsReplicationInstance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DmsReplicationInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DmsReplicationInstanceSpec   `json:"spec"`
	Status DmsReplicationInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DmsReplicationInstance contains a list of DmsReplicationInstanceList
type DmsReplicationInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DmsReplicationInstance `json:"items"`
}

// A DmsReplicationInstanceSpec defines the desired state of a DmsReplicationInstance
type DmsReplicationInstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DmsReplicationInstanceParameters `json:"forProvider"`
}

// A DmsReplicationInstanceParameters defines the desired state of a DmsReplicationInstance
type DmsReplicationInstanceParameters struct {
	Tags                       map[string]string `json:"tags"`
	VpcSecurityGroupIds        []string          `json:"vpc_security_group_ids"`
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	EngineVersion              string            `json:"engine_version"`
	MultiAz                    bool              `json:"multi_az"`
	AvailabilityZone           string            `json:"availability_zone"`
	PreferredMaintenanceWindow string            `json:"preferred_maintenance_window"`
	ReplicationSubnetGroupId   string            `json:"replication_subnet_group_id"`
	AllowMajorVersionUpgrade   bool              `json:"allow_major_version_upgrade"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	KmsKeyArn                  string            `json:"kms_key_arn"`
	PubliclyAccessible         bool              `json:"publicly_accessible"`
	ReplicationInstanceId      string            `json:"replication_instance_id"`
	AllocatedStorage           int64             `json:"allocated_storage"`
	ReplicationInstanceClass   string            `json:"replication_instance_class"`
	Timeouts                   Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A DmsReplicationInstanceStatus defines the observed state of a DmsReplicationInstance
type DmsReplicationInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DmsReplicationInstanceObservation `json:"atProvider"`
}

// A DmsReplicationInstanceObservation records the observed state of a DmsReplicationInstance
type DmsReplicationInstanceObservation struct {
	ReplicationInstancePrivateIps []string `json:"replication_instance_private_ips"`
	ReplicationInstanceArn        string   `json:"replication_instance_arn"`
	ReplicationInstancePublicIps  []string `json:"replication_instance_public_ips"`
}