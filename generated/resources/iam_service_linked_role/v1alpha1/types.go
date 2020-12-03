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

// IamServiceLinkedRole is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamServiceLinkedRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamServiceLinkedRoleSpec   `json:"spec"`
	Status IamServiceLinkedRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamServiceLinkedRole contains a list of IamServiceLinkedRoleList
type IamServiceLinkedRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamServiceLinkedRole `json:"items"`
}

// A IamServiceLinkedRoleSpec defines the desired state of a IamServiceLinkedRole
type IamServiceLinkedRoleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamServiceLinkedRoleParameters `json:",inline"`
}

// A IamServiceLinkedRoleParameters defines the desired state of a IamServiceLinkedRole
type IamServiceLinkedRoleParameters struct {
	AwsServiceName string `json:"aws_service_name"`
	Description    string `json:"description"`
	Id             string `json:"id"`
	CustomSuffix   string `json:"custom_suffix"`
}

// A IamServiceLinkedRoleStatus defines the observed state of a IamServiceLinkedRole
type IamServiceLinkedRoleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamServiceLinkedRoleObservation `json:",inline"`
}

// A IamServiceLinkedRoleObservation records the observed state of a IamServiceLinkedRole
type IamServiceLinkedRoleObservation struct {
	Arn        string `json:"arn"`
	CreateDate string `json:"create_date"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	UniqueId   string `json:"unique_id"`
}