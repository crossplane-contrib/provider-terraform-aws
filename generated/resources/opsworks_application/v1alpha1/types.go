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

// OpsworksApplication is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksApplicationSpec   `json:"spec"`
	Status OpsworksApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksApplication contains a list of OpsworksApplicationList
type OpsworksApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksApplication `json:"items"`
}

// A OpsworksApplicationSpec defines the desired state of a OpsworksApplication
type OpsworksApplicationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksApplicationParameters `json:",inline"`
}

// A OpsworksApplicationParameters defines the desired state of a OpsworksApplication
type OpsworksApplicationParameters struct {
	DataSourceType         string           `json:"data_source_type"`
	DocumentRoot           string           `json:"document_root"`
	Domains                []string         `json:"domains"`
	Name                   string           `json:"name"`
	Description            string           `json:"description"`
	EnableSsl              bool             `json:"enable_ssl"`
	Id                     string           `json:"id"`
	AwsFlowRubySettings    string           `json:"aws_flow_ruby_settings"`
	DataSourceDatabaseName string           `json:"data_source_database_name"`
	StackId                string           `json:"stack_id"`
	Type                   string           `json:"type"`
	AutoBundleOnDeploy     string           `json:"auto_bundle_on_deploy"`
	DataSourceArn          string           `json:"data_source_arn"`
	RailsEnv               string           `json:"rails_env"`
	ShortName              string           `json:"short_name"`
	Environment            Environment      `json:"environment"`
	SslConfiguration       SslConfiguration `json:"ssl_configuration"`
	AppSource              AppSource        `json:"app_source"`
}

type Environment struct {
	Key    string `json:"key"`
	Secure bool   `json:"secure"`
	Value  string `json:"value"`
}

type SslConfiguration struct {
	Chain       string `json:"chain"`
	PrivateKey  string `json:"private_key"`
	Certificate string `json:"certificate"`
}

type AppSource struct {
	SshKey   string `json:"ssh_key"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Revision string `json:"revision"`
}

// A OpsworksApplicationStatus defines the observed state of a OpsworksApplication
type OpsworksApplicationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksApplicationObservation `json:",inline"`
}

// A OpsworksApplicationObservation records the observed state of a OpsworksApplication
type OpsworksApplicationObservation struct{}