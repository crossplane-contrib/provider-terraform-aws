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

// WorkspacesDirectory is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WorkspacesDirectory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkspacesDirectorySpec   `json:"spec"`
	Status WorkspacesDirectoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspacesDirectory contains a list of WorkspacesDirectoryList
type WorkspacesDirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspacesDirectory `json:"items"`
}

// A WorkspacesDirectorySpec defines the desired state of a WorkspacesDirectory
type WorkspacesDirectorySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WorkspacesDirectoryParameters `json:",inline"`
}

// A WorkspacesDirectoryParameters defines the desired state of a WorkspacesDirectory
type WorkspacesDirectoryParameters struct {
	DirectoryId            string                 `json:"directory_id"`
	Tags                   map[string]string      `json:"tags"`
	SelfServicePermissions SelfServicePermissions `json:"self_service_permissions"`
}

type SelfServicePermissions struct {
	RestartWorkspace   bool `json:"restart_workspace"`
	SwitchRunningMode  bool `json:"switch_running_mode"`
	ChangeComputeType  bool `json:"change_compute_type"`
	IncreaseVolumeSize bool `json:"increase_volume_size"`
	RebuildWorkspace   bool `json:"rebuild_workspace"`
}

// A WorkspacesDirectoryStatus defines the observed state of a WorkspacesDirectory
type WorkspacesDirectoryStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WorkspacesDirectoryObservation `json:",inline"`
}

// A WorkspacesDirectoryObservation records the observed state of a WorkspacesDirectory
type WorkspacesDirectoryObservation struct {
	DirectoryName            string   `json:"directory_name"`
	DirectoryType            string   `json:"directory_type"`
	DnsIpAddresses           []string `json:"dns_ip_addresses"`
	Id                       string   `json:"id"`
	SubnetIds                []string `json:"subnet_ids"`
	WorkspaceSecurityGroupId string   `json:"workspace_security_group_id"`
	Alias                    string   `json:"alias"`
	IamRoleId                string   `json:"iam_role_id"`
	IpGroupIds               []string `json:"ip_group_ids"`
	RegistrationCode         string   `json:"registration_code"`
	CustomerUserName         string   `json:"customer_user_name"`
}