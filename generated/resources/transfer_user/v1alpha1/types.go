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

// TransferUser is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type TransferUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TransferUserSpec   `json:"spec"`
	Status TransferUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransferUser contains a list of TransferUserList
type TransferUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransferUser `json:"items"`
}

// A TransferUserSpec defines the desired state of a TransferUser
type TransferUserSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  TransferUserParameters `json:",inline"`
}

// A TransferUserParameters defines the desired state of a TransferUser
type TransferUserParameters struct {
	HomeDirectoryType     string                `json:"home_directory_type"`
	Role                  string                `json:"role"`
	Tags                  map[string]string     `json:"tags"`
	Id                    string                `json:"id"`
	Policy                string                `json:"policy"`
	ServerId              string                `json:"server_id"`
	UserName              string                `json:"user_name"`
	HomeDirectory         string                `json:"home_directory"`
	HomeDirectoryMappings HomeDirectoryMappings `json:"home_directory_mappings"`
}

type HomeDirectoryMappings struct {
	Entry  string `json:"entry"`
	Target string `json:"target"`
}

// A TransferUserStatus defines the observed state of a TransferUser
type TransferUserStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     TransferUserObservation `json:",inline"`
}

// A TransferUserObservation records the observed state of a TransferUser
type TransferUserObservation struct {
	Arn string `json:"arn"`
}