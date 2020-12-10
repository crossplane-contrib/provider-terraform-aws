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

// DmsCertificate is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DmsCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DmsCertificateSpec   `json:"spec"`
	Status DmsCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DmsCertificate contains a list of DmsCertificateList
type DmsCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DmsCertificate `json:"items"`
}

// A DmsCertificateSpec defines the desired state of a DmsCertificate
type DmsCertificateSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DmsCertificateParameters `json:",inline"`
}

// A DmsCertificateParameters defines the desired state of a DmsCertificate
type DmsCertificateParameters struct {
	CertificateWallet string `json:"certificate_wallet"`
	Id                string `json:"id"`
	CertificateId     string `json:"certificate_id"`
	CertificatePem    string `json:"certificate_pem"`
}

// A DmsCertificateStatus defines the observed state of a DmsCertificate
type DmsCertificateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DmsCertificateObservation `json:",inline"`
}

// A DmsCertificateObservation records the observed state of a DmsCertificate
type DmsCertificateObservation struct {
	CertificateArn string `json:"certificate_arn"`
}