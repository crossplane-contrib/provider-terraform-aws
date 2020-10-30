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

// CodedeployDeploymentGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodedeployDeploymentGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodedeployDeploymentGroupSpec   `json:"spec"`
	Status CodedeployDeploymentGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodedeployDeploymentGroup contains a list of CodedeployDeploymentGroupList
type CodedeployDeploymentGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodedeployDeploymentGroup `json:"items"`
}

// A CodedeployDeploymentGroupSpec defines the desired state of a CodedeployDeploymentGroup
type CodedeployDeploymentGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodedeployDeploymentGroupParameters `json:",inline"`
}

// A CodedeployDeploymentGroupParameters defines the desired state of a CodedeployDeploymentGroup
type CodedeployDeploymentGroupParameters struct {
	ServiceRoleArn       string   `json:"service_role_arn"`
	AppName              string   `json:"app_name"`
	AutoscalingGroups    []string `json:"autoscaling_groups"`
	DeploymentConfigName string   `json:"deployment_config_name"`
	DeploymentGroupName  string   `json:"deployment_group_name"`
}

// A CodedeployDeploymentGroupStatus defines the observed state of a CodedeployDeploymentGroup
type CodedeployDeploymentGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodedeployDeploymentGroupObservation `json:",inline"`
}

// A CodedeployDeploymentGroupObservation records the observed state of a CodedeployDeploymentGroup
type CodedeployDeploymentGroupObservation struct {
	Id string `json:"id"`
}