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

// AcmCertificateValidation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AcmCertificateValidation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AcmCertificateValidationSpec   `json:"spec"`
	Status AcmCertificateValidationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AcmCertificateValidation contains a list of AcmCertificateValidationList
type AcmCertificateValidationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AcmCertificateValidation `json:"items"`
}

// A AcmCertificateValidationSpec defines the desired state of a AcmCertificateValidation
type AcmCertificateValidationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AcmCertificateValidationParameters `json:",inline"`
}

// A AcmCertificateValidationParameters defines the desired state of a AcmCertificateValidation
type AcmCertificateValidationParameters struct {
	CertificateArn        string     `json:"certificate_arn"`
	Id                    string     `json:"id"`
	ValidationRecordFqdns []string   `json:"validation_record_fqdns"`
	Timeouts              []Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
}

// A AcmCertificateValidationStatus defines the observed state of a AcmCertificateValidation
type AcmCertificateValidationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AcmCertificateValidationObservation `json:",inline"`
}

// A AcmCertificateValidationObservation records the observed state of a AcmCertificateValidation
type AcmCertificateValidationObservation struct{}