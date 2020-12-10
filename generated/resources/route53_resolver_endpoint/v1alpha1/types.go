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

// Route53ResolverEndpoint is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Route53ResolverEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Route53ResolverEndpointSpec   `json:"spec"`
	Status Route53ResolverEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverEndpoint contains a list of Route53ResolverEndpointList
type Route53ResolverEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ResolverEndpoint `json:"items"`
}

// A Route53ResolverEndpointSpec defines the desired state of a Route53ResolverEndpoint
type Route53ResolverEndpointSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Route53ResolverEndpointParameters `json:",inline"`
}

// A Route53ResolverEndpointParameters defines the desired state of a Route53ResolverEndpoint
type Route53ResolverEndpointParameters struct {
	Tags             map[string]string `json:"tags"`
	Direction        string            `json:"direction"`
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	SecurityGroupIds []string          `json:"security_group_ids"`
	IpAddress        []IpAddress       `json:"ip_address"`
	Timeouts         Timeouts          `json:"timeouts"`
}

type IpAddress struct {
	IpId     string `json:"ip_id"`
	SubnetId string `json:"subnet_id"`
	Ip       string `json:"ip"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A Route53ResolverEndpointStatus defines the observed state of a Route53ResolverEndpoint
type Route53ResolverEndpointStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Route53ResolverEndpointObservation `json:",inline"`
}

// A Route53ResolverEndpointObservation records the observed state of a Route53ResolverEndpoint
type Route53ResolverEndpointObservation struct {
	Arn       string `json:"arn"`
	HostVpcId string `json:"host_vpc_id"`
}