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

// OpsworksHaproxyLayer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksHaproxyLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksHaproxyLayerSpec   `json:"spec"`
	Status OpsworksHaproxyLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksHaproxyLayer contains a list of OpsworksHaproxyLayerList
type OpsworksHaproxyLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksHaproxyLayer `json:"items"`
}

// A OpsworksHaproxyLayerSpec defines the desired state of a OpsworksHaproxyLayer
type OpsworksHaproxyLayerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksHaproxyLayerParameters `json:"forProvider"`
}

// A OpsworksHaproxyLayerParameters defines the desired state of a OpsworksHaproxyLayer
type OpsworksHaproxyLayerParameters struct {
	AutoAssignElasticIps     bool              `json:"auto_assign_elastic_ips"`
	AutoHealing              bool              `json:"auto_healing"`
	StatsEnabled             bool              `json:"stats_enabled"`
	DrainElbOnShutdown       bool              `json:"drain_elb_on_shutdown"`
	InstanceShutdownTimeout  int64             `json:"instance_shutdown_timeout"`
	CustomInstanceProfileArn string            `json:"custom_instance_profile_arn"`
	HealthcheckMethod        string            `json:"healthcheck_method"`
	StatsUser                string            `json:"stats_user"`
	CustomSetupRecipes       []string          `json:"custom_setup_recipes"`
	CustomShutdownRecipes    []string          `json:"custom_shutdown_recipes"`
	HealthcheckUrl           string            `json:"healthcheck_url"`
	Tags                     map[string]string `json:"tags"`
	UseEbsOptimizedInstances bool              `json:"use_ebs_optimized_instances"`
	CustomConfigureRecipes   []string          `json:"custom_configure_recipes"`
	Id                       string            `json:"id"`
	StackId                  string            `json:"stack_id"`
	AutoAssignPublicIps      bool              `json:"auto_assign_public_ips"`
	CustomUndeployRecipes    []string          `json:"custom_undeploy_recipes"`
	CustomDeployRecipes      []string          `json:"custom_deploy_recipes"`
	CustomJson               string            `json:"custom_json"`
	CustomSecurityGroupIds   []string          `json:"custom_security_group_ids"`
	ElasticLoadBalancer      string            `json:"elastic_load_balancer"`
	InstallUpdatesOnBoot     bool              `json:"install_updates_on_boot"`
	Name                     string            `json:"name"`
	StatsPassword            string            `json:"stats_password"`
	StatsUrl                 string            `json:"stats_url"`
	SystemPackages           []string          `json:"system_packages"`
	EbsVolume                EbsVolume         `json:"ebs_volume"`
}

type EbsVolume struct {
	Encrypted     bool   `json:"encrypted"`
	Iops          int64  `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int64  `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
	Size          int64  `json:"size"`
	Type          string `json:"type"`
}

// A OpsworksHaproxyLayerStatus defines the observed state of a OpsworksHaproxyLayer
type OpsworksHaproxyLayerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksHaproxyLayerObservation `json:"atProvider"`
}

// A OpsworksHaproxyLayerObservation records the observed state of a OpsworksHaproxyLayer
type OpsworksHaproxyLayerObservation struct {
	Arn string `json:"arn"`
}