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

// EmrCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EmrCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmrClusterSpec   `json:"spec"`
	Status EmrClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmrCluster contains a list of EmrClusterList
type EmrClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmrCluster `json:"items"`
}

// A EmrClusterSpec defines the desired state of a EmrCluster
type EmrClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EmrClusterParameters `json:",inline"`
}

// A EmrClusterParameters defines the desired state of a EmrCluster
type EmrClusterParameters struct {
	Configurations        string              `json:"configurations"`
	CustomAmiId           string              `json:"custom_ami_id"`
	StepConcurrencyLevel  int                 `json:"step_concurrency_level"`
	AdditionalInfo        string              `json:"additional_info"`
	Applications          []string            `json:"applications"`
	ServiceRole           string              `json:"service_role"`
	ConfigurationsJson    string              `json:"configurations_json"`
	LogUri                string              `json:"log_uri"`
	Name                  string              `json:"name"`
	ReleaseLabel          string              `json:"release_label"`
	SecurityConfiguration string              `json:"security_configuration"`
	Tags                  map[string]string   `json:"tags"`
	VisibleToAllUsers     bool                `json:"visible_to_all_users"`
	AutoscalingRole       string              `json:"autoscaling_role"`
	EbsRootVolumeSize     int                 `json:"ebs_root_volume_size"`
	KerberosAttributes    KerberosAttributes  `json:"kerberos_attributes"`
	MasterInstanceFleet   MasterInstanceFleet `json:"master_instance_fleet"`
	MasterInstanceGroup   MasterInstanceGroup `json:"master_instance_group"`
	BootstrapAction       []BootstrapAction   `json:"bootstrap_action"`
	CoreInstanceFleet     CoreInstanceFleet   `json:"core_instance_fleet"`
	CoreInstanceGroup     CoreInstanceGroup   `json:"core_instance_group"`
	Ec2Attributes         Ec2Attributes       `json:"ec2_attributes"`
}

type KerberosAttributes struct {
	AdDomainJoinPassword             string `json:"ad_domain_join_password"`
	AdDomainJoinUser                 string `json:"ad_domain_join_user"`
	CrossRealmTrustPrincipalPassword string `json:"cross_realm_trust_principal_password"`
	KdcAdminPassword                 string `json:"kdc_admin_password"`
	Realm                            string `json:"realm"`
}

type MasterInstanceFleet struct {
	Name                        string                `json:"name"`
	ProvisionedOnDemandCapacity int                   `json:"provisioned_on_demand_capacity"`
	ProvisionedSpotCapacity     int                   `json:"provisioned_spot_capacity"`
	TargetOnDemandCapacity      int                   `json:"target_on_demand_capacity"`
	TargetSpotCapacity          int                   `json:"target_spot_capacity"`
	Id                          string                `json:"id"`
	LaunchSpecifications        LaunchSpecifications  `json:"launch_specifications"`
	InstanceTypeConfigs         []InstanceTypeConfigs `json:"instance_type_configs"`
}

type LaunchSpecifications struct {
	OnDemandSpecification []OnDemandSpecification `json:"on_demand_specification"`
	SpotSpecification     []SpotSpecification     `json:"spot_specification"`
}

type OnDemandSpecification struct {
	AllocationStrategy string `json:"allocation_strategy"`
}

type SpotSpecification struct {
	AllocationStrategy     string `json:"allocation_strategy"`
	BlockDurationMinutes   int    `json:"block_duration_minutes"`
	TimeoutAction          string `json:"timeout_action"`
	TimeoutDurationMinutes int    `json:"timeout_duration_minutes"`
}

type InstanceTypeConfigs struct {
	InstanceType                        string           `json:"instance_type"`
	WeightedCapacity                    int              `json:"weighted_capacity"`
	BidPrice                            string           `json:"bid_price"`
	BidPriceAsPercentageOfOnDemandPrice int              `json:"bid_price_as_percentage_of_on_demand_price"`
	Configurations                      []Configurations `json:"configurations"`
	EbsConfig                           []EbsConfig      `json:"ebs_config"`
}

type Configurations struct {
	Classification string            `json:"classification"`
	Properties     map[string]string `json:"properties"`
}

type EbsConfig struct {
	Size               int    `json:"size"`
	Type               string `json:"type"`
	VolumesPerInstance int    `json:"volumes_per_instance"`
	Iops               int    `json:"iops"`
}

