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

// ApiGatewayMethodSettings is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayMethodSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayMethodSettingsSpec   `json:"spec"`
	Status ApiGatewayMethodSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayMethodSettings contains a list of ApiGatewayMethodSettingsList
type ApiGatewayMethodSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayMethodSettings `json:"items"`
}

// A ApiGatewayMethodSettingsSpec defines the desired state of a ApiGatewayMethodSettings
type ApiGatewayMethodSettingsSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayMethodSettingsParameters `json:",inline"`
}

// A ApiGatewayMethodSettingsParameters defines the desired state of a ApiGatewayMethodSettings
type ApiGatewayMethodSettingsParameters struct {
	MethodPath string `json:"method_path"`
	RestApiId  string `json:"rest_api_id"`
	StageName  string `json:"stage_name"`
}

// A ApiGatewayMethodSettingsStatus defines the observed state of a ApiGatewayMethodSettings
type ApiGatewayMethodSettingsStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayMethodSettingsObservation `json:",inline"`
}

// A ApiGatewayMethodSettingsObservation records the observed state of a ApiGatewayMethodSettings
type ApiGatewayMethodSettingsObservation struct {
	Id string `json:"id"`
}