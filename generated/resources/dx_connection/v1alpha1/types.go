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

// DxConnection is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxConnectionSpec   `json:"spec"`
	Status DxConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxConnection contains a list of DxConnectionList
type DxConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxConnection `json:"items"`
}

// A DxConnectionSpec defines the desired state of a DxConnection
type DxConnectionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxConnectionParameters `json:"forProvider"`
}

// A DxConnectionParameters defines the desired state of a DxConnection
type DxConnectionParameters struct {
	Name      string            `json:"name"`
	Location  string            `json:"location"`
	Bandwidth string            `json:"bandwidth"`
	Tags      map[string]string `json:"tags"`
}

// A DxConnectionStatus defines the observed state of a DxConnection
type DxConnectionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxConnectionObservation `json:"atProvider"`
}

// A DxConnectionObservation records the observed state of a DxConnection
type DxConnectionObservation struct {
	HasLogicalRedundancy string `json:"has_logical_redundancy"`
	JumboFrameCapable    bool   `json:"jumbo_frame_capable"`
	Arn                  string `json:"arn"`
	AwsDevice            string `json:"aws_device"`
}