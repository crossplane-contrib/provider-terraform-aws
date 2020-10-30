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

// WafWebAcl is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafWebAcl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafWebAclSpec   `json:"spec"`
	Status WafWebAclStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafWebAcl contains a list of WafWebAclList
type WafWebAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafWebAcl `json:"items"`
}

// A WafWebAclSpec defines the desired state of a WafWebAcl
type WafWebAclSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafWebAclParameters `json:",inline"`
}

// A WafWebAclParameters defines the desired state of a WafWebAcl
type WafWebAclParameters struct {
	MetricName string `json:"metric_name"`
	Name       string `json:"name"`
}

// A WafWebAclStatus defines the observed state of a WafWebAcl
type WafWebAclStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafWebAclObservation `json:",inline"`
}

// A WafWebAclObservation records the observed state of a WafWebAcl
type WafWebAclObservation struct {
	Arn string `json:"arn"`
	Id  string `json:"id"`
}