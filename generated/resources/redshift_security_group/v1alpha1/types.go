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

// RedshiftSecurityGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RedshiftSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedshiftSecurityGroupSpec   `json:"spec"`
	Status RedshiftSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftSecurityGroup contains a list of RedshiftSecurityGroupList
type RedshiftSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedshiftSecurityGroup `json:"items"`
}

// A RedshiftSecurityGroupSpec defines the desired state of a RedshiftSecurityGroup
type RedshiftSecurityGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RedshiftSecurityGroupParameters `json:"forProvider"`
}

// A RedshiftSecurityGroupParameters defines the desired state of a RedshiftSecurityGroup
type RedshiftSecurityGroupParameters struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Id          string    `json:"id"`
	Ingress     []Ingress `json:"ingress"`
}

type Ingress struct {
	SecurityGroupOwnerId string `json:"security_group_owner_id"`
	Cidr                 string `json:"cidr"`
	SecurityGroupName    string `json:"security_group_name"`
}

// A RedshiftSecurityGroupStatus defines the observed state of a RedshiftSecurityGroup
type RedshiftSecurityGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RedshiftSecurityGroupObservation `json:"atProvider"`
}

// A RedshiftSecurityGroupObservation records the observed state of a RedshiftSecurityGroup
type RedshiftSecurityGroupObservation struct{}