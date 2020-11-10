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
	ExcessCapacityTerminationPolicy  string                 `json:"excess_capacity_termination_policy"`
	IamFleetRole                     string                 `json:"iam_fleet_role"`
	AllocationStrategy               string                 `json:"allocation_strategy"`
	TargetCapacity                   int                    `json:"target_capacity"`
	ValidFrom                        string                 `json:"valid_from"`
	ValidUntil                       string                 `json:"valid_until"`
	FleetType                        string                 `json:"fleet_type"`
	InstancePoolsToUseCount          int                    `json:"instance_pools_to_use_count"`
	ReplaceUnhealthyInstances        bool                   `json:"replace_unhealthy_instances"`
	Tags                             map[string]string      `json:"tags"`
	TerminateInstancesWithExpiration bool                   `json:"terminate_instances_with_expiration"`
	InstanceInterruptionBehaviour    string                 `json:"instance_interruption_behaviour"`
	SpotPrice                        string                 `json:"spot_price"`
	WaitForFulfillment               bool                   `json:"wait_for_fulfillment"`
	LaunchSpecification              []LaunchSpecification  `json:"launch_specification"`
	LaunchTemplateConfig             []LaunchTemplateConfig `json:"launch_template_config"`
	Timeouts                         []Timeouts             `json:"timeouts"`
}

type LaunchSpecification struct {
	EbsOptimized             bool                   `json:"ebs_optimized"`
	InstanceType             string                 `json:"instance_type"`
	SpotPrice                string                 `json:"spot_price"`
	Monitoring               bool                   `json:"monitoring"`
	SubnetId                 string                 `json:"subnet_id"`
	Tags                     map[string]string      `json:"tags"`
	PlacementGroup           string                 `json:"placement_group"`
	UserData                 string                 `json:"user_data"`
	Ami                      string                 `json:"ami"`
	AssociatePublicIpAddress bool                   `json:"associate_public_ip_address"`
	AvailabilityZone         string                 `json:"availability_zone"`
	IamInstanceProfile       string                 `json:"iam_instance_profile"`
	IamInstanceProfileArn    string                 `json:"iam_instance_profile_arn"`
	KeyName                  string                 `json:"key_name"`
	VpcSecurityGroupIds      []string               `json:"vpc_security_group_ids"`
	PlacementTenancy         string                 `json:"placement_tenancy"`
	WeightedCapacity         string                 `json:"weighted_capacity"`
	RootBlockDevice          []RootBlockDevice      `json:"root_block_device"`
	EbsBlockDevice           []EbsBlockDevice       `json:"ebs_block_device"`
	EphemeralBlockDevice     []EphemeralBlockDevice `json:"ephemeral_block_device"`
}

type RootBlockDevice struct {
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type EbsBlockDevice struct {
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type EphemeralBlockDevice struct {
	VirtualName string `json:"virtual_name"`
	DeviceName  string `json:"device_name"`
}

type LaunchTemplateConfig struct {
	LaunchTemplateSpecification LaunchTemplateSpecification `json:"launch_template_specification"`
	Overrides                   []Overrides                 `json:"overrides"`
}

type LaunchTemplateSpecification struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Overrides struct {
	Priority         int    `json:"priority"`
	SpotPrice        string `json:"spot_price"`
	SubnetId         string `json:"subnet_id"`
	WeightedCapacity int    `json:"weighted_capacity"`
	AvailabilityZone string `json:"availability_zone"`
	InstanceType     string `json:"instance_type"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A SpotFleetRequestStatus defines the observed state of a SpotFleetRequest
type SpotFleetRequestStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SpotFleetRequestObservation `json:",inline"`
}

// A SpotFleetRequestObservation records the observed state of a SpotFleetRequest
type SpotFleetRequestObservation struct {
	ClientToken      string   `json:"client_token"`
	TargetGroupArns  []string `json:"target_group_arns"`
	SpotRequestState string   `json:"spot_request_state"`
	Id               string   `json:"id"`
	LoadBalancers    []string `json:"load_balancers"`
}