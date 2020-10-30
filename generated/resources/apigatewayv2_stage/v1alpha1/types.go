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

// Apigatewayv2Stage is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Apigatewayv2Stage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Apigatewayv2StageSpec   `json:"spec"`
	Status Apigatewayv2StageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2Stage contains a list of Apigatewayv2StageList
type Apigatewayv2StageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2Stage `json:"items"`
}

// A Apigatewayv2StageSpec defines the desired state of a Apigatewayv2Stage
type Apigatewayv2StageSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Apigatewayv2StageParameters `json:",inline"`
}

// A Apigatewayv2StageParameters defines the desired state of a Apigatewayv2Stage
type Apigatewayv2StageParameters struct {
	Description         string `json:"description"`
	Name                string `json:"name"`
	ApiId               string `json:"api_id"`
	AutoDeploy          bool   `json:"auto_deploy"`
	ClientCertificateId string `json:"client_certificate_id"`
}

// A Apigatewayv2StageStatus defines the observed state of a Apigatewayv2Stage
type Apigatewayv2StageStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Apigatewayv2StageObservation `json:",inline"`
}

// A Apigatewayv2StageObservation records the observed state of a Apigatewayv2Stage
type Apigatewayv2StageObservation struct {
	DeploymentId string `json:"deployment_id"`
	ExecutionArn string `json:"execution_arn"`
	Id           string `json:"id"`
	InvokeUrl    string `json:"invoke_url"`
	Arn          string `json:"arn"`
}