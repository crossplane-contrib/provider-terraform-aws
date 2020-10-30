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

// MqBroker is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type MqBroker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MqBrokerSpec   `json:"spec"`
	Status MqBrokerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MqBroker contains a list of MqBrokerList
type MqBrokerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MqBroker `json:"items"`
}

// A MqBrokerSpec defines the desired state of a MqBroker
type MqBrokerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  MqBrokerParameters `json:",inline"`
}

// A MqBrokerParameters defines the desired state of a MqBroker
type MqBrokerParameters struct {
	EngineType              string   `json:"engine_type"`
	AutoMinorVersionUpgrade bool     `json:"auto_minor_version_upgrade"`
	BrokerName              string   `json:"broker_name"`
	PubliclyAccessible      bool     `json:"publicly_accessible"`
	SecurityGroups          []string `json:"security_groups"`
	ApplyImmediately        bool     `json:"apply_immediately"`
	DeploymentMode          string   `json:"deployment_mode"`
	EngineVersion           string   `json:"engine_version"`
	HostInstanceType        string   `json:"host_instance_type"`
}

// A MqBrokerStatus defines the observed state of a MqBroker
type MqBrokerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     MqBrokerObservation `json:",inline"`
}

// A MqBrokerObservation records the observed state of a MqBroker
type MqBrokerObservation struct {
	Arn       string   `json:"arn"`
	SubnetIds []string `json:"subnet_ids"`
	Id        string   `json:"id"`
}