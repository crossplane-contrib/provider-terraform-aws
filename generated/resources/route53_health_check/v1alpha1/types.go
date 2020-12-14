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
	ForProvider                  Route53HealthCheckParameters `json:"forProvider"`
}

// A Route53HealthCheckParameters defines the desired state of a Route53HealthCheck
type Route53HealthCheckParameters struct {
	MeasureLatency               bool              `json:"measure_latency"`
	ResourcePath                 string            `json:"resource_path"`
	Tags                         map[string]string `json:"tags"`
	Type                         string            `json:"type"`
	ChildHealthchecks            []string          `json:"child_healthchecks"`
	InvertHealthcheck            bool              `json:"invert_healthcheck"`
	Regions                      []string          `json:"regions"`
	SearchString                 string            `json:"search_string"`
	ChildHealthThreshold         int64             `json:"child_health_threshold"`
	EnableSni                    bool              `json:"enable_sni"`
	FailureThreshold             int64             `json:"failure_threshold"`
	InsufficientDataHealthStatus string            `json:"insufficient_data_health_status"`
	Port                         int64             `json:"port"`
	ReferenceName                string            `json:"reference_name"`
	RequestInterval              int64             `json:"request_interval"`
	CloudwatchAlarmName          string            `json:"cloudwatch_alarm_name"`
	Disabled                     bool              `json:"disabled"`
	Id                           string            `json:"id"`
	IpAddress                    string            `json:"ip_address"`
	CloudwatchAlarmRegion        string            `json:"cloudwatch_alarm_region"`
	Fqdn                         string            `json:"fqdn"`
}

// A Route53HealthCheckStatus defines the observed state of a Route53HealthCheck
type Route53HealthCheckStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Route53HealthCheckObservation `json:"atProvider"`
}

// A Route53HealthCheckObservation records the observed state of a Route53HealthCheck
type Route53HealthCheckObservation struct{}