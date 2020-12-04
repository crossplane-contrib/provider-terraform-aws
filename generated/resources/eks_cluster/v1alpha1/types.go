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

// EksCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EksCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EksClusterSpec   `json:"spec"`
	Status EksClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EksCluster contains a list of EksClusterList
type EksClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EksCluster `json:"items"`
}

// A EksClusterSpec defines the desired state of a EksCluster
type EksClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EksClusterParameters `json:",inline"`
}

// A EksClusterParameters defines the desired state of a EksCluster
type EksClusterParameters struct {
	Name                   string            `json:"name"`
	Id                     string            `json:"id"`
	RoleArn                string            `json:"role_arn"`
	Tags                   map[string]string `json:"tags"`
	EnabledClusterLogTypes []string          `json:"enabled_cluster_log_types"`
	Version                string            `json:"version"`
	EncryptionConfig       EncryptionConfig  `json:"encryption_config"`
	Timeouts               Timeouts          `json:"timeouts"`
	VpcConfig              VpcConfig         `json:"vpc_config"`
}

type EncryptionConfig struct {
	Resources []string `json:"resources"`
	Provider  Provider `json:"provider"`
}

type Provider struct {
	KeyArn string `json:"key_arn"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type VpcConfig struct {
	SecurityGroupIds       []string `json:"security_group_ids"`
	SubnetIds              []string `json:"subnet_ids"`
	VpcId                  string   `json:"vpc_id"`
	ClusterSecurityGroupId string   `json:"cluster_security_group_id"`
	EndpointPrivateAccess  bool     `json:"endpoint_private_access"`
	EndpointPublicAccess   bool     `json:"endpoint_public_access"`
	PublicAccessCidrs      []string `json:"public_access_cidrs"`
}

// A EksClusterStatus defines the observed state of a EksCluster
type EksClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EksClusterObservation `json:",inline"`
}

// A EksClusterObservation records the observed state of a EksCluster
type EksClusterObservation struct {
	PlatformVersion      string                 `json:"platform_version"`
	Arn                  string                 `json:"arn"`
	Endpoint             string                 `json:"endpoint"`
	Identity             []Identity             `json:"identity"`
	Status               string                 `json:"status"`
	CertificateAuthority []CertificateAuthority `json:"certificate_authority"`
	CreatedAt            string                 `json:"created_at"`
}

type Identity struct {
	Oidc []Oidc `json:"oidc"`
}

type Oidc struct {
	Issuer string `json:"issuer"`
}

type CertificateAuthority struct {
	Data string `json:"data"`
}