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

// AlbTargetGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AlbTargetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlbTargetGroupSpec   `json:"spec"`
	Status AlbTargetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlbTargetGroup contains a list of AlbTargetGroupList
type AlbTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlbTargetGroup `json:"items"`
}

// A AlbTargetGroupSpec defines the desired state of a AlbTargetGroup
type AlbTargetGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AlbTargetGroupParameters `json:",inline"`
}

// A AlbTargetGroupParameters defines the desired state of a AlbTargetGroup
type AlbTargetGroupParameters struct {
	DeregistrationDelay            int               `json:"deregistration_delay"`
	ProxyProtocolV2                bool              `json:"proxy_protocol_v2"`
	Tags                           map[string]string `json:"tags"`
	LambdaMultiValueHeadersEnabled bool              `json:"lambda_multi_value_headers_enabled"`
	NamePrefix                     string            `json:"name_prefix"`
	Port                           int               `json:"port"`
	TargetType                     string            `json:"target_type"`
	VpcId                          string            `json:"vpc_id"`
	Protocol                       string            `json:"protocol"`
	SlowStart                      int               `json:"slow_start"`
	HealthCheck                    HealthCheck       `json:"health_check"`
	Stickiness                     Stickiness        `json:"stickiness"`
}

type HealthCheck struct {
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	Interval           int    `json:"interval"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	Matcher            string `json:"matcher"`
	Path               string `json:"path"`
	Port               string `json:"port"`
	Protocol           string `json:"protocol"`
	Timeout            int    `json:"timeout"`
	Enabled            bool   `json:"enabled"`
}

type Stickiness struct {
	CookieDuration int    `json:"cookie_duration"`
	Enabled        bool   `json:"enabled"`
	Type           string `json:"type"`
}

// A AlbTargetGroupStatus defines the observed state of a AlbTargetGroup
type AlbTargetGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AlbTargetGroupObservation `json:",inline"`
}

// A AlbTargetGroupObservation records the observed state of a AlbTargetGroup
type AlbTargetGroupObservation struct {
	ArnSuffix                  string `json:"arn_suffix"`
	LoadBalancingAlgorithmType string `json:"load_balancing_algorithm_type"`
	Arn                        string `json:"arn"`
	Id                         string `json:"id"`
	Name                       string `json:"name"`
}