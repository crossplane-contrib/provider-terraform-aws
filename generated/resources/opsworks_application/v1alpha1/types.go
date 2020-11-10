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
	Type                   string             `json:"type"`
	Description            string             `json:"description"`
	EnableSsl              bool               `json:"enable_ssl"`
	StackId                string             `json:"stack_id"`
	Domains                []string           `json:"domains"`
	Name                   string             `json:"name"`
	DataSourceType         string             `json:"data_source_type"`
	DocumentRoot           string             `json:"document_root"`
	AutoBundleOnDeploy     string             `json:"auto_bundle_on_deploy"`
	DataSourceDatabaseName string             `json:"data_source_database_name"`
	RailsEnv               string             `json:"rails_env"`
	AwsFlowRubySettings    string             `json:"aws_flow_ruby_settings"`
	DataSourceArn          string             `json:"data_source_arn"`
	AppSource              []AppSource        `json:"app_source"`
	Environment            []Environment      `json:"environment"`
	SslConfiguration       []SslConfiguration `json:"ssl_configuration"`
}

type AppSource struct {
	Password string `json:"password"`
	Revision string `json:"revision"`
	SshKey   string `json:"ssh_key"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	Username string `json:"username"`
}

type Environment struct {
	Secure bool   `json:"secure"`
	Value  string `json:"value"`
	Key    string `json:"key"`
}

type SslConfiguration struct {
	Certificate string `json:"certificate"`
	Chain       string `json:"chain"`
	PrivateKey  string `json:"private_key"`
}

// A OpsworksApplicationStatus defines the observed state of a OpsworksApplication
type OpsworksApplicationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksApplicationObservation `json:",inline"`
}

// A OpsworksApplicationObservation records the observed state of a OpsworksApplication
type OpsworksApplicationObservation struct {
	ShortName string `json:"short_name"`
	Id        string `json:"id"`
}