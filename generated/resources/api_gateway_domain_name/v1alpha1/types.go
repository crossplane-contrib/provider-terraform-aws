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
	CertificateChain        string                `json:"certificate_chain"`
	DomainName              string                `json:"domain_name"`
	Id                      string                `json:"id"`
	SecurityPolicy          string                `json:"security_policy"`
	Tags                    map[string]string     `json:"tags"`
	CertificateBody         string                `json:"certificate_body"`
	CertificatePrivateKey   string                `json:"certificate_private_key"`
	RegionalCertificateArn  string                `json:"regional_certificate_arn"`
	RegionalCertificateName string                `json:"regional_certificate_name"`
	CertificateName         string                `json:"certificate_name"`
	CertificateArn          string                `json:"certificate_arn"`
	EndpointConfiguration   EndpointConfiguration `json:"endpoint_configuration"`
}

type EndpointConfiguration struct {
	Types []string `json:"types"`
}

// A ApiGatewayDomainNameStatus defines the observed state of a ApiGatewayDomainName
type ApiGatewayDomainNameStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayDomainNameObservation `json:",inline"`
}

// A ApiGatewayDomainNameObservation records the observed state of a ApiGatewayDomainName
type ApiGatewayDomainNameObservation struct {
	CloudfrontZoneId      string `json:"cloudfront_zone_id"`
	RegionalDomainName    string `json:"regional_domain_name"`
	RegionalZoneId        string `json:"regional_zone_id"`
	CertificateUploadDate string `json:"certificate_upload_date"`
	CloudfrontDomainName  string `json:"cloudfront_domain_name"`
	Arn                   string `json:"arn"`
}