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

// EksFargateProfile is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EksFargateProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EksFargateProfileSpec   `json:"spec"`
	Status EksFargateProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EksFargateProfile contains a list of EksFargateProfileList
type EksFargateProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EksFargateProfile `json:"items"`
}

// A EksFargateProfileSpec defines the desired state of a EksFargateProfile
type EksFargateProfileSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EksFargateProfileParameters `json:"forProvider"`
}

// A EksFargateProfileParameters defines the desired state of a EksFargateProfile
type EksFargateProfileParameters struct {
	ClusterName         string            `json:"cluster_name"`
	FargateProfileName  string            `json:"fargate_profile_name"`
	Id                  string            `json:"id"`
	PodExecutionRoleArn string            `json:"pod_execution_role_arn"`
	SubnetIds           []string          `json:"subnet_ids"`
	Tags                map[string]string `json:"tags"`
	Selector            []Selector        `json:"selector"`
	Timeouts            Timeouts          `json:"timeouts"`
}

type Selector struct {
	Labels    map[string]string `json:"labels"`
	Namespace string            `json:"namespace"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Create string `json:"create"`
}

// A EksFargateProfileStatus defines the observed state of a EksFargateProfile
type EksFargateProfileStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EksFargateProfileObservation `json:"atProvider"`
}

// A EksFargateProfileObservation records the observed state of a EksFargateProfile
type EksFargateProfileObservation struct {
	Arn    string `json:"arn"`
	Status string `json:"status"`
}