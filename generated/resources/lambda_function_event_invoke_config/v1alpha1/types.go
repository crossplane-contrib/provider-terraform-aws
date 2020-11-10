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

// LambdaFunctionEventInvokeConfig is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LambdaFunctionEventInvokeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LambdaFunctionEventInvokeConfigSpec   `json:"spec"`
	Status LambdaFunctionEventInvokeConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaFunctionEventInvokeConfig contains a list of LambdaFunctionEventInvokeConfigList
type LambdaFunctionEventInvokeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaFunctionEventInvokeConfig `json:"items"`
}

// A LambdaFunctionEventInvokeConfigSpec defines the desired state of a LambdaFunctionEventInvokeConfig
type LambdaFunctionEventInvokeConfigSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LambdaFunctionEventInvokeConfigParameters `json:",inline"`
}

// A LambdaFunctionEventInvokeConfigParameters defines the desired state of a LambdaFunctionEventInvokeConfig
type LambdaFunctionEventInvokeConfigParameters struct {
	MaximumEventAgeInSeconds int               `json:"maximum_event_age_in_seconds"`
	MaximumRetryAttempts     int               `json:"maximum_retry_attempts"`
	Qualifier                string            `json:"qualifier"`
	FunctionName             string            `json:"function_name"`
	DestinationConfig        DestinationConfig `json:"destination_config"`
}

type DestinationConfig struct {
	OnFailure OnFailure `json:"on_failure"`
	OnSuccess OnSuccess `json:"on_success"`
}

type OnFailure struct {
	Destination string `json:"destination"`
}

type OnSuccess struct {
	Destination string `json:"destination"`
}

// A LambdaFunctionEventInvokeConfigStatus defines the observed state of a LambdaFunctionEventInvokeConfig
type LambdaFunctionEventInvokeConfigStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LambdaFunctionEventInvokeConfigObservation `json:",inline"`
}

// A LambdaFunctionEventInvokeConfigObservation records the observed state of a LambdaFunctionEventInvokeConfig
type LambdaFunctionEventInvokeConfigObservation struct {
	Id string `json:"id"`
}