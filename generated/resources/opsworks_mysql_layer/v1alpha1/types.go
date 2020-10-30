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

// OpsworksMysqlLayer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksMysqlLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksMysqlLayerSpec   `json:"spec"`
	Status OpsworksMysqlLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksMysqlLayer contains a list of OpsworksMysqlLayerList
type OpsworksMysqlLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksMysqlLayer `json:"items"`
}

// A OpsworksMysqlLayerSpec defines the desired state of a OpsworksMysqlLayer
type OpsworksMysqlLayerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksMysqlLayerParameters `json:",inline"`
}

// A OpsworksMysqlLayerParameters defines the desired state of a OpsworksMysqlLayer
type OpsworksMysqlLayerParameters struct {
	AutoAssignPublicIps        bool     `json:"auto_assign_public_ips"`
	CustomInstanceProfileArn   string   `json:"custom_instance_profile_arn"`
	InstallUpdatesOnBoot       bool     `json:"install_updates_on_boot"`
	InstanceShutdownTimeout    int      `json:"instance_shutdown_timeout"`
	UseEbsOptimizedInstances   bool     `json:"use_ebs_optimized_instances"`
	CustomDeployRecipes        []string `json:"custom_deploy_recipes"`
	CustomJson                 string   `json:"custom_json"`
	CustomSecurityGroupIds     []string `json:"custom_security_group_ids"`
	CustomUndeployRecipes      []string `json:"custom_undeploy_recipes"`
	RootPassword               string   `json:"root_password"`
	StackId                    string   `json:"stack_id"`
	AutoHealing                bool     `json:"auto_healing"`
	ElasticLoadBalancer        string   `json:"elastic_load_balancer"`
	RootPasswordOnAllInstances bool     `json:"root_password_on_all_instances"`
	DrainElbOnShutdown         bool     `json:"drain_elb_on_shutdown"`
	CustomConfigureRecipes     []string `json:"custom_configure_recipes"`
	CustomSetupRecipes         []string `json:"custom_setup_recipes"`
	CustomShutdownRecipes      []string `json:"custom_shutdown_recipes"`
	Name                       string   `json:"name"`
	SystemPackages             []string `json:"system_packages"`
	AutoAssignElasticIps       bool     `json:"auto_assign_elastic_ips"`
}

// A OpsworksMysqlLayerStatus defines the observed state of a OpsworksMysqlLayer
type OpsworksMysqlLayerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksMysqlLayerObservation `json:",inline"`
}

// A OpsworksMysqlLayerObservation records the observed state of a OpsworksMysqlLayer
type OpsworksMysqlLayerObservation struct {
	Arn string `json:"arn"`
	Id  string `json:"id"`
}