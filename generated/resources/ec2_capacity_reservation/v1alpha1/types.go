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

// Ec2CapacityReservation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2CapacityReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2CapacityReservationSpec   `json:"spec"`
	Status Ec2CapacityReservationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2CapacityReservation contains a list of Ec2CapacityReservationList
type Ec2CapacityReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2CapacityReservation `json:"items"`
}

// A Ec2CapacityReservationSpec defines the desired state of a Ec2CapacityReservation
type Ec2CapacityReservationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2CapacityReservationParameters `json:",inline"`
}

// A Ec2CapacityReservationParameters defines the desired state of a Ec2CapacityReservation
type Ec2CapacityReservationParameters struct {
	InstancePlatform      string            `json:"instance_platform"`
	InstanceType          string            `json:"instance_type"`
	Tags                  map[string]string `json:"tags"`
	Tenancy               string            `json:"tenancy"`
	EbsOptimized          bool              `json:"ebs_optimized"`
	EndDate               string            `json:"end_date"`
	EndDateType           string            `json:"end_date_type"`
	InstanceCount         int64             `json:"instance_count"`
	InstanceMatchCriteria string            `json:"instance_match_criteria"`
	AvailabilityZone      string            `json:"availability_zone"`
	EphemeralStorage      bool              `json:"ephemeral_storage"`
	Id                    string            `json:"id"`
}

// A Ec2CapacityReservationStatus defines the observed state of a Ec2CapacityReservation
type Ec2CapacityReservationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2CapacityReservationObservation `json:",inline"`
}

// A Ec2CapacityReservationObservation records the observed state of a Ec2CapacityReservation
type Ec2CapacityReservationObservation struct {
	Arn string `json:"arn"`
}