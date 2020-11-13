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
	ServiceType   string            `json:"service_type"`
	Version       int               `json:"version"`
	HttpMethod    string            `json:"http_method"`
	ResourceArn   string            `json:"resource_arn"`
	Id            string            `json:"id"`
	Host          string            `json:"host"`
	Priority      int               `json:"priority"`
	ReservoirSize int               `json:"reservoir_size"`
	RuleName      string            `json:"rule_name"`
	ServiceName   string            `json:"service_name"`
	Tags          map[string]string `json:"tags"`
	Attributes    map[string]string `json:"attributes"`
	FixedRate     int               `json:"fixed_rate"`
	UrlPath       string            `json:"url_path"`
}

// A XraySamplingRuleStatus defines the observed state of a XraySamplingRule
type XraySamplingRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     XraySamplingRuleObservation `json:",inline"`
}

// A XraySamplingRuleObservation records the observed state of a XraySamplingRule
type XraySamplingRuleObservation struct {
	Arn string `json:"arn"`
}