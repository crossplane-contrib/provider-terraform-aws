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

// GlueUserDefinedFunction is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlueUserDefinedFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlueUserDefinedFunctionSpec   `json:"spec"`
	Status GlueUserDefinedFunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueUserDefinedFunction contains a list of GlueUserDefinedFunctionList
type GlueUserDefinedFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueUserDefinedFunction `json:"items"`
}

// A GlueUserDefinedFunctionSpec defines the desired state of a GlueUserDefinedFunction
type GlueUserDefinedFunctionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlueUserDefinedFunctionParameters `json:",inline"`
}

// A GlueUserDefinedFunctionParameters defines the desired state of a GlueUserDefinedFunction
type GlueUserDefinedFunctionParameters struct {
	Name         string         `json:"name"`
	OwnerType    string         `json:"owner_type"`
	CatalogId    string         `json:"catalog_id"`
	ClassName    string         `json:"class_name"`
	DatabaseName string         `json:"database_name"`
	Id           string         `json:"id"`
	OwnerName    string         `json:"owner_name"`
	ResourceUris []ResourceUris `json:"resource_uris"`
}

type ResourceUris struct {
	ResourceType string `json:"resource_type"`
	Uri          string `json:"uri"`
}

// A GlueUserDefinedFunctionStatus defines the observed state of a GlueUserDefinedFunction
type GlueUserDefinedFunctionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlueUserDefinedFunctionObservation `json:",inline"`
}

// A GlueUserDefinedFunctionObservation records the observed state of a GlueUserDefinedFunction
type GlueUserDefinedFunctionObservation struct {
	CreateTime string `json:"create_time"`
	Arn        string `json:"arn"`
}