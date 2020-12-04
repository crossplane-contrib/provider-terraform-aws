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

// Route53Record is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Route53Record struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Route53RecordSpec   `json:"spec"`
	Status Route53RecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53Record contains a list of Route53RecordList
type Route53RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53Record `json:"items"`
}

// A Route53RecordSpec defines the desired state of a Route53Record
type Route53RecordSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Route53RecordParameters `json:",inline"`
}

// A Route53RecordParameters defines the desired state of a Route53Record
type Route53RecordParameters struct {
	HealthCheckId                 string                   `json:"health_check_id"`
	Id                            string                   `json:"id"`
	MultivalueAnswerRoutingPolicy bool                     `json:"multivalue_answer_routing_policy"`
	Name                          string                   `json:"name"`
	SetIdentifier                 string                   `json:"set_identifier"`
	ZoneId                        string                   `json:"zone_id"`
	AllowOverwrite                bool                     `json:"allow_overwrite"`
	Records                       []string                 `json:"records"`
	Ttl                           int64                    `json:"ttl"`
	Type                          string                   `json:"type"`
	Alias                         Alias                    `json:"alias"`
	FailoverRoutingPolicy         FailoverRoutingPolicy    `json:"failover_routing_policy"`
	GeolocationRoutingPolicy      GeolocationRoutingPolicy `json:"geolocation_routing_policy"`
	LatencyRoutingPolicy          LatencyRoutingPolicy     `json:"latency_routing_policy"`
	WeightedRoutingPolicy         WeightedRoutingPolicy    `json:"weighted_routing_policy"`
}

type Alias struct {
	EvaluateTargetHealth bool   `json:"evaluate_target_health"`
	Name                 string `json:"name"`
	ZoneId               string `json:"zone_id"`
}

type FailoverRoutingPolicy struct {
	Type string `json:"type"`
}

type GeolocationRoutingPolicy struct {
	Continent   string `json:"continent"`
	Country     string `json:"country"`
	Subdivision string `json:"subdivision"`
}

type LatencyRoutingPolicy struct {
	Region string `json:"region"`
}

type WeightedRoutingPolicy struct {
	Weight int64 `json:"weight"`
}

// A Route53RecordStatus defines the observed state of a Route53Record
type Route53RecordStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Route53RecordObservation `json:",inline"`
}

// A Route53RecordObservation records the observed state of a Route53Record
type Route53RecordObservation struct {
	Fqdn string `json:"fqdn"`
}