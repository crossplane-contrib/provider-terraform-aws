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

// CurReportDefinition is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CurReportDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CurReportDefinitionSpec   `json:"spec"`
	Status CurReportDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CurReportDefinition contains a list of CurReportDefinitionList
type CurReportDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CurReportDefinition `json:"items"`
}

// A CurReportDefinitionSpec defines the desired state of a CurReportDefinition
type CurReportDefinitionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CurReportDefinitionParameters `json:",inline"`
}

// A CurReportDefinitionParameters defines the desired state of a CurReportDefinition
type CurReportDefinitionParameters struct {
	RefreshClosedReports     bool     `json:"refresh_closed_reports"`
	ReportVersioning         string   `json:"report_versioning"`
	TimeUnit                 string   `json:"time_unit"`
	Compression              string   `json:"compression"`
	Format                   string   `json:"format"`
	Id                       string   `json:"id"`
	ReportName               string   `json:"report_name"`
	S3Bucket                 string   `json:"s3_bucket"`
	S3Prefix                 string   `json:"s3_prefix"`
	S3Region                 string   `json:"s3_region"`
	AdditionalArtifacts      []string `json:"additional_artifacts"`
	AdditionalSchemaElements []string `json:"additional_schema_elements"`
}

// A CurReportDefinitionStatus defines the observed state of a CurReportDefinition
type CurReportDefinitionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CurReportDefinitionObservation `json:",inline"`
}

// A CurReportDefinitionObservation records the observed state of a CurReportDefinition
type CurReportDefinitionObservation struct{}