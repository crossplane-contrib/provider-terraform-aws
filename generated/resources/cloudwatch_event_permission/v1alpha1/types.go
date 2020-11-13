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

// CloudwatchEventPermission is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudwatchEventPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudwatchEventPermissionSpec   `json:"spec"`
	Status CloudwatchEventPermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventPermission contains a list of CloudwatchEventPermissionList
type CloudwatchEventPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchEventPermission `json:"items"`
}

// A CloudwatchEventPermissionSpec defines the desired state of a CloudwatchEventPermission
type CloudwatchEventPermissionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudwatchEventPermissionParameters `json:",inline"`
}

// A CloudwatchEventPermissionParameters defines the desired state of a CloudwatchEventPermission
type CloudwatchEventPermissionParameters struct {
	StatementId string    `json:"statement_id"`
	Action      string    `json:"action"`
	Id          string    `json:"id"`
	Principal   string    `json:"principal"`
	Condition   Condition `json:"condition"`
}

type Condition struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Key   string `json:"key"`
}

// A CloudwatchEventPermissionStatus defines the observed state of a CloudwatchEventPermission
type CloudwatchEventPermissionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudwatchEventPermissionObservation `json:",inline"`
}

// A CloudwatchEventPermissionObservation records the observed state of a CloudwatchEventPermission
type CloudwatchEventPermissionObservation struct{}