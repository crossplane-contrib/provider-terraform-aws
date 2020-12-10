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

// VpcEndpointService is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpcEndpointService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpcEndpointServiceSpec   `json:"spec"`
	Status VpcEndpointServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcEndpointService contains a list of VpcEndpointServiceList
type VpcEndpointServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcEndpointService `json:"items"`
}

// A VpcEndpointServiceSpec defines the desired state of a VpcEndpointService
type VpcEndpointServiceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpcEndpointServiceParameters `json:",inline"`
}

// A VpcEndpointServiceParameters defines the desired state of a VpcEndpointService
type VpcEndpointServiceParameters struct {
	AllowedPrincipals       []string          `json:"allowed_principals"`
	Id                      string            `json:"id"`
	NetworkLoadBalancerArns []string          `json:"network_load_balancer_arns"`
	AcceptanceRequired      bool              `json:"acceptance_required"`
	Tags                    map[string]string `json:"tags"`
}

// A VpcEndpointServiceStatus defines the observed state of a VpcEndpointService
type VpcEndpointServiceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpcEndpointServiceObservation `json:",inline"`
}

// A VpcEndpointServiceObservation records the observed state of a VpcEndpointService
type VpcEndpointServiceObservation struct {
	ManagesVpcEndpoints  bool     `json:"manages_vpc_endpoints"`
	PrivateDnsName       string   `json:"private_dns_name"`
	ServiceName          string   `json:"service_name"`
	BaseEndpointDnsNames []string `json:"base_endpoint_dns_names"`
	Arn                  string   `json:"arn"`
	AvailabilityZones    []string `json:"availability_zones"`
	ServiceType          string   `json:"service_type"`
	State                string   `json:"state"`
}