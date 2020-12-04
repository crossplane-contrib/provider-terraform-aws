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

// AmiLaunchPermission is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AmiLaunchPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AmiLaunchPermissionSpec   `json:"spec"`
	Status AmiLaunchPermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AmiLaunchPermission contains a list of AmiLaunchPermissionList
type AmiLaunchPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AmiLaunchPermission `json:"items"`
}

// A AmiLaunchPermissionSpec defines the desired state of a AmiLaunchPermission
type AmiLaunchPermissionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AmiLaunchPermissionParameters `json:",inline"`
}

// A AmiLaunchPermissionParameters defines the desired state of a AmiLaunchPermission
type AmiLaunchPermissionParameters struct {
	Id        string `json:"id"`
	ImageId   string `json:"image_id"`
	AccountId string `json:"account_id"`
}

// A AmiLaunchPermissionStatus defines the observed state of a AmiLaunchPermission
type AmiLaunchPermissionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AmiLaunchPermissionObservation `json:",inline"`
}

// A AmiLaunchPermissionObservation records the observed state of a AmiLaunchPermission
type AmiLaunchPermissionObservation struct{}