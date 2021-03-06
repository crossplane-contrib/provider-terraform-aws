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

// SesTemplate is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SesTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SesTemplateSpec   `json:"spec"`
	Status SesTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesTemplate contains a list of SesTemplateList
type SesTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesTemplate `json:"items"`
}

// A SesTemplateSpec defines the desired state of a SesTemplate
type SesTemplateSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SesTemplateParameters `json:"forProvider"`
}

// A SesTemplateParameters defines the desired state of a SesTemplate
type SesTemplateParameters struct {
	Html    string `json:"html"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}

// A SesTemplateStatus defines the observed state of a SesTemplate
type SesTemplateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SesTemplateObservation `json:"atProvider"`
}

// A SesTemplateObservation records the observed state of a SesTemplate
type SesTemplateObservation struct{}