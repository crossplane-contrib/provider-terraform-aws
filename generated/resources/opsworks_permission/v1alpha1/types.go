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

// OpsworksPermission is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksPermissionSpec   `json:"spec"`
	Status OpsworksPermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksPermission contains a list of OpsworksPermissionList
type OpsworksPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksPermission `json:"items"`
}

// A OpsworksPermissionSpec defines the desired state of a OpsworksPermission
type OpsworksPermissionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksPermissionParameters `json:",inline"`
}

// A OpsworksPermissionParameters defines the desired state of a OpsworksPermission
type OpsworksPermissionParameters struct {
	UserArn string `json:"user_arn"`
}

// A OpsworksPermissionStatus defines the observed state of a OpsworksPermission
type OpsworksPermissionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksPermissionObservation `json:",inline"`
}

// A OpsworksPermissionObservation records the observed state of a OpsworksPermission
type OpsworksPermissionObservation struct {
	AllowSsh  bool   `json:"allow_ssh"`
	AllowSudo bool   `json:"allow_sudo"`
	Id        string `json:"id"`
	Level     string `json:"level"`
	StackId   string `json:"stack_id"`
}