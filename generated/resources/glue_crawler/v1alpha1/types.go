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

// GlueCrawler is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlueCrawler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlueCrawlerSpec   `json:"spec"`
	Status GlueCrawlerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueCrawler contains a list of GlueCrawlerList
type GlueCrawlerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueCrawler `json:"items"`
}

// A GlueCrawlerSpec defines the desired state of a GlueCrawler
type GlueCrawlerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlueCrawlerParameters `json:",inline"`
}

// A GlueCrawlerParameters defines the desired state of a GlueCrawler
type GlueCrawlerParameters struct {
	Classifiers           []string           `json:"classifiers"`
	Description           string             `json:"description"`
	TablePrefix           string             `json:"table_prefix"`
	Schedule              string             `json:"schedule"`
	SecurityConfiguration string             `json:"security_configuration"`
	Tags                  map[string]string  `json:"tags"`
	Configuration         string             `json:"configuration"`
	DatabaseName          string             `json:"database_name"`
	Id                    string             `json:"id"`
	Name                  string             `json:"name"`
	Role                  string             `json:"role"`
	CatalogTarget         []CatalogTarget    `json:"catalog_target"`
	DynamodbTarget        []DynamodbTarget   `json:"dynamodb_target"`
	JdbcTarget            []JdbcTarget       `json:"jdbc_target"`
	S3Target              []S3Target         `json:"s3_target"`
	SchemaChangePolicy    SchemaChangePolicy `json:"schema_change_policy"`
}

type CatalogTarget struct {
	DatabaseName string   `json:"database_name"`
	Tables       []string `json:"tables"`
}

type DynamodbTarget struct {
	Path     string `json:"path"`
	ScanAll  bool   `json:"scan_all"`
	ScanRate int    `json:"scan_rate"`
}

type JdbcTarget struct {
	Exclusions     []string `json:"exclusions"`
	Path           string   `json:"path"`
	ConnectionName string   `json:"connection_name"`
}

type S3Target struct {
	ConnectionName string   `json:"connection_name"`
	Exclusions     []string `json:"exclusions"`
	Path           string   `json:"path"`
}

type SchemaChangePolicy struct {
	DeleteBehavior string `json:"delete_behavior"`
	UpdateBehavior string `json:"update_behavior"`
}

// A GlueCrawlerStatus defines the observed state of a GlueCrawler
type GlueCrawlerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlueCrawlerObservation `json:",inline"`
}

// A GlueCrawlerObservation records the observed state of a GlueCrawler
type GlueCrawlerObservation struct {
	Arn string `json:"arn"`
}