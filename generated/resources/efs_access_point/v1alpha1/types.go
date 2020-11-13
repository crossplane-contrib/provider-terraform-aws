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
	ForProvider                  EfsAccessPointParameters `json:",inline"`
}

// A EfsAccessPointParameters defines the desired state of a EfsAccessPoint
type EfsAccessPointParameters struct {
	FileSystemId  string            `json:"file_system_id"`
	Id            string            `json:"id"`
	Tags          map[string]string `json:"tags"`
	PosixUser     PosixUser         `json:"posix_user"`
	RootDirectory RootDirectory     `json:"root_directory"`
}

type PosixUser struct {
	SecondaryGids []int `json:"secondary_gids"`
	Uid           int   `json:"uid"`
	Gid           int   `json:"gid"`
}

type RootDirectory struct {
	Path         string       `json:"path"`
	CreationInfo CreationInfo `json:"creation_info"`
}

type CreationInfo struct {
	OwnerGid    int    `json:"owner_gid"`
	OwnerUid    int    `json:"owner_uid"`
	Permissions string `json:"permissions"`
}

// A EfsAccessPointStatus defines the observed state of a EfsAccessPoint
type EfsAccessPointStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EfsAccessPointObservation `json:",inline"`
}

// A EfsAccessPointObservation records the observed state of a EfsAccessPoint
type EfsAccessPointObservation struct {
	Arn           string `json:"arn"`
	FileSystemArn string `json:"file_system_arn"`
	OwnerId       string `json:"owner_id"`
}