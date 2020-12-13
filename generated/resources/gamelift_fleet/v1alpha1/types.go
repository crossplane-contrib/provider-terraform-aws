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

// GameliftFleet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GameliftFleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GameliftFleetSpec   `json:"spec"`
	Status GameliftFleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GameliftFleet contains a list of GameliftFleetList
type GameliftFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GameliftFleet `json:"items"`
}

// A GameliftFleetSpec defines the desired state of a GameliftFleet
type GameliftFleetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GameliftFleetParameters `json:"forProvider"`
}

// A GameliftFleetParameters defines the desired state of a GameliftFleet
type GameliftFleetParameters struct {
	Description                    string                      `json:"description"`
	FleetType                      string                      `json:"fleet_type"`
	Id                             string                      `json:"id"`
	NewGameSessionProtectionPolicy string                      `json:"new_game_session_protection_policy"`
	Name                           string                      `json:"name"`
	Tags                           map[string]string           `json:"tags"`
	BuildId                        string                      `json:"build_id"`
	Ec2InstanceType                string                      `json:"ec2_instance_type"`
	InstanceRoleArn                string                      `json:"instance_role_arn"`
	MetricGroups                   []string                    `json:"metric_groups"`
	Ec2InboundPermission           []Ec2InboundPermission      `json:"ec2_inbound_permission"`
	ResourceCreationLimitPolicy    ResourceCreationLimitPolicy `json:"resource_creation_limit_policy"`
	RuntimeConfiguration           RuntimeConfiguration        `json:"runtime_configuration"`
	Timeouts                       Timeouts                    `json:"timeouts"`
}

type Ec2InboundPermission struct {
	FromPort int64  `json:"from_port"`
	IpRange  string `json:"ip_range"`
	Protocol string `json:"protocol"`
	ToPort   int64  `json:"to_port"`
}

type ResourceCreationLimitPolicy struct {
	NewGameSessionsPerCreator int64 `json:"new_game_sessions_per_creator"`
	PolicyPeriodInMinutes     int64 `json:"policy_period_in_minutes"`
}

type RuntimeConfiguration struct {
	GameSessionActivationTimeoutSeconds int64           `json:"game_session_activation_timeout_seconds"`
	MaxConcurrentGameSessionActivations int64           `json:"max_concurrent_game_session_activations"`
	ServerProcess                       []ServerProcess `json:"server_process"`
}

type ServerProcess struct {
	ConcurrentExecutions int64  `json:"concurrent_executions"`
	LaunchPath           string `json:"launch_path"`
	Parameters           string `json:"parameters"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A GameliftFleetStatus defines the observed state of a GameliftFleet
type GameliftFleetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GameliftFleetObservation `json:"atProvider"`
}

// A GameliftFleetObservation records the observed state of a GameliftFleet
type GameliftFleetObservation struct {
	OperatingSystem string   `json:"operating_system"`
	Arn             string   `json:"arn"`
	LogPaths        []string `json:"log_paths"`
}