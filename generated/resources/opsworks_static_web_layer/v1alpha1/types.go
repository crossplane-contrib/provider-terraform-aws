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

// OpsworksStaticWebLayer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksStaticWebLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksStaticWebLayerSpec   `json:"spec"`
	Status OpsworksStaticWebLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksStaticWebLayer contains a list of OpsworksStaticWebLayerList
type OpsworksStaticWebLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksStaticWebLayer `json:"items"`
}

// A OpsworksStaticWebLayerSpec defines the desired state of a OpsworksStaticWebLayer
type OpsworksStaticWebLayerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksStaticWebLayerParameters `json:",inline"`
}

// A OpsworksStaticWebLayerParameters defines the desired state of a OpsworksStaticWebLayer
type OpsworksStaticWebLayerParameters struct {
	AutoHealing              bool              `json:"auto_healing"`
	CustomInstanceProfileArn string            `json:"custom_instance_profile_arn"`
	ElasticLoadBalancer      string            `json:"elastic_load_balancer"`
	Id                       string            `json:"id"`
	Tags                     map[string]string `json:"tags"`
	AutoAssignElasticIps     bool              `json:"auto_assign_elastic_ips"`
	CustomJson               string            `json:"custom_json"`
	CustomSecurityGroupIds   []string          `json:"custom_security_group_ids"`
	DrainElbOnShutdown       bool              `json:"drain_elb_on_shutdown"`
	AutoAssignPublicIps      bool              `json:"auto_assign_public_ips"`
	CustomConfigureRecipes   []string          `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string          `json:"custom_deploy_recipes"`
	Name                     string            `json:"name"`
	SystemPackages           []string          `json:"system_packages"`
	UseEbsOptimizedInstances bool              `json:"use_ebs_optimized_instances"`
	CustomSetupRecipes       []string          `json:"custom_setup_recipes"`
	CustomShutdownRecipes    []string          `json:"custom_shutdown_recipes"`
	CustomUndeployRecipes    []string          `json:"custom_undeploy_recipes"`
	InstallUpdatesOnBoot     bool              `json:"install_updates_on_boot"`
	InstanceShutdownTimeout  int               `json:"instance_shutdown_timeout"`
	StackId                  string            `json:"stack_id"`
	EbsVolume                []EbsVolume       `json:"ebs_volume"`
}

type EbsVolume struct {
	Type          string `json:"type"`
	Encrypted     bool   `json:"encrypted"`
	Iops          int    `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int    `json:"size"`
}

// A OpsworksStaticWebLayerStatus defines the observed state of a OpsworksStaticWebLayer
type OpsworksStaticWebLayerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksStaticWebLayerObservation `json:",inline"`
}

// A OpsworksStaticWebLayerObservation records the observed state of a OpsworksStaticWebLayer
type OpsworksStaticWebLayerObservation struct {
	Arn string `json:"arn"`
}