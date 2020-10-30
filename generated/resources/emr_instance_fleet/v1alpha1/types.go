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

// EmrInstanceFleet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EmrInstanceFleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmrInstanceFleetSpec   `json:"spec"`
	Status EmrInstanceFleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmrInstanceFleet contains a list of EmrInstanceFleetList
type EmrInstanceFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmrInstanceFleet `json:"items"`
}

// A EmrInstanceFleetSpec defines the desired state of a EmrInstanceFleet
type EmrInstanceFleetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EmrInstanceFleetParameters `json:",inline"`
}

// A EmrInstanceFleetParameters defines the desired state of a EmrInstanceFleet
type EmrInstanceFleetParameters struct {
	ClusterId              string `json:"cluster_id"`
	Name                   string `json:"name"`
	TargetOnDemandCapacity int    `json:"target_on_demand_capacity"`
	TargetSpotCapacity     int    `json:"target_spot_capacity"`
}

// A EmrInstanceFleetStatus defines the observed state of a EmrInstanceFleet
type EmrInstanceFleetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EmrInstanceFleetObservation `json:",inline"`
}

// A EmrInstanceFleetObservation records the observed state of a EmrInstanceFleet
type EmrInstanceFleetObservation struct {
	Id                          string `json:"id"`
	ProvisionedOnDemandCapacity int    `json:"provisioned_on_demand_capacity"`
	ProvisionedSpotCapacity     int    `json:"provisioned_spot_capacity"`
}