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

// EipAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EipAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EipAssociationSpec   `json:"spec"`
	Status EipAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EipAssociation contains a list of EipAssociationList
type EipAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EipAssociation `json:"items"`
}

// A EipAssociationSpec defines the desired state of a EipAssociation
type EipAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EipAssociationParameters `json:"forProvider"`
}

// A EipAssociationParameters defines the desired state of a EipAssociation
type EipAssociationParameters struct {
	AllowReassociation bool   `json:"allow_reassociation"`
	Id                 string `json:"id"`
	InstanceId         string `json:"instance_id"`
	NetworkInterfaceId string `json:"network_interface_id"`
	PrivateIpAddress   string `json:"private_ip_address"`
	PublicIp           string `json:"public_ip"`
	AllocationId       string `json:"allocation_id"`
}

// A EipAssociationStatus defines the observed state of a EipAssociation
type EipAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EipAssociationObservation `json:"atProvider"`
}

// A EipAssociationObservation records the observed state of a EipAssociation
type EipAssociationObservation struct{}