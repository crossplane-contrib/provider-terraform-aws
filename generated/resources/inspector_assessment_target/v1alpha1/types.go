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

// InspectorAssessmentTarget is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type InspectorAssessmentTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InspectorAssessmentTargetSpec   `json:"spec"`
	Status InspectorAssessmentTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InspectorAssessmentTarget contains a list of InspectorAssessmentTargetList
type InspectorAssessmentTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InspectorAssessmentTarget `json:"items"`
}

// A InspectorAssessmentTargetSpec defines the desired state of a InspectorAssessmentTarget
type InspectorAssessmentTargetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  InspectorAssessmentTargetParameters `json:"forProvider"`
}

// A InspectorAssessmentTargetParameters defines the desired state of a InspectorAssessmentTarget
type InspectorAssessmentTargetParameters struct {
	Name             string `json:"name"`
	ResourceGroupArn string `json:"resource_group_arn"`
}

// A InspectorAssessmentTargetStatus defines the observed state of a InspectorAssessmentTarget
type InspectorAssessmentTargetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     InspectorAssessmentTargetObservation `json:"atProvider"`
}

// A InspectorAssessmentTargetObservation records the observed state of a InspectorAssessmentTarget
type InspectorAssessmentTargetObservation struct {
	Arn string `json:"arn"`
}