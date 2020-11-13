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

// StoragegatewayWorkingStorage is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type StoragegatewayWorkingStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StoragegatewayWorkingStorageSpec   `json:"spec"`
	Status StoragegatewayWorkingStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayWorkingStorage contains a list of StoragegatewayWorkingStorageList
type StoragegatewayWorkingStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoragegatewayWorkingStorage `json:"items"`
}

// A StoragegatewayWorkingStorageSpec defines the desired state of a StoragegatewayWorkingStorage
type StoragegatewayWorkingStorageSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  StoragegatewayWorkingStorageParameters `json:",inline"`
}

// A StoragegatewayWorkingStorageParameters defines the desired state of a StoragegatewayWorkingStorage
type StoragegatewayWorkingStorageParameters struct {
	DiskId     string `json:"disk_id"`
	GatewayArn string `json:"gateway_arn"`
	Id         string `json:"id"`
}

// A StoragegatewayWorkingStorageStatus defines the observed state of a StoragegatewayWorkingStorage
type StoragegatewayWorkingStorageStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     StoragegatewayWorkingStorageObservation `json:",inline"`
}

// A StoragegatewayWorkingStorageObservation records the observed state of a StoragegatewayWorkingStorage
type StoragegatewayWorkingStorageObservation struct{}