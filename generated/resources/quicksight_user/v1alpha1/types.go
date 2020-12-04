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

// QuicksightUser is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type QuicksightUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuicksightUserSpec   `json:"spec"`
	Status QuicksightUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QuicksightUser contains a list of QuicksightUserList
type QuicksightUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuicksightUser `json:"items"`
}

// A QuicksightUserSpec defines the desired state of a QuicksightUser
type QuicksightUserSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  QuicksightUserParameters `json:",inline"`
}

// A QuicksightUserParameters defines the desired state of a QuicksightUser
type QuicksightUserParameters struct {
	Id           string `json:"id"`
	UserName     string `json:"user_name"`
	IdentityType string `json:"identity_type"`
	Namespace    string `json:"namespace"`
	SessionName  string `json:"session_name"`
	UserRole     string `json:"user_role"`
	AwsAccountId string `json:"aws_account_id"`
	Email        string `json:"email"`
	IamArn       string `json:"iam_arn"`
}

// A QuicksightUserStatus defines the observed state of a QuicksightUser
type QuicksightUserStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     QuicksightUserObservation `json:",inline"`
}

// A QuicksightUserObservation records the observed state of a QuicksightUser
type QuicksightUserObservation struct {
	Arn string `json:"arn"`
}