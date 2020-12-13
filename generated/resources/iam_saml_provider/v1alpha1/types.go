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

// IamSamlProvider is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamSamlProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamSamlProviderSpec   `json:"spec"`
	Status IamSamlProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamSamlProvider contains a list of IamSamlProviderList
type IamSamlProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamSamlProvider `json:"items"`
}

// A IamSamlProviderSpec defines the desired state of a IamSamlProvider
type IamSamlProviderSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamSamlProviderParameters `json:"forProvider"`
}

// A IamSamlProviderParameters defines the desired state of a IamSamlProvider
type IamSamlProviderParameters struct {
	Name                 string `json:"name"`
	SamlMetadataDocument string `json:"saml_metadata_document"`
	Id                   string `json:"id"`
}

// A IamSamlProviderStatus defines the observed state of a IamSamlProvider
type IamSamlProviderStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamSamlProviderObservation `json:"atProvider"`
}

// A IamSamlProviderObservation records the observed state of a IamSamlProvider
type IamSamlProviderObservation struct {
	ValidUntil string `json:"valid_until"`
	Arn        string `json:"arn"`
}