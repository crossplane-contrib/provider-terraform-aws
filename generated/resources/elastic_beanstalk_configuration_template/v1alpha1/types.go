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

// ElasticBeanstalkConfigurationTemplate is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElasticBeanstalkConfigurationTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticBeanstalkConfigurationTemplateSpec   `json:"spec"`
	Status ElasticBeanstalkConfigurationTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticBeanstalkConfigurationTemplate contains a list of ElasticBeanstalkConfigurationTemplateList
type ElasticBeanstalkConfigurationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticBeanstalkConfigurationTemplate `json:"items"`
}

// A ElasticBeanstalkConfigurationTemplateSpec defines the desired state of a ElasticBeanstalkConfigurationTemplate
type ElasticBeanstalkConfigurationTemplateSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElasticBeanstalkConfigurationTemplateParameters `json:",inline"`
}

// A ElasticBeanstalkConfigurationTemplateParameters defines the desired state of a ElasticBeanstalkConfigurationTemplate
type ElasticBeanstalkConfigurationTemplateParameters struct {
	Application       string  `json:"application"`
	Description       string  `json:"description"`
	EnvironmentId     string  `json:"environment_id"`
	Id                string  `json:"id"`
	Name              string  `json:"name"`
	SolutionStackName string  `json:"solution_stack_name"`
	Setting           Setting `json:"setting"`
}

type Setting struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Resource  string `json:"resource"`
	Value     string `json:"value"`
}

// A ElasticBeanstalkConfigurationTemplateStatus defines the observed state of a ElasticBeanstalkConfigurationTemplate
type ElasticBeanstalkConfigurationTemplateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElasticBeanstalkConfigurationTemplateObservation `json:",inline"`
}

// A ElasticBeanstalkConfigurationTemplateObservation records the observed state of a ElasticBeanstalkConfigurationTemplate
type ElasticBeanstalkConfigurationTemplateObservation struct{}