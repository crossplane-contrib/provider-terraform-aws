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

// Apigatewayv2Deployment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Apigatewayv2Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Apigatewayv2DeploymentSpec   `json:"spec"`
	Status Apigatewayv2DeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2Deployment contains a list of Apigatewayv2DeploymentList
type Apigatewayv2DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2Deployment `json:"items"`
}

// A Apigatewayv2DeploymentSpec defines the desired state of a Apigatewayv2Deployment
type Apigatewayv2DeploymentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Apigatewayv2DeploymentParameters `json:"forProvider"`
}

// A Apigatewayv2DeploymentParameters defines the desired state of a Apigatewayv2Deployment
type Apigatewayv2DeploymentParameters struct {
	ApiId       string            `json:"api_id"`
	Description string            `json:"description"`
	Triggers    map[string]string `json:"triggers"`
}

// A Apigatewayv2DeploymentStatus defines the observed state of a Apigatewayv2Deployment
type Apigatewayv2DeploymentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Apigatewayv2DeploymentObservation `json:"atProvider"`
}

// A Apigatewayv2DeploymentObservation records the observed state of a Apigatewayv2Deployment
type Apigatewayv2DeploymentObservation struct {
	AutoDeployed bool `json:"auto_deployed"`
}