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

// AcmCertificate is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AcmCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AcmCertificateSpec   `json:"spec"`
	Status AcmCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AcmCertificate contains a list of AcmCertificateList
type AcmCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AcmCertificate `json:"items"`
}

// A AcmCertificateSpec defines the desired state of a AcmCertificate
type AcmCertificateSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AcmCertificateParameters `json:",inline"`
}

// A AcmCertificateParameters defines the desired state of a AcmCertificate
type AcmCertificateParameters struct {
	DomainName              string            `json:"domain_name"`
	ValidationMethod        string            `json:"validation_method"`
	CertificateAuthorityArn string            `json:"certificate_authority_arn"`
	CertificateChain        string            `json:"certificate_chain"`
	Id                      string            `json:"id"`
	PrivateKey              string            `json:"private_key"`
	SubjectAlternativeNames []string          `json:"subject_alternative_names"`
	Tags                    map[string]string `json:"tags"`
	CertificateBody         string            `json:"certificate_body"`
	Options                 Options           `json:"options"`
}

type Options struct {
	CertificateTransparencyLoggingPreference string `json:"certificate_transparency_logging_preference"`
}

// A AcmCertificateStatus defines the observed state of a AcmCertificate
type AcmCertificateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AcmCertificateObservation `json:",inline"`
}

// A AcmCertificateObservation records the observed state of a AcmCertificate
type AcmCertificateObservation struct {
	DomainValidationOptions []DomainValidationOptions `json:"domain_validation_options"`
	Status                  string                    `json:"status"`
	ValidationEmails        []string                  `json:"validation_emails"`
	Arn                     string                    `json:"arn"`
}

type DomainValidationOptions struct {
	DomainName          string `json:"domain_name"`
	ResourceRecordName  string `json:"resource_record_name"`
	ResourceRecordType  string `json:"resource_record_type"`
	ResourceRecordValue string `json:"resource_record_value"`
}