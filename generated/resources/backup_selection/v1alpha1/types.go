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

// BackupSelection is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type BackupSelection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupSelectionSpec   `json:"spec"`
	Status BackupSelectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupSelection contains a list of BackupSelectionList
type BackupSelectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupSelection `json:"items"`
}

// A BackupSelectionSpec defines the desired state of a BackupSelection
type BackupSelectionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  BackupSelectionParameters `json:"forProvider"`
}

// A BackupSelectionParameters defines the desired state of a BackupSelection
type BackupSelectionParameters struct {
	IamRoleArn   string       `json:"iam_role_arn"`
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	PlanId       string       `json:"plan_id"`
	Resources    []string     `json:"resources"`
	SelectionTag SelectionTag `json:"selection_tag"`
}

type SelectionTag struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// A BackupSelectionStatus defines the observed state of a BackupSelection
type BackupSelectionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     BackupSelectionObservation `json:"atProvider"`
}

// A BackupSelectionObservation records the observed state of a BackupSelection
type BackupSelectionObservation struct{}