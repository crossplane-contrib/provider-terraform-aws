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

// CloudformationStackSetInstance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudformationStackSetInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudformationStackSetInstanceSpec   `json:"spec"`
	Status CloudformationStackSetInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudformationStackSetInstance contains a list of CloudformationStackSetInstanceList
type CloudformationStackSetInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudformationStackSetInstance `json:"items"`
}

// A CloudformationStackSetInstanceSpec defines the desired state of a CloudformationStackSetInstance
type CloudformationStackSetInstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudformationStackSetInstanceParameters `json:"forProvider"`
}

// A CloudformationStackSetInstanceParameters defines the desired state of a CloudformationStackSetInstance
type CloudformationStackSetInstanceParameters struct {
	AccountId          string            `json:"account_id"`
	ParameterOverrides map[string]string `json:"parameter_overrides"`
	Region             string            `json:"region"`
	RetainStack        bool              `json:"retain_stack"`
	StackSetName       string            `json:"stack_set_name"`
	Timeouts           Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A CloudformationStackSetInstanceStatus defines the observed state of a CloudformationStackSetInstance
type CloudformationStackSetInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudformationStackSetInstanceObservation `json:"atProvider"`
}

// A CloudformationStackSetInstanceObservation records the observed state of a CloudformationStackSetInstance
type CloudformationStackSetInstanceObservation struct {
	StackId string `json:"stack_id"`
}