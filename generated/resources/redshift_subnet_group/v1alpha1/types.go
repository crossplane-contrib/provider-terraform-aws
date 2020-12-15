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

// RedshiftSubnetGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RedshiftSubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedshiftSubnetGroupSpec   `json:"spec"`
	Status RedshiftSubnetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftSubnetGroup contains a list of RedshiftSubnetGroupList
type RedshiftSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedshiftSubnetGroup `json:"items"`
}

// A RedshiftSubnetGroupSpec defines the desired state of a RedshiftSubnetGroup
type RedshiftSubnetGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RedshiftSubnetGroupParameters `json:"forProvider"`
}

// A RedshiftSubnetGroupParameters defines the desired state of a RedshiftSubnetGroup
type RedshiftSubnetGroupParameters struct {
	Tags        map[string]string `json:"tags"`
	Description string            `json:"description"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	SubnetIds   []string          `json:"subnet_ids"`
}

// A RedshiftSubnetGroupStatus defines the observed state of a RedshiftSubnetGroup
type RedshiftSubnetGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RedshiftSubnetGroupObservation `json:"atProvider"`
}

// A RedshiftSubnetGroupObservation records the observed state of a RedshiftSubnetGroup
type RedshiftSubnetGroupObservation struct {
	Arn string `json:"arn"`
}