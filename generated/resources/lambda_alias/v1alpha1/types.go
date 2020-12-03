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

// LambdaAlias is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LambdaAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LambdaAliasSpec   `json:"spec"`
	Status LambdaAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaAlias contains a list of LambdaAliasList
type LambdaAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaAlias `json:"items"`
}

// A LambdaAliasSpec defines the desired state of a LambdaAlias
type LambdaAliasSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LambdaAliasParameters `json:",inline"`
}

// A LambdaAliasParameters defines the desired state of a LambdaAlias
type LambdaAliasParameters struct {
	FunctionVersion string        `json:"function_version"`
	Id              string        `json:"id"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	FunctionName    string        `json:"function_name"`
	RoutingConfig   RoutingConfig `json:"routing_config"`
}

type RoutingConfig struct {
	AdditionalVersionWeights map[string]int `json:"additional_version_weights"`
}

// A LambdaAliasStatus defines the observed state of a LambdaAlias
type LambdaAliasStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LambdaAliasObservation `json:",inline"`
}

// A LambdaAliasObservation records the observed state of a LambdaAlias
type LambdaAliasObservation struct {
	InvokeArn string `json:"invoke_arn"`
	Arn       string `json:"arn"`
}