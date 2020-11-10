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

// OpsworksCustomLayer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksCustomLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksCustomLayerSpec   `json:"spec"`
	Status OpsworksCustomLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksCustomLayer contains a list of OpsworksCustomLayerList
type OpsworksCustomLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksCustomLayer `json:"items"`
}

// A OpsworksCustomLayerSpec defines the desired state of a OpsworksCustomLayer
type OpsworksCustomLayerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksCustomLayerParameters `json:",inline"`
}

// A OpsworksCustomLayerParameters defines the desired state of a OpsworksCustomLayer
type OpsworksCustomLayerParameters struct {
	CustomSecurityGroupIds   []string          `json:"custom_security_group_ids"`
	InstallUpdatesOnBoot     bool              `json:"install_updates_on_boot"`
	Name                     string            `json:"name"`
	CustomDeployRecipes      []string          `json:"custom_deploy_recipes"`
	CustomJson               string            `json:"custom_json"`
	StackId                  string            `json:"stack_id"`
	Tags                     map[string]string `json:"tags"`
	UseEbsOptimizedInstances bool              `json:"use_ebs_optimized_instances"`
	AutoAssignPublicIps      bool              `json:"auto_assign_public_ips"`
	CustomConfigureRecipes   []string          `json:"custom_configure_recipes"`
	CustomInstanceProfileArn string            `json:"custom_instance_profile_arn"`
	ShortName                string            `json:"short_name"`
	SystemPackages           []string          `json:"system_packages"`
	CustomShutdownRecipes    []string          `json:"custom_shutdown_recipes"`
	CustomUndeployRecipes    []string          `json:"custom_undeploy_recipes"`
	InstanceShutdownTimeout  int               `json:"instance_shutdown_timeout"`
	DrainElbOnShutdown       bool              `json:"drain_elb_on_shutdown"`
	ElasticLoadBalancer      string            `json:"elastic_load_balancer"`
	AutoAssignElasticIps     bool              `json:"auto_assign_elastic_ips"`
	AutoHealing              bool              `json:"auto_healing"`
	CustomSetupRecipes       []string          `json:"custom_setup_recipes"`
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

// A OpsworksCustomLayerStatus defines the observed state of a OpsworksCustomLayer
type OpsworksCustomLayerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksCustomLayerObservation `json:",inline"`
}

// A OpsworksCustomLayerObservation records the observed state of a OpsworksCustomLayer
type OpsworksCustomLayerObservation struct {
	Arn string `json:"arn"`
	Id  string `json:"id"`
}