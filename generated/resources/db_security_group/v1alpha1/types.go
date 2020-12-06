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

// DbSecurityGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DbSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbSecurityGroupSpec   `json:"spec"`
	Status DbSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbSecurityGroup contains a list of DbSecurityGroupList
type DbSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbSecurityGroup `json:"items"`
}

// A DbSecurityGroupSpec defines the desired state of a DbSecurityGroup
type DbSecurityGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DbSecurityGroupParameters `json:",inline"`
}

// A DbSecurityGroupParameters defines the desired state of a DbSecurityGroup
type DbSecurityGroupParameters struct {
	Tags        map[string]string `json:"tags"`
	Description string            `json:"description"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Ingress     []Ingress         `json:"ingress"`
}

type Ingress struct {
	Cidr                 string `json:"cidr"`
	SecurityGroupId      string `json:"security_group_id"`
	SecurityGroupName    string `json:"security_group_name"`
	SecurityGroupOwnerId string `json:"security_group_owner_id"`
}

// A DbSecurityGroupStatus defines the observed state of a DbSecurityGroup
type DbSecurityGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DbSecurityGroupObservation `json:",inline"`
}

// A DbSecurityGroupObservation records the observed state of a DbSecurityGroup
type DbSecurityGroupObservation struct {
	Arn string `json:"arn"`
}