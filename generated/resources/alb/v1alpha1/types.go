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

// Alb is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Alb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlbSpec   `json:"spec"`
	Status AlbStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Alb contains a list of AlbList
type AlbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alb `json:"items"`
}

// A AlbSpec defines the desired state of a Alb
type AlbSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AlbParameters `json:"forProvider"`
}

// A AlbParameters defines the desired state of a Alb
type AlbParameters struct {
	Id                           string            `json:"id"`
	IdleTimeout                  int64             `json:"idle_timeout"`
	Internal                     bool              `json:"internal"`
	IpAddressType                string            `json:"ip_address_type"`
	Name                         string            `json:"name"`
	NamePrefix                   string            `json:"name_prefix"`
	EnableCrossZoneLoadBalancing bool              `json:"enable_cross_zone_load_balancing"`
	EnableDeletionProtection     bool              `json:"enable_deletion_protection"`
	SecurityGroups               []string          `json:"security_groups"`
	Subnets                      []string          `json:"subnets"`
	Tags                         map[string]string `json:"tags"`
	CustomerOwnedIpv4Pool        string            `json:"customer_owned_ipv4_pool"`
	DropInvalidHeaderFields      bool              `json:"drop_invalid_header_fields"`
	EnableHttp2                  bool              `json:"enable_http2"`
	LoadBalancerType             string            `json:"load_balancer_type"`
	AccessLogs                   AccessLogs        `json:"access_logs"`
	SubnetMapping                SubnetMapping     `json:"subnet_mapping"`
	Timeouts                     Timeouts          `json:"timeouts"`
}

type AccessLogs struct {
	Enabled bool   `json:"enabled"`
	Prefix  string `json:"prefix"`
	Bucket  string `json:"bucket"`
}

type SubnetMapping struct {
	AllocationId       string `json:"allocation_id"`
	OutpostId          string `json:"outpost_id"`
	PrivateIpv4Address string `json:"private_ipv4_address"`
	SubnetId           string `json:"subnet_id"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A AlbStatus defines the observed state of a Alb
type AlbStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AlbObservation `json:"atProvider"`
}

// A AlbObservation records the observed state of a Alb
type AlbObservation struct {
	DnsName   string `json:"dns_name"`
	ZoneId    string `json:"zone_id"`
	Arn       string `json:"arn"`
	ArnSuffix string `json:"arn_suffix"`
	VpcId     string `json:"vpc_id"`
}