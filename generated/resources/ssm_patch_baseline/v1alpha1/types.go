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

// SsmPatchBaseline is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SsmPatchBaseline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SsmPatchBaselineSpec   `json:"spec"`
	Status SsmPatchBaselineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmPatchBaseline contains a list of SsmPatchBaselineList
type SsmPatchBaselineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmPatchBaseline `json:"items"`
}

// A SsmPatchBaselineSpec defines the desired state of a SsmPatchBaseline
type SsmPatchBaselineSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SsmPatchBaselineParameters `json:",inline"`
}

// A SsmPatchBaselineParameters defines the desired state of a SsmPatchBaseline
type SsmPatchBaselineParameters struct {
	RejectedPatches                []string          `json:"rejected_patches"`
	Tags                           map[string]string `json:"tags"`
	ApprovedPatches                []string          `json:"approved_patches"`
	ApprovedPatchesComplianceLevel string            `json:"approved_patches_compliance_level"`
	Description                    string            `json:"description"`
	Id                             string            `json:"id"`
	Name                           string            `json:"name"`
	OperatingSystem                string            `json:"operating_system"`
	GlobalFilter                   []GlobalFilter    `json:"global_filter"`
	ApprovalRule                   ApprovalRule      `json:"approval_rule"`
}

type GlobalFilter struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type ApprovalRule struct {
	ApproveAfterDays  int64         `json:"approve_after_days"`
	ComplianceLevel   string        `json:"compliance_level"`
	EnableNonSecurity bool          `json:"enable_non_security"`
	PatchFilter       []PatchFilter `json:"patch_filter"`
}

type PatchFilter struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// A SsmPatchBaselineStatus defines the observed state of a SsmPatchBaseline
type SsmPatchBaselineStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SsmPatchBaselineObservation `json:",inline"`
}

// A SsmPatchBaselineObservation records the observed state of a SsmPatchBaseline
type SsmPatchBaselineObservation struct{}