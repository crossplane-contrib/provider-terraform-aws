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

// AutoscalingNotification is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AutoscalingNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AutoscalingNotificationSpec   `json:"spec"`
	Status AutoscalingNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingNotification contains a list of AutoscalingNotificationList
type AutoscalingNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingNotification `json:"items"`
}

// A AutoscalingNotificationSpec defines the desired state of a AutoscalingNotification
type AutoscalingNotificationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AutoscalingNotificationParameters `json:",inline"`
}

// A AutoscalingNotificationParameters defines the desired state of a AutoscalingNotification
type AutoscalingNotificationParameters struct {
	TopicArn      string   `json:"topic_arn"`
	GroupNames    []string `json:"group_names"`
	Notifications []string `json:"notifications"`
}

// A AutoscalingNotificationStatus defines the observed state of a AutoscalingNotification
type AutoscalingNotificationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AutoscalingNotificationObservation `json:",inline"`
}

// A AutoscalingNotificationObservation records the observed state of a AutoscalingNotification
type AutoscalingNotificationObservation struct {
	Id string `json:"id"`
}