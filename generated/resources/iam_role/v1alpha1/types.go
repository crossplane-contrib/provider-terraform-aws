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

// IamRole is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamRoleSpec   `json:"spec"`
	Status IamRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamRole contains a list of IamRoleList
type IamRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamRole `json:"items"`
}

// A IamRoleSpec defines the desired state of a IamRole
type IamRoleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamRoleParameters `json:"forProvider"`
}

// A IamRoleParameters defines the desired state of a IamRole
type IamRoleParameters struct {
	ForceDetachPolicies bool              `json:"force_detach_policies"`
	MaxSessionDuration  int64             `json:"max_session_duration"`
	Name                string            `json:"name"`
	Path                string            `json:"path"`
	Tags                map[string]string `json:"tags"`
	AssumeRolePolicy    string            `json:"assume_role_policy"`
	Description         string            `json:"description"`
	Id                  string            `json:"id"`
	NamePrefix          string            `json:"name_prefix"`
	PermissionsBoundary string            `json:"permissions_boundary"`
}

// A IamRoleStatus defines the observed state of a IamRole
type IamRoleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamRoleObservation `json:"atProvider"`
}

// A IamRoleObservation records the observed state of a IamRole
type IamRoleObservation struct {
	Arn        string `json:"arn"`
	CreateDate string `json:"create_date"`
	UniqueId   string `json:"unique_id"`
}