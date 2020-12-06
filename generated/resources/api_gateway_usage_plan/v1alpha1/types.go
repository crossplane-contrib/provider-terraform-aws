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

// ApiGatewayUsagePlan is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ApiGatewayUsagePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiGatewayUsagePlanSpec   `json:"spec"`
	Status ApiGatewayUsagePlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayUsagePlan contains a list of ApiGatewayUsagePlanList
type ApiGatewayUsagePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayUsagePlan `json:"items"`
}

// A ApiGatewayUsagePlanSpec defines the desired state of a ApiGatewayUsagePlan
type ApiGatewayUsagePlanSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ApiGatewayUsagePlanParameters `json:",inline"`
}

// A ApiGatewayUsagePlanParameters defines the desired state of a ApiGatewayUsagePlan
type ApiGatewayUsagePlanParameters struct {
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	ProductCode      string            `json:"product_code"`
	Tags             map[string]string `json:"tags"`
	Description      string            `json:"description"`
	ApiStages        ApiStages         `json:"api_stages"`
	QuotaSettings    QuotaSettings     `json:"quota_settings"`
	ThrottleSettings ThrottleSettings  `json:"throttle_settings"`
}

type ApiStages struct {
	ApiId string `json:"api_id"`
	Stage string `json:"stage"`
}

type QuotaSettings struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Period string `json:"period"`
}

type ThrottleSettings struct {
	BurstLimit int64 `json:"burst_limit"`
	RateLimit  int64 `json:"rate_limit"`
}

// A ApiGatewayUsagePlanStatus defines the observed state of a ApiGatewayUsagePlan
type ApiGatewayUsagePlanStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ApiGatewayUsagePlanObservation `json:",inline"`
}

// A ApiGatewayUsagePlanObservation records the observed state of a ApiGatewayUsagePlan
type ApiGatewayUsagePlanObservation struct {
	Arn string `json:"arn"`
}