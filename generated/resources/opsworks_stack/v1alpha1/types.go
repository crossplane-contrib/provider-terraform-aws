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
	ForProvider                  OpsworksStackParameters `json:"forProvider"`
}

// A OpsworksStackParameters defines the desired state of a OpsworksStack
type OpsworksStackParameters struct {
	ServiceRoleArn              string                `json:"service_role_arn"`
	CustomJson                  string                `json:"custom_json"`
	DefaultOs                   string                `json:"default_os"`
	DefaultSubnetId             string                `json:"default_subnet_id"`
	HostnameTheme               string                `json:"hostname_theme"`
	Name                        string                `json:"name"`
	BerkshelfVersion            string                `json:"berkshelf_version"`
	DefaultInstanceProfileArn   string                `json:"default_instance_profile_arn"`
	VpcId                       string                `json:"vpc_id"`
	UseCustomCookbooks          bool                  `json:"use_custom_cookbooks"`
	AgentVersion                string                `json:"agent_version"`
	Color                       string                `json:"color"`
	ManageBerkshelf             bool                  `json:"manage_berkshelf"`
	Region                      string                `json:"region"`
	Tags                        map[string]string     `json:"tags"`
	Id                          string                `json:"id"`
	UseOpsworksSecurityGroups   bool                  `json:"use_opsworks_security_groups"`
	ConfigurationManagerName    string                `json:"configuration_manager_name"`
	ConfigurationManagerVersion string                `json:"configuration_manager_version"`
	DefaultAvailabilityZone     string                `json:"default_availability_zone"`
	DefaultRootDeviceType       string                `json:"default_root_device_type"`
	DefaultSshKeyName           string                `json:"default_ssh_key_name"`
	CustomCookbooksSource       CustomCookbooksSource `json:"custom_cookbooks_source"`
}

type CustomCookbooksSource struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Revision string `json:"revision"`
	SshKey   string `json:"ssh_key"`
	Type     string `json:"type"`
	Url      string `json:"url"`
}

// A OpsworksStackStatus defines the observed state of a OpsworksStack
type OpsworksStackStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksStackObservation `json:"atProvider"`
}

// A OpsworksStackObservation records the observed state of a OpsworksStack
type OpsworksStackObservation struct {
	StackEndpoint string `json:"stack_endpoint"`
	Arn           string `json:"arn"`
}