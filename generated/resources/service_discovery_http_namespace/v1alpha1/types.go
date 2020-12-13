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

// ServiceDiscoveryHttpNamespace is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ServiceDiscoveryHttpNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceDiscoveryHttpNamespaceSpec   `json:"spec"`
	Status ServiceDiscoveryHttpNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceDiscoveryHttpNamespace contains a list of ServiceDiscoveryHttpNamespaceList
type ServiceDiscoveryHttpNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDiscoveryHttpNamespace `json:"items"`
}

// A ServiceDiscoveryHttpNamespaceSpec defines the desired state of a ServiceDiscoveryHttpNamespace
type ServiceDiscoveryHttpNamespaceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ServiceDiscoveryHttpNamespaceParameters `json:"forProvider"`
}

// A ServiceDiscoveryHttpNamespaceParameters defines the desired state of a ServiceDiscoveryHttpNamespace
type ServiceDiscoveryHttpNamespaceParameters struct {
	Description string            `json:"description"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
}

// A ServiceDiscoveryHttpNamespaceStatus defines the observed state of a ServiceDiscoveryHttpNamespace
type ServiceDiscoveryHttpNamespaceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ServiceDiscoveryHttpNamespaceObservation `json:"atProvider"`
}

// A ServiceDiscoveryHttpNamespaceObservation records the observed state of a ServiceDiscoveryHttpNamespace
type ServiceDiscoveryHttpNamespaceObservation struct {
	Arn string `json:"arn"`
}