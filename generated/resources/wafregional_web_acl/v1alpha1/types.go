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

// WafregionalWebAcl is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafregionalWebAcl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafregionalWebAclSpec   `json:"spec"`
	Status WafregionalWebAclStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalWebAcl contains a list of WafregionalWebAclList
type WafregionalWebAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalWebAcl `json:"items"`
}

// A WafregionalWebAclSpec defines the desired state of a WafregionalWebAcl
type WafregionalWebAclSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafregionalWebAclParameters `json:",inline"`
}

// A WafregionalWebAclParameters defines the desired state of a WafregionalWebAcl
type WafregionalWebAclParameters struct {
	Name       string `json:"name"`
	MetricName string `json:"metric_name"`
}

// A WafregionalWebAclStatus defines the observed state of a WafregionalWebAcl
type WafregionalWebAclStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafregionalWebAclObservation `json:",inline"`
}

// A WafregionalWebAclObservation records the observed state of a WafregionalWebAcl
type WafregionalWebAclObservation struct {
	Arn string `json:"arn"`
	Id  string `json:"id"`
}