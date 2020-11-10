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

// ElasticBeanstalkEnvironment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElasticBeanstalkEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticBeanstalkEnvironmentSpec   `json:"spec"`
	Status ElasticBeanstalkEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticBeanstalkEnvironment contains a list of ElasticBeanstalkEnvironmentList
type ElasticBeanstalkEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticBeanstalkEnvironment `json:"items"`
}

// A ElasticBeanstalkEnvironmentSpec defines the desired state of a ElasticBeanstalkEnvironment
type ElasticBeanstalkEnvironmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElasticBeanstalkEnvironmentParameters `json:",inline"`
}

// A ElasticBeanstalkEnvironmentParameters defines the desired state of a ElasticBeanstalkEnvironment
type ElasticBeanstalkEnvironmentParameters struct {
	Application         string            `json:"application"`
	Name                string            `json:"name"`
	TemplateName        string            `json:"template_name"`
	WaitForReadyTimeout string            `json:"wait_for_ready_timeout"`
	Tags                map[string]string `json:"tags"`
	PollInterval        string            `json:"poll_interval"`
	Description         string            `json:"description"`
	Tier                string            `json:"tier"`
	Setting             []Setting         `json:"setting"`
}

type Setting struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Resource  string `json:"resource"`
	Value     string `json:"value"`
}

// A ElasticBeanstalkEnvironmentStatus defines the observed state of a ElasticBeanstalkEnvironment
type ElasticBeanstalkEnvironmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElasticBeanstalkEnvironmentObservation `json:",inline"`
}

// A ElasticBeanstalkEnvironmentObservation records the observed state of a ElasticBeanstalkEnvironment
type ElasticBeanstalkEnvironmentObservation struct {
	AutoscalingGroups    []string `json:"autoscaling_groups"`
	Id                   string   `json:"id"`
	LoadBalancers        []string `json:"load_balancers"`
	Instances            []string `json:"instances"`
	PlatformArn          string   `json:"platform_arn"`
	Queues               []string `json:"queues"`
	Arn                  string   `json:"arn"`
	Cname                string   `json:"cname"`
	EndpointUrl          string   `json:"endpoint_url"`
	SolutionStackName    string   `json:"solution_stack_name"`
	Triggers             []string `json:"triggers"`
	VersionLabel         string   `json:"version_label"`
	CnamePrefix          string   `json:"cname_prefix"`
	LaunchConfigurations []string `json:"launch_configurations"`
}