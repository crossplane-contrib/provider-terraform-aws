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

// CognitoResourceServer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CognitoResourceServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CognitoResourceServerSpec   `json:"spec"`
	Status CognitoResourceServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoResourceServer contains a list of CognitoResourceServerList
type CognitoResourceServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoResourceServer `json:"items"`
}

// A CognitoResourceServerSpec defines the desired state of a CognitoResourceServer
type CognitoResourceServerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CognitoResourceServerParameters `json:",inline"`
}

// A CognitoResourceServerParameters defines the desired state of a CognitoResourceServer
type CognitoResourceServerParameters struct {
	Id         string  `json:"id"`
	Identifier string  `json:"identifier"`
	Name       string  `json:"name"`
	UserPoolId string  `json:"user_pool_id"`
	Scope      []Scope `json:"scope"`
}

type Scope struct {
	ScopeDescription string `json:"scope_description"`
	ScopeName        string `json:"scope_name"`
}

// A CognitoResourceServerStatus defines the observed state of a CognitoResourceServer
type CognitoResourceServerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CognitoResourceServerObservation `json:",inline"`
}

// A CognitoResourceServerObservation records the observed state of a CognitoResourceServer
type CognitoResourceServerObservation struct {
	ScopeIdentifiers []string `json:"scope_identifiers"`
}