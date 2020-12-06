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

// SsmMaintenanceWindowTarget is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SsmMaintenanceWindowTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SsmMaintenanceWindowTargetSpec   `json:"spec"`
	Status SsmMaintenanceWindowTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmMaintenanceWindowTarget contains a list of SsmMaintenanceWindowTargetList
type SsmMaintenanceWindowTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmMaintenanceWindowTarget `json:"items"`
}

// A SsmMaintenanceWindowTargetSpec defines the desired state of a SsmMaintenanceWindowTarget
type SsmMaintenanceWindowTargetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SsmMaintenanceWindowTargetParameters `json:",inline"`
}

// A SsmMaintenanceWindowTargetParameters defines the desired state of a SsmMaintenanceWindowTarget
type SsmMaintenanceWindowTargetParameters struct {
	Name             string    `json:"name"`
	OwnerInformation string    `json:"owner_information"`
	ResourceType     string    `json:"resource_type"`
	WindowId         string    `json:"window_id"`
	Description      string    `json:"description"`
	Id               string    `json:"id"`
	Targets          []Targets `json:"targets"`
}

type Targets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// A SsmMaintenanceWindowTargetStatus defines the observed state of a SsmMaintenanceWindowTarget
type SsmMaintenanceWindowTargetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SsmMaintenanceWindowTargetObservation `json:",inline"`
}

// A SsmMaintenanceWindowTargetObservation records the observed state of a SsmMaintenanceWindowTarget
type SsmMaintenanceWindowTargetObservation struct{}