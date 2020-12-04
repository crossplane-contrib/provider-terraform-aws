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

// IotCertificate is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IotCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IotCertificateSpec   `json:"spec"`
	Status IotCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IotCertificate contains a list of IotCertificateList
type IotCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IotCertificate `json:"items"`
}

// A IotCertificateSpec defines the desired state of a IotCertificate
type IotCertificateSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IotCertificateParameters `json:",inline"`
}

// A IotCertificateParameters defines the desired state of a IotCertificate
type IotCertificateParameters struct {
	Csr    string `json:"csr"`
	Id     string `json:"id"`
	Active bool   `json:"active"`
}

// A IotCertificateStatus defines the observed state of a IotCertificate
type IotCertificateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IotCertificateObservation `json:",inline"`
}

// A IotCertificateObservation records the observed state of a IotCertificate
type IotCertificateObservation struct {
	PrivateKey     string `json:"private_key"`
	PublicKey      string `json:"public_key"`
	Arn            string `json:"arn"`
	CertificatePem string `json:"certificate_pem"`
}