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

// GlobalacceleratorAccelerator is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlobalacceleratorAccelerator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalacceleratorAcceleratorSpec   `json:"spec"`
	Status GlobalacceleratorAcceleratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalacceleratorAccelerator contains a list of GlobalacceleratorAcceleratorList
type GlobalacceleratorAcceleratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalacceleratorAccelerator `json:"items"`
}

// A GlobalacceleratorAcceleratorSpec defines the desired state of a GlobalacceleratorAccelerator
type GlobalacceleratorAcceleratorSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlobalacceleratorAcceleratorParameters `json:"forProvider"`
}

// A GlobalacceleratorAcceleratorParameters defines the desired state of a GlobalacceleratorAccelerator
type GlobalacceleratorAcceleratorParameters struct {
	Name          string            `json:"name"`
	Tags          map[string]string `json:"tags"`
	Enabled       bool              `json:"enabled"`
	Id            string            `json:"id"`
	IpAddressType string            `json:"ip_address_type"`
	Attributes    Attributes        `json:"attributes"`
}

type Attributes struct {
	FlowLogsS3Prefix string `json:"flow_logs_s3_prefix"`
	FlowLogsEnabled  bool   `json:"flow_logs_enabled"`
	FlowLogsS3Bucket string `json:"flow_logs_s3_bucket"`
}

// A GlobalacceleratorAcceleratorStatus defines the observed state of a GlobalacceleratorAccelerator
type GlobalacceleratorAcceleratorStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlobalacceleratorAcceleratorObservation `json:"atProvider"`
}

// A GlobalacceleratorAcceleratorObservation records the observed state of a GlobalacceleratorAccelerator
type GlobalacceleratorAcceleratorObservation struct {
	IpSets       []IpSets `json:"ip_sets"`
	DnsName      string   `json:"dns_name"`
	HostedZoneId string   `json:"hosted_zone_id"`
}

type IpSets struct {
	IpAddresses []string `json:"ip_addresses"`
	IpFamily    string   `json:"ip_family"`
}