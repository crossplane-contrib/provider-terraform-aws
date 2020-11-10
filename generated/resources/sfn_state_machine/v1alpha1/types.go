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

// SfnStateMachine is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SfnStateMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SfnStateMachineSpec   `json:"spec"`
	Status SfnStateMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SfnStateMachine contains a list of SfnStateMachineList
type SfnStateMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SfnStateMachine `json:"items"`
}

// A SfnStateMachineSpec defines the desired state of a SfnStateMachine
type SfnStateMachineSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SfnStateMachineParameters `json:",inline"`
}

// A SfnStateMachineParameters defines the desired state of a SfnStateMachine
type SfnStateMachineParameters struct {
	Definition string            `json:"definition"`
	Name       string            `json:"name"`
	RoleArn    string            `json:"role_arn"`
	Tags       map[string]string `json:"tags"`
}

// A SfnStateMachineStatus defines the observed state of a SfnStateMachine
type SfnStateMachineStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SfnStateMachineObservation `json:",inline"`
}

// A SfnStateMachineObservation records the observed state of a SfnStateMachine
type SfnStateMachineObservation struct {
	Id           string `json:"id"`
	Status       string `json:"status"`
	Arn          string `json:"arn"`
	CreationDate string `json:"creation_date"`
}