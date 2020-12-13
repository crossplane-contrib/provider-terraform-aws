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

// CloudhsmV2Hsm is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudhsmV2Hsm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudhsmV2HsmSpec   `json:"spec"`
	Status CloudhsmV2HsmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudhsmV2Hsm contains a list of CloudhsmV2HsmList
type CloudhsmV2HsmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudhsmV2Hsm `json:"items"`
}

// A CloudhsmV2HsmSpec defines the desired state of a CloudhsmV2Hsm
type CloudhsmV2HsmSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudhsmV2HsmParameters `json:"forProvider"`
}

// A CloudhsmV2HsmParameters defines the desired state of a CloudhsmV2Hsm
type CloudhsmV2HsmParameters struct {
	Id               string   `json:"id"`
	IpAddress        string   `json:"ip_address"`
	SubnetId         string   `json:"subnet_id"`
	AvailabilityZone string   `json:"availability_zone"`
	ClusterId        string   `json:"cluster_id"`
	Timeouts         Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A CloudhsmV2HsmStatus defines the observed state of a CloudhsmV2Hsm
type CloudhsmV2HsmStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudhsmV2HsmObservation `json:"atProvider"`
}

// A CloudhsmV2HsmObservation records the observed state of a CloudhsmV2Hsm
type CloudhsmV2HsmObservation struct {
	HsmState string `json:"hsm_state"`
	HsmEniId string `json:"hsm_eni_id"`
	HsmId    string `json:"hsm_id"`
}