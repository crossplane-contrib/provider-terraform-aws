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
	StreamViewType       string               `json:"stream_view_type"`
	HashKey              string               `json:"hash_key"`
	Name                 string               `json:"name"`
	ReadCapacity         int64                `json:"read_capacity"`
	StreamEnabled        bool                 `json:"stream_enabled"`
	Tags                 map[string]string    `json:"tags"`
	WriteCapacity        int64                `json:"write_capacity"`
	BillingMode          string               `json:"billing_mode"`
	Id                   string               `json:"id"`
	RangeKey             string               `json:"range_key"`
	Replica              Replica              `json:"replica"`
	ServerSideEncryption ServerSideEncryption `json:"server_side_encryption"`
	Timeouts             Timeouts             `json:"timeouts"`
	Ttl                  Ttl                  `json:"ttl"`
	Attribute            []Attribute          `json:"attribute"`
	GlobalSecondaryIndex GlobalSecondaryIndex `json:"global_secondary_index"`
	LocalSecondaryIndex  LocalSecondaryIndex  `json:"local_secondary_index"`
	PointInTimeRecovery  PointInTimeRecovery  `json:"point_in_time_recovery"`
}

type Replica struct {
	RegionName string `json:"region_name"`
}

type ServerSideEncryption struct {
	Enabled   bool   `json:"enabled"`
	KmsKeyArn string `json:"kms_key_arn"`
}

type Timeouts struct {
	Update string `json:"update"`
	Create string `json:"create"`
	Delete string `json:"delete"`
}

type Ttl struct {
	AttributeName string `json:"attribute_name"`
	Enabled       bool   `json:"enabled"`
}

type Attribute struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type GlobalSecondaryIndex struct {
	NonKeyAttributes []string `json:"non_key_attributes"`
	ProjectionType   string   `json:"projection_type"`
	RangeKey         string   `json:"range_key"`
	ReadCapacity     int64    `json:"read_capacity"`
	WriteCapacity    int64    `json:"write_capacity"`
	HashKey          string   `json:"hash_key"`
	Name             string   `json:"name"`
}

type LocalSecondaryIndex struct {
	RangeKey         string   `json:"range_key"`
	Name             string   `json:"name"`
	NonKeyAttributes []string `json:"non_key_attributes"`
	ProjectionType   string   `json:"projection_type"`
}

type PointInTimeRecovery struct {
	Enabled bool `json:"enabled"`
}

// A DynamodbTableStatus defines the observed state of a DynamodbTable
type DynamodbTableStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DynamodbTableObservation `json:",inline"`
}

// A DynamodbTableObservation records the observed state of a DynamodbTable
type DynamodbTableObservation struct {
	StreamArn   string `json:"stream_arn"`
	Arn         string `json:"arn"`
	StreamLabel string `json:"stream_label"`
}