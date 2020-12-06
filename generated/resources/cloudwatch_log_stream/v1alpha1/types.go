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

// CloudwatchLogStream is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudwatchLogStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudwatchLogStreamSpec   `json:"spec"`
	Status CloudwatchLogStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchLogStream contains a list of CloudwatchLogStreamList
type CloudwatchLogStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchLogStream `json:"items"`
}

// A CloudwatchLogStreamSpec defines the desired state of a CloudwatchLogStream
type CloudwatchLogStreamSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudwatchLogStreamParameters `json:",inline"`
}

// A CloudwatchLogStreamParameters defines the desired state of a CloudwatchLogStream
type CloudwatchLogStreamParameters struct {
	LogGroupName string `json:"log_group_name"`
	Name         string `json:"name"`
	Id           string `json:"id"`
}

// A CloudwatchLogStreamStatus defines the observed state of a CloudwatchLogStream
type CloudwatchLogStreamStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudwatchLogStreamObservation `json:",inline"`
}

// A CloudwatchLogStreamObservation records the observed state of a CloudwatchLogStream
type CloudwatchLogStreamObservation struct {
	Arn string `json:"arn"`
}