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

// CloudwatchDashboard is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudwatchDashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudwatchDashboardSpec   `json:"spec"`
	Status CloudwatchDashboardStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchDashboard contains a list of CloudwatchDashboardList
type CloudwatchDashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchDashboard `json:"items"`
}

// A CloudwatchDashboardSpec defines the desired state of a CloudwatchDashboard
type CloudwatchDashboardSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudwatchDashboardParameters `json:"forProvider"`
}

// A CloudwatchDashboardParameters defines the desired state of a CloudwatchDashboard
type CloudwatchDashboardParameters struct {
	DashboardName string `json:"dashboard_name"`
	Id            string `json:"id"`
	DashboardBody string `json:"dashboard_body"`
}

// A CloudwatchDashboardStatus defines the observed state of a CloudwatchDashboard
type CloudwatchDashboardStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudwatchDashboardObservation `json:"atProvider"`
}

// A CloudwatchDashboardObservation records the observed state of a CloudwatchDashboard
type CloudwatchDashboardObservation struct {
	DashboardArn string `json:"dashboard_arn"`
}