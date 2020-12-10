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

// OpsworksNodejsAppLayer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksNodejsAppLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksNodejsAppLayerSpec   `json:"spec"`
	Status OpsworksNodejsAppLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksNodejsAppLayer contains a list of OpsworksNodejsAppLayerList
type OpsworksNodejsAppLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksNodejsAppLayer `json:"items"`
}

// A OpsworksNodejsAppLayerSpec defines the desired state of a OpsworksNodejsAppLayer
type OpsworksNodejsAppLayerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksNodejsAppLayerParameters `json:",inline"`
}

// A OpsworksNodejsAppLayerParameters defines the desired state of a OpsworksNodejsAppLayer
type OpsworksNodejsAppLayerParameters struct {
	SystemPackages           []string          `json:"system_packages"`
	AutoAssignElasticIps     bool              `json:"auto_assign_elastic_ips"`
	AutoHealing              bool              `json:"auto_healing"`
	CustomInstanceProfileArn string            `json:"custom_instance_profile_arn"`
	CustomSetupRecipes       []string          `json:"custom_setup_recipes"`
	CustomShutdownRecipes    []string          `json:"custom_shutdown_recipes"`
	Id                       string            `json:"id"`
	InstallUpdatesOnBoot     bool              `json:"install_updates_on_boot"`
	UseEbsOptimizedInstances bool              `json:"use_ebs_optimized_instances"`
	AutoAssignPublicIps      bool              `json:"auto_assign_public_ips"`
	CustomDeployRecipes      []string          `json:"custom_deploy_recipes"`
	CustomJson               string            `json:"custom_json"`
	DrainElbOnShutdown       bool              `json:"drain_elb_on_shutdown"`
	ElasticLoadBalancer      string            `json:"elastic_load_balancer"`
	InstanceShutdownTimeout  int64             `json:"instance_shutdown_timeout"`
	StackId                  string            `json:"stack_id"`
	NodejsVersion            string            `json:"nodejs_version"`
	CustomConfigureRecipes   []string          `json:"custom_configure_recipes"`
	CustomSecurityGroupIds   []string          `json:"custom_security_group_ids"`
	CustomUndeployRecipes    []string          `json:"custom_undeploy_recipes"`
	Name                     string            `json:"name"`
	Tags                     map[string]string `json:"tags"`
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

// A OpsworksNodejsAppLayerStatus defines the observed state of a OpsworksNodejsAppLayer
type OpsworksNodejsAppLayerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksNodejsAppLayerObservation `json:",inline"`
}

// A OpsworksNodejsAppLayerObservation records the observed state of a OpsworksNodejsAppLayer
type OpsworksNodejsAppLayerObservation struct {
	Arn string `json:"arn"`
}