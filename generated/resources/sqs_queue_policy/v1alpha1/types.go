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

// SqsQueuePolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SqsQueuePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SqsQueuePolicySpec   `json:"spec"`
	Status SqsQueuePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqsQueuePolicy contains a list of SqsQueuePolicyList
type SqsQueuePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqsQueuePolicy `json:"items"`
}

// A SqsQueuePolicySpec defines the desired state of a SqsQueuePolicy
type SqsQueuePolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SqsQueuePolicyParameters `json:"forProvider"`
}

// A SqsQueuePolicyParameters defines the desired state of a SqsQueuePolicy
type SqsQueuePolicyParameters struct {
	Id       string `json:"id"`
	Policy   string `json:"policy"`
	QueueUrl string `json:"queue_url"`
}

// A SqsQueuePolicyStatus defines the observed state of a SqsQueuePolicy
type SqsQueuePolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SqsQueuePolicyObservation `json:"atProvider"`
}

// A SqsQueuePolicyObservation records the observed state of a SqsQueuePolicy
type SqsQueuePolicyObservation struct{}