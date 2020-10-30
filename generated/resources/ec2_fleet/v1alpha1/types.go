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

// Ec2Fleet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2Fleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2FleetSpec   `json:"spec"`
	Status Ec2FleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2Fleet contains a list of Ec2FleetList
type Ec2FleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2Fleet `json:"items"`
}

// A Ec2FleetSpec defines the desired state of a Ec2Fleet
type Ec2FleetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2FleetParameters `json:",inline"`
}

// A Ec2FleetParameters defines the desired state of a Ec2Fleet
type Ec2FleetParameters struct {
	Type                             string `json:"type"`
	ExcessCapacityTerminationPolicy  string `json:"excess_capacity_termination_policy"`
	ReplaceUnhealthyInstances        bool   `json:"replace_unhealthy_instances"`
	TerminateInstances               bool   `json:"terminate_instances"`
	TerminateInstancesWithExpiration bool   `json:"terminate_instances_with_expiration"`
}

// A Ec2FleetStatus defines the observed state of a Ec2Fleet
type Ec2FleetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2FleetObservation `json:",inline"`
}

// A Ec2FleetObservation records the observed state of a Ec2Fleet
type Ec2FleetObservation struct {
	Id string `json:"id"`
}