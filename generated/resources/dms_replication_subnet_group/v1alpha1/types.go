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

// DmsReplicationSubnetGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DmsReplicationSubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DmsReplicationSubnetGroupSpec   `json:"spec"`
	Status DmsReplicationSubnetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DmsReplicationSubnetGroup contains a list of DmsReplicationSubnetGroupList
type DmsReplicationSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DmsReplicationSubnetGroup `json:"items"`
}

// A DmsReplicationSubnetGroupSpec defines the desired state of a DmsReplicationSubnetGroup
type DmsReplicationSubnetGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DmsReplicationSubnetGroupParameters `json:"forProvider"`
}

// A DmsReplicationSubnetGroupParameters defines the desired state of a DmsReplicationSubnetGroup
type DmsReplicationSubnetGroupParameters struct {
	ReplicationSubnetGroupDescription string            `json:"replication_subnet_group_description"`
	ReplicationSubnetGroupId          string            `json:"replication_subnet_group_id"`
	SubnetIds                         []string          `json:"subnet_ids"`
	Tags                              map[string]string `json:"tags"`
}

// A DmsReplicationSubnetGroupStatus defines the observed state of a DmsReplicationSubnetGroup
type DmsReplicationSubnetGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DmsReplicationSubnetGroupObservation `json:"atProvider"`
}

// A DmsReplicationSubnetGroupObservation records the observed state of a DmsReplicationSubnetGroup
type DmsReplicationSubnetGroupObservation struct {
	VpcId                     string `json:"vpc_id"`
	ReplicationSubnetGroupArn string `json:"replication_subnet_group_arn"`
}