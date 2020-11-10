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

// CognitoUserPoolDomain is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CognitoUserPoolDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CognitoUserPoolDomainSpec   `json:"spec"`
	Status CognitoUserPoolDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPoolDomain contains a list of CognitoUserPoolDomainList
type CognitoUserPoolDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoUserPoolDomain `json:"items"`
}

// A CognitoUserPoolDomainSpec defines the desired state of a CognitoUserPoolDomain
type CognitoUserPoolDomainSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CognitoUserPoolDomainParameters `json:",inline"`
}

// A CognitoUserPoolDomainParameters defines the desired state of a CognitoUserPoolDomain
type CognitoUserPoolDomainParameters struct {
	CertificateArn string `json:"certificate_arn"`
	Domain         string `json:"domain"`
	UserPoolId     string `json:"user_pool_id"`
}

// A CognitoUserPoolDomainStatus defines the observed state of a CognitoUserPoolDomain
type CognitoUserPoolDomainStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CognitoUserPoolDomainObservation `json:",inline"`
}

// A CognitoUserPoolDomainObservation records the observed state of a CognitoUserPoolDomain
type CognitoUserPoolDomainObservation struct {
	Version                   string `json:"version"`
	AwsAccountId              string `json:"aws_account_id"`
	CloudfrontDistributionArn string `json:"cloudfront_distribution_arn"`
	Id                        string `json:"id"`
	S3Bucket                  string `json:"s3_bucket"`
}