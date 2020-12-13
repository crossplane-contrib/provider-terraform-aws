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

// LightsailStaticIpAttachment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LightsailStaticIpAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LightsailStaticIpAttachmentSpec   `json:"spec"`
	Status LightsailStaticIpAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LightsailStaticIpAttachment contains a list of LightsailStaticIpAttachmentList
type LightsailStaticIpAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LightsailStaticIpAttachment `json:"items"`
}

// A LightsailStaticIpAttachmentSpec defines the desired state of a LightsailStaticIpAttachment
type LightsailStaticIpAttachmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LightsailStaticIpAttachmentParameters `json:"forProvider"`
}

// A LightsailStaticIpAttachmentParameters defines the desired state of a LightsailStaticIpAttachment
type LightsailStaticIpAttachmentParameters struct {
	InstanceName string `json:"instance_name"`
	StaticIpName string `json:"static_ip_name"`
	Id           string `json:"id"`
}

// A LightsailStaticIpAttachmentStatus defines the observed state of a LightsailStaticIpAttachment
type LightsailStaticIpAttachmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LightsailStaticIpAttachmentObservation `json:"atProvider"`
}

// A LightsailStaticIpAttachmentObservation records the observed state of a LightsailStaticIpAttachment
type LightsailStaticIpAttachmentObservation struct {
	IpAddress string `json:"ip_address"`
}