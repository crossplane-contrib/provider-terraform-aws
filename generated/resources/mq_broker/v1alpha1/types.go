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
	Tags                       map[string]string          `json:"tags"`
	AutoMinorVersionUpgrade    bool                       `json:"auto_minor_version_upgrade"`
	EngineVersion              string                     `json:"engine_version"`
	HostInstanceType           string                     `json:"host_instance_type"`
	PubliclyAccessible         bool                       `json:"publicly_accessible"`
	ApplyImmediately           bool                       `json:"apply_immediately"`
	BrokerName                 string                     `json:"broker_name"`
	DeploymentMode             string                     `json:"deployment_mode"`
	EngineType                 string                     `json:"engine_type"`
	SecurityGroups             []string                   `json:"security_groups"`
	Configuration              Configuration              `json:"configuration"`
	EncryptionOptions          EncryptionOptions          `json:"encryption_options"`
	Logs                       Logs                       `json:"logs"`
	MaintenanceWindowStartTime MaintenanceWindowStartTime `json:"maintenance_window_start_time"`
	User                       []User                     `json:"user"`
}

type Configuration struct {
	Id       string `json:"id"`
	Revision int    `json:"revision"`
}

type EncryptionOptions struct {
	KmsKeyId       string `json:"kms_key_id"`
	UseAwsOwnedKey bool   `json:"use_aws_owned_key"`
}

type Logs struct {
	Audit   bool `json:"audit"`
	General bool `json:"general"`
}

type MaintenanceWindowStartTime struct {
	DayOfWeek string `json:"day_of_week"`
	TimeOfDay string `json:"time_of_day"`
	TimeZone  string `json:"time_zone"`
}

type User struct {
	ConsoleAccess bool     `json:"console_access"`
	Groups        []string `json:"groups"`
	Password      string   `json:"password"`
	Username      string   `json:"username"`
}

// A MqBrokerStatus defines the observed state of a MqBroker
type MqBrokerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     MqBrokerObservation `json:",inline"`
}

// A MqBrokerObservation records the observed state of a MqBroker
type MqBrokerObservation struct {
	SubnetIds []string `json:"subnet_ids"`
	Arn       string   `json:"arn"`
	Id        string   `json:"id"`
}