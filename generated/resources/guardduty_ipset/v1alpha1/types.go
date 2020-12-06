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

// GuarddutyIpset is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GuarddutyIpset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GuarddutyIpsetSpec   `json:"spec"`
	Status GuarddutyIpsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuarddutyIpset contains a list of GuarddutyIpsetList
type GuarddutyIpsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GuarddutyIpset `json:"items"`
}

// A GuarddutyIpsetSpec defines the desired state of a GuarddutyIpset
type GuarddutyIpsetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GuarddutyIpsetParameters `json:",inline"`
}

// A GuarddutyIpsetParameters defines the desired state of a GuarddutyIpset
type GuarddutyIpsetParameters struct {
	Format     string            `json:"format"`
	Id         string            `json:"id"`
	Location   string            `json:"location"`
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	Activate   bool              `json:"activate"`
	DetectorId string            `json:"detector_id"`
}

// A GuarddutyIpsetStatus defines the observed state of a GuarddutyIpset
type GuarddutyIpsetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GuarddutyIpsetObservation `json:",inline"`
}

// A GuarddutyIpsetObservation records the observed state of a GuarddutyIpset
type GuarddutyIpsetObservation struct {
	Arn string `json:"arn"`
}