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

// DatasyncTask is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DatasyncTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatasyncTaskSpec   `json:"spec"`
	Status DatasyncTaskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasyncTask contains a list of DatasyncTaskList
type DatasyncTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatasyncTask `json:"items"`
}

// A DatasyncTaskSpec defines the desired state of a DatasyncTask
type DatasyncTaskSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DatasyncTaskParameters `json:",inline"`
}

// A DatasyncTaskParameters defines the desired state of a DatasyncTask
type DatasyncTaskParameters struct {
	CloudwatchLogGroupArn  string            `json:"cloudwatch_log_group_arn"`
	DestinationLocationArn string            `json:"destination_location_arn"`
	Id                     string            `json:"id"`
	Name                   string            `json:"name"`
	SourceLocationArn      string            `json:"source_location_arn"`
	Tags                   map[string]string `json:"tags"`
	Timeouts               []Timeouts        `json:"timeouts"`
	Options                Options           `json:"options"`
}

type Timeouts struct {
	Create string `json:"create"`
}

type Options struct {
	BytesPerSecond       int    `json:"bytes_per_second"`
	Gid                  string `json:"gid"`
	PosixPermissions     string `json:"posix_permissions"`
	VerifyMode           string `json:"verify_mode"`
	Atime                string `json:"atime"`
	Mtime                string `json:"mtime"`
	PreserveDeletedFiles string `json:"preserve_deleted_files"`
	PreserveDevices      string `json:"preserve_devices"`
	Uid                  string `json:"uid"`
}

// A DatasyncTaskStatus defines the observed state of a DatasyncTask
type DatasyncTaskStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DatasyncTaskObservation `json:",inline"`
}

// A DatasyncTaskObservation records the observed state of a DatasyncTask
type DatasyncTaskObservation struct {
	Arn string `json:"arn"`
}