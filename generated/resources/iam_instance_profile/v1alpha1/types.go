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

// IamInstanceProfile is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamInstanceProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamInstanceProfileSpec   `json:"spec"`
	Status IamInstanceProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamInstanceProfile contains a list of IamInstanceProfileList
type IamInstanceProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamInstanceProfile `json:"items"`
}

// A IamInstanceProfileSpec defines the desired state of a IamInstanceProfile
type IamInstanceProfileSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamInstanceProfileParameters `json:",inline"`
}

// A IamInstanceProfileParameters defines the desired state of a IamInstanceProfile
type IamInstanceProfileParameters struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	NamePrefix string `json:"name_prefix"`
	Path       string `json:"path"`
	Role       string `json:"role"`
}

// A IamInstanceProfileStatus defines the observed state of a IamInstanceProfile
type IamInstanceProfileStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamInstanceProfileObservation `json:",inline"`
}

// A IamInstanceProfileObservation records the observed state of a IamInstanceProfile
type IamInstanceProfileObservation struct {
	UniqueId   string `json:"unique_id"`
	Arn        string `json:"arn"`
	CreateDate string `json:"create_date"`
}