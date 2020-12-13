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

// SesDomainIdentityVerification is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SesDomainIdentityVerification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SesDomainIdentityVerificationSpec   `json:"spec"`
	Status SesDomainIdentityVerificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesDomainIdentityVerification contains a list of SesDomainIdentityVerificationList
type SesDomainIdentityVerificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesDomainIdentityVerification `json:"items"`
}

// A SesDomainIdentityVerificationSpec defines the desired state of a SesDomainIdentityVerification
type SesDomainIdentityVerificationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SesDomainIdentityVerificationParameters `json:"forProvider"`
}

// A SesDomainIdentityVerificationParameters defines the desired state of a SesDomainIdentityVerification
type SesDomainIdentityVerificationParameters struct {
	Domain   string   `json:"domain"`
	Id       string   `json:"id"`
	Timeouts Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
}

// A SesDomainIdentityVerificationStatus defines the observed state of a SesDomainIdentityVerification
type SesDomainIdentityVerificationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SesDomainIdentityVerificationObservation `json:"atProvider"`
}

// A SesDomainIdentityVerificationObservation records the observed state of a SesDomainIdentityVerification
type SesDomainIdentityVerificationObservation struct {
	Arn string `json:"arn"`
}