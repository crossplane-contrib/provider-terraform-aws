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

// DefaultVpcDhcpOptions is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DefaultVpcDhcpOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DefaultVpcDhcpOptionsSpec   `json:"spec"`
	Status DefaultVpcDhcpOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultVpcDhcpOptions contains a list of DefaultVpcDhcpOptionsList
type DefaultVpcDhcpOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultVpcDhcpOptions `json:"items"`
}

// A DefaultVpcDhcpOptionsSpec defines the desired state of a DefaultVpcDhcpOptions
type DefaultVpcDhcpOptionsSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DefaultVpcDhcpOptionsParameters `json:",inline"`
}

// A DefaultVpcDhcpOptionsParameters defines the desired state of a DefaultVpcDhcpOptions
type DefaultVpcDhcpOptionsParameters struct {
	Id                 string            `json:"id"`
	NetbiosNodeType    string            `json:"netbios_node_type"`
	Tags               map[string]string `json:"tags"`
	NetbiosNameServers []string          `json:"netbios_name_servers"`
}

// A DefaultVpcDhcpOptionsStatus defines the observed state of a DefaultVpcDhcpOptions
type DefaultVpcDhcpOptionsStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DefaultVpcDhcpOptionsObservation `json:",inline"`
}

// A DefaultVpcDhcpOptionsObservation records the observed state of a DefaultVpcDhcpOptions
type DefaultVpcDhcpOptionsObservation struct {
	Arn               string `json:"arn"`
	DomainName        string `json:"domain_name"`
	NtpServers        string `json:"ntp_servers"`
	OwnerId           string `json:"owner_id"`
	DomainNameServers string `json:"domain_name_servers"`
}