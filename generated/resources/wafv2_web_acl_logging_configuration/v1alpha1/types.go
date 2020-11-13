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

// Wafv2WebAclLoggingConfiguration is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Wafv2WebAclLoggingConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Wafv2WebAclLoggingConfigurationSpec   `json:"spec"`
	Status Wafv2WebAclLoggingConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2WebAclLoggingConfiguration contains a list of Wafv2WebAclLoggingConfigurationList
type Wafv2WebAclLoggingConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wafv2WebAclLoggingConfiguration `json:"items"`
}

// A Wafv2WebAclLoggingConfigurationSpec defines the desired state of a Wafv2WebAclLoggingConfiguration
type Wafv2WebAclLoggingConfigurationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Wafv2WebAclLoggingConfigurationParameters `json:",inline"`
}

// A Wafv2WebAclLoggingConfigurationParameters defines the desired state of a Wafv2WebAclLoggingConfiguration
type Wafv2WebAclLoggingConfigurationParameters struct {
	LogDestinationConfigs []string         `json:"log_destination_configs"`
	ResourceArn           string           `json:"resource_arn"`
	Id                    string           `json:"id"`
	RedactedFields        []RedactedFields `json:"redacted_fields"`
}

type RedactedFields struct {
	QueryString         QueryString         `json:"query_string"`
	SingleHeader        SingleHeader        `json:"single_header"`
	SingleQueryArgument SingleQueryArgument `json:"single_query_argument"`
	UriPath             UriPath             `json:"uri_path"`
	AllQueryArguments   AllQueryArguments   `json:"all_query_arguments"`
	Body                Body                `json:"body"`
	Method              Method              `json:"method"`
}

type QueryString struct{}

type SingleHeader struct {
	Name string `json:"name"`
}

type SingleQueryArgument struct {
	Name string `json:"name"`
}

type UriPath struct{}

type AllQueryArguments struct{}

type Body struct{}

type Method struct{}

// A Wafv2WebAclLoggingConfigurationStatus defines the observed state of a Wafv2WebAclLoggingConfiguration
type Wafv2WebAclLoggingConfigurationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Wafv2WebAclLoggingConfigurationObservation `json:",inline"`
}

// A Wafv2WebAclLoggingConfigurationObservation records the observed state of a Wafv2WebAclLoggingConfiguration
type Wafv2WebAclLoggingConfigurationObservation struct{}