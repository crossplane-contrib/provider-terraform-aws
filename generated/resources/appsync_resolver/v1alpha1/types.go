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

// AppsyncResolver is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppsyncResolver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppsyncResolverSpec   `json:"spec"`
	Status AppsyncResolverStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppsyncResolver contains a list of AppsyncResolverList
type AppsyncResolverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppsyncResolver `json:"items"`
}

// A AppsyncResolverSpec defines the desired state of a AppsyncResolver
type AppsyncResolverSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppsyncResolverParameters `json:",inline"`
}

// A AppsyncResolverParameters defines the desired state of a AppsyncResolver
type AppsyncResolverParameters struct {
	Kind             string         `json:"kind"`
	ResponseTemplate string         `json:"response_template"`
	Id               string         `json:"id"`
	RequestTemplate  string         `json:"request_template"`
	Type             string         `json:"type"`
	ApiId            string         `json:"api_id"`
	DataSource       string         `json:"data_source"`
	Field            string         `json:"field"`
	CachingConfig    CachingConfig  `json:"caching_config"`
	PipelineConfig   PipelineConfig `json:"pipeline_config"`
}

type CachingConfig struct {
	CachingKeys []string `json:"caching_keys"`
	Ttl         int64    `json:"ttl"`
}

type PipelineConfig struct {
	Functions []string `json:"functions"`
}

// A AppsyncResolverStatus defines the observed state of a AppsyncResolver
type AppsyncResolverStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppsyncResolverObservation `json:",inline"`
}

// A AppsyncResolverObservation records the observed state of a AppsyncResolver
type AppsyncResolverObservation struct {
	Arn string `json:"arn"`
}