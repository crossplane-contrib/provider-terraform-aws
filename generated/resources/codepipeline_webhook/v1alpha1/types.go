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

// CodepipelineWebhook is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodepipelineWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodepipelineWebhookSpec   `json:"spec"`
	Status CodepipelineWebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodepipelineWebhook contains a list of CodepipelineWebhookList
type CodepipelineWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodepipelineWebhook `json:"items"`
}

// A CodepipelineWebhookSpec defines the desired state of a CodepipelineWebhook
type CodepipelineWebhookSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodepipelineWebhookParameters `json:",inline"`
}

// A CodepipelineWebhookParameters defines the desired state of a CodepipelineWebhook
type CodepipelineWebhookParameters struct {
	Authentication              string                      `json:"authentication"`
	Id                          string                      `json:"id"`
	Name                        string                      `json:"name"`
	Tags                        map[string]string           `json:"tags"`
	TargetAction                string                      `json:"target_action"`
	TargetPipeline              string                      `json:"target_pipeline"`
	AuthenticationConfiguration AuthenticationConfiguration `json:"authentication_configuration"`
	Filter                      []Filter                    `json:"filter"`
}

type AuthenticationConfiguration struct {
	AllowedIpRange string `json:"allowed_ip_range"`
	SecretToken    string `json:"secret_token"`
}

type Filter struct {
	JsonPath    string `json:"json_path"`
	MatchEquals string `json:"match_equals"`
}

// A CodepipelineWebhookStatus defines the observed state of a CodepipelineWebhook
type CodepipelineWebhookStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodepipelineWebhookObservation `json:",inline"`
}

// A CodepipelineWebhookObservation records the observed state of a CodepipelineWebhook
type CodepipelineWebhookObservation struct {
	Url string `json:"url"`
}