type MasterInstanceGroup struct {
	Id            string      `json:"id"`
	InstanceCount int         `json:"instance_count"`
	InstanceType  string      `json:"instance_type"`
	Name          string      `json:"name"`
	BidPrice      string      `json:"bid_price"`
	EbsConfig     []EbsConfig `json:"ebs_config"`
}

type EbsConfig struct {
	Type               string `json:"type"`
	VolumesPerInstance int    `json:"volumes_per_instance"`
	Iops               int    `json:"iops"`
	Size               int    `json:"size"`
}

type BootstrapAction struct {
	Args []string `json:"args"`
	Name string   `json:"name"`
	Path string   `json:"path"`
}

type CoreInstanceFleet struct {
	Name                        string                `json:"name"`
	ProvisionedOnDemandCapacity int                   `json:"provisioned_on_demand_capacity"`
	ProvisionedSpotCapacity     int                   `json:"provisioned_spot_capacity"`
	TargetOnDemandCapacity      int                   `json:"target_on_demand_capacity"`
	TargetSpotCapacity          int                   `json:"target_spot_capacity"`
	Id                          string                `json:"id"`
	InstanceTypeConfigs         []InstanceTypeConfigs `json:"instance_type_configs"`
	LaunchSpecifications        LaunchSpecifications  `json:"launch_specifications"`
}

type InstanceTypeConfigs struct {
	WeightedCapacity                    int              `json:"weighted_capacity"`
	BidPrice                            string           `json:"bid_price"`
	BidPriceAsPercentageOfOnDemandPrice int              `json:"bid_price_as_percentage_of_on_demand_price"`
	InstanceType                        string           `json:"instance_type"`
	Configurations                      []Configurations `json:"configurations"`
	EbsConfig                           []EbsConfig      `json:"ebs_config"`
}

type Configurations struct {
	Classification string            `json:"classification"`
	Properties     map[string]string `json:"properties"`
}

type EbsConfig struct {
	VolumesPerInstance int    `json:"volumes_per_instance"`
	Iops               int    `json:"iops"`
	Size               int    `json:"size"`
	Type               string `json:"type"`
}

type LaunchSpecifications struct {
	SpotSpecification     []SpotSpecification     `json:"spot_specification"`
	OnDemandSpecification []OnDemandSpecification `json:"on_demand_specification"`
}

type SpotSpecification struct {
	TimeoutAction          string `json:"timeout_action"`
	TimeoutDurationMinutes int    `json:"timeout_duration_minutes"`
	AllocationStrategy     string `json:"allocation_strategy"`
	BlockDurationMinutes   int    `json:"block_duration_minutes"`
}

type OnDemandSpecification struct {
	AllocationStrategy string `json:"allocation_strategy"`
}

type CoreInstanceGroup struct {
	BidPrice          string      `json:"bid_price"`
	Id                string      `json:"id"`
	InstanceCount     int         `json:"instance_count"`
	InstanceType      string      `json:"instance_type"`
	Name              string      `json:"name"`
	AutoscalingPolicy string      `json:"autoscaling_policy"`
	EbsConfig         []EbsConfig `json:"ebs_config"`
}

type EbsConfig struct {
	Iops               int    `json:"iops"`
	Size               int    `json:"size"`
	Type               string `json:"type"`
	VolumesPerInstance int    `json:"volumes_per_instance"`
}

type Ec2Attributes struct {
	AdditionalSlaveSecurityGroups  string `json:"additional_slave_security_groups"`
	EmrManagedMasterSecurityGroup  string `json:"emr_managed_master_security_group"`
	EmrManagedSlaveSecurityGroup   string `json:"emr_managed_slave_security_group"`
	InstanceProfile                string `json:"instance_profile"`
	KeyName                        string `json:"key_name"`
	ServiceAccessSecurityGroup     string `json:"service_access_security_group"`
	SubnetId                       string `json:"subnet_id"`
	AdditionalMasterSecurityGroups string `json:"additional_master_security_groups"`
}

// A EmrClusterStatus defines the observed state of a EmrCluster
type EmrClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EmrClusterObservation `json:",inline"`
}

// A EmrClusterObservation records the observed state of a EmrCluster
type EmrClusterObservation struct {
	Id                          string `json:"id"`
	KeepJobFlowAliveWhenNoSteps bool   `json:"keep_job_flow_alive_when_no_steps"`
	MasterPublicDns             string `json:"master_public_dns"`
	Arn                         string `json:"arn"`
	ScaleDownBehavior           string `json:"scale_down_behavior"`
	TerminationProtection       bool   `json:"termination_protection"`
	ClusterState                string `json:"cluster_state"`
}