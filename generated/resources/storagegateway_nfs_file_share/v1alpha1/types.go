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

// StoragegatewayNfsFileShare is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type StoragegatewayNfsFileShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StoragegatewayNfsFileShareSpec   `json:"spec"`
	Status StoragegatewayNfsFileShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayNfsFileShare contains a list of StoragegatewayNfsFileShareList
type StoragegatewayNfsFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoragegatewayNfsFileShare `json:"items"`
}

// A StoragegatewayNfsFileShareSpec defines the desired state of a StoragegatewayNfsFileShare
type StoragegatewayNfsFileShareSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  StoragegatewayNfsFileShareParameters `json:",inline"`
}

// A StoragegatewayNfsFileShareParameters defines the desired state of a StoragegatewayNfsFileShare
type StoragegatewayNfsFileShareParameters struct {
	ClientList           []string             `json:"client_list"`
	GatewayArn           string               `json:"gateway_arn"`
	KmsEncrypted         bool                 `json:"kms_encrypted"`
	LocationArn          string               `json:"location_arn"`
	RoleArn              string               `json:"role_arn"`
	Squash               string               `json:"squash"`
	Tags                 map[string]string    `json:"tags"`
	ReadOnly             bool                 `json:"read_only"`
	RequesterPays        bool                 `json:"requester_pays"`
	DefaultStorageClass  string               `json:"default_storage_class"`
	ObjectAcl            string               `json:"object_acl"`
	Id                   string               `json:"id"`
	KmsKeyArn            string               `json:"kms_key_arn"`
	GuessMimeTypeEnabled bool                 `json:"guess_mime_type_enabled"`
	CacheAttributes      CacheAttributes      `json:"cache_attributes"`
	NfsFileShareDefaults NfsFileShareDefaults `json:"nfs_file_share_defaults"`
	Timeouts             Timeouts             `json:"timeouts"`
}

type CacheAttributes struct {
	CacheStaleTimeoutInSeconds int64 `json:"cache_stale_timeout_in_seconds"`
}

type NfsFileShareDefaults struct {
	DirectoryMode string `json:"directory_mode"`
	FileMode      string `json:"file_mode"`
	GroupId       int64  `json:"group_id"`
	OwnerId       int64  `json:"owner_id"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A StoragegatewayNfsFileShareStatus defines the observed state of a StoragegatewayNfsFileShare
type StoragegatewayNfsFileShareStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     StoragegatewayNfsFileShareObservation `json:",inline"`
}

// A StoragegatewayNfsFileShareObservation records the observed state of a StoragegatewayNfsFileShare
type StoragegatewayNfsFileShareObservation struct {
	Arn         string `json:"arn"`
	FileshareId string `json:"fileshare_id"`
	Path        string `json:"path"`
}