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

// Route53HealthCheck is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Route53HealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Route53HealthCheckSpec   `json:"spec"`
	Status Route53HealthCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53HealthCheck contains a list of Route53HealthCheckList
type Route53HealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53HealthCheck `json:"items"`
}

// A Route53HealthCheckSpec defines the desired state of a Route53HealthCheck
type Route53HealthCheckSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Route53HealthCheckParameters `json:",inline"`
}

// A Route53HealthCheckParameters defines the desired state of a Route53HealthCheck
type Route53HealthCheckParameters struct {
	SearchString                 string            `json:"search_string"`
	Type                         string            `json:"type"`
	CloudwatchAlarmRegion        string            `json:"cloudwatch_alarm_region"`
	InsufficientDataHealthStatus string            `json:"insufficient_data_health_status"`
	IpAddress                    string            `json:"ip_address"`
	RequestInterval              int               `json:"request_interval"`
	ChildHealthThreshold         int               `json:"child_health_threshold"`
	Disabled                     bool              `json:"disabled"`
	InvertHealthcheck            bool              `json:"invert_healthcheck"`
	MeasureLatency               bool              `json:"measure_latency"`
	Port                         int               `json:"port"`
	ResourcePath                 string            `json:"resource_path"`
	Tags                         map[string]string `json:"tags"`
	CloudwatchAlarmName          string            `json:"cloudwatch_alarm_name"`
	Fqdn                         string            `json:"fqdn"`
	ReferenceName                string            `json:"reference_name"`
	Regions                      []string          `json:"regions"`
	ChildHealthchecks            []string          `json:"child_healthchecks"`
	FailureThreshold             int               `json:"failure_threshold"`
}

// A Route53HealthCheckStatus defines the observed state of a Route53HealthCheck
type Route53HealthCheckStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Route53HealthCheckObservation `json:",inline"`
}

// A Route53HealthCheckObservation records the observed state of a Route53HealthCheck
type Route53HealthCheckObservation struct {
	Id        string `json:"id"`
	EnableSni bool   `json:"enable_sni"`
}