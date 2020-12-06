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

// OpsworksStack is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksStack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksStackSpec   `json:"spec"`
	Status OpsworksStackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksStack contains a list of OpsworksStackList
type OpsworksStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksStack `json:"items"`
}

// A OpsworksStackSpec defines the desired state of a OpsworksStack
type OpsworksStackSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksStackParameters `json:",inline"`
}

// A OpsworksStackParameters defines the desired state of a OpsworksStack
type OpsworksStackParameters struct {
	Id                          string                `json:"id"`
	Name                        string                `json:"name"`
	CustomJson                  string                `json:"custom_json"`
	DefaultInstanceProfileArn   string                `json:"default_instance_profile_arn"`
	DefaultOs                   string                `json:"default_os"`
	DefaultSubnetId             string                `json:"default_subnet_id"`
	UseOpsworksSecurityGroups   bool                  `json:"use_opsworks_security_groups"`
	VpcId                       string                `json:"vpc_id"`
	AgentVersion                string                `json:"agent_version"`
	Color                       string                `json:"color"`
	DefaultAvailabilityZone     string                `json:"default_availability_zone"`
	DefaultRootDeviceType       string                `json:"default_root_device_type"`
	UseCustomCookbooks          bool                  `json:"use_custom_cookbooks"`
	ConfigurationManagerName    string                `json:"configuration_manager_name"`
	ManageBerkshelf             bool                  `json:"manage_berkshelf"`
	Region                      string                `json:"region"`
	ServiceRoleArn              string                `json:"service_role_arn"`
	Tags                        map[string]string     `json:"tags"`
	BerkshelfVersion            string                `json:"berkshelf_version"`
	ConfigurationManagerVersion string                `json:"configuration_manager_version"`
	DefaultSshKeyName           string                `json:"default_ssh_key_name"`
	HostnameTheme               string                `json:"hostname_theme"`
	CustomCookbooksSource       CustomCookbooksSource `json:"custom_cookbooks_source"`
}

type CustomCookbooksSource struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Revision string `json:"revision"`
	SshKey   string `json:"ssh_key"`
	Type     string `json:"type"`
}

// A OpsworksStackStatus defines the observed state of a OpsworksStack
type OpsworksStackStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksStackObservation `json:",inline"`
}

// A OpsworksStackObservation records the observed state of a OpsworksStack
type OpsworksStackObservation struct {
	StackEndpoint string `json:"stack_endpoint"`
	Arn           string `json:"arn"`
}