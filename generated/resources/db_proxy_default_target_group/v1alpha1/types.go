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

// DbProxyDefaultTargetGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DbProxyDefaultTargetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbProxyDefaultTargetGroupSpec   `json:"spec"`
	Status DbProxyDefaultTargetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbProxyDefaultTargetGroup contains a list of DbProxyDefaultTargetGroupList
type DbProxyDefaultTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbProxyDefaultTargetGroup `json:"items"`
}

// A DbProxyDefaultTargetGroupSpec defines the desired state of a DbProxyDefaultTargetGroup
type DbProxyDefaultTargetGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DbProxyDefaultTargetGroupParameters `json:",inline"`
}

// A DbProxyDefaultTargetGroupParameters defines the desired state of a DbProxyDefaultTargetGroup
type DbProxyDefaultTargetGroupParameters struct {
	Id                   string               `json:"id"`
	DbProxyName          string               `json:"db_proxy_name"`
	ConnectionPoolConfig ConnectionPoolConfig `json:"connection_pool_config"`
	Timeouts             Timeouts             `json:"timeouts"`
}

type ConnectionPoolConfig struct {
	ConnectionBorrowTimeout   int64    `json:"connection_borrow_timeout"`
	InitQuery                 string   `json:"init_query"`
	MaxConnectionsPercent     int64    `json:"max_connections_percent"`
	MaxIdleConnectionsPercent int64    `json:"max_idle_connections_percent"`
	SessionPinningFilters     []string `json:"session_pinning_filters"`
}

type Timeouts struct {
	Create string `json:"create"`
	Update string `json:"update"`
}

// A DbProxyDefaultTargetGroupStatus defines the observed state of a DbProxyDefaultTargetGroup
type DbProxyDefaultTargetGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DbProxyDefaultTargetGroupObservation `json:",inline"`
}

// A DbProxyDefaultTargetGroupObservation records the observed state of a DbProxyDefaultTargetGroup
type DbProxyDefaultTargetGroupObservation struct {
	Name string `json:"name"`
	Arn  string `json:"arn"`
}