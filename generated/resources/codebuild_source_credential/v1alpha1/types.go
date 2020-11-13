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

// CodebuildSourceCredential is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodebuildSourceCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodebuildSourceCredentialSpec   `json:"spec"`
	Status CodebuildSourceCredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodebuildSourceCredential contains a list of CodebuildSourceCredentialList
type CodebuildSourceCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodebuildSourceCredential `json:"items"`
}

// A CodebuildSourceCredentialSpec defines the desired state of a CodebuildSourceCredential
type CodebuildSourceCredentialSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodebuildSourceCredentialParameters `json:",inline"`
}

// A CodebuildSourceCredentialParameters defines the desired state of a CodebuildSourceCredential
type CodebuildSourceCredentialParameters struct {
	AuthType   string `json:"auth_type"`
	Id         string `json:"id"`
	ServerType string `json:"server_type"`
	Token      string `json:"token"`
	UserName   string `json:"user_name"`
}

// A CodebuildSourceCredentialStatus defines the observed state of a CodebuildSourceCredential
type CodebuildSourceCredentialStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodebuildSourceCredentialObservation `json:",inline"`
}

// A CodebuildSourceCredentialObservation records the observed state of a CodebuildSourceCredential
type CodebuildSourceCredentialObservation struct {
	Arn string `json:"arn"`
}