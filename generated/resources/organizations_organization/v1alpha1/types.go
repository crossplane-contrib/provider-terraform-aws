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

// OrganizationsOrganization is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OrganizationsOrganization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrganizationsOrganizationSpec   `json:"spec"`
	Status OrganizationsOrganizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationsOrganization contains a list of OrganizationsOrganizationList
type OrganizationsOrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationsOrganization `json:"items"`
}

// A OrganizationsOrganizationSpec defines the desired state of a OrganizationsOrganization
type OrganizationsOrganizationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OrganizationsOrganizationParameters `json:"forProvider"`
}

// A OrganizationsOrganizationParameters defines the desired state of a OrganizationsOrganization
type OrganizationsOrganizationParameters struct {
	AwsServiceAccessPrincipals []string `json:"aws_service_access_principals"`
	EnabledPolicyTypes         []string `json:"enabled_policy_types"`
	FeatureSet                 string   `json:"feature_set"`
	Id                         string   `json:"id"`
}

// A OrganizationsOrganizationStatus defines the observed state of a OrganizationsOrganization
type OrganizationsOrganizationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OrganizationsOrganizationObservation `json:"atProvider"`
}

// A OrganizationsOrganizationObservation records the observed state of a OrganizationsOrganization
type OrganizationsOrganizationObservation struct {
	MasterAccountId    string              `json:"master_account_id"`
	NonMasterAccounts  []NonMasterAccounts `json:"non_master_accounts"`
	Arn                string              `json:"arn"`
	MasterAccountArn   string              `json:"master_account_arn"`
	MasterAccountEmail string              `json:"master_account_email"`
	Roots              []Roots             `json:"roots"`
	Accounts           []Accounts          `json:"accounts"`
}

type NonMasterAccounts struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Arn    string `json:"arn"`
	Email  string `json:"email"`
	Id     string `json:"id"`
}

type Roots struct {
	Name        string        `json:"name"`
	PolicyTypes []PolicyTypes `json:"policy_types"`
	Arn         string        `json:"arn"`
	Id          string        `json:"id"`
}

type PolicyTypes struct {
	Status string `json:"status"`
	Type   string `json:"type"`
}

type Accounts struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Arn    string `json:"arn"`
	Email  string `json:"email"`
}