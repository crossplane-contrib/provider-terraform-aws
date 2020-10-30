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

// Lb is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Lb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LbSpec   `json:"spec"`
	Status LbStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Lb contains a list of LbList
type LbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Lb `json:"items"`
}

// A LbSpec defines the desired state of a Lb
type LbSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LbParameters `json:",inline"`
}

// A LbParameters defines the desired state of a Lb
type LbParameters struct {
	EnableHttp2                  bool   `json:"enable_http2"`
	DropInvalidHeaderFields      bool   `json:"drop_invalid_header_fields"`
	IdleTimeout                  int    `json:"idle_timeout"`
	EnableCrossZoneLoadBalancing bool   `json:"enable_cross_zone_load_balancing"`
	EnableDeletionProtection     bool   `json:"enable_deletion_protection"`
	NamePrefix                   string `json:"name_prefix"`
	CustomerOwnedIpv4Pool        string `json:"customer_owned_ipv4_pool"`
	LoadBalancerType             string `json:"load_balancer_type"`
}

// A LbStatus defines the observed state of a Lb
type LbStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LbObservation `json:",inline"`
}

// A LbObservation records the observed state of a Lb
type LbObservation struct {
	DnsName        string   `json:"dns_name"`
	IpAddressType  string   `json:"ip_address_type"`
	Name           string   `json:"name"`
	VpcId          string   `json:"vpc_id"`
	Arn            string   `json:"arn"`
	ArnSuffix      string   `json:"arn_suffix"`
	Id             string   `json:"id"`
	ZoneId         string   `json:"zone_id"`
	Internal       bool     `json:"internal"`
	Subnets        []string `json:"subnets"`
	SecurityGroups []string `json:"security_groups"`
}