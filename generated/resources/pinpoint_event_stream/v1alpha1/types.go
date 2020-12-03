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

// PinpointEventStream is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type PinpointEventStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PinpointEventStreamSpec   `json:"spec"`
	Status PinpointEventStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointEventStream contains a list of PinpointEventStreamList
type PinpointEventStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointEventStream `json:"items"`
}

// A PinpointEventStreamSpec defines the desired state of a PinpointEventStream
type PinpointEventStreamSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  PinpointEventStreamParameters `json:",inline"`
}

// A PinpointEventStreamParameters defines the desired state of a PinpointEventStream
type PinpointEventStreamParameters struct {
	Id                   string `json:"id"`
	RoleArn              string `json:"role_arn"`
	ApplicationId        string `json:"application_id"`
	DestinationStreamArn string `json:"destination_stream_arn"`
}

// A PinpointEventStreamStatus defines the observed state of a PinpointEventStream
type PinpointEventStreamStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     PinpointEventStreamObservation `json:",inline"`
}

// A PinpointEventStreamObservation records the observed state of a PinpointEventStream
type PinpointEventStreamObservation struct{}