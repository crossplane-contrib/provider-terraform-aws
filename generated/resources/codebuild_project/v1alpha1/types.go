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

// CodebuildProject is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodebuildProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodebuildProjectSpec   `json:"spec"`
	Status CodebuildProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodebuildProject contains a list of CodebuildProjectList
type CodebuildProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodebuildProject `json:"items"`
}

// A CodebuildProjectSpec defines the desired state of a CodebuildProject
type CodebuildProjectSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodebuildProjectParameters `json:",inline"`
}

// A CodebuildProjectParameters defines the desired state of a CodebuildProject
type CodebuildProjectParameters struct {
	BadgeEnabled  bool   `json:"badge_enabled"`
	Name          string `json:"name"`
	SourceVersion string `json:"source_version"`
	BuildTimeout  int    `json:"build_timeout"`
	QueuedTimeout int    `json:"queued_timeout"`
	ServiceRole   string `json:"service_role"`
}

// A CodebuildProjectStatus defines the observed state of a CodebuildProject
type CodebuildProjectStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodebuildProjectObservation `json:",inline"`
}

// A CodebuildProjectObservation records the observed state of a CodebuildProject
type CodebuildProjectObservation struct {
	Description   string `json:"description"`
	EncryptionKey string `json:"encryption_key"`
	Id            string `json:"id"`
	Arn           string `json:"arn"`
	BadgeUrl      string `json:"badge_url"`
}