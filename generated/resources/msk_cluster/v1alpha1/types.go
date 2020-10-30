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

// MskCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type MskCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MskClusterSpec   `json:"spec"`
	Status MskClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MskCluster contains a list of MskClusterList
type MskClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MskCluster `json:"items"`
}

// A MskClusterSpec defines the desired state of a MskCluster
type MskClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  MskClusterParameters `json:",inline"`
}

// A MskClusterParameters defines the desired state of a MskCluster
type MskClusterParameters struct {
	ClusterName         string `json:"cluster_name"`
	EnhancedMonitoring  string `json:"enhanced_monitoring"`
	KafkaVersion        string `json:"kafka_version"`
	NumberOfBrokerNodes int    `json:"number_of_broker_nodes"`
}

// A MskClusterStatus defines the observed state of a MskCluster
type MskClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     MskClusterObservation `json:",inline"`
}

// A MskClusterObservation records the observed state of a MskCluster
type MskClusterObservation struct {
	BootstrapBrokersTls    string `json:"bootstrap_brokers_tls"`
	CurrentVersion         string `json:"current_version"`
	Id                     string `json:"id"`
	ZookeeperConnectString string `json:"zookeeper_connect_string"`
	Arn                    string `json:"arn"`
	BootstrapBrokers       string `json:"bootstrap_brokers"`
}