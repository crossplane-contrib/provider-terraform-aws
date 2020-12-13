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
	ForProvider                  EksClusterParameters `json:"forProvider"`
}

// A EksClusterParameters defines the desired state of a EksCluster
type EksClusterParameters struct {
	Name                   string            `json:"name"`
	RoleArn                string            `json:"role_arn"`
	Version                string            `json:"version"`
	EnabledClusterLogTypes []string          `json:"enabled_cluster_log_types"`
	Id                     string            `json:"id"`
	Tags                   map[string]string `json:"tags"`
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
	Update string `json:"update"`
	Create string `json:"create"`
	Delete string `json:"delete"`
}

type VpcConfig struct {
	SubnetIds              []string `json:"subnet_ids"`
	VpcId                  string   `json:"vpc_id"`
	ClusterSecurityGroupId string   `json:"cluster_security_group_id"`
	EndpointPrivateAccess  bool     `json:"endpoint_private_access"`
	EndpointPublicAccess   bool     `json:"endpoint_public_access"`
	PublicAccessCidrs      []string `json:"public_access_cidrs"`
	SecurityGroupIds       []string `json:"security_group_ids"`
}

// A EksClusterStatus defines the observed state of a EksCluster
type EksClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EksClusterObservation `json:"atProvider"`
}

// A EksClusterObservation records the observed state of a EksCluster
type EksClusterObservation struct {
	PlatformVersion      string                 `json:"platform_version"`
	CreatedAt            string                 `json:"created_at"`
	Endpoint             string                 `json:"endpoint"`
	Identity             []Identity             `json:"identity"`
	Status               string                 `json:"status"`
	Arn                  string                 `json:"arn"`
	CertificateAuthority []CertificateAuthority `json:"certificate_authority"`
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