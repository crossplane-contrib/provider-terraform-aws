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
	Name                 string               `json:"name"`
	StageVariables       map[string]string    `json:"stage_variables"`
	AutoDeploy           bool                 `json:"auto_deploy"`
	DeploymentId         string               `json:"deployment_id"`
	Id                   string               `json:"id"`
	Tags                 map[string]string    `json:"tags"`
	ApiId                string               `json:"api_id"`
	ClientCertificateId  string               `json:"client_certificate_id"`
	Description          string               `json:"description"`
	AccessLogSettings    AccessLogSettings    `json:"access_log_settings"`
	DefaultRouteSettings DefaultRouteSettings `json:"default_route_settings"`
	RouteSettings        []RouteSettings      `json:"route_settings"`
}

type AccessLogSettings struct {
	DestinationArn string `json:"destination_arn"`
	Format         string `json:"format"`
}

type DefaultRouteSettings struct {
	ThrottlingBurstLimit   int    `json:"throttling_burst_limit"`
	ThrottlingRateLimit    int    `json:"throttling_rate_limit"`
	DataTraceEnabled       bool   `json:"data_trace_enabled"`
	DetailedMetricsEnabled bool   `json:"detailed_metrics_enabled"`
	LoggingLevel           string `json:"logging_level"`
}

type RouteSettings struct {
	LoggingLevel           string `json:"logging_level"`
	RouteKey               string `json:"route_key"`
	ThrottlingBurstLimit   int    `json:"throttling_burst_limit"`
	ThrottlingRateLimit    int    `json:"throttling_rate_limit"`
	DataTraceEnabled       bool   `json:"data_trace_enabled"`
	DetailedMetricsEnabled bool   `json:"detailed_metrics_enabled"`
}

// A Apigatewayv2StageStatus defines the observed state of a Apigatewayv2Stage
type Apigatewayv2StageStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Apigatewayv2StageObservation `json:",inline"`
}

// A Apigatewayv2StageObservation records the observed state of a Apigatewayv2Stage
type Apigatewayv2StageObservation struct {
	Arn          string `json:"arn"`
	ExecutionArn string `json:"execution_arn"`
	InvokeUrl    string `json:"invoke_url"`
}