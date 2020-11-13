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

// S3AccessPoint is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type S3AccessPoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S3AccessPointSpec   `json:"spec"`
	Status S3AccessPointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3AccessPoint contains a list of S3AccessPointList
type S3AccessPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3AccessPoint `json:"items"`
}

// A S3AccessPointSpec defines the desired state of a S3AccessPoint
type S3AccessPointSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  S3AccessPointParameters `json:",inline"`
}

// A S3AccessPointParameters defines the desired state of a S3AccessPoint
type S3AccessPointParameters struct {
	AccountId                      string                         `json:"account_id"`
	Id                             string                         `json:"id"`
	Name                           string                         `json:"name"`
	Policy                         string                         `json:"policy"`
	Bucket                         string                         `json:"bucket"`
	PublicAccessBlockConfiguration PublicAccessBlockConfiguration `json:"public_access_block_configuration"`
	VpcConfiguration               VpcConfiguration               `json:"vpc_configuration"`
}

type PublicAccessBlockConfiguration struct {
	RestrictPublicBuckets bool `json:"restrict_public_buckets"`
	BlockPublicAcls       bool `json:"block_public_acls"`
	BlockPublicPolicy     bool `json:"block_public_policy"`
	IgnorePublicAcls      bool `json:"ignore_public_acls"`
}

type VpcConfiguration struct {
	VpcId string `json:"vpc_id"`
}

// A S3AccessPointStatus defines the observed state of a S3AccessPoint
type S3AccessPointStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     S3AccessPointObservation `json:",inline"`
}

// A S3AccessPointObservation records the observed state of a S3AccessPoint
type S3AccessPointObservation struct {
	Arn                   string `json:"arn"`
	DomainName            string `json:"domain_name"`
	HasPublicAccessPolicy bool   `json:"has_public_access_policy"`
	NetworkOrigin         string `json:"network_origin"`
}