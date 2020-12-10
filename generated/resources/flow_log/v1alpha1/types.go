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

// FlowLog is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type FlowLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FlowLogSpec   `json:"spec"`
	Status FlowLogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlowLog contains a list of FlowLogList
type FlowLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlowLog `json:"items"`
}

// A FlowLogSpec defines the desired state of a FlowLog
type FlowLogSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  FlowLogParameters `json:",inline"`
}

// A FlowLogParameters defines the desired state of a FlowLog
type FlowLogParameters struct {
	EniId                  string            `json:"eni_id"`
	IamRoleArn             string            `json:"iam_role_arn"`
	Id                     string            `json:"id"`
	LogDestination         string            `json:"log_destination"`
	LogFormat              string            `json:"log_format"`
	MaxAggregationInterval int64             `json:"max_aggregation_interval"`
	SubnetId               string            `json:"subnet_id"`
	Tags                   map[string]string `json:"tags"`
	LogDestinationType     string            `json:"log_destination_type"`
	LogGroupName           string            `json:"log_group_name"`
	TrafficType            string            `json:"traffic_type"`
	VpcId                  string            `json:"vpc_id"`
}

// A FlowLogStatus defines the observed state of a FlowLog
type FlowLogStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     FlowLogObservation `json:",inline"`
}

// A FlowLogObservation records the observed state of a FlowLog
type FlowLogObservation struct {
	Arn string `json:"arn"`
}