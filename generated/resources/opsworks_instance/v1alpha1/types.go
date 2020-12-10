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

// OpsworksInstance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksInstanceSpec   `json:"spec"`
	Status OpsworksInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksInstance contains a list of OpsworksInstanceList
type OpsworksInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksInstance `json:"items"`
}

// A OpsworksInstanceSpec defines the desired state of a OpsworksInstance
type OpsworksInstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksInstanceParameters `json:",inline"`
}

// A OpsworksInstanceParameters defines the desired state of a OpsworksInstance
type OpsworksInstanceParameters struct {
	Tenancy                  string               `json:"tenancy"`
	InstanceType             string               `json:"instance_type"`
	PrivateDns               string               `json:"private_dns"`
	ReportedOsFamily         string               `json:"reported_os_family"`
	SshHostRsaKeyFingerprint string               `json:"ssh_host_rsa_key_fingerprint"`
	DeleteEip                bool                 `json:"delete_eip"`
	InstanceProfileArn       string               `json:"instance_profile_arn"`
	Platform                 string               `json:"platform"`
	PublicDns                string               `json:"public_dns"`
	RootDeviceType           string               `json:"root_device_type"`
	Architecture             string               `json:"architecture"`
	AutoScalingType          string               `json:"auto_scaling_type"`
	Hostname                 string               `json:"hostname"`
	LayerIds                 []string             `json:"layer_ids"`
	RegisteredBy             string               `json:"registered_by"`
	ReportedAgentVersion     string               `json:"reported_agent_version"`
	ReportedOsName           string               `json:"reported_os_name"`
	ReportedOsVersion        string               `json:"reported_os_version"`
	SecurityGroupIds         []string             `json:"security_group_ids"`
	SubnetId                 string               `json:"subnet_id"`
	AmiId                    string               `json:"ami_id"`
	CreatedAt                string               `json:"created_at"`
	DeleteEbs                bool                 `json:"delete_ebs"`
	RootDeviceVolumeId       string               `json:"root_device_volume_id"`
	StackId                  string               `json:"stack_id"`
	AvailabilityZone         string               `json:"availability_zone"`
	EbsOptimized             bool                 `json:"ebs_optimized"`
	EcsClusterArn            string               `json:"ecs_cluster_arn"`
	InfrastructureClass      string               `json:"infrastructure_class"`
	LastServiceErrorId       string               `json:"last_service_error_id"`
	Status                   string               `json:"status"`
	ElasticIp                string               `json:"elastic_ip"`
	Os                       string               `json:"os"`
	PublicIp                 string               `json:"public_ip"`
	SshKeyName               string               `json:"ssh_key_name"`
	State                    string               `json:"state"`
	VirtualizationType       string               `json:"virtualization_type"`
	AgentVersion             string               `json:"agent_version"`
	Id                       string               `json:"id"`
	InstallUpdatesOnBoot     bool                 `json:"install_updates_on_boot"`
	PrivateIp                string               `json:"private_ip"`
	SshHostDsaKeyFingerprint string               `json:"ssh_host_dsa_key_fingerprint"`
	EphemeralBlockDevice     EphemeralBlockDevice `json:"ephemeral_block_device"`
	RootBlockDevice          RootBlockDevice      `json:"root_block_device"`
	Timeouts                 Timeouts             `json:"timeouts"`
	EbsBlockDevice           EbsBlockDevice       `json:"ebs_block_device"`
}

type EphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type RootBlockDevice struct {
	Iops                int64  `json:"iops"`
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type EbsBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Iops                int64  `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

// A OpsworksInstanceStatus defines the observed state of a OpsworksInstance
type OpsworksInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksInstanceObservation `json:",inline"`
}

// A OpsworksInstanceObservation records the observed state of a OpsworksInstance
type OpsworksInstanceObservation struct {
	Ec2InstanceId string `json:"ec2_instance_id"`
}