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
	Name                        string              `json:"name"`
	TerminationProtection       bool                `json:"termination_protection"`
	ConfigurationsJson          string              `json:"configurations_json"`
	CustomAmiId                 string              `json:"custom_ami_id"`
	EbsRootVolumeSize           int64               `json:"ebs_root_volume_size"`
	Id                          string              `json:"id"`
	ReleaseLabel                string              `json:"release_label"`
	ScaleDownBehavior           string              `json:"scale_down_behavior"`
	AdditionalInfo              string              `json:"additional_info"`
	AutoscalingRole             string              `json:"autoscaling_role"`
	VisibleToAllUsers           bool                `json:"visible_to_all_users"`
	Step                        []Step              `json:"step"`
	StepConcurrencyLevel        int64               `json:"step_concurrency_level"`
	Configurations              string              `json:"configurations"`
	KeepJobFlowAliveWhenNoSteps bool                `json:"keep_job_flow_alive_when_no_steps"`
	LogUri                      string              `json:"log_uri"`
	SecurityConfiguration       string              `json:"security_configuration"`
	ServiceRole                 string              `json:"service_role"`
	Applications                []string            `json:"applications"`
	Tags                        map[string]string   `json:"tags"`
	CoreInstanceFleet           CoreInstanceFleet   `json:"core_instance_fleet"`
	CoreInstanceGroup           CoreInstanceGroup   `json:"core_instance_group"`
	Ec2Attributes               Ec2Attributes       `json:"ec2_attributes"`
	KerberosAttributes          KerberosAttributes  `json:"kerberos_attributes"`
	MasterInstanceFleet         MasterInstanceFleet `json:"master_instance_fleet"`
	MasterInstanceGroup         MasterInstanceGroup `json:"master_instance_group"`
	BootstrapAction             BootstrapAction     `json:"bootstrap_action"`
}

type Step struct {
	ActionOnFailure string          `json:"action_on_failure"`
	HadoopJarStep   []HadoopJarStep `json:"hadoop_jar_step"`
	Name            string          `json:"name"`
}

type HadoopJarStep struct {
	MainClass  string            `json:"main_class"`
	Properties map[string]string `json:"properties"`
	Args       []string          `json:"args"`
	Jar        string            `json:"jar"`
}

type CoreInstanceFleet struct {
	TargetSpotCapacity          int64                `json:"target_spot_capacity"`
	Id                          string               `json:"id"`
	Name                        string               `json:"name"`
	ProvisionedOnDemandCapacity int64                `json:"provisioned_on_demand_capacity"`
	ProvisionedSpotCapacity     int64                `json:"provisioned_spot_capacity"`
	TargetOnDemandCapacity      int64                `json:"target_on_demand_capacity"`
	InstanceTypeConfigs         InstanceTypeConfigs  `json:"instance_type_configs"`
	LaunchSpecifications        LaunchSpecifications `json:"launch_specifications"`
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
	Properties     map[string]string `json:"properties"`
	Classification string            `json:"classification"`
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
	TimeoutAction          string `json:"timeout_action"`
	TimeoutDurationMinutes int64  `json:"timeout_duration_minutes"`
	AllocationStrategy     string `json:"allocation_strategy"`
	BlockDurationMinutes   int64  `json:"block_duration_minutes"`
}

type CoreInstanceGroup struct {
	BidPrice          string    `json:"bid_price"`
	Id                string    `json:"id"`
	InstanceCount     int64     `json:"instance_count"`
	InstanceType      string    `json:"instance_type"`
	Name              string    `json:"name"`
	AutoscalingPolicy string    `json:"autoscaling_policy"`
	EbsConfig         EbsConfig `json:"ebs_config"`
}

type EbsConfig struct {
	Iops               int64  `json:"iops"`
	Size               int64  `json:"size"`
	Type               string `json:"type"`
	VolumesPerInstance int64  `json:"volumes_per_instance"`
}

type Ec2Attributes struct {
	ServiceAccessSecurityGroup     string `json:"service_access_security_group"`
	SubnetId                       string `json:"subnet_id"`
	AdditionalMasterSecurityGroups string `json:"additional_master_security_groups"`
	AdditionalSlaveSecurityGroups  string `json:"additional_slave_security_groups"`
	EmrManagedMasterSecurityGroup  string `json:"emr_managed_master_security_group"`
	EmrManagedSlaveSecurityGroup   string `json:"emr_managed_slave_security_group"`
	InstanceProfile                string `json:"instance_profile"`
	KeyName                        string `json:"key_name"`
}

type KerberosAttributes struct {
	Realm                            string `json:"realm"`
	AdDomainJoinPassword             string `json:"ad_domain_join_password"`
	AdDomainJoinUser                 string `json:"ad_domain_join_user"`
	CrossRealmTrustPrincipalPassword string `json:"cross_realm_trust_principal_password"`
	KdcAdminPassword                 string `json:"kdc_admin_password"`
}

type MasterInstanceFleet struct {
	ProvisionedSpotCapacity     int64                `json:"provisioned_spot_capacity"`
	TargetOnDemandCapacity      int64                `json:"target_on_demand_capacity"`
	TargetSpotCapacity          int64                `json:"target_spot_capacity"`
	Id                          string               `json:"id"`
	Name                        string               `json:"name"`
	ProvisionedOnDemandCapacity int64                `json:"provisioned_on_demand_capacity"`
	InstanceTypeConfigs         InstanceTypeConfigs  `json:"instance_type_configs"`
	LaunchSpecifications        LaunchSpecifications `json:"launch_specifications"`
}

type InstanceTypeConfigs struct {
	InstanceType                        string         `json:"instance_type"`
	WeightedCapacity                    int64          `json:"weighted_capacity"`
	BidPrice                            string         `json:"bid_price"`
	BidPriceAsPercentageOfOnDemandPrice int64          `json:"bid_price_as_percentage_of_on_demand_price"`
	EbsConfig                           EbsConfig      `json:"ebs_config"`
	Configurations                      Configurations `json:"configurations"`
}

type EbsConfig struct {
	Iops               int64  `json:"iops"`
	Size               int64  `json:"size"`
	Type               string `json:"type"`
	VolumesPerInstance int64  `json:"volumes_per_instance"`
}

type Configurations struct {
	Classification string            `json:"classification"`
	Properties     map[string]string `json:"properties"`
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

type MasterInstanceGroup struct {
	InstanceType  string    `json:"instance_type"`
	Name          string    `json:"name"`
	BidPrice      string    `json:"bid_price"`
	Id            string    `json:"id"`
	InstanceCount int64     `json:"instance_count"`
	EbsConfig     EbsConfig `json:"ebs_config"`
}

type EbsConfig struct {
	Size               int64  `json:"size"`
	Type               string `json:"type"`
	VolumesPerInstance int64  `json:"volumes_per_instance"`
	Iops               int64  `json:"iops"`
}

type BootstrapAction struct {
	Args []string `json:"args"`
	Name string   `json:"name"`
	Path string   `json:"path"`
}

// A EmrClusterStatus defines the observed state of a EmrCluster
type EmrClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EmrClusterObservation `json:",inline"`
}

// A EmrClusterObservation records the observed state of a EmrCluster
type EmrClusterObservation struct {
	ClusterState    string `json:"cluster_state"`
	Arn             string `json:"arn"`
	MasterPublicDns string `json:"master_public_dns"`
}