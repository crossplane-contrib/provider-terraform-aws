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

// Codepipeline is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Codepipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodepipelineSpec   `json:"spec"`
	Status CodepipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Codepipeline contains a list of CodepipelineList
type CodepipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Codepipeline `json:"items"`
}

// A CodepipelineSpec defines the desired state of a Codepipeline
type CodepipelineSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodepipelineParameters `json:",inline"`
}

// A CodepipelineParameters defines the desired state of a Codepipeline
type CodepipelineParameters struct {
	Name          string            `json:"name"`
	RoleArn       string            `json:"role_arn"`
	Tags          map[string]string `json:"tags"`
	Id            string            `json:"id"`
	ArtifactStore []ArtifactStore   `json:"artifact_store"`
	Stage         []Stage           `json:"stage"`
}

type ArtifactStore struct {
	Location      string        `json:"location"`
	Region        string        `json:"region"`
	Type          string        `json:"type"`
	EncryptionKey EncryptionKey `json:"encryption_key"`
}

type EncryptionKey struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Stage struct {
	Name   string   `json:"name"`
	Action []Action `json:"action"`
}

type Action struct {
	Region          string            `json:"region"`
	RoleArn         string            `json:"role_arn"`
	Version         string            `json:"version"`
	Category        string            `json:"category"`
	InputArtifacts  []string          `json:"input_artifacts"`
	Name            string            `json:"name"`
	OutputArtifacts []string          `json:"output_artifacts"`
	Owner           string            `json:"owner"`
	Configuration   map[string]string `json:"configuration"`
	Namespace       string            `json:"namespace"`
	Provider        string            `json:"provider"`
	RunOrder        int64             `json:"run_order"`
}

// A CodepipelineStatus defines the observed state of a Codepipeline
type CodepipelineStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodepipelineObservation `json:",inline"`
}

// A CodepipelineObservation records the observed state of a Codepipeline
type CodepipelineObservation struct {
	Arn string `json:"arn"`
}