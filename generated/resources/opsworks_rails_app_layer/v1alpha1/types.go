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

// OpsworksRailsAppLayer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksRailsAppLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksRailsAppLayerSpec   `json:"spec"`
	Status OpsworksRailsAppLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksRailsAppLayer contains a list of OpsworksRailsAppLayerList
type OpsworksRailsAppLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksRailsAppLayer `json:"items"`
}

// A OpsworksRailsAppLayerSpec defines the desired state of a OpsworksRailsAppLayer
type OpsworksRailsAppLayerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksRailsAppLayerParameters `json:",inline"`
}

// A OpsworksRailsAppLayerParameters defines the desired state of a OpsworksRailsAppLayer
type OpsworksRailsAppLayerParameters struct {
	CustomSecurityGroupIds   []string          `json:"custom_security_group_ids"`
	Id                       string            `json:"id"`
	InstanceShutdownTimeout  int               `json:"instance_shutdown_timeout"`
	RubyVersion              string            `json:"ruby_version"`
	CustomShutdownRecipes    []string          `json:"custom_shutdown_recipes"`
	RubygemsVersion          string            `json:"rubygems_version"`
	Tags                     map[string]string `json:"tags"`
	AutoAssignElasticIps     bool              `json:"auto_assign_elastic_ips"`
	AutoHealing              bool              `json:"auto_healing"`
	BundlerVersion           string            `json:"bundler_version"`
	CustomInstanceProfileArn string            `json:"custom_instance_profile_arn"`
	CustomConfigureRecipes   []string          `json:"custom_configure_recipes"`
	Name                     string            `json:"name"`
	UseEbsOptimizedInstances bool              `json:"use_ebs_optimized_instances"`
	ElasticLoadBalancer      string            `json:"elastic_load_balancer"`
	InstallUpdatesOnBoot     bool              `json:"install_updates_on_boot"`
	StackId                  string            `json:"stack_id"`
	AppServer                string            `json:"app_server"`
	AutoAssignPublicIps      bool              `json:"auto_assign_public_ips"`
	CustomSetupRecipes       []string          `json:"custom_setup_recipes"`
	DrainElbOnShutdown       bool              `json:"drain_elb_on_shutdown"`
	ManageBundler            bool              `json:"manage_bundler"`
	PassengerVersion         string            `json:"passenger_version"`
	SystemPackages           []string          `json:"system_packages"`
	CustomDeployRecipes      []string          `json:"custom_deploy_recipes"`
	CustomJson               string            `json:"custom_json"`
	CustomUndeployRecipes    []string          `json:"custom_undeploy_recipes"`
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

// A OpsworksRailsAppLayerStatus defines the observed state of a OpsworksRailsAppLayer
type OpsworksRailsAppLayerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksRailsAppLayerObservation `json:",inline"`
}

// A OpsworksRailsAppLayerObservation records the observed state of a OpsworksRailsAppLayer
type OpsworksRailsAppLayerObservation struct {
	Arn string `json:"arn"`
}