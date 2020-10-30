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

// ApiGatewayStage is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayStage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayStageSpec   `json:"spec"`
	Status ApiGatewayStageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayStage contains a list of ApiGatewayStageList
type ApiGatewayStageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayStage `json:"items"`
}

// A ApiGatewayStageSpec defines the desired state of a ApiGatewayStage
type ApiGatewayStageSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayStageParameters `json:",inline"`
}

// A ApiGatewayStageParameters defines the desired state of a ApiGatewayStage
type ApiGatewayStageParameters struct {
	DeploymentId         string `json:"deployment_id"`
	Description          string `json:"description"`
	XrayTracingEnabled   bool   `json:"xray_tracing_enabled"`
	CacheClusterSize     string `json:"cache_cluster_size"`
	DocumentationVersion string `json:"documentation_version"`
	RestApiId            string `json:"rest_api_id"`
	CacheClusterEnabled  bool   `json:"cache_cluster_enabled"`
	ClientCertificateId  string `json:"client_certificate_id"`
	StageName            string `json:"stage_name"`
}

// A ApiGatewayStageStatus defines the observed state of a ApiGatewayStage
type ApiGatewayStageStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayStageObservation `json:",inline"`
}

// A ApiGatewayStageObservation records the observed state of a ApiGatewayStage
type ApiGatewayStageObservation struct {
	Arn          string `json:"arn"`
	InvokeUrl    string `json:"invoke_url"`
	ExecutionArn string `json:"execution_arn"`
	Id           string `json:"id"`
}