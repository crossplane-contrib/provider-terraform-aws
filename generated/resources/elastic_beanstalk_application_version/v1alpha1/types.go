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

// ElasticBeanstalkApplicationVersion is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElasticBeanstalkApplicationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticBeanstalkApplicationVersionSpec   `json:"spec"`
	Status ElasticBeanstalkApplicationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticBeanstalkApplicationVersion contains a list of ElasticBeanstalkApplicationVersionList
type ElasticBeanstalkApplicationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticBeanstalkApplicationVersion `json:"items"`
}

// A ElasticBeanstalkApplicationVersionSpec defines the desired state of a ElasticBeanstalkApplicationVersion
type ElasticBeanstalkApplicationVersionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElasticBeanstalkApplicationVersionParameters `json:",inline"`
}

// A ElasticBeanstalkApplicationVersionParameters defines the desired state of a ElasticBeanstalkApplicationVersion
type ElasticBeanstalkApplicationVersionParameters struct {
	Bucket      string            `json:"bucket"`
	Id          string            `json:"id"`
	Key         string            `json:"key"`
	Application string            `json:"application"`
	Description string            `json:"description"`
	ForceDelete bool              `json:"force_delete"`
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
}

// A ElasticBeanstalkApplicationVersionStatus defines the observed state of a ElasticBeanstalkApplicationVersion
type ElasticBeanstalkApplicationVersionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElasticBeanstalkApplicationVersionObservation `json:",inline"`
}

// A ElasticBeanstalkApplicationVersionObservation records the observed state of a ElasticBeanstalkApplicationVersion
type ElasticBeanstalkApplicationVersionObservation struct {
	Arn string `json:"arn"`
}