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

// SsmDocument is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SsmDocument struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SsmDocumentSpec   `json:"spec"`
	Status SsmDocumentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmDocument contains a list of SsmDocumentList
type SsmDocumentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmDocument `json:"items"`
}

// A SsmDocumentSpec defines the desired state of a SsmDocument
type SsmDocumentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SsmDocumentParameters `json:"forProvider"`
}

// A SsmDocumentParameters defines the desired state of a SsmDocument
type SsmDocumentParameters struct {
	Id                string            `json:"id"`
	Content           string            `json:"content"`
	DocumentFormat    string            `json:"document_format"`
	Permissions       map[string]string `json:"permissions"`
	Name              string            `json:"name"`
	DocumentType      string            `json:"document_type"`
	Tags              map[string]string `json:"tags"`
	TargetType        string            `json:"target_type"`
	AttachmentsSource AttachmentsSource `json:"attachments_source"`
}

type AttachmentsSource struct {
	Key    string   `json:"key"`
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

// A SsmDocumentStatus defines the observed state of a SsmDocument
type SsmDocumentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SsmDocumentObservation `json:"atProvider"`
}

// A SsmDocumentObservation records the observed state of a SsmDocument
type SsmDocumentObservation struct {
	LatestVersion   string      `json:"latest_version"`
	CreatedDate     string      `json:"created_date"`
	Description     string      `json:"description"`
	Parameter       []Parameter `json:"parameter"`
	Status          string      `json:"status"`
	HashType        string      `json:"hash_type"`
	Owner           string      `json:"owner"`
	Arn             string      `json:"arn"`
	DefaultVersion  string      `json:"default_version"`
	DocumentVersion string      `json:"document_version"`
	Hash            string      `json:"hash"`
	PlatformTypes   []string    `json:"platform_types"`
	SchemaVersion   string      `json:"schema_version"`
}

type Parameter struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	DefaultValue string `json:"default_value"`
	Description  string `json:"description"`
}