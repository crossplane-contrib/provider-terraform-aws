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

// EfsAccessPoint is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EfsAccessPoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EfsAccessPointSpec   `json:"spec"`
	Status EfsAccessPointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EfsAccessPoint contains a list of EfsAccessPointList
type EfsAccessPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EfsAccessPoint `json:"items"`
}

// A EfsAccessPointSpec defines the desired state of a EfsAccessPoint
type EfsAccessPointSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EfsAccessPointParameters `json:"forProvider"`
}

// A EfsAccessPointParameters defines the desired state of a EfsAccessPoint
type EfsAccessPointParameters struct {
	Tags          map[string]string `json:"tags"`
	FileSystemId  string            `json:"file_system_id"`
	Id            string            `json:"id"`
	RootDirectory RootDirectory     `json:"root_directory"`
	PosixUser     PosixUser         `json:"posix_user"`
}

type RootDirectory struct {
	Path         string       `json:"path"`
	CreationInfo CreationInfo `json:"creation_info"`
}

type CreationInfo struct {
	Permissions string `json:"permissions"`
	OwnerGid    int64  `json:"owner_gid"`
	OwnerUid    int64  `json:"owner_uid"`
}

type PosixUser struct {
	Uid           int64   `json:"uid"`
	Gid           int64   `json:"gid"`
	SecondaryGids []int64 `json:"secondary_gids"`
}

// A EfsAccessPointStatus defines the observed state of a EfsAccessPoint
type EfsAccessPointStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EfsAccessPointObservation `json:"atProvider"`
}

// A EfsAccessPointObservation records the observed state of a EfsAccessPoint
type EfsAccessPointObservation struct {
	OwnerId       string `json:"owner_id"`
	Arn           string `json:"arn"`
	FileSystemArn string `json:"file_system_arn"`
}