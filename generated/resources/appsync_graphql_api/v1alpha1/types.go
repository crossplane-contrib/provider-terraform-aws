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

// AppsyncGraphqlApi is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppsyncGraphqlApi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppsyncGraphqlApiSpec   `json:"spec"`
	Status AppsyncGraphqlApiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppsyncGraphqlApi contains a list of AppsyncGraphqlApiList
type AppsyncGraphqlApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppsyncGraphqlApi `json:"items"`
}

// A AppsyncGraphqlApiSpec defines the desired state of a AppsyncGraphqlApi
type AppsyncGraphqlApiSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppsyncGraphqlApiParameters `json:",inline"`
}

// A AppsyncGraphqlApiParameters defines the desired state of a AppsyncGraphqlApi
type AppsyncGraphqlApiParameters struct {
	Name                             string                           `json:"name"`
	Schema                           string                           `json:"schema"`
	Tags                             map[string]string                `json:"tags"`
	XrayEnabled                      bool                             `json:"xray_enabled"`
	AuthenticationType               string                           `json:"authentication_type"`
	Id                               string                           `json:"id"`
	AdditionalAuthenticationProvider AdditionalAuthenticationProvider `json:"additional_authentication_provider"`
	LogConfig                        LogConfig                        `json:"log_config"`
	OpenidConnectConfig              OpenidConnectConfig              `json:"openid_connect_config"`
	UserPoolConfig                   UserPoolConfig                   `json:"user_pool_config"`
}

type AdditionalAuthenticationProvider struct {
	AuthenticationType  string              `json:"authentication_type"`
	OpenidConnectConfig OpenidConnectConfig `json:"openid_connect_config"`
	UserPoolConfig      UserPoolConfig      `json:"user_pool_config"`
}

type OpenidConnectConfig struct {
	AuthTtl  int64  `json:"auth_ttl"`
	ClientId string `json:"client_id"`
	IatTtl   int64  `json:"iat_ttl"`
	Issuer   string `json:"issuer"`
}

type UserPoolConfig struct {
	AppIdClientRegex string `json:"app_id_client_regex"`
	AwsRegion        string `json:"aws_region"`
	UserPoolId       string `json:"user_pool_id"`
}

type LogConfig struct {
	CloudwatchLogsRoleArn string `json:"cloudwatch_logs_role_arn"`
	ExcludeVerboseContent bool   `json:"exclude_verbose_content"`
	FieldLogLevel         string `json:"field_log_level"`
}

type OpenidConnectConfig struct {
	AuthTtl  int64  `json:"auth_ttl"`
	ClientId string `json:"client_id"`
	IatTtl   int64  `json:"iat_ttl"`
	Issuer   string `json:"issuer"`
}

type UserPoolConfig struct {
	DefaultAction    string `json:"default_action"`
	UserPoolId       string `json:"user_pool_id"`
	AppIdClientRegex string `json:"app_id_client_regex"`
	AwsRegion        string `json:"aws_region"`
}

// A AppsyncGraphqlApiStatus defines the observed state of a AppsyncGraphqlApi
type AppsyncGraphqlApiStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppsyncGraphqlApiObservation `json:",inline"`
}

// A AppsyncGraphqlApiObservation records the observed state of a AppsyncGraphqlApi
type AppsyncGraphqlApiObservation struct {
	Uris map[string]string `json:"uris"`
	Arn  string            `json:"arn"`
}