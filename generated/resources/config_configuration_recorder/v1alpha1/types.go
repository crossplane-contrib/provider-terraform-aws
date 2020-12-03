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

// ConfigConfigurationRecorder is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ConfigConfigurationRecorder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigConfigurationRecorderSpec   `json:"spec"`
	Status ConfigConfigurationRecorderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigConfigurationRecorder contains a list of ConfigConfigurationRecorderList
type ConfigConfigurationRecorderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigConfigurationRecorder `json:"items"`
}

// A ConfigConfigurationRecorderSpec defines the desired state of a ConfigConfigurationRecorder
type ConfigConfigurationRecorderSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ConfigConfigurationRecorderParameters `json:",inline"`
}

// A ConfigConfigurationRecorderParameters defines the desired state of a ConfigConfigurationRecorder
type ConfigConfigurationRecorderParameters struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	RoleArn        string         `json:"role_arn"`
	RecordingGroup RecordingGroup `json:"recording_group"`
}

type RecordingGroup struct {
	AllSupported               bool     `json:"all_supported"`
	IncludeGlobalResourceTypes bool     `json:"include_global_resource_types"`
	ResourceTypes              []string `json:"resource_types"`
}

// A ConfigConfigurationRecorderStatus defines the observed state of a ConfigConfigurationRecorder
type ConfigConfigurationRecorderStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ConfigConfigurationRecorderObservation `json:",inline"`
}

// A ConfigConfigurationRecorderObservation records the observed state of a ConfigConfigurationRecorder
type ConfigConfigurationRecorderObservation struct{}