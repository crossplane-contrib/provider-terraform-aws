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

// DatasyncLocationFsxWindowsFileSystem is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DatasyncLocationFsxWindowsFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatasyncLocationFsxWindowsFileSystemSpec   `json:"spec"`
	Status DatasyncLocationFsxWindowsFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasyncLocationFsxWindowsFileSystem contains a list of DatasyncLocationFsxWindowsFileSystemList
type DatasyncLocationFsxWindowsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatasyncLocationFsxWindowsFileSystem `json:"items"`
}

// A DatasyncLocationFsxWindowsFileSystemSpec defines the desired state of a DatasyncLocationFsxWindowsFileSystem
type DatasyncLocationFsxWindowsFileSystemSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DatasyncLocationFsxWindowsFileSystemParameters `json:"forProvider"`
}

// A DatasyncLocationFsxWindowsFileSystemParameters defines the desired state of a DatasyncLocationFsxWindowsFileSystem
type DatasyncLocationFsxWindowsFileSystemParameters struct {
	User              string            `json:"user"`
	FsxFilesystemArn  string            `json:"fsx_filesystem_arn"`
	Password          string            `json:"password"`
	Subdirectory      string            `json:"subdirectory"`
	Domain            string            `json:"domain"`
	Id                string            `json:"id"`
	SecurityGroupArns []string          `json:"security_group_arns"`
	Tags              map[string]string `json:"tags"`
}

// A DatasyncLocationFsxWindowsFileSystemStatus defines the observed state of a DatasyncLocationFsxWindowsFileSystem
type DatasyncLocationFsxWindowsFileSystemStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DatasyncLocationFsxWindowsFileSystemObservation `json:"atProvider"`
}

// A DatasyncLocationFsxWindowsFileSystemObservation records the observed state of a DatasyncLocationFsxWindowsFileSystem
type DatasyncLocationFsxWindowsFileSystemObservation struct {
	Uri          string `json:"uri"`
	Arn          string `json:"arn"`
	CreationTime string `json:"creation_time"`
}