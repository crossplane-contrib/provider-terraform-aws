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

// VpnGatewayAttachment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpnGatewayAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpnGatewayAttachmentSpec   `json:"spec"`
	Status VpnGatewayAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpnGatewayAttachment contains a list of VpnGatewayAttachmentList
type VpnGatewayAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpnGatewayAttachment `json:"items"`
}

// A VpnGatewayAttachmentSpec defines the desired state of a VpnGatewayAttachment
type VpnGatewayAttachmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpnGatewayAttachmentParameters `json:"forProvider"`
}

// A VpnGatewayAttachmentParameters defines the desired state of a VpnGatewayAttachment
type VpnGatewayAttachmentParameters struct {
	VpcId        string `json:"vpc_id"`
	VpnGatewayId string `json:"vpn_gateway_id"`
}

// A VpnGatewayAttachmentStatus defines the observed state of a VpnGatewayAttachment
type VpnGatewayAttachmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpnGatewayAttachmentObservation `json:"atProvider"`
}

// A VpnGatewayAttachmentObservation records the observed state of a VpnGatewayAttachment
type VpnGatewayAttachmentObservation struct{}