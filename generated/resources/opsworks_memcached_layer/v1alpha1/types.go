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

// OpsworksMemcachedLayer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksMemcachedLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksMemcachedLayerSpec   `json:"spec"`
	Status OpsworksMemcachedLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksMemcachedLayer contains a list of OpsworksMemcachedLayerList
type OpsworksMemcachedLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksMemcachedLayer `json:"items"`
}

// A OpsworksMemcachedLayerSpec defines the desired state of a OpsworksMemcachedLayer
type OpsworksMemcachedLayerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksMemcachedLayerParameters `json:",inline"`
}

// A OpsworksMemcachedLayerParameters defines the desired state of a OpsworksMemcachedLayer
type OpsworksMemcachedLayerParameters struct {
	CustomJson               string            `json:"custom_json"`
	CustomUndeployRecipes    []string          `json:"custom_undeploy_recipes"`
	SystemPackages           []string          `json:"system_packages"`
	Tags                     map[string]string `json:"tags"`
	AllocatedMemory          int               `json:"allocated_memory"`
	AutoAssignElasticIps     bool              `json:"auto_assign_elastic_ips"`
	CustomDeployRecipes      []string          `json:"custom_deploy_recipes"`
	ElasticLoadBalancer      string            `json:"elastic_load_balancer"`
	StackId                  string            `json:"stack_id"`
	UseEbsOptimizedInstances bool              `json:"use_ebs_optimized_instances"`
	CustomInstanceProfileArn string            `json:"custom_instance_profile_arn"`
	CustomSetupRecipes       []string          `json:"custom_setup_recipes"`
	DrainElbOnShutdown       bool              `json:"drain_elb_on_shutdown"`
	CustomSecurityGroupIds   []string          `json:"custom_security_group_ids"`
	InstanceShutdownTimeout  int               `json:"instance_shutdown_timeout"`
	Name                     string            `json:"name"`
	AutoAssignPublicIps      bool              `json:"auto_assign_public_ips"`
	AutoHealing              bool              `json:"auto_healing"`
	CustomConfigureRecipes   []string          `json:"custom_configure_recipes"`
	CustomShutdownRecipes    []string          `json:"custom_shutdown_recipes"`
	InstallUpdatesOnBoot     bool              `json:"install_updates_on_boot"`
	EbsVolume                []EbsVolume       `json:"ebs_volume"`
}

type EbsVolume struct {
	Encrypted     bool   `json:"encrypted"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
	Type          string `json:"type"`
}

// A OpsworksMemcachedLayerStatus defines the observed state of a OpsworksMemcachedLayer
type OpsworksMemcachedLayerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksMemcachedLayerObservation `json:",inline"`
}

// A OpsworksMemcachedLayerObservation records the observed state of a OpsworksMemcachedLayer
type OpsworksMemcachedLayerObservation struct {
	Id  string `json:"id"`
	Arn string `json:"arn"`
}