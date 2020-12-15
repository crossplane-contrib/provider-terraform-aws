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

// QuicksightGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type QuicksightGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuicksightGroupSpec   `json:"spec"`
	Status QuicksightGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QuicksightGroup contains a list of QuicksightGroupList
type QuicksightGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuicksightGroup `json:"items"`
}

// A QuicksightGroupSpec defines the desired state of a QuicksightGroup
type QuicksightGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  QuicksightGroupParameters `json:"forProvider"`
}

// A QuicksightGroupParameters defines the desired state of a QuicksightGroup
type QuicksightGroupParameters struct {
	AwsAccountId string `json:"aws_account_id"`
	Description  string `json:"description"`
	GroupName    string `json:"group_name"`
	Namespace    string `json:"namespace"`
}

// A QuicksightGroupStatus defines the observed state of a QuicksightGroup
type QuicksightGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     QuicksightGroupObservation `json:"atProvider"`
}

// A QuicksightGroupObservation records the observed state of a QuicksightGroup
type QuicksightGroupObservation struct {
	Arn string `json:"arn"`
}