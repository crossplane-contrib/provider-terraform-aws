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

// InspectorAssessmentTemplate is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type InspectorAssessmentTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InspectorAssessmentTemplateSpec   `json:"spec"`
	Status InspectorAssessmentTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InspectorAssessmentTemplate contains a list of InspectorAssessmentTemplateList
type InspectorAssessmentTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InspectorAssessmentTemplate `json:"items"`
}

// A InspectorAssessmentTemplateSpec defines the desired state of a InspectorAssessmentTemplate
type InspectorAssessmentTemplateSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  InspectorAssessmentTemplateParameters `json:",inline"`
}

// A InspectorAssessmentTemplateParameters defines the desired state of a InspectorAssessmentTemplate
type InspectorAssessmentTemplateParameters struct {
	TargetArn        string            `json:"target_arn"`
	Duration         int               `json:"duration"`
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	RulesPackageArns []string          `json:"rules_package_arns"`
	Tags             map[string]string `json:"tags"`
}

// A InspectorAssessmentTemplateStatus defines the observed state of a InspectorAssessmentTemplate
type InspectorAssessmentTemplateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     InspectorAssessmentTemplateObservation `json:",inline"`
}

// A InspectorAssessmentTemplateObservation records the observed state of a InspectorAssessmentTemplate
type InspectorAssessmentTemplateObservation struct {
	Arn string `json:"arn"`
}