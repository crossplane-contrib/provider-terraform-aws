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

// SsmAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SsmAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SsmAssociationSpec   `json:"spec"`
	Status SsmAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmAssociation contains a list of SsmAssociationList
type SsmAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmAssociation `json:"items"`
}

// A SsmAssociationSpec defines the desired state of a SsmAssociation
type SsmAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SsmAssociationParameters `json:",inline"`
}

// A SsmAssociationParameters defines the desired state of a SsmAssociation
type SsmAssociationParameters struct {
	AssociationName               string            `json:"association_name"`
	AutomationTargetParameterName string            `json:"automation_target_parameter_name"`
	ComplianceSeverity            string            `json:"compliance_severity"`
	DocumentVersion               string            `json:"document_version"`
	Id                            string            `json:"id"`
	InstanceId                    string            `json:"instance_id"`
	Name                          string            `json:"name"`
	Parameters                    map[string]string `json:"parameters"`
	ScheduleExpression            string            `json:"schedule_expression"`
	MaxConcurrency                string            `json:"max_concurrency"`
	MaxErrors                     string            `json:"max_errors"`
	OutputLocation                OutputLocation    `json:"output_location"`
	Targets                       []Targets         `json:"targets"`
}

type OutputLocation struct {
	S3KeyPrefix  string `json:"s3_key_prefix"`
	S3BucketName string `json:"s3_bucket_name"`
}

type Targets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// A SsmAssociationStatus defines the observed state of a SsmAssociation
type SsmAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SsmAssociationObservation `json:",inline"`
}

// A SsmAssociationObservation records the observed state of a SsmAssociation
type SsmAssociationObservation struct {
	AssociationId string `json:"association_id"`
}