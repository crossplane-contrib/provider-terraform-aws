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

// DynamodbTableItem is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DynamodbTableItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DynamodbTableItemSpec   `json:"spec"`
	Status DynamodbTableItemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DynamodbTableItem contains a list of DynamodbTableItemList
type DynamodbTableItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DynamodbTableItem `json:"items"`
}

// A DynamodbTableItemSpec defines the desired state of a DynamodbTableItem
type DynamodbTableItemSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DynamodbTableItemParameters `json:"forProvider"`
}

// A DynamodbTableItemParameters defines the desired state of a DynamodbTableItem
type DynamodbTableItemParameters struct {
	HashKey   string `json:"hash_key"`
	Item      string `json:"item"`
	RangeKey  string `json:"range_key"`
	TableName string `json:"table_name"`
}

// A DynamodbTableItemStatus defines the observed state of a DynamodbTableItem
type DynamodbTableItemStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DynamodbTableItemObservation `json:"atProvider"`
}

// A DynamodbTableItemObservation records the observed state of a DynamodbTableItem
type DynamodbTableItemObservation struct{}