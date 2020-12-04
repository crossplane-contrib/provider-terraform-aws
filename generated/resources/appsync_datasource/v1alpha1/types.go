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

// AppsyncDatasource is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppsyncDatasource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppsyncDatasourceSpec   `json:"spec"`
	Status AppsyncDatasourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppsyncDatasource contains a list of AppsyncDatasourceList
type AppsyncDatasourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppsyncDatasource `json:"items"`
}

// A AppsyncDatasourceSpec defines the desired state of a AppsyncDatasource
type AppsyncDatasourceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppsyncDatasourceParameters `json:",inline"`
}

// A AppsyncDatasourceParameters defines the desired state of a AppsyncDatasource
type AppsyncDatasourceParameters struct {
	Type                string              `json:"type"`
	ApiId               string              `json:"api_id"`
	Description         string              `json:"description"`
	Id                  string              `json:"id"`
	Name                string              `json:"name"`
	ServiceRoleArn      string              `json:"service_role_arn"`
	DynamodbConfig      DynamodbConfig      `json:"dynamodb_config"`
	ElasticsearchConfig ElasticsearchConfig `json:"elasticsearch_config"`
	HttpConfig          HttpConfig          `json:"http_config"`
	LambdaConfig        LambdaConfig        `json:"lambda_config"`
}

type DynamodbConfig struct {
	Region               string `json:"region"`
	TableName            string `json:"table_name"`
	UseCallerCredentials bool   `json:"use_caller_credentials"`
}

type ElasticsearchConfig struct {
	Endpoint string `json:"endpoint"`
	Region   string `json:"region"`
}

type HttpConfig struct {
	Endpoint string `json:"endpoint"`
}

type LambdaConfig struct {
	FunctionArn string `json:"function_arn"`
}

// A AppsyncDatasourceStatus defines the observed state of a AppsyncDatasource
type AppsyncDatasourceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppsyncDatasourceObservation `json:",inline"`
}

// A AppsyncDatasourceObservation records the observed state of a AppsyncDatasource
type AppsyncDatasourceObservation struct {
	Arn string `json:"arn"`
}