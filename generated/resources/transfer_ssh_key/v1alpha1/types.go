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

// TransferSshKey is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type TransferSshKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TransferSshKeySpec   `json:"spec"`
	Status TransferSshKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransferSshKey contains a list of TransferSshKeyList
type TransferSshKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransferSshKey `json:"items"`
}

// A TransferSshKeySpec defines the desired state of a TransferSshKey
type TransferSshKeySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  TransferSshKeyParameters `json:",inline"`
}

// A TransferSshKeyParameters defines the desired state of a TransferSshKey
type TransferSshKeyParameters struct {
	Body     string `json:"body"`
	Id       string `json:"id"`
	ServerId string `json:"server_id"`
	UserName string `json:"user_name"`
}

// A TransferSshKeyStatus defines the observed state of a TransferSshKey
type TransferSshKeyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     TransferSshKeyObservation `json:",inline"`
}

// A TransferSshKeyObservation records the observed state of a TransferSshKey
type TransferSshKeyObservation struct{}