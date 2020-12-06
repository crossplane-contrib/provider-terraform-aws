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

// AppsyncFunction is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppsyncFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppsyncFunctionSpec   `json:"spec"`
	Status AppsyncFunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppsyncFunction contains a list of AppsyncFunctionList
type AppsyncFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppsyncFunction `json:"items"`
}

// A AppsyncFunctionSpec defines the desired state of a AppsyncFunction
type AppsyncFunctionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppsyncFunctionParameters `json:",inline"`
}

// A AppsyncFunctionParameters defines the desired state of a AppsyncFunction
type AppsyncFunctionParameters struct {
	DataSource              string `json:"data_source"`
	Description             string `json:"description"`
	FunctionVersion         string `json:"function_version"`
	Id                      string `json:"id"`
	Name                    string `json:"name"`
	ApiId                   string `json:"api_id"`
	ResponseMappingTemplate string `json:"response_mapping_template"`
	RequestMappingTemplate  string `json:"request_mapping_template"`
}

// A AppsyncFunctionStatus defines the observed state of a AppsyncFunction
type AppsyncFunctionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppsyncFunctionObservation `json:",inline"`
}

// A AppsyncFunctionObservation records the observed state of a AppsyncFunction
type AppsyncFunctionObservation struct {
	Arn        string `json:"arn"`
	FunctionId string `json:"function_id"`
}