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

// VpcEndpointConnectionNotification is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpcEndpointConnectionNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpcEndpointConnectionNotificationSpec   `json:"spec"`
	Status VpcEndpointConnectionNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcEndpointConnectionNotification contains a list of VpcEndpointConnectionNotificationList
type VpcEndpointConnectionNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcEndpointConnectionNotification `json:"items"`
}

// A VpcEndpointConnectionNotificationSpec defines the desired state of a VpcEndpointConnectionNotification
type VpcEndpointConnectionNotificationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpcEndpointConnectionNotificationParameters `json:"forProvider"`
}

// A VpcEndpointConnectionNotificationParameters defines the desired state of a VpcEndpointConnectionNotification
type VpcEndpointConnectionNotificationParameters struct {
	VpcEndpointId             string   `json:"vpc_endpoint_id"`
	VpcEndpointServiceId      string   `json:"vpc_endpoint_service_id"`
	ConnectionEvents          []string `json:"connection_events"`
	ConnectionNotificationArn string   `json:"connection_notification_arn"`
}

// A VpcEndpointConnectionNotificationStatus defines the observed state of a VpcEndpointConnectionNotification
type VpcEndpointConnectionNotificationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpcEndpointConnectionNotificationObservation `json:"atProvider"`
}

// A VpcEndpointConnectionNotificationObservation records the observed state of a VpcEndpointConnectionNotification
type VpcEndpointConnectionNotificationObservation struct {
	NotificationType string `json:"notification_type"`
	State            string `json:"state"`
}