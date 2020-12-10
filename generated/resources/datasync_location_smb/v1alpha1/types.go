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

// DatasyncLocationSmb is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DatasyncLocationSmb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatasyncLocationSmbSpec   `json:"spec"`
	Status DatasyncLocationSmbStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasyncLocationSmb contains a list of DatasyncLocationSmbList
type DatasyncLocationSmbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatasyncLocationSmb `json:"items"`
}

// A DatasyncLocationSmbSpec defines the desired state of a DatasyncLocationSmb
type DatasyncLocationSmbSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DatasyncLocationSmbParameters `json:",inline"`
}

// A DatasyncLocationSmbParameters defines the desired state of a DatasyncLocationSmb
type DatasyncLocationSmbParameters struct {
	AgentArns      []string          `json:"agent_arns"`
	Id             string            `json:"id"`
	User           string            `json:"user"`
	Tags           map[string]string `json:"tags"`
	Domain         string            `json:"domain"`
	Password       string            `json:"password"`
	ServerHostname string            `json:"server_hostname"`
	Subdirectory   string            `json:"subdirectory"`
	MountOptions   MountOptions      `json:"mount_options"`
}

type MountOptions struct {
	Version string `json:"version"`
}

// A DatasyncLocationSmbStatus defines the observed state of a DatasyncLocationSmb
type DatasyncLocationSmbStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DatasyncLocationSmbObservation `json:",inline"`
}

// A DatasyncLocationSmbObservation records the observed state of a DatasyncLocationSmb
type DatasyncLocationSmbObservation struct {
	Uri string `json:"uri"`
	Arn string `json:"arn"`
}