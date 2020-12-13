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

// IamOpenidConnectProvider is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamOpenidConnectProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamOpenidConnectProviderSpec   `json:"spec"`
	Status IamOpenidConnectProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamOpenidConnectProvider contains a list of IamOpenidConnectProviderList
type IamOpenidConnectProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamOpenidConnectProvider `json:"items"`
}

// A IamOpenidConnectProviderSpec defines the desired state of a IamOpenidConnectProvider
type IamOpenidConnectProviderSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamOpenidConnectProviderParameters `json:"forProvider"`
}

// A IamOpenidConnectProviderParameters defines the desired state of a IamOpenidConnectProvider
type IamOpenidConnectProviderParameters struct {
	ThumbprintList []string `json:"thumbprint_list"`
	Url            string   `json:"url"`
	ClientIdList   []string `json:"client_id_list"`
	Id             string   `json:"id"`
}

// A IamOpenidConnectProviderStatus defines the observed state of a IamOpenidConnectProvider
type IamOpenidConnectProviderStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamOpenidConnectProviderObservation `json:"atProvider"`
}

// A IamOpenidConnectProviderObservation records the observed state of a IamOpenidConnectProvider
type IamOpenidConnectProviderObservation struct {
	Arn string `json:"arn"`
}