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

// AccessanalyzerAnalyzer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AccessanalyzerAnalyzer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessanalyzerAnalyzerSpec   `json:"spec"`
	Status AccessanalyzerAnalyzerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessanalyzerAnalyzer contains a list of AccessanalyzerAnalyzerList
type AccessanalyzerAnalyzerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessanalyzerAnalyzer `json:"items"`
}

// A AccessanalyzerAnalyzerSpec defines the desired state of a AccessanalyzerAnalyzer
type AccessanalyzerAnalyzerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AccessanalyzerAnalyzerParameters `json:",inline"`
}

// A AccessanalyzerAnalyzerParameters defines the desired state of a AccessanalyzerAnalyzer
type AccessanalyzerAnalyzerParameters struct {
	Tags         map[string]string `json:"tags"`
	Type         string            `json:"type"`
	AnalyzerName string            `json:"analyzer_name"`
	Id           string            `json:"id"`
}

// A AccessanalyzerAnalyzerStatus defines the observed state of a AccessanalyzerAnalyzer
type AccessanalyzerAnalyzerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AccessanalyzerAnalyzerObservation `json:",inline"`
}

// A AccessanalyzerAnalyzerObservation records the observed state of a AccessanalyzerAnalyzer
type AccessanalyzerAnalyzerObservation struct {
	Arn string `json:"arn"`
}