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

// CodebuildWebhook is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodebuildWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodebuildWebhookSpec   `json:"spec"`
	Status CodebuildWebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodebuildWebhook contains a list of CodebuildWebhookList
type CodebuildWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodebuildWebhook `json:"items"`
}

// A CodebuildWebhookSpec defines the desired state of a CodebuildWebhook
type CodebuildWebhookSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodebuildWebhookParameters `json:",inline"`
}

// A CodebuildWebhookParameters defines the desired state of a CodebuildWebhook
type CodebuildWebhookParameters struct {
	ProjectName  string `json:"project_name"`
	BranchFilter string `json:"branch_filter"`
}

// A CodebuildWebhookStatus defines the observed state of a CodebuildWebhook
type CodebuildWebhookStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodebuildWebhookObservation `json:",inline"`
}

// A CodebuildWebhookObservation records the observed state of a CodebuildWebhook
type CodebuildWebhookObservation struct {
	PayloadUrl string `json:"payload_url"`
	Secret     string `json:"secret"`
	Url        string `json:"url"`
	Id         string `json:"id"`
}