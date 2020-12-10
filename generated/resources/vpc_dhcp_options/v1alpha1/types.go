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

// VpcDhcpOptions is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpcDhcpOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpcDhcpOptionsSpec   `json:"spec"`
	Status VpcDhcpOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcDhcpOptions contains a list of VpcDhcpOptionsList
type VpcDhcpOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcDhcpOptions `json:"items"`
}

// A VpcDhcpOptionsSpec defines the desired state of a VpcDhcpOptions
type VpcDhcpOptionsSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpcDhcpOptionsParameters `json:",inline"`
}

// A VpcDhcpOptionsParameters defines the desired state of a VpcDhcpOptions
type VpcDhcpOptionsParameters struct {
	NetbiosNodeType    string            `json:"netbios_node_type"`
	DomainNameServers  []string          `json:"domain_name_servers"`
	Id                 string            `json:"id"`
	NetbiosNameServers []string          `json:"netbios_name_servers"`
	NtpServers         []string          `json:"ntp_servers"`
	Tags               map[string]string `json:"tags"`
	DomainName         string            `json:"domain_name"`
}

// A VpcDhcpOptionsStatus defines the observed state of a VpcDhcpOptions
type VpcDhcpOptionsStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpcDhcpOptionsObservation `json:",inline"`
}

// A VpcDhcpOptionsObservation records the observed state of a VpcDhcpOptions
type VpcDhcpOptionsObservation struct {
	OwnerId string `json:"owner_id"`
	Arn     string `json:"arn"`
}