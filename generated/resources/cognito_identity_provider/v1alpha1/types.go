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

// CognitoIdentityProvider is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CognitoIdentityProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CognitoIdentityProviderSpec   `json:"spec"`
	Status CognitoIdentityProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoIdentityProvider contains a list of CognitoIdentityProviderList
type CognitoIdentityProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoIdentityProvider `json:"items"`
}

// A CognitoIdentityProviderSpec defines the desired state of a CognitoIdentityProvider
type CognitoIdentityProviderSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CognitoIdentityProviderParameters `json:"forProvider"`
}

// A CognitoIdentityProviderParameters defines the desired state of a CognitoIdentityProvider
type CognitoIdentityProviderParameters struct {
	ProviderName     string            `json:"provider_name"`
	ProviderType     string            `json:"provider_type"`
	UserPoolId       string            `json:"user_pool_id"`
	AttributeMapping map[string]string `json:"attribute_mapping"`
	Id               string            `json:"id"`
	IdpIdentifiers   []string          `json:"idp_identifiers"`
	ProviderDetails  map[string]string `json:"provider_details"`
}

// A CognitoIdentityProviderStatus defines the observed state of a CognitoIdentityProvider
type CognitoIdentityProviderStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CognitoIdentityProviderObservation `json:"atProvider"`
}

// A CognitoIdentityProviderObservation records the observed state of a CognitoIdentityProvider
type CognitoIdentityProviderObservation struct{}