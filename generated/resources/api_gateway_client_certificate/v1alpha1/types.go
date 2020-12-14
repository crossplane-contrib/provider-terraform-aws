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

// ApiGatewayClientCertificate is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayClientCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayClientCertificateSpec   `json:"spec"`
	Status ApiGatewayClientCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayClientCertificate contains a list of ApiGatewayClientCertificateList
type ApiGatewayClientCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayClientCertificate `json:"items"`
}

// A ApiGatewayClientCertificateSpec defines the desired state of a ApiGatewayClientCertificate
type ApiGatewayClientCertificateSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayClientCertificateParameters `json:"forProvider"`
}

// A ApiGatewayClientCertificateParameters defines the desired state of a ApiGatewayClientCertificate
type ApiGatewayClientCertificateParameters struct {
	Id          string            `json:"id"`
	Tags        map[string]string `json:"tags"`
	Description string            `json:"description"`
}

// A ApiGatewayClientCertificateStatus defines the observed state of a ApiGatewayClientCertificate
type ApiGatewayClientCertificateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayClientCertificateObservation `json:"atProvider"`
}

// A ApiGatewayClientCertificateObservation records the observed state of a ApiGatewayClientCertificate
type ApiGatewayClientCertificateObservation struct {
	PemEncodedCertificate string `json:"pem_encoded_certificate"`
	Arn                   string `json:"arn"`
	CreatedDate           string `json:"created_date"`
	ExpirationDate        string `json:"expiration_date"`
}