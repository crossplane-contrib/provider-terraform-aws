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

// EmrCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EmrCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmrClusterSpec   `json:"spec"`
	Status EmrClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmrCluster contains a list of EmrClusterList
type EmrClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmrCluster `json:"items"`
}

// A EmrClusterSpec defines the desired state of a EmrCluster
type EmrClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EmrClusterParameters `json:",inline"`
}

// A EmrClusterParameters defines the desired state of a EmrCluster
type EmrClusterParameters struct {
	AutoscalingRole       string   `json:"autoscaling_role"`
	Name                  string   `json:"name"`
	ServiceRole           string   `json:"service_role"`
	AdditionalInfo        string   `json:"additional_info"`
	LogUri                string   `json:"log_uri"`
	SecurityConfiguration string   `json:"security_configuration"`
	Applications          []string `json:"applications"`
	ConfigurationsJson    string   `json:"configurations_json"`
	EbsRootVolumeSize     int      `json:"ebs_root_volume_size"`
	Configurations        string   `json:"configurations"`
	CustomAmiId           string   `json:"custom_ami_id"`
	ReleaseLabel          string   `json:"release_label"`
	StepConcurrencyLevel  int      `json:"step_concurrency_level"`
	VisibleToAllUsers     bool     `json:"visible_to_all_users"`
}

// A EmrClusterStatus defines the observed state of a EmrCluster
type EmrClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EmrClusterObservation `json:",inline"`
}

// A EmrClusterObservation records the observed state of a EmrCluster
type EmrClusterObservation struct {
	MasterPublicDns             string `json:"master_public_dns"`
	Id                          string `json:"id"`
	TerminationProtection       bool   `json:"termination_protection"`
	Arn                         string `json:"arn"`
	ClusterState                string `json:"cluster_state"`
	KeepJobFlowAliveWhenNoSteps bool   `json:"keep_job_flow_alive_when_no_steps"`
	ScaleDownBehavior           string `json:"scale_down_behavior"`
}