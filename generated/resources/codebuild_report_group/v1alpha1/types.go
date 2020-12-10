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

// CodebuildReportGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodebuildReportGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodebuildReportGroupSpec   `json:"spec"`
	Status CodebuildReportGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodebuildReportGroup contains a list of CodebuildReportGroupList
type CodebuildReportGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodebuildReportGroup `json:"items"`
}

// A CodebuildReportGroupSpec defines the desired state of a CodebuildReportGroup
type CodebuildReportGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodebuildReportGroupParameters `json:",inline"`
}

// A CodebuildReportGroupParameters defines the desired state of a CodebuildReportGroup
type CodebuildReportGroupParameters struct {
	Id           string            `json:"id"`
	Name         string            `json:"name"`
	Tags         map[string]string `json:"tags"`
	Type         string            `json:"type"`
	ExportConfig ExportConfig      `json:"export_config"`
}

type ExportConfig struct {
	Type          string        `json:"type"`
	S3Destination S3Destination `json:"s3_destination"`
}

type S3Destination struct {
	EncryptionDisabled bool   `json:"encryption_disabled"`
	EncryptionKey      string `json:"encryption_key"`
	Packaging          string `json:"packaging"`
	Path               string `json:"path"`
	Bucket             string `json:"bucket"`
}

// A CodebuildReportGroupStatus defines the observed state of a CodebuildReportGroup
type CodebuildReportGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodebuildReportGroupObservation `json:",inline"`
}

// A CodebuildReportGroupObservation records the observed state of a CodebuildReportGroup
type CodebuildReportGroupObservation struct {
	Arn     string `json:"arn"`
	Created string `json:"created"`
}