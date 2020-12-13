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

// OpsworksGangliaLayer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksGangliaLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksGangliaLayerSpec   `json:"spec"`
	Status OpsworksGangliaLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksGangliaLayer contains a list of OpsworksGangliaLayerList
type OpsworksGangliaLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksGangliaLayer `json:"items"`
}

// A OpsworksGangliaLayerSpec defines the desired state of a OpsworksGangliaLayer
type OpsworksGangliaLayerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksGangliaLayerParameters `json:"forProvider"`
}

// A OpsworksGangliaLayerParameters defines the desired state of a OpsworksGangliaLayer
type OpsworksGangliaLayerParameters struct {
	Url                      string            `json:"url"`
	AutoHealing              bool              `json:"auto_healing"`
	InstanceShutdownTimeout  int64             `json:"instance_shutdown_timeout"`
	Name                     string            `json:"name"`
	Password                 string            `json:"password"`
	StackId                  string            `json:"stack_id"`
	InstallUpdatesOnBoot     bool              `json:"install_updates_on_boot"`
	AutoAssignPublicIps      bool              `json:"auto_assign_public_ips"`
	CustomSecurityGroupIds   []string          `json:"custom_security_group_ids"`
	CustomSetupRecipes       []string          `json:"custom_setup_recipes"`
	CustomUndeployRecipes    []string          `json:"custom_undeploy_recipes"`
	Id                       string            `json:"id"`
	UseEbsOptimizedInstances bool              `json:"use_ebs_optimized_instances"`
	Username                 string            `json:"username"`
	CustomInstanceProfileArn string            `json:"custom_instance_profile_arn"`
	CustomShutdownRecipes    []string          `json:"custom_shutdown_recipes"`
	DrainElbOnShutdown       bool              `json:"drain_elb_on_shutdown"`
	ElasticLoadBalancer      string            `json:"elastic_load_balancer"`
	Tags                     map[string]string `json:"tags"`
	SystemPackages           []string          `json:"system_packages"`
	AutoAssignElasticIps     bool              `json:"auto_assign_elastic_ips"`
	CustomConfigureRecipes   []string          `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string          `json:"custom_deploy_recipes"`
	CustomJson               string            `json:"custom_json"`
	EbsVolume                EbsVolume         `json:"ebs_volume"`
}

type EbsVolume struct {
	Size          int64  `json:"size"`
	Type          string `json:"type"`
	Encrypted     bool   `json:"encrypted"`
	Iops          int64  `json:"iops"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int64  `json:"number_of_disks"`
	RaidLevel     string `json:"raid_level"`
}

// A OpsworksGangliaLayerStatus defines the observed state of a OpsworksGangliaLayer
type OpsworksGangliaLayerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksGangliaLayerObservation `json:"atProvider"`
}

// A OpsworksGangliaLayerObservation records the observed state of a OpsworksGangliaLayer
type OpsworksGangliaLayerObservation struct {
	Arn string `json:"arn"`
}