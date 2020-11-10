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

// LexSlotType is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LexSlotType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LexSlotTypeSpec   `json:"spec"`
	Status LexSlotTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LexSlotType contains a list of LexSlotTypeList
type LexSlotTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LexSlotType `json:"items"`
}

// A LexSlotTypeSpec defines the desired state of a LexSlotType
type LexSlotTypeSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LexSlotTypeParameters `json:",inline"`
}

// A LexSlotTypeParameters defines the desired state of a LexSlotType
type LexSlotTypeParameters struct {
	Name                   string             `json:"name"`
	ValueSelectionStrategy string             `json:"value_selection_strategy"`
	CreateVersion          bool               `json:"create_version"`
	Description            string             `json:"description"`
	EnumerationValue       []EnumerationValue `json:"enumeration_value"`
	Timeouts               []Timeouts         `json:"timeouts"`
}

type EnumerationValue struct {
	Value    string   `json:"value"`
	Synonyms []string `json:"synonyms"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A LexSlotTypeStatus defines the observed state of a LexSlotType
type LexSlotTypeStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LexSlotTypeObservation `json:",inline"`
}

// A LexSlotTypeObservation records the observed state of a LexSlotType
type LexSlotTypeObservation struct {
	Id              string `json:"id"`
	LastUpdatedDate string `json:"last_updated_date"`
	Version         string `json:"version"`
	Checksum        string `json:"checksum"`
	CreatedDate     string `json:"created_date"`
}