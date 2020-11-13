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
	ForProvider                  DmsReplicationInstanceParameters `json:",inline"`
}

// A DmsReplicationInstanceParameters defines the desired state of a DmsReplicationInstance
type DmsReplicationInstanceParameters struct {
	AvailabilityZone           string            `json:"availability_zone"`
	MultiAz                    bool              `json:"multi_az"`
	VpcSecurityGroupIds        []string          `json:"vpc_security_group_ids"`
	AllowMajorVersionUpgrade   bool              `json:"allow_major_version_upgrade"`
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	Id                         string            `json:"id"`
	ReplicationInstanceClass   string            `json:"replication_instance_class"`
	Tags                       map[string]string `json:"tags"`
	AllocatedStorage           int               `json:"allocated_storage"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	KmsKeyArn                  string            `json:"kms_key_arn"`
	PreferredMaintenanceWindow string            `json:"preferred_maintenance_window"`
	ReplicationInstanceId      string            `json:"replication_instance_id"`
	ReplicationSubnetGroupId   string            `json:"replication_subnet_group_id"`
	EngineVersion              string            `json:"engine_version"`
	PubliclyAccessible         bool              `json:"publicly_accessible"`
	Timeouts                   []Timeouts        `json:"timeouts"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Update string `json:"update"`
	Create string `json:"create"`
}

// A DmsReplicationInstanceStatus defines the observed state of a DmsReplicationInstance
type DmsReplicationInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DmsReplicationInstanceObservation `json:",inline"`
}

// A DmsReplicationInstanceObservation records the observed state of a DmsReplicationInstance
type DmsReplicationInstanceObservation struct {
	ReplicationInstancePrivateIps []string `json:"replication_instance_private_ips"`
	ReplicationInstancePublicIps  []string `json:"replication_instance_public_ips"`
	ReplicationInstanceArn        string   `json:"replication_instance_arn"`
}