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

// FmsAdminAccount is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type FmsAdminAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FmsAdminAccountSpec   `json:"spec"`
	Status FmsAdminAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FmsAdminAccount contains a list of FmsAdminAccountList
type FmsAdminAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FmsAdminAccount `json:"items"`
}

// A FmsAdminAccountSpec defines the desired state of a FmsAdminAccount
type FmsAdminAccountSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  FmsAdminAccountParameters `json:"forProvider"`
}

// A FmsAdminAccountParameters defines the desired state of a FmsAdminAccount
type FmsAdminAccountParameters struct {
	AccountId string `json:"account_id"`
}

// A FmsAdminAccountStatus defines the observed state of a FmsAdminAccount
type FmsAdminAccountStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     FmsAdminAccountObservation `json:"atProvider"`
}

// A FmsAdminAccountObservation records the observed state of a FmsAdminAccount
type FmsAdminAccountObservation struct{}