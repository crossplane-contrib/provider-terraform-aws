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

// SecurityhubAccount is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SecurityhubAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecurityhubAccountSpec   `json:"spec"`
	Status SecurityhubAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityhubAccount contains a list of SecurityhubAccountList
type SecurityhubAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityhubAccount `json:"items"`
}

// A SecurityhubAccountSpec defines the desired state of a SecurityhubAccount
type SecurityhubAccountSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SecurityhubAccountParameters `json:"forProvider"`
}

// A SecurityhubAccountParameters defines the desired state of a SecurityhubAccount
type SecurityhubAccountParameters struct {
	Id string `json:"id"`
}

// A SecurityhubAccountStatus defines the observed state of a SecurityhubAccount
type SecurityhubAccountStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SecurityhubAccountObservation `json:"atProvider"`
}

// A SecurityhubAccountObservation records the observed state of a SecurityhubAccount
type SecurityhubAccountObservation struct{}