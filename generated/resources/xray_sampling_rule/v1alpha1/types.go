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

// XraySamplingRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type XraySamplingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   XraySamplingRuleSpec   `json:"spec"`
	Status XraySamplingRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// XraySamplingRule contains a list of XraySamplingRuleList
type XraySamplingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []XraySamplingRule `json:"items"`
}

// A XraySamplingRuleSpec defines the desired state of a XraySamplingRule
type XraySamplingRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  XraySamplingRuleParameters `json:",inline"`
}

// A XraySamplingRuleParameters defines the desired state of a XraySamplingRule
type XraySamplingRuleParameters struct {
	RuleName      string            `json:"rule_name"`
	Version       int               `json:"version"`
	Tags          map[string]string `json:"tags"`
	UrlPath       string            `json:"url_path"`
	FixedRate     int               `json:"fixed_rate"`
	Host          string            `json:"host"`
	ReservoirSize int               `json:"reservoir_size"`
	ResourceArn   string            `json:"resource_arn"`
	ServiceName   string            `json:"service_name"`
	ServiceType   string            `json:"service_type"`
	Attributes    map[string]string `json:"attributes"`
	HttpMethod    string            `json:"http_method"`
	Priority      int               `json:"priority"`
}

// A XraySamplingRuleStatus defines the observed state of a XraySamplingRule
type XraySamplingRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     XraySamplingRuleObservation `json:",inline"`
}

// A XraySamplingRuleObservation records the observed state of a XraySamplingRule
type XraySamplingRuleObservation struct {
	Id  string `json:"id"`
	Arn string `json:"arn"`
}