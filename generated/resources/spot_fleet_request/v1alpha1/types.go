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

// SpotFleetRequest is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SpotFleetRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpotFleetRequestSpec   `json:"spec"`
	Status SpotFleetRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpotFleetRequest contains a list of SpotFleetRequestList
type SpotFleetRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotFleetRequest `json:"items"`
}

// A SpotFleetRequestSpec defines the desired state of a SpotFleetRequest
type SpotFleetRequestSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SpotFleetRequestParameters `json:",inline"`
}

// A SpotFleetRequestParameters defines the desired state of a SpotFleetRequest
type SpotFleetRequestParameters struct {
	TerminateInstancesWithExpiration bool   `json:"terminate_instances_with_expiration"`
	InstanceInterruptionBehaviour    string `json:"instance_interruption_behaviour"`
	InstancePoolsToUseCount          int    `json:"instance_pools_to_use_count"`
	ValidFrom                        string `json:"valid_from"`
	ExcessCapacityTerminationPolicy  string `json:"excess_capacity_termination_policy"`
	FleetType                        string `json:"fleet_type"`
	WaitForFulfillment               bool   `json:"wait_for_fulfillment"`
	SpotPrice                        string `json:"spot_price"`
	ReplaceUnhealthyInstances        bool   `json:"replace_unhealthy_instances"`
	TargetCapacity                   int    `json:"target_capacity"`
	ValidUntil                       string `json:"valid_until"`
	AllocationStrategy               string `json:"allocation_strategy"`
	IamFleetRole                     string `json:"iam_fleet_role"`
}

// A SpotFleetRequestStatus defines the observed state of a SpotFleetRequest
type SpotFleetRequestStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SpotFleetRequestObservation `json:",inline"`
}

// A SpotFleetRequestObservation records the observed state of a SpotFleetRequest
type SpotFleetRequestObservation struct {
	LoadBalancers    []string `json:"load_balancers"`
	SpotRequestState string   `json:"spot_request_state"`
	ClientToken      string   `json:"client_token"`
	Id               string   `json:"id"`
	TargetGroupArns  []string `json:"target_group_arns"`
}