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

// DbProxyTarget is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DbProxyTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbProxyTargetSpec   `json:"spec"`
	Status DbProxyTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbProxyTarget contains a list of DbProxyTargetList
type DbProxyTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbProxyTarget `json:"items"`
}

// A DbProxyTargetSpec defines the desired state of a DbProxyTarget
type DbProxyTargetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DbProxyTargetParameters `json:"forProvider"`
}

// A DbProxyTargetParameters defines the desired state of a DbProxyTarget
type DbProxyTargetParameters struct {
	DbProxyName          string `json:"db_proxy_name"`
	Id                   string `json:"id"`
	TargetGroupName      string `json:"target_group_name"`
	DbClusterIdentifier  string `json:"db_cluster_identifier"`
	DbInstanceIdentifier string `json:"db_instance_identifier"`
}

// A DbProxyTargetStatus defines the observed state of a DbProxyTarget
type DbProxyTargetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DbProxyTargetObservation `json:"atProvider"`
}

// A DbProxyTargetObservation records the observed state of a DbProxyTarget
type DbProxyTargetObservation struct {
	Type             string `json:"type"`
	Endpoint         string `json:"endpoint"`
	TrackedClusterId string `json:"tracked_cluster_id"`
	Port             int64  `json:"port"`
	RdsResourceId    string `json:"rds_resource_id"`
	TargetArn        string `json:"target_arn"`
}