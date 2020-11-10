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

// GuarddutyThreatintelset is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GuarddutyThreatintelset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GuarddutyThreatintelsetSpec   `json:"spec"`
	Status GuarddutyThreatintelsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuarddutyThreatintelset contains a list of GuarddutyThreatintelsetList
type GuarddutyThreatintelsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GuarddutyThreatintelset `json:"items"`
}

// A GuarddutyThreatintelsetSpec defines the desired state of a GuarddutyThreatintelset
type GuarddutyThreatintelsetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GuarddutyThreatintelsetParameters `json:",inline"`
}

// A GuarddutyThreatintelsetParameters defines the desired state of a GuarddutyThreatintelset
type GuarddutyThreatintelsetParameters struct {
	Tags       map[string]string `json:"tags"`
	Activate   bool              `json:"activate"`
	DetectorId string            `json:"detector_id"`
	Format     string            `json:"format"`
	Location   string            `json:"location"`
	Name       string            `json:"name"`
}

// A GuarddutyThreatintelsetStatus defines the observed state of a GuarddutyThreatintelset
type GuarddutyThreatintelsetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GuarddutyThreatintelsetObservation `json:",inline"`
}

// A GuarddutyThreatintelsetObservation records the observed state of a GuarddutyThreatintelset
type GuarddutyThreatintelsetObservation struct {
	Arn string `json:"arn"`
	Id  string `json:"id"`
}