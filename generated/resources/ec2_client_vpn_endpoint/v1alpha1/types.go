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

// Ec2ClientVpnEndpoint is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2ClientVpnEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2ClientVpnEndpointSpec   `json:"spec"`
	Status Ec2ClientVpnEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2ClientVpnEndpoint contains a list of Ec2ClientVpnEndpointList
type Ec2ClientVpnEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2ClientVpnEndpoint `json:"items"`
}

// A Ec2ClientVpnEndpointSpec defines the desired state of a Ec2ClientVpnEndpoint
type Ec2ClientVpnEndpointSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2ClientVpnEndpointParameters `json:",inline"`
}

// A Ec2ClientVpnEndpointParameters defines the desired state of a Ec2ClientVpnEndpoint
type Ec2ClientVpnEndpointParameters struct {
	Id                    string                  `json:"id"`
	SplitTunnel           bool                    `json:"split_tunnel"`
	Tags                  map[string]string       `json:"tags"`
	TransportProtocol     string                  `json:"transport_protocol"`
	ClientCidrBlock       string                  `json:"client_cidr_block"`
	Description           string                  `json:"description"`
	DnsServers            []string                `json:"dns_servers"`
	ServerCertificateArn  string                  `json:"server_certificate_arn"`
	AuthenticationOptions []AuthenticationOptions `json:"authentication_options"`
	ConnectionLogOptions  ConnectionLogOptions    `json:"connection_log_options"`
}

type AuthenticationOptions struct {
	ActiveDirectoryId       string `json:"active_directory_id"`
	RootCertificateChainArn string `json:"root_certificate_chain_arn"`
	SamlProviderArn         string `json:"saml_provider_arn"`
	Type                    string `json:"type"`
}

type ConnectionLogOptions struct {
	Enabled             bool   `json:"enabled"`
	CloudwatchLogGroup  string `json:"cloudwatch_log_group"`
	CloudwatchLogStream string `json:"cloudwatch_log_stream"`
}

// A Ec2ClientVpnEndpointStatus defines the observed state of a Ec2ClientVpnEndpoint
type Ec2ClientVpnEndpointStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2ClientVpnEndpointObservation `json:",inline"`
}

// A Ec2ClientVpnEndpointObservation records the observed state of a Ec2ClientVpnEndpoint
type Ec2ClientVpnEndpointObservation struct {
	Arn     string `json:"arn"`
	DnsName string `json:"dns_name"`
	Status  string `json:"status"`
}