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

// GlobalacceleratorEndpointGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlobalacceleratorEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalacceleratorEndpointGroupSpec   `json:"spec"`
	Status GlobalacceleratorEndpointGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalacceleratorEndpointGroup contains a list of GlobalacceleratorEndpointGroupList
type GlobalacceleratorEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalacceleratorEndpointGroup `json:"items"`
}

// A GlobalacceleratorEndpointGroupSpec defines the desired state of a GlobalacceleratorEndpointGroup
type GlobalacceleratorEndpointGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlobalacceleratorEndpointGroupParameters `json:",inline"`
}

// A GlobalacceleratorEndpointGroupParameters defines the desired state of a GlobalacceleratorEndpointGroup
type GlobalacceleratorEndpointGroupParameters struct {
	HealthCheckProtocol        string                  `json:"health_check_protocol"`
	ListenerArn                string                  `json:"listener_arn"`
	EndpointGroupRegion        string                  `json:"endpoint_group_region"`
	HealthCheckIntervalSeconds int64                   `json:"health_check_interval_seconds"`
	HealthCheckPath            string                  `json:"health_check_path"`
	HealthCheckPort            int64                   `json:"health_check_port"`
	Id                         string                  `json:"id"`
	ThresholdCount             int64                   `json:"threshold_count"`
	TrafficDialPercentage      int64                   `json:"traffic_dial_percentage"`
	EndpointConfiguration      []EndpointConfiguration `json:"endpoint_configuration"`
}

type EndpointConfiguration struct {
	EndpointId                  string `json:"endpoint_id"`
	Weight                      int64  `json:"weight"`
	ClientIpPreservationEnabled bool   `json:"client_ip_preservation_enabled"`
}

// A GlobalacceleratorEndpointGroupStatus defines the observed state of a GlobalacceleratorEndpointGroup
type GlobalacceleratorEndpointGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlobalacceleratorEndpointGroupObservation `json:",inline"`
}

// A GlobalacceleratorEndpointGroupObservation records the observed state of a GlobalacceleratorEndpointGroup
type GlobalacceleratorEndpointGroupObservation struct{}