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

// DynamodbTable is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DynamodbTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DynamodbTableSpec   `json:"spec"`
	Status DynamodbTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DynamodbTable contains a list of DynamodbTableList
type DynamodbTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DynamodbTable `json:"items"`
}

// A DynamodbTableSpec defines the desired state of a DynamodbTable
type DynamodbTableSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DynamodbTableParameters `json:",inline"`
}

// A DynamodbTableParameters defines the desired state of a DynamodbTable
type DynamodbTableParameters struct {
	BillingMode   string `json:"billing_mode"`
	Name          string `json:"name"`
	RangeKey      string `json:"range_key"`
	ReadCapacity  int    `json:"read_capacity"`
	WriteCapacity int    `json:"write_capacity"`
	HashKey       string `json:"hash_key"`
	StreamEnabled bool   `json:"stream_enabled"`
}

// A DynamodbTableStatus defines the observed state of a DynamodbTable
type DynamodbTableStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DynamodbTableObservation `json:",inline"`
}

// A DynamodbTableObservation records the observed state of a DynamodbTable
type DynamodbTableObservation struct {
	Id             string `json:"id"`
	StreamLabel    string `json:"stream_label"`
	Arn            string `json:"arn"`
	StreamArn      string `json:"stream_arn"`
	StreamViewType string `json:"stream_view_type"`
}