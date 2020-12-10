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

// LambdaProvisionedConcurrencyConfig is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LambdaProvisionedConcurrencyConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LambdaProvisionedConcurrencyConfigSpec   `json:"spec"`
	Status LambdaProvisionedConcurrencyConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaProvisionedConcurrencyConfig contains a list of LambdaProvisionedConcurrencyConfigList
type LambdaProvisionedConcurrencyConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaProvisionedConcurrencyConfig `json:"items"`
}

// A LambdaProvisionedConcurrencyConfigSpec defines the desired state of a LambdaProvisionedConcurrencyConfig
type LambdaProvisionedConcurrencyConfigSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LambdaProvisionedConcurrencyConfigParameters `json:",inline"`
}

// A LambdaProvisionedConcurrencyConfigParameters defines the desired state of a LambdaProvisionedConcurrencyConfig
type LambdaProvisionedConcurrencyConfigParameters struct {
	FunctionName                    string   `json:"function_name"`
	Id                              string   `json:"id"`
	ProvisionedConcurrentExecutions int64    `json:"provisioned_concurrent_executions"`
	Qualifier                       string   `json:"qualifier"`
	Timeouts                        Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Update string `json:"update"`
	Create string `json:"create"`
}

// A LambdaProvisionedConcurrencyConfigStatus defines the observed state of a LambdaProvisionedConcurrencyConfig
type LambdaProvisionedConcurrencyConfigStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LambdaProvisionedConcurrencyConfigObservation `json:",inline"`
}

// A LambdaProvisionedConcurrencyConfigObservation records the observed state of a LambdaProvisionedConcurrencyConfig
type LambdaProvisionedConcurrencyConfigObservation struct{}