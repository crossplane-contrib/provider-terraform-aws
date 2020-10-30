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
	Architecture         string   `json:"architecture"`
	InstallUpdatesOnBoot bool     `json:"install_updates_on_boot"`
	EbsOptimized         bool     `json:"ebs_optimized"`
	DeleteEbs            bool     `json:"delete_ebs"`
	DeleteEip            bool     `json:"delete_eip"`
	LayerIds             []string `json:"layer_ids"`
	State                string   `json:"state"`
	AutoScalingType      string   `json:"auto_scaling_type"`
	InstanceType         string   `json:"instance_type"`
	StackId              string   `json:"stack_id"`
	AgentVersion         string   `json:"agent_version"`
}

// A OpsworksInstanceStatus defines the observed state of a OpsworksInstance
type OpsworksInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksInstanceObservation `json:",inline"`
}

// A OpsworksInstanceObservation records the observed state of a OpsworksInstance
type OpsworksInstanceObservation struct {
	EcsClusterArn            string   `json:"ecs_cluster_arn"`
	SshHostDsaKeyFingerprint string   `json:"ssh_host_dsa_key_fingerprint"`
	VirtualizationType       string   `json:"virtualization_type"`
	RootDeviceType           string   `json:"root_device_type"`
	InstanceProfileArn       string   `json:"instance_profile_arn"`
	Os                       string   `json:"os"`
	ReportedOsFamily         string   `json:"reported_os_family"`
	ReportedOsVersion        string   `json:"reported_os_version"`
	ReportedOsName           string   `json:"reported_os_name"`
	PrivateDns               string   `json:"private_dns"`
	PrivateIp                string   `json:"private_ip"`
	RegisteredBy             string   `json:"registered_by"`
	AvailabilityZone         string   `json:"availability_zone"`
	Ec2InstanceId            string   `json:"ec2_instance_id"`
	SshKeyName               string   `json:"ssh_key_name"`
	SecurityGroupIds         []string `json:"security_group_ids"`
	Status                   string   `json:"status"`
	AmiId                    string   `json:"ami_id"`
	ReportedAgentVersion     string   `json:"reported_agent_version"`
	RootDeviceVolumeId       string   `json:"root_device_volume_id"`
	Tenancy                  string   `json:"tenancy"`
	CreatedAt                string   `json:"created_at"`
	Hostname                 string   `json:"hostname"`
	Id                       string   `json:"id"`
	LastServiceErrorId       string   `json:"last_service_error_id"`
	SshHostRsaKeyFingerprint string   `json:"ssh_host_rsa_key_fingerprint"`
	SubnetId                 string   `json:"subnet_id"`
	ElasticIp                string   `json:"elastic_ip"`
	InfrastructureClass      string   `json:"infrastructure_class"`
	Platform                 string   `json:"platform"`
	PublicIp                 string   `json:"public_ip"`
	PublicDns                string   `json:"public_dns"`
}