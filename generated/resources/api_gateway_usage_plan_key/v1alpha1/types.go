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

// ApiGatewayUsagePlanKey is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayUsagePlanKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayUsagePlanKeySpec   `json:"spec"`
	Status ApiGatewayUsagePlanKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayUsagePlanKey contains a list of ApiGatewayUsagePlanKeyList
type ApiGatewayUsagePlanKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayUsagePlanKey `json:"items"`
}

// A ApiGatewayUsagePlanKeySpec defines the desired state of a ApiGatewayUsagePlanKey
type ApiGatewayUsagePlanKeySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayUsagePlanKeyParameters `json:",inline"`
}

// A ApiGatewayUsagePlanKeyParameters defines the desired state of a ApiGatewayUsagePlanKey
type ApiGatewayUsagePlanKeyParameters struct {
	Id          string `json:"id"`
	KeyId       string `json:"key_id"`
	KeyType     string `json:"key_type"`
	UsagePlanId string `json:"usage_plan_id"`
}

// A ApiGatewayUsagePlanKeyStatus defines the observed state of a ApiGatewayUsagePlanKey
type ApiGatewayUsagePlanKeyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayUsagePlanKeyObservation `json:",inline"`
}

// A ApiGatewayUsagePlanKeyObservation records the observed state of a ApiGatewayUsagePlanKey
type ApiGatewayUsagePlanKeyObservation struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}