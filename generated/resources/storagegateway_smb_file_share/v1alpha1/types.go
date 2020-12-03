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

// StoragegatewaySmbFileShare is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type StoragegatewaySmbFileShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StoragegatewaySmbFileShareSpec   `json:"spec"`
	Status StoragegatewaySmbFileShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewaySmbFileShare contains a list of StoragegatewaySmbFileShareList
type StoragegatewaySmbFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoragegatewaySmbFileShare `json:"items"`
}

// A StoragegatewaySmbFileShareSpec defines the desired state of a StoragegatewaySmbFileShare
type StoragegatewaySmbFileShareSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  StoragegatewaySmbFileShareParameters `json:",inline"`
}

// A StoragegatewaySmbFileShareParameters defines the desired state of a StoragegatewaySmbFileShare
type StoragegatewaySmbFileShareParameters struct {
	GuessMimeTypeEnabled bool              `json:"guess_mime_type_enabled"`
	KmsEncrypted         bool              `json:"kms_encrypted"`
	AdminUserList        []string          `json:"admin_user_list"`
	AuditDestinationArn  string            `json:"audit_destination_arn"`
	Authentication       string            `json:"authentication"`
	DefaultStorageClass  string            `json:"default_storage_class"`
	KmsKeyArn            string            `json:"kms_key_arn"`
	RequesterPays        bool              `json:"requester_pays"`
	SmbAclEnabled        bool              `json:"smb_acl_enabled"`
	Tags                 map[string]string `json:"tags"`
	CaseSensitivity      string            `json:"case_sensitivity"`
	Id                   string            `json:"id"`
	InvalidUserList      []string          `json:"invalid_user_list"`
	ObjectAcl            string            `json:"object_acl"`
	ValidUserList        []string          `json:"valid_user_list"`
	GatewayArn           string            `json:"gateway_arn"`
	LocationArn          string            `json:"location_arn"`
	ReadOnly             bool              `json:"read_only"`
	RoleArn              string            `json:"role_arn"`
	CacheAttributes      CacheAttributes   `json:"cache_attributes"`
	Timeouts             []Timeouts        `json:"timeouts"`
}

type CacheAttributes struct {
	CacheStaleTimeoutInSeconds int `json:"cache_stale_timeout_in_seconds"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Update string `json:"update"`
	Create string `json:"create"`
}

// A StoragegatewaySmbFileShareStatus defines the observed state of a StoragegatewaySmbFileShare
type StoragegatewaySmbFileShareStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     StoragegatewaySmbFileShareObservation `json:",inline"`
}

// A StoragegatewaySmbFileShareObservation records the observed state of a StoragegatewaySmbFileShare
type StoragegatewaySmbFileShareObservation struct {
	FileshareId string `json:"fileshare_id"`
	Path        string `json:"path"`
	Arn         string `json:"arn"`
}