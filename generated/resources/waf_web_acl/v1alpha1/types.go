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

// WafWebAcl is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafWebAcl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafWebAclSpec   `json:"spec"`
	Status WafWebAclStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafWebAcl contains a list of WafWebAclList
type WafWebAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafWebAcl `json:"items"`
}

// A WafWebAclSpec defines the desired state of a WafWebAcl
type WafWebAclSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafWebAclParameters `json:"forProvider"`
}

// A WafWebAclParameters defines the desired state of a WafWebAcl
type WafWebAclParameters struct {
	Name                 string               `json:"name"`
	Tags                 map[string]string    `json:"tags"`
	Id                   string               `json:"id"`
	MetricName           string               `json:"metric_name"`
	DefaultAction        DefaultAction        `json:"default_action"`
	LoggingConfiguration LoggingConfiguration `json:"logging_configuration"`
	Rules                Rules                `json:"rules"`
}

type DefaultAction struct {
	Type string `json:"type"`
}

type LoggingConfiguration struct {
	LogDestination string         `json:"log_destination"`
	RedactedFields RedactedFields `json:"redacted_fields"`
}

type RedactedFields struct {
	FieldToMatch []FieldToMatch `json:"field_to_match"`
}

type FieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type Rules struct {
	RuleId         string         `json:"rule_id"`
	Type           string         `json:"type"`
	Priority       int64          `json:"priority"`
	Action         Action         `json:"action"`
	OverrideAction OverrideAction `json:"override_action"`
}

type Action struct {
	Type string `json:"type"`
}

type OverrideAction struct {
	Type string `json:"type"`
}

// A WafWebAclStatus defines the observed state of a WafWebAcl
type WafWebAclStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafWebAclObservation `json:"atProvider"`
}

// A WafWebAclObservation records the observed state of a WafWebAcl
type WafWebAclObservation struct {
	Arn string `json:"arn"`
}