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

// XrayGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type XrayGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   XrayGroupSpec   `json:"spec"`
	Status XrayGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// XrayGroup contains a list of XrayGroupList
type XrayGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []XrayGroup `json:"items"`
}

// A XrayGroupSpec defines the desired state of a XrayGroup
type XrayGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  XrayGroupParameters `json:",inline"`
}

// A XrayGroupParameters defines the desired state of a XrayGroup
type XrayGroupParameters struct {
	Tags             map[string]string `json:"tags"`
	FilterExpression string            `json:"filter_expression"`
	GroupName        string            `json:"group_name"`
	Id               string            `json:"id"`
}

// A XrayGroupStatus defines the observed state of a XrayGroup
type XrayGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     XrayGroupObservation `json:",inline"`
}

// A XrayGroupObservation records the observed state of a XrayGroup
type XrayGroupObservation struct {
	Arn string `json:"arn"`
}