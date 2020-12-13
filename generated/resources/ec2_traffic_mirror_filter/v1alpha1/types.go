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

// Ec2TrafficMirrorFilter is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2TrafficMirrorFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2TrafficMirrorFilterSpec   `json:"spec"`
	Status Ec2TrafficMirrorFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TrafficMirrorFilter contains a list of Ec2TrafficMirrorFilterList
type Ec2TrafficMirrorFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TrafficMirrorFilter `json:"items"`
}

// A Ec2TrafficMirrorFilterSpec defines the desired state of a Ec2TrafficMirrorFilter
type Ec2TrafficMirrorFilterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2TrafficMirrorFilterParameters `json:"forProvider"`
}

// A Ec2TrafficMirrorFilterParameters defines the desired state of a Ec2TrafficMirrorFilter
type Ec2TrafficMirrorFilterParameters struct {
	Description     string            `json:"description"`
	Id              string            `json:"id"`
	NetworkServices []string          `json:"network_services"`
	Tags            map[string]string `json:"tags"`
}

// A Ec2TrafficMirrorFilterStatus defines the observed state of a Ec2TrafficMirrorFilter
type Ec2TrafficMirrorFilterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2TrafficMirrorFilterObservation `json:"atProvider"`
}

// A Ec2TrafficMirrorFilterObservation records the observed state of a Ec2TrafficMirrorFilter
type Ec2TrafficMirrorFilterObservation struct{}