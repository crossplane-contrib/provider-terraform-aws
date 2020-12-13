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

// SesDomainDkim is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SesDomainDkim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SesDomainDkimSpec   `json:"spec"`
	Status SesDomainDkimStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesDomainDkim contains a list of SesDomainDkimList
type SesDomainDkimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesDomainDkim `json:"items"`
}

// A SesDomainDkimSpec defines the desired state of a SesDomainDkim
type SesDomainDkimSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SesDomainDkimParameters `json:"forProvider"`
}

// A SesDomainDkimParameters defines the desired state of a SesDomainDkim
type SesDomainDkimParameters struct {
	Domain string `json:"domain"`
	Id     string `json:"id"`
}

// A SesDomainDkimStatus defines the observed state of a SesDomainDkim
type SesDomainDkimStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SesDomainDkimObservation `json:"atProvider"`
}

// A SesDomainDkimObservation records the observed state of a SesDomainDkim
type SesDomainDkimObservation struct {
	DkimTokens []string `json:"dkim_tokens"`
}