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

// ServiceDiscoveryPublicDnsNamespace is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ServiceDiscoveryPublicDnsNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceDiscoveryPublicDnsNamespaceSpec   `json:"spec"`
	Status ServiceDiscoveryPublicDnsNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceDiscoveryPublicDnsNamespace contains a list of ServiceDiscoveryPublicDnsNamespaceList
type ServiceDiscoveryPublicDnsNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDiscoveryPublicDnsNamespace `json:"items"`
}

// A ServiceDiscoveryPublicDnsNamespaceSpec defines the desired state of a ServiceDiscoveryPublicDnsNamespace
type ServiceDiscoveryPublicDnsNamespaceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ServiceDiscoveryPublicDnsNamespaceParameters `json:",inline"`
}

// A ServiceDiscoveryPublicDnsNamespaceParameters defines the desired state of a ServiceDiscoveryPublicDnsNamespace
type ServiceDiscoveryPublicDnsNamespaceParameters struct {
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
	Description string            `json:"description"`
}

// A ServiceDiscoveryPublicDnsNamespaceStatus defines the observed state of a ServiceDiscoveryPublicDnsNamespace
type ServiceDiscoveryPublicDnsNamespaceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ServiceDiscoveryPublicDnsNamespaceObservation `json:",inline"`
}

// A ServiceDiscoveryPublicDnsNamespaceObservation records the observed state of a ServiceDiscoveryPublicDnsNamespace
type ServiceDiscoveryPublicDnsNamespaceObservation struct {
	HostedZone string `json:"hosted_zone"`
	Id         string `json:"id"`
	Arn        string `json:"arn"`
}