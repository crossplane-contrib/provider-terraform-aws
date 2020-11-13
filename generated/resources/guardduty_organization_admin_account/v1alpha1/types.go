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

// GuarddutyOrganizationAdminAccount is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GuarddutyOrganizationAdminAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GuarddutyOrganizationAdminAccountSpec   `json:"spec"`
	Status GuarddutyOrganizationAdminAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuarddutyOrganizationAdminAccount contains a list of GuarddutyOrganizationAdminAccountList
type GuarddutyOrganizationAdminAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GuarddutyOrganizationAdminAccount `json:"items"`
}

// A GuarddutyOrganizationAdminAccountSpec defines the desired state of a GuarddutyOrganizationAdminAccount
type GuarddutyOrganizationAdminAccountSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GuarddutyOrganizationAdminAccountParameters `json:",inline"`
}

// A GuarddutyOrganizationAdminAccountParameters defines the desired state of a GuarddutyOrganizationAdminAccount
type GuarddutyOrganizationAdminAccountParameters struct {
	AdminAccountId string `json:"admin_account_id"`
	Id             string `json:"id"`
}

// A GuarddutyOrganizationAdminAccountStatus defines the observed state of a GuarddutyOrganizationAdminAccount
type GuarddutyOrganizationAdminAccountStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GuarddutyOrganizationAdminAccountObservation `json:",inline"`
}

// A GuarddutyOrganizationAdminAccountObservation records the observed state of a GuarddutyOrganizationAdminAccount
type GuarddutyOrganizationAdminAccountObservation struct{}