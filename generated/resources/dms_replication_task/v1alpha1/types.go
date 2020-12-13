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

// DmsReplicationTask is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DmsReplicationTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DmsReplicationTaskSpec   `json:"spec"`
	Status DmsReplicationTaskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DmsReplicationTask contains a list of DmsReplicationTaskList
type DmsReplicationTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DmsReplicationTask `json:"items"`
}

// A DmsReplicationTaskSpec defines the desired state of a DmsReplicationTask
type DmsReplicationTaskSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DmsReplicationTaskParameters `json:"forProvider"`
}

// A DmsReplicationTaskParameters defines the desired state of a DmsReplicationTask
type DmsReplicationTaskParameters struct {
	TableMappings           string            `json:"table_mappings"`
	Tags                    map[string]string `json:"tags"`
	TargetEndpointArn       string            `json:"target_endpoint_arn"`
	ReplicationTaskId       string            `json:"replication_task_id"`
	ReplicationTaskSettings string            `json:"replication_task_settings"`
	SourceEndpointArn       string            `json:"source_endpoint_arn"`
	CdcStartTime            string            `json:"cdc_start_time"`
	Id                      string            `json:"id"`
	MigrationType           string            `json:"migration_type"`
	ReplicationInstanceArn  string            `json:"replication_instance_arn"`
}

// A DmsReplicationTaskStatus defines the observed state of a DmsReplicationTask
type DmsReplicationTaskStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DmsReplicationTaskObservation `json:"atProvider"`
}

// A DmsReplicationTaskObservation records the observed state of a DmsReplicationTask
type DmsReplicationTaskObservation struct {
	ReplicationTaskArn string `json:"replication_task_arn"`
}