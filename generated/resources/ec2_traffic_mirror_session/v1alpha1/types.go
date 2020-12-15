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

// Ec2TrafficMirrorSession is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2TrafficMirrorSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2TrafficMirrorSessionSpec   `json:"spec"`
	Status Ec2TrafficMirrorSessionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TrafficMirrorSession contains a list of Ec2TrafficMirrorSessionList
type Ec2TrafficMirrorSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TrafficMirrorSession `json:"items"`
}

// A Ec2TrafficMirrorSessionSpec defines the desired state of a Ec2TrafficMirrorSession
type Ec2TrafficMirrorSessionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2TrafficMirrorSessionParameters `json:"forProvider"`
}

// A Ec2TrafficMirrorSessionParameters defines the desired state of a Ec2TrafficMirrorSession
type Ec2TrafficMirrorSessionParameters struct {
	Description           string            `json:"description"`
	SessionNumber         int64             `json:"session_number"`
	TrafficMirrorTargetId string            `json:"traffic_mirror_target_id"`
	NetworkInterfaceId    string            `json:"network_interface_id"`
	PacketLength          int64             `json:"packet_length"`
	Tags                  map[string]string `json:"tags"`
	TrafficMirrorFilterId string            `json:"traffic_mirror_filter_id"`
	VirtualNetworkId      int64             `json:"virtual_network_id"`
}

// A Ec2TrafficMirrorSessionStatus defines the observed state of a Ec2TrafficMirrorSession
type Ec2TrafficMirrorSessionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2TrafficMirrorSessionObservation `json:"atProvider"`
}

// A Ec2TrafficMirrorSessionObservation records the observed state of a Ec2TrafficMirrorSession
type Ec2TrafficMirrorSessionObservation struct {
	Arn string `json:"arn"`
}