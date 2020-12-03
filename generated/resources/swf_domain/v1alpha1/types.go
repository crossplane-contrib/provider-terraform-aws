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

// SwfDomain is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SwfDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SwfDomainSpec   `json:"spec"`
	Status SwfDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SwfDomain contains a list of SwfDomainList
type SwfDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SwfDomain `json:"items"`
}

// A SwfDomainSpec defines the desired state of a SwfDomain
type SwfDomainSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SwfDomainParameters `json:",inline"`
}

// A SwfDomainParameters defines the desired state of a SwfDomain
type SwfDomainParameters struct {
	Id                                     string            `json:"id"`
	Name                                   string            `json:"name"`
	NamePrefix                             string            `json:"name_prefix"`
	Tags                                   map[string]string `json:"tags"`
	WorkflowExecutionRetentionPeriodInDays string            `json:"workflow_execution_retention_period_in_days"`
	Description                            string            `json:"description"`
}

// A SwfDomainStatus defines the observed state of a SwfDomain
type SwfDomainStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SwfDomainObservation `json:",inline"`
}

// A SwfDomainObservation records the observed state of a SwfDomain
type SwfDomainObservation struct {
	Arn string `json:"arn"`
}