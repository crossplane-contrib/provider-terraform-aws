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

// NeptuneSubnetGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type NeptuneSubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NeptuneSubnetGroupSpec   `json:"spec"`
	Status NeptuneSubnetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneSubnetGroup contains a list of NeptuneSubnetGroupList
type NeptuneSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NeptuneSubnetGroup `json:"items"`
}

// A NeptuneSubnetGroupSpec defines the desired state of a NeptuneSubnetGroup
type NeptuneSubnetGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  NeptuneSubnetGroupParameters `json:"forProvider"`
}

// A NeptuneSubnetGroupParameters defines the desired state of a NeptuneSubnetGroup
type NeptuneSubnetGroupParameters struct {
	NamePrefix  string            `json:"name_prefix"`
	SubnetIds   []string          `json:"subnet_ids"`
	Tags        map[string]string `json:"tags"`
	Description string            `json:"description"`
	Name        string            `json:"name"`
}

// A NeptuneSubnetGroupStatus defines the observed state of a NeptuneSubnetGroup
type NeptuneSubnetGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     NeptuneSubnetGroupObservation `json:"atProvider"`
}

// A NeptuneSubnetGroupObservation records the observed state of a NeptuneSubnetGroup
type NeptuneSubnetGroupObservation struct {
	Arn string `json:"arn"`
}