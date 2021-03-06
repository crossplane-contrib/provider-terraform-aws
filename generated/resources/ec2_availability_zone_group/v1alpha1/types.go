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

// Ec2AvailabilityZoneGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2AvailabilityZoneGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2AvailabilityZoneGroupSpec   `json:"spec"`
	Status Ec2AvailabilityZoneGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2AvailabilityZoneGroup contains a list of Ec2AvailabilityZoneGroupList
type Ec2AvailabilityZoneGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2AvailabilityZoneGroup `json:"items"`
}

// A Ec2AvailabilityZoneGroupSpec defines the desired state of a Ec2AvailabilityZoneGroup
type Ec2AvailabilityZoneGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2AvailabilityZoneGroupParameters `json:"forProvider"`
}

// A Ec2AvailabilityZoneGroupParameters defines the desired state of a Ec2AvailabilityZoneGroup
type Ec2AvailabilityZoneGroupParameters struct {
	GroupName   string `json:"group_name"`
	OptInStatus string `json:"opt_in_status"`
}

// A Ec2AvailabilityZoneGroupStatus defines the observed state of a Ec2AvailabilityZoneGroup
type Ec2AvailabilityZoneGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2AvailabilityZoneGroupObservation `json:"atProvider"`
}

// A Ec2AvailabilityZoneGroupObservation records the observed state of a Ec2AvailabilityZoneGroup
type Ec2AvailabilityZoneGroupObservation struct{}