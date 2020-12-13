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

// EfsFileSystem is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EfsFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EfsFileSystemSpec   `json:"spec"`
	Status EfsFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EfsFileSystem contains a list of EfsFileSystemList
type EfsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EfsFileSystem `json:"items"`
}

// A EfsFileSystemSpec defines the desired state of a EfsFileSystem
type EfsFileSystemSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EfsFileSystemParameters `json:"forProvider"`
}

// A EfsFileSystemParameters defines the desired state of a EfsFileSystem
type EfsFileSystemParameters struct {
	Encrypted                    bool              `json:"encrypted"`
	Id                           string            `json:"id"`
	KmsKeyId                     string            `json:"kms_key_id"`
	ProvisionedThroughputInMibps int64             `json:"provisioned_throughput_in_mibps"`
	Tags                         map[string]string `json:"tags"`
	CreationToken                string            `json:"creation_token"`
	PerformanceMode              string            `json:"performance_mode"`
	ThroughputMode               string            `json:"throughput_mode"`
	LifecyclePolicy              LifecyclePolicy   `json:"lifecycle_policy"`
}

type LifecyclePolicy struct {
	TransitionToIa string `json:"transition_to_ia"`
}

// A EfsFileSystemStatus defines the observed state of a EfsFileSystem
type EfsFileSystemStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EfsFileSystemObservation `json:"atProvider"`
}

// A EfsFileSystemObservation records the observed state of a EfsFileSystem
type EfsFileSystemObservation struct {
	Arn     string `json:"arn"`
	DnsName string `json:"dns_name"`
}