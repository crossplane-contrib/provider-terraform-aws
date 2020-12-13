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

// DocdbSubnetGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DocdbSubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DocdbSubnetGroupSpec   `json:"spec"`
	Status DocdbSubnetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocdbSubnetGroup contains a list of DocdbSubnetGroupList
type DocdbSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocdbSubnetGroup `json:"items"`
}

// A DocdbSubnetGroupSpec defines the desired state of a DocdbSubnetGroup
type DocdbSubnetGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DocdbSubnetGroupParameters `json:"forProvider"`
}

// A DocdbSubnetGroupParameters defines the desired state of a DocdbSubnetGroup
type DocdbSubnetGroupParameters struct {
	Description string            `json:"description"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	NamePrefix  string            `json:"name_prefix"`
	SubnetIds   []string          `json:"subnet_ids"`
	Tags        map[string]string `json:"tags"`
}

// A DocdbSubnetGroupStatus defines the observed state of a DocdbSubnetGroup
type DocdbSubnetGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DocdbSubnetGroupObservation `json:"atProvider"`
}

// A DocdbSubnetGroupObservation records the observed state of a DocdbSubnetGroup
type DocdbSubnetGroupObservation struct {
	Arn string `json:"arn"`
}