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

// NetworkInterface is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type NetworkInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkInterfaceSpec   `json:"spec"`
	Status NetworkInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterface contains a list of NetworkInterfaceList
type NetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterface `json:"items"`
}

// A NetworkInterfaceSpec defines the desired state of a NetworkInterface
type NetworkInterfaceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  NetworkInterfaceParameters `json:"forProvider"`
}

// A NetworkInterfaceParameters defines the desired state of a NetworkInterface
type NetworkInterfaceParameters struct {
	Tags            map[string]string `json:"tags"`
	Id              string            `json:"id"`
	PrivateIps      []string          `json:"private_ips"`
	SecurityGroups  []string          `json:"security_groups"`
	SourceDestCheck bool              `json:"source_dest_check"`
	Description     string            `json:"description"`
	PrivateIp       string            `json:"private_ip"`
	PrivateIpsCount int64             `json:"private_ips_count"`
	SubnetId        string            `json:"subnet_id"`
	Attachment      Attachment        `json:"attachment"`
}

type Attachment struct {
	Instance     string `json:"instance"`
	AttachmentId string `json:"attachment_id"`
	DeviceIndex  int64  `json:"device_index"`
}

// A NetworkInterfaceStatus defines the observed state of a NetworkInterface
type NetworkInterfaceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     NetworkInterfaceObservation `json:"atProvider"`
}

// A NetworkInterfaceObservation records the observed state of a NetworkInterface
type NetworkInterfaceObservation struct {
	MacAddress     string `json:"mac_address"`
	OutpostArn     string `json:"outpost_arn"`
	PrivateDnsName string `json:"private_dns_name"`
}