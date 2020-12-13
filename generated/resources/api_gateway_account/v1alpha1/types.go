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

// ApiGatewayAccount is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayAccountSpec   `json:"spec"`
	Status ApiGatewayAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayAccount contains a list of ApiGatewayAccountList
type ApiGatewayAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayAccount `json:"items"`
}

// A ApiGatewayAccountSpec defines the desired state of a ApiGatewayAccount
type ApiGatewayAccountSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayAccountParameters `json:"forProvider"`
}

// A ApiGatewayAccountParameters defines the desired state of a ApiGatewayAccount
type ApiGatewayAccountParameters struct {
	CloudwatchRoleArn string `json:"cloudwatch_role_arn"`
	Id                string `json:"id"`
}

// A ApiGatewayAccountStatus defines the observed state of a ApiGatewayAccount
type ApiGatewayAccountStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayAccountObservation `json:"atProvider"`
}

// A ApiGatewayAccountObservation records the observed state of a ApiGatewayAccount
type ApiGatewayAccountObservation struct {
	ThrottleSettings []ThrottleSettings `json:"throttle_settings"`
}

type ThrottleSettings struct {
	BurstLimit int64 `json:"burst_limit"`
	RateLimit  int64 `json:"rate_limit"`
}