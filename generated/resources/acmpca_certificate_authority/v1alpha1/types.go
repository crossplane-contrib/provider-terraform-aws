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

// AcmpcaCertificateAuthority is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AcmpcaCertificateAuthority struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AcmpcaCertificateAuthoritySpec   `json:"spec"`
	Status AcmpcaCertificateAuthorityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AcmpcaCertificateAuthority contains a list of AcmpcaCertificateAuthorityList
type AcmpcaCertificateAuthorityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AcmpcaCertificateAuthority `json:"items"`
}

// A AcmpcaCertificateAuthoritySpec defines the desired state of a AcmpcaCertificateAuthority
type AcmpcaCertificateAuthoritySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AcmpcaCertificateAuthorityParameters `json:",inline"`
}

// A AcmpcaCertificateAuthorityParameters defines the desired state of a AcmpcaCertificateAuthority
type AcmpcaCertificateAuthorityParameters struct {
	PermanentDeletionTimeInDays       int                               `json:"permanent_deletion_time_in_days"`
	Enabled                           bool                              `json:"enabled"`
	Tags                              map[string]string                 `json:"tags"`
	Type                              string                            `json:"type"`
	CertificateAuthorityConfiguration CertificateAuthorityConfiguration `json:"certificate_authority_configuration"`
	RevocationConfiguration           RevocationConfiguration           `json:"revocation_configuration"`
	Timeouts                          []Timeouts                        `json:"timeouts"`
}

type CertificateAuthorityConfiguration struct {
	KeyAlgorithm     string  `json:"key_algorithm"`
	SigningAlgorithm string  `json:"signing_algorithm"`
	Subject          Subject `json:"subject"`
}

type Subject struct {
	Initials                   string `json:"initials"`
	Organization               string `json:"organization"`
	OrganizationalUnit         string `json:"organizational_unit"`
	Pseudonym                  string `json:"pseudonym"`
	Title                      string `json:"title"`
	CommonName                 string `json:"common_name"`
	Country                    string `json:"country"`
	GivenName                  string `json:"given_name"`
	Locality                   string `json:"locality"`
	State                      string `json:"state"`
	Surname                    string `json:"surname"`
	DistinguishedNameQualifier string `json:"distinguished_name_qualifier"`
	GenerationQualifier        string `json:"generation_qualifier"`
}

type RevocationConfiguration struct {
	CrlConfiguration CrlConfiguration `json:"crl_configuration"`
}

type CrlConfiguration struct {
	CustomCname      string `json:"custom_cname"`
	Enabled          bool   `json:"enabled"`
	ExpirationInDays int    `json:"expiration_in_days"`
	S3BucketName     string `json:"s3_bucket_name"`
}

type Timeouts struct {
	Create string `json:"create"`
}

// A AcmpcaCertificateAuthorityStatus defines the observed state of a AcmpcaCertificateAuthority
type AcmpcaCertificateAuthorityStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AcmpcaCertificateAuthorityObservation `json:",inline"`
}

// A AcmpcaCertificateAuthorityObservation records the observed state of a AcmpcaCertificateAuthority
type AcmpcaCertificateAuthorityObservation struct {
	Status                    string `json:"status"`
	Arn                       string `json:"arn"`
	CertificateChain          string `json:"certificate_chain"`
	Id                        string `json:"id"`
	NotAfter                  string `json:"not_after"`
	NotBefore                 string `json:"not_before"`
	Certificate               string `json:"certificate"`
	CertificateSigningRequest string `json:"certificate_signing_request"`
	Serial                    string `json:"serial"`
}