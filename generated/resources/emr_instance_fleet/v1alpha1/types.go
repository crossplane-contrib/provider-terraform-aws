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

// EmrInstanceFleet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EmrInstanceFleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmrInstanceFleetSpec   `json:"spec"`
	Status EmrInstanceFleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmrInstanceFleet contains a list of EmrInstanceFleetList
type EmrInstanceFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmrInstanceFleet `json:"items"`
}

// A EmrInstanceFleetSpec defines the desired state of a EmrInstanceFleet
type EmrInstanceFleetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EmrInstanceFleetParameters `json:"forProvider"`
}

// A EmrInstanceFleetParameters defines the desired state of a EmrInstanceFleet
type EmrInstanceFleetParameters struct {
	ClusterId              string               `json:"cluster_id"`
	Id                     string               `json:"id"`
	Name                   string               `json:"name"`
	TargetOnDemandCapacity int64                `json:"target_on_demand_capacity"`
	TargetSpotCapacity     int64                `json:"target_spot_capacity"`
	InstanceTypeConfigs    InstanceTypeConfigs  `json:"instance_type_configs"`
	LaunchSpecifications   LaunchSpecifications `json:"launch_specifications"`
}

type InstanceTypeConfigs struct {
	BidPrice                            string         `json:"bid_price"`
	BidPriceAsPercentageOfOnDemandPrice int64          `json:"bid_price_as_percentage_of_on_demand_price"`
	InstanceType                        string         `json:"instance_type"`
	WeightedCapacity                    int64          `json:"weighted_capacity"`
	Configurations                      Configurations `json:"configurations"`
	EbsConfig                           EbsConfig      `json:"ebs_config"`
}

type Configurations struct {
	Classification string            `json:"classification"`
	Properties     map[string]string `json:"properties"`
}

type EbsConfig struct {
	Iops               int64  `json:"iops"`
	Size               int64  `json:"size"`
	Type               string `json:"type"`
	VolumesPerInstance int64  `json:"volumes_per_instance"`
}

type LaunchSpecifications struct {
	OnDemandSpecification OnDemandSpecification `json:"on_demand_specification"`
	SpotSpecification     SpotSpecification     `json:"spot_specification"`
}

type OnDemandSpecification struct {
	AllocationStrategy string `json:"allocation_strategy"`
}

type SpotSpecification struct {
	AllocationStrategy     string `json:"allocation_strategy"`
	BlockDurationMinutes   int64  `json:"block_duration_minutes"`
	TimeoutAction          string `json:"timeout_action"`
	TimeoutDurationMinutes int64  `json:"timeout_duration_minutes"`
}

// A EmrInstanceFleetStatus defines the observed state of a EmrInstanceFleet
type EmrInstanceFleetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EmrInstanceFleetObservation `json:"atProvider"`
}

// A EmrInstanceFleetObservation records the observed state of a EmrInstanceFleet
type EmrInstanceFleetObservation struct {
	ProvisionedOnDemandCapacity int64 `json:"provisioned_on_demand_capacity"`
	ProvisionedSpotCapacity     int64 `json:"provisioned_spot_capacity"`
}