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

// SesReceiptRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SesReceiptRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SesReceiptRuleSpec   `json:"spec"`
	Status SesReceiptRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesReceiptRule contains a list of SesReceiptRuleList
type SesReceiptRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesReceiptRule `json:"items"`
}

// A SesReceiptRuleSpec defines the desired state of a SesReceiptRule
type SesReceiptRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SesReceiptRuleParameters `json:",inline"`
}

// A SesReceiptRuleParameters defines the desired state of a SesReceiptRule
type SesReceiptRuleParameters struct {
	After       string   `json:"after"`
	Name        string   `json:"name"`
	Recipients  []string `json:"recipients"`
	RuleSetName string   `json:"rule_set_name"`
}

// A SesReceiptRuleStatus defines the observed state of a SesReceiptRule
type SesReceiptRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SesReceiptRuleObservation `json:",inline"`
}

// A SesReceiptRuleObservation records the observed state of a SesReceiptRule
type SesReceiptRuleObservation struct {
	TlsPolicy   string `json:"tls_policy"`
	Enabled     bool   `json:"enabled"`
	Id          string `json:"id"`
	ScanEnabled bool   `json:"scan_enabled"`
}