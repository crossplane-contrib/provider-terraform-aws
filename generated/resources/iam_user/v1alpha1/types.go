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

// IamUser is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamUserSpec   `json:"spec"`
	Status IamUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamUser contains a list of IamUserList
type IamUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamUser `json:"items"`
}

// A IamUserSpec defines the desired state of a IamUser
type IamUserSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamUserParameters `json:"forProvider"`
}

// A IamUserParameters defines the desired state of a IamUser
type IamUserParameters struct {
	Tags                map[string]string `json:"tags"`
	ForceDestroy        bool              `json:"force_destroy"`
	Id                  string            `json:"id"`
	Name                string            `json:"name"`
	Path                string            `json:"path"`
	PermissionsBoundary string            `json:"permissions_boundary"`
}

// A IamUserStatus defines the observed state of a IamUser
type IamUserStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamUserObservation `json:"atProvider"`
}

// A IamUserObservation records the observed state of a IamUser
type IamUserObservation struct {
	UniqueId string `json:"unique_id"`
	Arn      string `json:"arn"`
}