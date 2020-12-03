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

// OpsworksJavaAppLayer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksJavaAppLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksJavaAppLayerSpec   `json:"spec"`
	Status OpsworksJavaAppLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksJavaAppLayer contains a list of OpsworksJavaAppLayerList
type OpsworksJavaAppLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksJavaAppLayer `json:"items"`
}

// A OpsworksJavaAppLayerSpec defines the desired state of a OpsworksJavaAppLayer
type OpsworksJavaAppLayerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksJavaAppLayerParameters `json:",inline"`
}

// A OpsworksJavaAppLayerParameters defines the desired state of a OpsworksJavaAppLayer
type OpsworksJavaAppLayerParameters struct {
	CustomJson               string            `json:"custom_json"`
	CustomInstanceProfileArn string            `json:"custom_instance_profile_arn"`
	CustomShutdownRecipes    []string          `json:"custom_shutdown_recipes"`
	Tags                     map[string]string `json:"tags"`
	AppServerVersion         string            `json:"app_server_version"`
	DrainElbOnShutdown       bool              `json:"drain_elb_on_shutdown"`
	JvmType                  string            `json:"jvm_type"`
	StackId                  string            `json:"stack_id"`
	Name                     string            `json:"name"`
	JvmVersion               string            `json:"jvm_version"`
	UseEbsOptimizedInstances bool              `json:"use_ebs_optimized_instances"`
	InstanceShutdownTimeout  int               `json:"instance_shutdown_timeout"`
	CustomSecurityGroupIds   []string          `json:"custom_security_group_ids"`
	JvmOptions               string            `json:"jvm_options"`
	AutoAssignElasticIps     bool              `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps      bool              `json:"auto_assign_public_ips"`
	ElasticLoadBalancer      string            `json:"elastic_load_balancer"`
	SystemPackages           []string          `json:"system_packages"`
	CustomConfigureRecipes   []string          `json:"custom_configure_recipes"`
	CustomDeployRecipes      []string          `json:"custom_deploy_recipes"`
	CustomUndeployRecipes    []string          `json:"custom_undeploy_recipes"`
	Id                       string            `json:"id"`
	InstallUpdatesOnBoot     bool              `json:"install_updates_on_boot"`
	AppServer                string            `json:"app_server"`
	AutoHealing              bool              `json:"auto_healing"`
	CustomSetupRecipes       []string          `json:"custom_setup_recipes"`
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

// A OpsworksJavaAppLayerStatus defines the observed state of a OpsworksJavaAppLayer
type OpsworksJavaAppLayerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksJavaAppLayerObservation `json:",inline"`
}

// A OpsworksJavaAppLayerObservation records the observed state of a OpsworksJavaAppLayer
type OpsworksJavaAppLayerObservation struct {
	Arn string `json:"arn"`
}