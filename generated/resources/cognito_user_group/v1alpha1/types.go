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

// CognitoUserGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CognitoUserGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CognitoUserGroupSpec   `json:"spec"`
	Status CognitoUserGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserGroup contains a list of CognitoUserGroupList
type CognitoUserGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoUserGroup `json:"items"`
}

// A CognitoUserGroupSpec defines the desired state of a CognitoUserGroup
type CognitoUserGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CognitoUserGroupParameters `json:"forProvider"`
}

// A CognitoUserGroupParameters defines the desired state of a CognitoUserGroup
type CognitoUserGroupParameters struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Precedence  int64  `json:"precedence"`
	RoleArn     string `json:"role_arn"`
	UserPoolId  string `json:"user_pool_id"`
	Description string `json:"description"`
}

// A CognitoUserGroupStatus defines the observed state of a CognitoUserGroup
type CognitoUserGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CognitoUserGroupObservation `json:"atProvider"`
}

// A CognitoUserGroupObservation records the observed state of a CognitoUserGroup
type CognitoUserGroupObservation struct{}