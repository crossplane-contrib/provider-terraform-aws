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
	FleetType                        string               `json:"fleet_type"`
	IamFleetRole                     string               `json:"iam_fleet_role"`
	SpotPrice                        string               `json:"spot_price"`
	Id                               string               `json:"id"`
	InstanceInterruptionBehaviour    string               `json:"instance_interruption_behaviour"`
	ReplaceUnhealthyInstances        bool                 `json:"replace_unhealthy_instances"`
	ValidFrom                        string               `json:"valid_from"`
	ValidUntil                       string               `json:"valid_until"`
	AllocationStrategy               string               `json:"allocation_strategy"`
	LoadBalancers                    []string             `json:"load_balancers"`
	Tags                             map[string]string    `json:"tags"`
	TargetCapacity                   int64                `json:"target_capacity"`
	TargetGroupArns                  []string             `json:"target_group_arns"`
	TerminateInstancesWithExpiration bool                 `json:"terminate_instances_with_expiration"`
	WaitForFulfillment               bool                 `json:"wait_for_fulfillment"`
	ExcessCapacityTerminationPolicy  string               `json:"excess_capacity_termination_policy"`
	InstancePoolsToUseCount          int64                `json:"instance_pools_to_use_count"`
	LaunchSpecification              LaunchSpecification  `json:"launch_specification"`
	LaunchTemplateConfig             LaunchTemplateConfig `json:"launch_template_config"`
	Timeouts                         Timeouts             `json:"timeouts"`
}

type LaunchSpecification struct {
	Ami                      string               `json:"ami"`
	PlacementGroup           string               `json:"placement_group"`
	VpcSecurityGroupIds      []string             `json:"vpc_security_group_ids"`
	Monitoring               bool                 `json:"monitoring"`
	SpotPrice                string               `json:"spot_price"`
	Tags                     map[string]string    `json:"tags"`
	AssociatePublicIpAddress bool                 `json:"associate_public_ip_address"`
	IamInstanceProfile       string               `json:"iam_instance_profile"`
	IamInstanceProfileArn    string               `json:"iam_instance_profile_arn"`
	SubnetId                 string               `json:"subnet_id"`
	UserData                 string               `json:"user_data"`
	KeyName                  string               `json:"key_name"`
	PlacementTenancy         string               `json:"placement_tenancy"`
	WeightedCapacity         string               `json:"weighted_capacity"`
	AvailabilityZone         string               `json:"availability_zone"`
	EbsOptimized             bool                 `json:"ebs_optimized"`
	InstanceType             string               `json:"instance_type"`
	EphemeralBlockDevice     EphemeralBlockDevice `json:"ephemeral_block_device"`
	RootBlockDevice          RootBlockDevice      `json:"root_block_device"`
	EbsBlockDevice           EbsBlockDevice       `json:"ebs_block_device"`
}

type EphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type RootBlockDevice struct {
	Encrypted           bool   `json:"encrypted"`
	Iops                int64  `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type EbsBlockDevice struct {
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int64  `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
	SnapshotId          string `json:"snapshot_id"`
}

type LaunchTemplateConfig struct {
	Overrides                   Overrides                   `json:"overrides"`
	LaunchTemplateSpecification LaunchTemplateSpecification `json:"launch_template_specification"`
}

type Overrides struct {
	AvailabilityZone string `json:"availability_zone"`
	InstanceType     string `json:"instance_type"`
	Priority         int64  `json:"priority"`
	SpotPrice        string `json:"spot_price"`
	SubnetId         string `json:"subnet_id"`
	WeightedCapacity int64  `json:"weighted_capacity"`
}

type LaunchTemplateSpecification struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Id      string `json:"id"`
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
	SpotRequestState string `json:"spot_request_state"`
	ClientToken      string `json:"client_token"`
}