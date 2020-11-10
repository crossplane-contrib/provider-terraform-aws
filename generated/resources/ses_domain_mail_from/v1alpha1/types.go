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

// SesDomainMailFrom is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SesDomainMailFrom struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SesDomainMailFromSpec   `json:"spec"`
	Status SesDomainMailFromStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesDomainMailFrom contains a list of SesDomainMailFromList
type SesDomainMailFromList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesDomainMailFrom `json:"items"`
}

// A SesDomainMailFromSpec defines the desired state of a SesDomainMailFrom
type SesDomainMailFromSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SesDomainMailFromParameters `json:",inline"`
}

// A SesDomainMailFromParameters defines the desired state of a SesDomainMailFrom
type SesDomainMailFromParameters struct {
	MailFromDomain      string `json:"mail_from_domain"`
	BehaviorOnMxFailure string `json:"behavior_on_mx_failure"`
	Domain              string `json:"domain"`
}

// A SesDomainMailFromStatus defines the observed state of a SesDomainMailFrom
type SesDomainMailFromStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SesDomainMailFromObservation `json:",inline"`
}

// A SesDomainMailFromObservation records the observed state of a SesDomainMailFrom
type SesDomainMailFromObservation struct {
	Id string `json:"id"`
}