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

// ElasticacheSecurityGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElasticacheSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticacheSecurityGroupSpec   `json:"spec"`
	Status ElasticacheSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheSecurityGroup contains a list of ElasticacheSecurityGroupList
type ElasticacheSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticacheSecurityGroup `json:"items"`
}

// A ElasticacheSecurityGroupSpec defines the desired state of a ElasticacheSecurityGroup
type ElasticacheSecurityGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElasticacheSecurityGroupParameters `json:"forProvider"`
}

// A ElasticacheSecurityGroupParameters defines the desired state of a ElasticacheSecurityGroup
type ElasticacheSecurityGroupParameters struct {
	Description        string   `json:"description"`
	Name               string   `json:"name"`
	SecurityGroupNames []string `json:"security_group_names"`
}

// A ElasticacheSecurityGroupStatus defines the observed state of a ElasticacheSecurityGroup
type ElasticacheSecurityGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElasticacheSecurityGroupObservation `json:"atProvider"`
}

// A ElasticacheSecurityGroupObservation records the observed state of a ElasticacheSecurityGroup
type ElasticacheSecurityGroupObservation struct{}