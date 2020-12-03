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

// DbOptionGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DbOptionGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbOptionGroupSpec   `json:"spec"`
	Status DbOptionGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbOptionGroup contains a list of DbOptionGroupList
type DbOptionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbOptionGroup `json:"items"`
}

// A DbOptionGroupSpec defines the desired state of a DbOptionGroup
type DbOptionGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DbOptionGroupParameters `json:",inline"`
}

// A DbOptionGroupParameters defines the desired state of a DbOptionGroup
type DbOptionGroupParameters struct {
	Id                     string            `json:"id"`
	MajorEngineVersion     string            `json:"major_engine_version"`
	Name                   string            `json:"name"`
	NamePrefix             string            `json:"name_prefix"`
	OptionGroupDescription string            `json:"option_group_description"`
	Tags                   map[string]string `json:"tags"`
	EngineName             string            `json:"engine_name"`
	Option                 []Option          `json:"option"`
	Timeouts               []Timeouts        `json:"timeouts"`
}

type Option struct {
	DbSecurityGroupMemberships  []string         `json:"db_security_group_memberships"`
	OptionName                  string           `json:"option_name"`
	Port                        int              `json:"port"`
	Version                     string           `json:"version"`
	VpcSecurityGroupMemberships []string         `json:"vpc_security_group_memberships"`
	OptionSettings              []OptionSettings `json:"option_settings"`
}

type OptionSettings struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Timeouts struct {
	Delete string `json:"delete"`
}

// A DbOptionGroupStatus defines the observed state of a DbOptionGroup
type DbOptionGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DbOptionGroupObservation `json:",inline"`
}

// A DbOptionGroupObservation records the observed state of a DbOptionGroup
type DbOptionGroupObservation struct {
	Arn string `json:"arn"`
}