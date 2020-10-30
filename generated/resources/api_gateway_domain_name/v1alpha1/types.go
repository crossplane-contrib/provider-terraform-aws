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

// ApiGatewayDomainName is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayDomainName struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayDomainNameSpec   `json:"spec"`
	Status ApiGatewayDomainNameStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayDomainName contains a list of ApiGatewayDomainNameList
type ApiGatewayDomainNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayDomainName `json:"items"`
}

// A ApiGatewayDomainNameSpec defines the desired state of a ApiGatewayDomainName
type ApiGatewayDomainNameSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayDomainNameParameters `json:",inline"`
}

// A ApiGatewayDomainNameParameters defines the desired state of a ApiGatewayDomainName
type ApiGatewayDomainNameParameters struct {
	CertificateName         string `json:"certificate_name"`
	DomainName              string `json:"domain_name"`
	CertificateArn          string `json:"certificate_arn"`
	CertificateBody         string `json:"certificate_body"`
	RegionalCertificateArn  string `json:"regional_certificate_arn"`
	CertificateChain        string `json:"certificate_chain"`
	RegionalCertificateName string `json:"regional_certificate_name"`
	CertificatePrivateKey   string `json:"certificate_private_key"`
}

// A ApiGatewayDomainNameStatus defines the observed state of a ApiGatewayDomainName
type ApiGatewayDomainNameStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayDomainNameObservation `json:",inline"`
}

// A ApiGatewayDomainNameObservation records the observed state of a ApiGatewayDomainName
type ApiGatewayDomainNameObservation struct {
	Arn                   string `json:"arn"`
	CloudfrontDomainName  string `json:"cloudfront_domain_name"`
	Id                    string `json:"id"`
	SecurityPolicy        string `json:"security_policy"`
	RegionalDomainName    string `json:"regional_domain_name"`
	CertificateUploadDate string `json:"certificate_upload_date"`
	CloudfrontZoneId      string `json:"cloudfront_zone_id"`
	RegionalZoneId        string `json:"regional_zone_id"`
}