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

// DbProxy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DbProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbProxySpec   `json:"spec"`
	Status DbProxyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbProxy contains a list of DbProxyList
type DbProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbProxy `json:"items"`
}

// A DbProxySpec defines the desired state of a DbProxy
type DbProxySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DbProxyParameters `json:",inline"`
}

// A DbProxyParameters defines the desired state of a DbProxy
type DbProxyParameters struct {
	EngineFamily        string            `json:"engine_family"`
	IdleClientTimeout   int64             `json:"idle_client_timeout"`
	VpcSecurityGroupIds []string          `json:"vpc_security_group_ids"`
	VpcSubnetIds        []string          `json:"vpc_subnet_ids"`
	RoleArn             string            `json:"role_arn"`
	Tags                map[string]string `json:"tags"`
	DebugLogging        bool              `json:"debug_logging"`
	Id                  string            `json:"id"`
	Name                string            `json:"name"`
	RequireTls          bool              `json:"require_tls"`
	Auth                []Auth            `json:"auth"`
	Timeouts            Timeouts          `json:"timeouts"`
}

type Auth struct {
	SecretArn   string `json:"secret_arn"`
	AuthScheme  string `json:"auth_scheme"`
	Description string `json:"description"`
	IamAuth     string `json:"iam_auth"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A DbProxyStatus defines the observed state of a DbProxy
type DbProxyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DbProxyObservation `json:",inline"`
}

// A DbProxyObservation records the observed state of a DbProxy
type DbProxyObservation struct {
	Arn      string `json:"arn"`
	Endpoint string `json:"endpoint"`
}