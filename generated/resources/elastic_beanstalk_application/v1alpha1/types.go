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

// ElasticBeanstalkApplication is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElasticBeanstalkApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticBeanstalkApplicationSpec   `json:"spec"`
	Status ElasticBeanstalkApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticBeanstalkApplication contains a list of ElasticBeanstalkApplicationList
type ElasticBeanstalkApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticBeanstalkApplication `json:"items"`
}

// A ElasticBeanstalkApplicationSpec defines the desired state of a ElasticBeanstalkApplication
type ElasticBeanstalkApplicationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElasticBeanstalkApplicationParameters `json:",inline"`
}

// A ElasticBeanstalkApplicationParameters defines the desired state of a ElasticBeanstalkApplication
type ElasticBeanstalkApplicationParameters struct {
	Name                string              `json:"name"`
	Tags                map[string]string   `json:"tags"`
	Description         string              `json:"description"`
	Id                  string              `json:"id"`
	AppversionLifecycle AppversionLifecycle `json:"appversion_lifecycle"`
}

type AppversionLifecycle struct {
	MaxCount           int64  `json:"max_count"`
	ServiceRole        string `json:"service_role"`
	DeleteSourceFromS3 bool   `json:"delete_source_from_s3"`
	MaxAgeInDays       int64  `json:"max_age_in_days"`
}

// A ElasticBeanstalkApplicationStatus defines the observed state of a ElasticBeanstalkApplication
type ElasticBeanstalkApplicationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElasticBeanstalkApplicationObservation `json:",inline"`
}

// A ElasticBeanstalkApplicationObservation records the observed state of a ElasticBeanstalkApplication
type ElasticBeanstalkApplicationObservation struct {
	Arn string `json:"arn"`
}