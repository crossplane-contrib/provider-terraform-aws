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

// InspectorResourceGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type InspectorResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InspectorResourceGroupSpec   `json:"spec"`
	Status InspectorResourceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InspectorResourceGroup contains a list of InspectorResourceGroupList
type InspectorResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InspectorResourceGroup `json:"items"`
}

// A InspectorResourceGroupSpec defines the desired state of a InspectorResourceGroup
type InspectorResourceGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  InspectorResourceGroupParameters `json:",inline"`
}

// A InspectorResourceGroupParameters defines the desired state of a InspectorResourceGroup
type InspectorResourceGroupParameters struct {
	Tags map[string]string `json:"tags"`
	Id   string            `json:"id"`
}

// A InspectorResourceGroupStatus defines the observed state of a InspectorResourceGroup
type InspectorResourceGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     InspectorResourceGroupObservation `json:",inline"`
}

// A InspectorResourceGroupObservation records the observed state of a InspectorResourceGroup
type InspectorResourceGroupObservation struct {
	Arn string `json:"arn"`
}