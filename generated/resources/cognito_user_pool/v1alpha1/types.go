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

// CognitoUserPool is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CognitoUserPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CognitoUserPoolSpec   `json:"spec"`
	Status CognitoUserPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPool contains a list of CognitoUserPoolList
type CognitoUserPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoUserPool `json:"items"`
}

// A CognitoUserPoolSpec defines the desired state of a CognitoUserPool
type CognitoUserPoolSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CognitoUserPoolParameters `json:",inline"`
}

// A CognitoUserPoolParameters defines the desired state of a CognitoUserPool
type CognitoUserPoolParameters struct {
	AliasAttributes          []string `json:"alias_attributes"`
	Name                     string   `json:"name"`
	SmsAuthenticationMessage string   `json:"sms_authentication_message"`
	AutoVerifiedAttributes   []string `json:"auto_verified_attributes"`
	MfaConfiguration         string   `json:"mfa_configuration"`
	UsernameAttributes       []string `json:"username_attributes"`
}

// A CognitoUserPoolStatus defines the observed state of a CognitoUserPool
type CognitoUserPoolStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CognitoUserPoolObservation `json:",inline"`
}

// A CognitoUserPoolObservation records the observed state of a CognitoUserPool
type CognitoUserPoolObservation struct {
	Arn                      string `json:"arn"`
	CreationDate             string `json:"creation_date"`
	LastModifiedDate         string `json:"last_modified_date"`
	SmsVerificationMessage   string `json:"sms_verification_message"`
	Id                       string `json:"id"`
	EmailVerificationMessage string `json:"email_verification_message"`
	EmailVerificationSubject string `json:"email_verification_subject"`
	Endpoint                 string `json:"endpoint"`
}