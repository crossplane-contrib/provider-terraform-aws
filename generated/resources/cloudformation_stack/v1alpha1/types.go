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

// CloudformationStack is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudformationStack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudformationStackSpec   `json:"spec"`
	Status CloudformationStackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudformationStack contains a list of CloudformationStackList
type CloudformationStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudformationStack `json:"items"`
}

// A CloudformationStackSpec defines the desired state of a CloudformationStack
type CloudformationStackSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudformationStackParameters `json:",inline"`
}

// A CloudformationStackParameters defines the desired state of a CloudformationStack
type CloudformationStackParameters struct {
	Tags             map[string]string `json:"tags"`
	DisableRollback  bool              `json:"disable_rollback"`
	IamRoleArn       string            `json:"iam_role_arn"`
	NotificationArns []string          `json:"notification_arns"`
	OnFailure        string            `json:"on_failure"`
	PolicyUrl        string            `json:"policy_url"`
	TemplateUrl      string            `json:"template_url"`
	TimeoutInMinutes int               `json:"timeout_in_minutes"`
	Name             string            `json:"name"`
	Capabilities     []string          `json:"capabilities"`
	Timeouts         []Timeouts        `json:"timeouts"`
}

type Timeouts struct {
	Update string `json:"update"`
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A CloudformationStackStatus defines the observed state of a CloudformationStack
type CloudformationStackStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudformationStackObservation `json:",inline"`
}

// A CloudformationStackObservation records the observed state of a CloudformationStack
type CloudformationStackObservation struct {
	Id           string            `json:"id"`
	Parameters   map[string]string `json:"parameters"`
	PolicyBody   string            `json:"policy_body"`
	Outputs      map[string]string `json:"outputs"`
	TemplateBody string            `json:"template_body"`
}