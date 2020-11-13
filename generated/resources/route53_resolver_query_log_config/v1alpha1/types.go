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

// Route53ResolverQueryLogConfig is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Route53ResolverQueryLogConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Route53ResolverQueryLogConfigSpec   `json:"spec"`
	Status Route53ResolverQueryLogConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverQueryLogConfig contains a list of Route53ResolverQueryLogConfigList
type Route53ResolverQueryLogConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ResolverQueryLogConfig `json:"items"`
}

// A Route53ResolverQueryLogConfigSpec defines the desired state of a Route53ResolverQueryLogConfig
type Route53ResolverQueryLogConfigSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Route53ResolverQueryLogConfigParameters `json:",inline"`
}

// A Route53ResolverQueryLogConfigParameters defines the desired state of a Route53ResolverQueryLogConfig
type Route53ResolverQueryLogConfigParameters struct {
	Name           string            `json:"name"`
	Tags           map[string]string `json:"tags"`
	DestinationArn string            `json:"destination_arn"`
	Id             string            `json:"id"`
}

// A Route53ResolverQueryLogConfigStatus defines the observed state of a Route53ResolverQueryLogConfig
type Route53ResolverQueryLogConfigStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Route53ResolverQueryLogConfigObservation `json:",inline"`
}

// A Route53ResolverQueryLogConfigObservation records the observed state of a Route53ResolverQueryLogConfig
type Route53ResolverQueryLogConfigObservation struct {
	OwnerId     string `json:"owner_id"`
	ShareStatus string `json:"share_status"`
	Arn         string `json:"arn"`
}