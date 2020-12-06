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

// StoragegatewayCache is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type StoragegatewayCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StoragegatewayCacheSpec   `json:"spec"`
	Status StoragegatewayCacheStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayCache contains a list of StoragegatewayCacheList
type StoragegatewayCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoragegatewayCache `json:"items"`
}

// A StoragegatewayCacheSpec defines the desired state of a StoragegatewayCache
type StoragegatewayCacheSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  StoragegatewayCacheParameters `json:",inline"`
}

// A StoragegatewayCacheParameters defines the desired state of a StoragegatewayCache
type StoragegatewayCacheParameters struct {
	Id         string `json:"id"`
	DiskId     string `json:"disk_id"`
	GatewayArn string `json:"gateway_arn"`
}

// A StoragegatewayCacheStatus defines the observed state of a StoragegatewayCache
type StoragegatewayCacheStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     StoragegatewayCacheObservation `json:",inline"`
}

// A StoragegatewayCacheObservation records the observed state of a StoragegatewayCache
type StoragegatewayCacheObservation struct{}