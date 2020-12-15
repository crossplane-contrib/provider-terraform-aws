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

// Ec2TrafficMirrorTarget is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2TrafficMirrorTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2TrafficMirrorTargetSpec   `json:"spec"`
	Status Ec2TrafficMirrorTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TrafficMirrorTarget contains a list of Ec2TrafficMirrorTargetList
type Ec2TrafficMirrorTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TrafficMirrorTarget `json:"items"`
}

// A Ec2TrafficMirrorTargetSpec defines the desired state of a Ec2TrafficMirrorTarget
type Ec2TrafficMirrorTargetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2TrafficMirrorTargetParameters `json:"forProvider"`
}

// A Ec2TrafficMirrorTargetParameters defines the desired state of a Ec2TrafficMirrorTarget
type Ec2TrafficMirrorTargetParameters struct {
	Tags                   map[string]string `json:"tags"`
	Description            string            `json:"description"`
	NetworkInterfaceId     string            `json:"network_interface_id"`
	NetworkLoadBalancerArn string            `json:"network_load_balancer_arn"`
}

// A Ec2TrafficMirrorTargetStatus defines the observed state of a Ec2TrafficMirrorTarget
type Ec2TrafficMirrorTargetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2TrafficMirrorTargetObservation `json:"atProvider"`
}

// A Ec2TrafficMirrorTargetObservation records the observed state of a Ec2TrafficMirrorTarget
type Ec2TrafficMirrorTargetObservation struct {
	Arn string `json:"arn"`
}