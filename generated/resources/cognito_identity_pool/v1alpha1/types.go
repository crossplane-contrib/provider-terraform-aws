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

// CognitoIdentityPool is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CognitoIdentityPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CognitoIdentityPoolSpec   `json:"spec"`
	Status CognitoIdentityPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoIdentityPool contains a list of CognitoIdentityPoolList
type CognitoIdentityPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoIdentityPool `json:"items"`
}

// A CognitoIdentityPoolSpec defines the desired state of a CognitoIdentityPool
type CognitoIdentityPoolSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CognitoIdentityPoolParameters `json:",inline"`
}

// A CognitoIdentityPoolParameters defines the desired state of a CognitoIdentityPool
type CognitoIdentityPoolParameters struct {
	Tags                           map[string]string        `json:"tags"`
	DeveloperProviderName          string                   `json:"developer_provider_name"`
	Id                             string                   `json:"id"`
	OpenidConnectProviderArns      []string                 `json:"openid_connect_provider_arns"`
	SamlProviderArns               []string                 `json:"saml_provider_arns"`
	SupportedLoginProviders        map[string]string        `json:"supported_login_providers"`
	AllowUnauthenticatedIdentities bool                     `json:"allow_unauthenticated_identities"`
	IdentityPoolName               string                   `json:"identity_pool_name"`
	CognitoIdentityProviders       CognitoIdentityProviders `json:"cognito_identity_providers"`
}

type CognitoIdentityProviders struct {
	ClientId             string `json:"client_id"`
	ProviderName         string `json:"provider_name"`
	ServerSideTokenCheck bool   `json:"server_side_token_check"`
}

// A CognitoIdentityPoolStatus defines the observed state of a CognitoIdentityPool
type CognitoIdentityPoolStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CognitoIdentityPoolObservation `json:",inline"`
}

// A CognitoIdentityPoolObservation records the observed state of a CognitoIdentityPool
type CognitoIdentityPoolObservation struct {
	Arn string `json:"arn"`
}