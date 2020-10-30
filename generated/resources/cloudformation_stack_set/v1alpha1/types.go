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

// CloudformationStackSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudformationStackSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudformationStackSetSpec   `json:"spec"`
	Status CloudformationStackSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudformationStackSet contains a list of CloudformationStackSetList
type CloudformationStackSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudformationStackSet `json:"items"`
}

// A CloudformationStackSetSpec defines the desired state of a CloudformationStackSet
type CloudformationStackSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudformationStackSetParameters `json:",inline"`
}

// A CloudformationStackSetParameters defines the desired state of a CloudformationStackSet
type CloudformationStackSetParameters struct {
	Description           string   `json:"description"`
	Capabilities          []string `json:"capabilities"`
	ExecutionRoleName     string   `json:"execution_role_name"`
	Name                  string   `json:"name"`
	AdministrationRoleArn string   `json:"administration_role_arn"`
	TemplateUrl           string   `json:"template_url"`
}

// A CloudformationStackSetStatus defines the observed state of a CloudformationStackSet
type CloudformationStackSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudformationStackSetObservation `json:",inline"`
}

// A CloudformationStackSetObservation records the observed state of a CloudformationStackSet
type CloudformationStackSetObservation struct {
	Id           string `json:"id"`
	TemplateBody string `json:"template_body"`
	Arn          string `json:"arn"`
	StackSetId   string `json:"stack_set_id"`
}