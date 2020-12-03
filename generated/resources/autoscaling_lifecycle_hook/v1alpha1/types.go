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

// AutoscalingLifecycleHook is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AutoscalingLifecycleHook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AutoscalingLifecycleHookSpec   `json:"spec"`
	Status AutoscalingLifecycleHookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingLifecycleHook contains a list of AutoscalingLifecycleHookList
type AutoscalingLifecycleHookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingLifecycleHook `json:"items"`
}

// A AutoscalingLifecycleHookSpec defines the desired state of a AutoscalingLifecycleHook
type AutoscalingLifecycleHookSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AutoscalingLifecycleHookParameters `json:",inline"`
}

// A AutoscalingLifecycleHookParameters defines the desired state of a AutoscalingLifecycleHook
type AutoscalingLifecycleHookParameters struct {
	Id                    string `json:"id"`
	Name                  string `json:"name"`
	NotificationMetadata  string `json:"notification_metadata"`
	HeartbeatTimeout      int    `json:"heartbeat_timeout"`
	DefaultResult         string `json:"default_result"`
	LifecycleTransition   string `json:"lifecycle_transition"`
	NotificationTargetArn string `json:"notification_target_arn"`
	RoleArn               string `json:"role_arn"`
	AutoscalingGroupName  string `json:"autoscaling_group_name"`
}

// A AutoscalingLifecycleHookStatus defines the observed state of a AutoscalingLifecycleHook
type AutoscalingLifecycleHookStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AutoscalingLifecycleHookObservation `json:",inline"`
}

// A AutoscalingLifecycleHookObservation records the observed state of a AutoscalingLifecycleHook
type AutoscalingLifecycleHookObservation struct{}