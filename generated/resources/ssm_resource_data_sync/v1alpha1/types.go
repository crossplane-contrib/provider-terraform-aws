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

// SsmResourceDataSync is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SsmResourceDataSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SsmResourceDataSyncSpec   `json:"spec"`
	Status SsmResourceDataSyncStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmResourceDataSync contains a list of SsmResourceDataSyncList
type SsmResourceDataSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmResourceDataSync `json:"items"`
}

// A SsmResourceDataSyncSpec defines the desired state of a SsmResourceDataSync
type SsmResourceDataSyncSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SsmResourceDataSyncParameters `json:",inline"`
}

// A SsmResourceDataSyncParameters defines the desired state of a SsmResourceDataSync
type SsmResourceDataSyncParameters struct {
	Name string `json:"name"`
}

// A SsmResourceDataSyncStatus defines the observed state of a SsmResourceDataSync
type SsmResourceDataSyncStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SsmResourceDataSyncObservation `json:",inline"`
}

// A SsmResourceDataSyncObservation records the observed state of a SsmResourceDataSync
type SsmResourceDataSyncObservation struct {
	Id string `json:"id"`
}