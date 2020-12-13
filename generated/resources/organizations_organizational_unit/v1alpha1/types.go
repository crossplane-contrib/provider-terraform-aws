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

// OrganizationsOrganizationalUnit is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OrganizationsOrganizationalUnit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrganizationsOrganizationalUnitSpec   `json:"spec"`
	Status OrganizationsOrganizationalUnitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationsOrganizationalUnit contains a list of OrganizationsOrganizationalUnitList
type OrganizationsOrganizationalUnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationsOrganizationalUnit `json:"items"`
}

// A OrganizationsOrganizationalUnitSpec defines the desired state of a OrganizationsOrganizationalUnit
type OrganizationsOrganizationalUnitSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OrganizationsOrganizationalUnitParameters `json:"forProvider"`
}

// A OrganizationsOrganizationalUnitParameters defines the desired state of a OrganizationsOrganizationalUnit
type OrganizationsOrganizationalUnitParameters struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

// A OrganizationsOrganizationalUnitStatus defines the observed state of a OrganizationsOrganizationalUnit
type OrganizationsOrganizationalUnitStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OrganizationsOrganizationalUnitObservation `json:"atProvider"`
}

// A OrganizationsOrganizationalUnitObservation records the observed state of a OrganizationsOrganizationalUnit
type OrganizationsOrganizationalUnitObservation struct {
	Accounts []Accounts `json:"accounts"`
	Arn      string     `json:"arn"`
}

type Accounts struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	Name  string `json:"name"`
	Arn   string `json:"arn"`
}