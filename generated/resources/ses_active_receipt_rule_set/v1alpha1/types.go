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

// SesActiveReceiptRuleSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SesActiveReceiptRuleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SesActiveReceiptRuleSetSpec   `json:"spec"`
	Status SesActiveReceiptRuleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesActiveReceiptRuleSet contains a list of SesActiveReceiptRuleSetList
type SesActiveReceiptRuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesActiveReceiptRuleSet `json:"items"`
}

// A SesActiveReceiptRuleSetSpec defines the desired state of a SesActiveReceiptRuleSet
type SesActiveReceiptRuleSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SesActiveReceiptRuleSetParameters `json:"forProvider"`
}

// A SesActiveReceiptRuleSetParameters defines the desired state of a SesActiveReceiptRuleSet
type SesActiveReceiptRuleSetParameters struct {
	Id          string `json:"id"`
	RuleSetName string `json:"rule_set_name"`
}

// A SesActiveReceiptRuleSetStatus defines the observed state of a SesActiveReceiptRuleSet
type SesActiveReceiptRuleSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SesActiveReceiptRuleSetObservation `json:"atProvider"`
}

// A SesActiveReceiptRuleSetObservation records the observed state of a SesActiveReceiptRuleSet
type SesActiveReceiptRuleSetObservation struct{}