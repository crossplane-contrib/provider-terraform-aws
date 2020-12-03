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
	AllocatedMemory          int               `json:"allocated_memory"`
	AutoAssignElasticIps     bool              `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool              `json:"auto_assign_public_ips"`
	InstanceShutdownTimeout  int               `json:"instance_shutdown_timeout"`
	SystemPackages           []string          `json:"system_packages"`
	CustomConfigureRecipes   []string          `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string          `json:"custom_deploy_recipes"`
	CustomInstanceProfileArn string            `json:"custom_instance_profile_arn"`
	CustomSetupRecipes       []string          `json:"custom_setup_recipes"`
	CustomUndeployRecipes    []string          `json:"custom_undeploy_recipes"`
	Id                       string            `json:"id"`
	Name                     string            `json:"name"`
	CustomSecurityGroupIds   []string          `json:"custom_security_group_ids"`
	ElasticLoadBalancer      string            `json:"elastic_load_balancer"`
	Tags                     map[string]string `json:"tags"`
	AutoHealing              bool              `json:"auto_healing"`
	CustomJson               string            `json:"custom_json"`
	CustomShutdownRecipes    []string          `json:"custom_shutdown_recipes"`
	DrainElbOnShutdown       bool              `json:"drain_elb_on_shutdown"`
	InstallUpdatesOnBoot     bool              `json:"install_updates_on_boot"`
	StackId                  string            `json:"stack_id"`
	UseEbsOptimizedInstances bool              `json:"use_ebs_optimized_instances"`
	EbsVolume                []EbsVolume       `json:"ebs_volume"`
}

type EbsVolume struct {
	Size          int    `json:"size"`
	Type          string `json:"type"`
	Encrypted     bool   `json:"encrypted"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
}

// A OpsworksMemcachedLayerStatus defines the observed state of a OpsworksMemcachedLayer
type OpsworksMemcachedLayerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksMemcachedLayerObservation `json:",inline"`
}

// A OpsworksMemcachedLayerObservation records the observed state of a OpsworksMemcachedLayer
type OpsworksMemcachedLayerObservation struct {
	Arn string `json:"arn"`
}