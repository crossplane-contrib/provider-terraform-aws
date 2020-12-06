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

// Cloud9EnvironmentEc2 is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Cloud9EnvironmentEc2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Cloud9EnvironmentEc2Spec   `json:"spec"`
	Status Cloud9EnvironmentEc2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Cloud9EnvironmentEc2 contains a list of Cloud9EnvironmentEc2List
type Cloud9EnvironmentEc2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cloud9EnvironmentEc2 `json:"items"`
}

// A Cloud9EnvironmentEc2Spec defines the desired state of a Cloud9EnvironmentEc2
type Cloud9EnvironmentEc2Spec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Cloud9EnvironmentEc2Parameters `json:",inline"`
}

// A Cloud9EnvironmentEc2Parameters defines the desired state of a Cloud9EnvironmentEc2
type Cloud9EnvironmentEc2Parameters struct {
	Description              string            `json:"description"`
	Id                       string            `json:"id"`
	InstanceType             string            `json:"instance_type"`
	OwnerArn                 string            `json:"owner_arn"`
	SubnetId                 string            `json:"subnet_id"`
	AutomaticStopTimeMinutes int64             `json:"automatic_stop_time_minutes"`
	Name                     string            `json:"name"`
	Tags                     map[string]string `json:"tags"`
}

// A Cloud9EnvironmentEc2Status defines the observed state of a Cloud9EnvironmentEc2
type Cloud9EnvironmentEc2Status struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Cloud9EnvironmentEc2Observation `json:",inline"`
}

// A Cloud9EnvironmentEc2Observation records the observed state of a Cloud9EnvironmentEc2
type Cloud9EnvironmentEc2Observation struct {
	Arn  string `json:"arn"`
	Type string `json:"type"`
}