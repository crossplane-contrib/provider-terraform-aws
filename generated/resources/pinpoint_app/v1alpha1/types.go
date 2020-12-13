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

// PinpointApp is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type PinpointApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PinpointAppSpec   `json:"spec"`
	Status PinpointAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointApp contains a list of PinpointAppList
type PinpointAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointApp `json:"items"`
}

// A PinpointAppSpec defines the desired state of a PinpointApp
type PinpointAppSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  PinpointAppParameters `json:"forProvider"`
}

// A PinpointAppParameters defines the desired state of a PinpointApp
type PinpointAppParameters struct {
	Id           string            `json:"id"`
	Name         string            `json:"name"`
	NamePrefix   string            `json:"name_prefix"`
	Tags         map[string]string `json:"tags"`
	CampaignHook CampaignHook      `json:"campaign_hook"`
	Limits       Limits            `json:"limits"`
	QuietTime    QuietTime         `json:"quiet_time"`
}

type CampaignHook struct {
	LambdaFunctionName string `json:"lambda_function_name"`
	Mode               string `json:"mode"`
	WebUrl             string `json:"web_url"`
}

type Limits struct {
	MessagesPerSecond int64 `json:"messages_per_second"`
	Total             int64 `json:"total"`
	Daily             int64 `json:"daily"`
	MaximumDuration   int64 `json:"maximum_duration"`
}

type QuietTime struct {
	End   string `json:"end"`
	Start string `json:"start"`
}

// A PinpointAppStatus defines the observed state of a PinpointApp
type PinpointAppStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     PinpointAppObservation `json:"atProvider"`
}

// A PinpointAppObservation records the observed state of a PinpointApp
type PinpointAppObservation struct {
	ApplicationId string `json:"application_id"`
	Arn           string `json:"arn"`
}