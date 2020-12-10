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

// DatasyncAgent is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DatasyncAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatasyncAgentSpec   `json:"spec"`
	Status DatasyncAgentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasyncAgent contains a list of DatasyncAgentList
type DatasyncAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatasyncAgent `json:"items"`
}

// A DatasyncAgentSpec defines the desired state of a DatasyncAgent
type DatasyncAgentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DatasyncAgentParameters `json:",inline"`
}

// A DatasyncAgentParameters defines the desired state of a DatasyncAgent
type DatasyncAgentParameters struct {
	ActivationKey string            `json:"activation_key"`
	Id            string            `json:"id"`
	IpAddress     string            `json:"ip_address"`
	Name          string            `json:"name"`
	Tags          map[string]string `json:"tags"`
	Timeouts      Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
}

// A DatasyncAgentStatus defines the observed state of a DatasyncAgent
type DatasyncAgentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DatasyncAgentObservation `json:",inline"`
}

// A DatasyncAgentObservation records the observed state of a DatasyncAgent
type DatasyncAgentObservation struct {
	Arn string `json:"arn"`
}