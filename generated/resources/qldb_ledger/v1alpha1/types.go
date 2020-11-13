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

// QldbLedger is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type QldbLedger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QldbLedgerSpec   `json:"spec"`
	Status QldbLedgerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QldbLedger contains a list of QldbLedgerList
type QldbLedgerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QldbLedger `json:"items"`
}

// A QldbLedgerSpec defines the desired state of a QldbLedger
type QldbLedgerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  QldbLedgerParameters `json:",inline"`
}

// A QldbLedgerParameters defines the desired state of a QldbLedger
type QldbLedgerParameters struct {
	Tags               map[string]string `json:"tags"`
	DeletionProtection bool              `json:"deletion_protection"`
	Id                 string            `json:"id"`
	Name               string            `json:"name"`
}

// A QldbLedgerStatus defines the observed state of a QldbLedger
type QldbLedgerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     QldbLedgerObservation `json:",inline"`
}

// A QldbLedgerObservation records the observed state of a QldbLedger
type QldbLedgerObservation struct {
	Arn string `json:"arn"`
}