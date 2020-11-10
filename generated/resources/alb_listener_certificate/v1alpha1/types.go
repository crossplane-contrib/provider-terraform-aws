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

// AlbListenerCertificate is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AlbListenerCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlbListenerCertificateSpec   `json:"spec"`
	Status AlbListenerCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlbListenerCertificate contains a list of AlbListenerCertificateList
type AlbListenerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlbListenerCertificate `json:"items"`
}

// A AlbListenerCertificateSpec defines the desired state of a AlbListenerCertificate
type AlbListenerCertificateSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AlbListenerCertificateParameters `json:",inline"`
}

// A AlbListenerCertificateParameters defines the desired state of a AlbListenerCertificate
type AlbListenerCertificateParameters struct {
	CertificateArn string `json:"certificate_arn"`
	ListenerArn    string `json:"listener_arn"`
}

// A AlbListenerCertificateStatus defines the observed state of a AlbListenerCertificate
type AlbListenerCertificateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AlbListenerCertificateObservation `json:",inline"`
}

// A AlbListenerCertificateObservation records the observed state of a AlbListenerCertificate
type AlbListenerCertificateObservation struct {
	Id string `json:"id"`
}