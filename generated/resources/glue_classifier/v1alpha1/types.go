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

// GlueClassifier is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlueClassifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlueClassifierSpec   `json:"spec"`
	Status GlueClassifierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueClassifier contains a list of GlueClassifierList
type GlueClassifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueClassifier `json:"items"`
}

// A GlueClassifierSpec defines the desired state of a GlueClassifier
type GlueClassifierSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlueClassifierParameters `json:",inline"`
}

// A GlueClassifierParameters defines the desired state of a GlueClassifier
type GlueClassifierParameters struct {
	Name           string         `json:"name"`
	CsvClassifier  CsvClassifier  `json:"csv_classifier"`
	GrokClassifier GrokClassifier `json:"grok_classifier"`
	JsonClassifier JsonClassifier `json:"json_classifier"`
	XmlClassifier  XmlClassifier  `json:"xml_classifier"`
}

type CsvClassifier struct {
	QuoteSymbol          string   `json:"quote_symbol"`
	AllowSingleColumn    bool     `json:"allow_single_column"`
	ContainsHeader       string   `json:"contains_header"`
	Delimiter            string   `json:"delimiter"`
	DisableValueTrimming bool     `json:"disable_value_trimming"`
	Header               []string `json:"header"`
}

type GrokClassifier struct {
	Classification string `json:"classification"`
	CustomPatterns string `json:"custom_patterns"`
	GrokPattern    string `json:"grok_pattern"`
}

type JsonClassifier struct {
	JsonPath string `json:"json_path"`
}

type XmlClassifier struct {
	Classification string `json:"classification"`
	RowTag         string `json:"row_tag"`
}

// A GlueClassifierStatus defines the observed state of a GlueClassifier
type GlueClassifierStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlueClassifierObservation `json:",inline"`
}

// A GlueClassifierObservation records the observed state of a GlueClassifier
type GlueClassifierObservation struct {
	Id string `json:"id"`
}