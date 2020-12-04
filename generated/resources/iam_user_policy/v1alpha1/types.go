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

// IamUserPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamUserPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamUserPolicySpec   `json:"spec"`
	Status IamUserPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamUserPolicy contains a list of IamUserPolicyList
type IamUserPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamUserPolicy `json:"items"`
}

// A IamUserPolicySpec defines the desired state of a IamUserPolicy
type IamUserPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamUserPolicyParameters `json:",inline"`
}

// A IamUserPolicyParameters defines the desired state of a IamUserPolicy
type IamUserPolicyParameters struct {
	User       string `json:"user"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	NamePrefix string `json:"name_prefix"`
	Policy     string `json:"policy"`
}

// A IamUserPolicyStatus defines the observed state of a IamUserPolicy
type IamUserPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamUserPolicyObservation `json:",inline"`
}

// A IamUserPolicyObservation records the observed state of a IamUserPolicy
type IamUserPolicyObservation struct{}