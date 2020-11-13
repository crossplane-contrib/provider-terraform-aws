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

// SesEventDestination is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SesEventDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SesEventDestinationSpec   `json:"spec"`
	Status SesEventDestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesEventDestination contains a list of SesEventDestinationList
type SesEventDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesEventDestination `json:"items"`
}

// A SesEventDestinationSpec defines the desired state of a SesEventDestination
type SesEventDestinationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SesEventDestinationParameters `json:",inline"`
}

// A SesEventDestinationParameters defines the desired state of a SesEventDestination
type SesEventDestinationParameters struct {
	ConfigurationSetName  string                  `json:"configuration_set_name"`
	Enabled               bool                    `json:"enabled"`
	Id                    string                  `json:"id"`
	MatchingTypes         []string                `json:"matching_types"`
	Name                  string                  `json:"name"`
	CloudwatchDestination []CloudwatchDestination `json:"cloudwatch_destination"`
	KinesisDestination    KinesisDestination      `json:"kinesis_destination"`
	SnsDestination        SnsDestination          `json:"sns_destination"`
}

type CloudwatchDestination struct {
	DefaultValue  string `json:"default_value"`
	DimensionName string `json:"dimension_name"`
	ValueSource   string `json:"value_source"`
}

type KinesisDestination struct {
	RoleArn   string `json:"role_arn"`
	StreamArn string `json:"stream_arn"`
}

type SnsDestination struct {
	TopicArn string `json:"topic_arn"`
}

// A SesEventDestinationStatus defines the observed state of a SesEventDestination
type SesEventDestinationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SesEventDestinationObservation `json:",inline"`
}

// A SesEventDestinationObservation records the observed state of a SesEventDestination
type SesEventDestinationObservation struct{}