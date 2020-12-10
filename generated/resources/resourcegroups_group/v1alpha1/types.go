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

// ResourcegroupsGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ResourcegroupsGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourcegroupsGroupSpec   `json:"spec"`
	Status ResourcegroupsGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcegroupsGroup contains a list of ResourcegroupsGroupList
type ResourcegroupsGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourcegroupsGroup `json:"items"`
}

// A ResourcegroupsGroupSpec defines the desired state of a ResourcegroupsGroup
type ResourcegroupsGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ResourcegroupsGroupParameters `json:",inline"`
}

// A ResourcegroupsGroupParameters defines the desired state of a ResourcegroupsGroup
type ResourcegroupsGroupParameters struct {
	Id            string            `json:"id"`
	Name          string            `json:"name"`
	Tags          map[string]string `json:"tags"`
	Description   string            `json:"description"`
	ResourceQuery ResourceQuery     `json:"resource_query"`
}

type ResourceQuery struct {
	Query string `json:"query"`
	Type  string `json:"type"`
}

// A ResourcegroupsGroupStatus defines the observed state of a ResourcegroupsGroup
type ResourcegroupsGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ResourcegroupsGroupObservation `json:",inline"`
}

// A ResourcegroupsGroupObservation records the observed state of a ResourcegroupsGroup
type ResourcegroupsGroupObservation struct {
	Arn string `json:"arn"`
}