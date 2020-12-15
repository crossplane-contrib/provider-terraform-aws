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

// GuarddutyOrganizationConfiguration is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GuarddutyOrganizationConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GuarddutyOrganizationConfigurationSpec   `json:"spec"`
	Status GuarddutyOrganizationConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuarddutyOrganizationConfiguration contains a list of GuarddutyOrganizationConfigurationList
type GuarddutyOrganizationConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GuarddutyOrganizationConfiguration `json:"items"`
}

// A GuarddutyOrganizationConfigurationSpec defines the desired state of a GuarddutyOrganizationConfiguration
type GuarddutyOrganizationConfigurationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GuarddutyOrganizationConfigurationParameters `json:"forProvider"`
}

// A GuarddutyOrganizationConfigurationParameters defines the desired state of a GuarddutyOrganizationConfiguration
type GuarddutyOrganizationConfigurationParameters struct {
	AutoEnable bool   `json:"auto_enable"`
	DetectorId string `json:"detector_id"`
}

// A GuarddutyOrganizationConfigurationStatus defines the observed state of a GuarddutyOrganizationConfiguration
type GuarddutyOrganizationConfigurationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GuarddutyOrganizationConfigurationObservation `json:"atProvider"`
}

// A GuarddutyOrganizationConfigurationObservation records the observed state of a GuarddutyOrganizationConfiguration
type GuarddutyOrganizationConfigurationObservation struct{}