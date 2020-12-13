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

// GlueTrigger is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlueTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlueTriggerSpec   `json:"spec"`
	Status GlueTriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueTrigger contains a list of GlueTriggerList
type GlueTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueTrigger `json:"items"`
}

// A GlueTriggerSpec defines the desired state of a GlueTrigger
type GlueTriggerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlueTriggerParameters `json:"forProvider"`
}

// A GlueTriggerParameters defines the desired state of a GlueTrigger
type GlueTriggerParameters struct {
	Schedule     string            `json:"schedule"`
	Tags         map[string]string `json:"tags"`
	Type         string            `json:"type"`
	Description  string            `json:"description"`
	Name         string            `json:"name"`
	Id           string            `json:"id"`
	WorkflowName string            `json:"workflow_name"`
	Enabled      bool              `json:"enabled"`
	Actions      []Actions         `json:"actions"`
	Predicate    Predicate         `json:"predicate"`
	Timeouts     Timeouts          `json:"timeouts"`
}

type Actions struct {
	Arguments   map[string]string `json:"arguments"`
	CrawlerName string            `json:"crawler_name"`
	JobName     string            `json:"job_name"`
	Timeout     int64             `json:"timeout"`
}

type Predicate struct {
	Logical    string       `json:"logical"`
	Conditions []Conditions `json:"conditions"`
}

type Conditions struct {
	CrawlState      string `json:"crawl_state"`
	CrawlerName     string `json:"crawler_name"`
	JobName         string `json:"job_name"`
	LogicalOperator string `json:"logical_operator"`
	State           string `json:"state"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A GlueTriggerStatus defines the observed state of a GlueTrigger
type GlueTriggerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlueTriggerObservation `json:"atProvider"`
}

// A GlueTriggerObservation records the observed state of a GlueTrigger
type GlueTriggerObservation struct {
	Arn string `json:"arn"`
}