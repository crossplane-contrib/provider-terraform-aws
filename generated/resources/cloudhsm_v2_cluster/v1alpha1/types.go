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

// CloudhsmV2Cluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudhsmV2Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudhsmV2ClusterSpec   `json:"spec"`
	Status CloudhsmV2ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudhsmV2Cluster contains a list of CloudhsmV2ClusterList
type CloudhsmV2ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudhsmV2Cluster `json:"items"`
}

// A CloudhsmV2ClusterSpec defines the desired state of a CloudhsmV2Cluster
type CloudhsmV2ClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudhsmV2ClusterParameters `json:"forProvider"`
}

// A CloudhsmV2ClusterParameters defines the desired state of a CloudhsmV2Cluster
type CloudhsmV2ClusterParameters struct {
	HsmType                string            `json:"hsm_type"`
	Id                     string            `json:"id"`
	SourceBackupIdentifier string            `json:"source_backup_identifier"`
	SubnetIds              []string          `json:"subnet_ids"`
	Tags                   map[string]string `json:"tags"`
	Timeouts               Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Update string `json:"update"`
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A CloudhsmV2ClusterStatus defines the observed state of a CloudhsmV2Cluster
type CloudhsmV2ClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudhsmV2ClusterObservation `json:"atProvider"`
}

// A CloudhsmV2ClusterObservation records the observed state of a CloudhsmV2Cluster
type CloudhsmV2ClusterObservation struct {
	ClusterCertificates []ClusterCertificates `json:"cluster_certificates"`
	ClusterId           string                `json:"cluster_id"`
	ClusterState        string                `json:"cluster_state"`
	SecurityGroupId     string                `json:"security_group_id"`
	VpcId               string                `json:"vpc_id"`
}

type ClusterCertificates struct {
	ManufacturerHardwareCertificate string `json:"manufacturer_hardware_certificate"`
	AwsHardwareCertificate          string `json:"aws_hardware_certificate"`
	ClusterCertificate              string `json:"cluster_certificate"`
	ClusterCsr                      string `json:"cluster_csr"`
	HsmCertificate                  string `json:"hsm_certificate"`
}