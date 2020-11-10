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

// ServiceDiscoveryService is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ServiceDiscoveryService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceDiscoveryServiceSpec   `json:"spec"`
	Status ServiceDiscoveryServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceDiscoveryService contains a list of ServiceDiscoveryServiceList
type ServiceDiscoveryServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDiscoveryService `json:"items"`
}

// A ServiceDiscoveryServiceSpec defines the desired state of a ServiceDiscoveryService
type ServiceDiscoveryServiceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ServiceDiscoveryServiceParameters `json:",inline"`
}

// A ServiceDiscoveryServiceParameters defines the desired state of a ServiceDiscoveryService
type ServiceDiscoveryServiceParameters struct {
	Description             string                  `json:"description"`
	Name                    string                  `json:"name"`
	Tags                    map[string]string       `json:"tags"`
	HealthCheckConfig       HealthCheckConfig       `json:"health_check_config"`
	HealthCheckCustomConfig HealthCheckCustomConfig `json:"health_check_custom_config"`
	DnsConfig               DnsConfig               `json:"dns_config"`
}

type HealthCheckConfig struct {
	FailureThreshold int    `json:"failure_threshold"`
	ResourcePath     string `json:"resource_path"`
	Type             string `json:"type"`
}

type HealthCheckCustomConfig struct {
	FailureThreshold int `json:"failure_threshold"`
}

type DnsConfig struct {
	NamespaceId   string       `json:"namespace_id"`
	RoutingPolicy string       `json:"routing_policy"`
	DnsRecords    []DnsRecords `json:"dns_records"`
}

type DnsRecords struct {
	Type string `json:"type"`
	Ttl  int    `json:"ttl"`
}

// A ServiceDiscoveryServiceStatus defines the observed state of a ServiceDiscoveryService
type ServiceDiscoveryServiceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ServiceDiscoveryServiceObservation `json:",inline"`
}

// A ServiceDiscoveryServiceObservation records the observed state of a ServiceDiscoveryService
type ServiceDiscoveryServiceObservation struct {
	Id          string `json:"id"`
	NamespaceId string `json:"namespace_id"`
	Arn         string `json:"arn"`
}