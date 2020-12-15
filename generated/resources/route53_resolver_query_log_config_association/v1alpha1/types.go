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

// Route53ResolverQueryLogConfigAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Route53ResolverQueryLogConfigAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Route53ResolverQueryLogConfigAssociationSpec   `json:"spec"`
	Status Route53ResolverQueryLogConfigAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverQueryLogConfigAssociation contains a list of Route53ResolverQueryLogConfigAssociationList
type Route53ResolverQueryLogConfigAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ResolverQueryLogConfigAssociation `json:"items"`
}

// A Route53ResolverQueryLogConfigAssociationSpec defines the desired state of a Route53ResolverQueryLogConfigAssociation
type Route53ResolverQueryLogConfigAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Route53ResolverQueryLogConfigAssociationParameters `json:"forProvider"`
}

// A Route53ResolverQueryLogConfigAssociationParameters defines the desired state of a Route53ResolverQueryLogConfigAssociation
type Route53ResolverQueryLogConfigAssociationParameters struct {
	ResolverQueryLogConfigId string `json:"resolver_query_log_config_id"`
	ResourceId               string `json:"resource_id"`
}

// A Route53ResolverQueryLogConfigAssociationStatus defines the observed state of a Route53ResolverQueryLogConfigAssociation
type Route53ResolverQueryLogConfigAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Route53ResolverQueryLogConfigAssociationObservation `json:"atProvider"`
}

// A Route53ResolverQueryLogConfigAssociationObservation records the observed state of a Route53ResolverQueryLogConfigAssociation
type Route53ResolverQueryLogConfigAssociationObservation struct{}