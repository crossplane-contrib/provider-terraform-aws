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

// LbTargetGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LbTargetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LbTargetGroupSpec   `json:"spec"`
	Status LbTargetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbTargetGroup contains a list of LbTargetGroupList
type LbTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbTargetGroup `json:"items"`
}

// A LbTargetGroupSpec defines the desired state of a LbTargetGroup
type LbTargetGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LbTargetGroupParameters `json:",inline"`
}

// A LbTargetGroupParameters defines the desired state of a LbTargetGroup
type LbTargetGroupParameters struct {
	TargetType                     string            `json:"target_type"`
	VpcId                          string            `json:"vpc_id"`
	DeregistrationDelay            int64             `json:"deregistration_delay"`
	LambdaMultiValueHeadersEnabled bool              `json:"lambda_multi_value_headers_enabled"`
	Name                           string            `json:"name"`
	Protocol                       string            `json:"protocol"`
	ProxyProtocolV2                bool              `json:"proxy_protocol_v2"`
	Tags                           map[string]string `json:"tags"`
	LoadBalancingAlgorithmType     string            `json:"load_balancing_algorithm_type"`
	NamePrefix                     string            `json:"name_prefix"`
	SlowStart                      int64             `json:"slow_start"`
	Id                             string            `json:"id"`
	Port                           int64             `json:"port"`
	HealthCheck                    HealthCheck       `json:"health_check"`
	Stickiness                     Stickiness        `json:"stickiness"`
}

type HealthCheck struct {
	Interval           int64  `json:"interval"`
	Port               string `json:"port"`
	Timeout            int64  `json:"timeout"`
	UnhealthyThreshold int64  `json:"unhealthy_threshold"`
	Enabled            bool   `json:"enabled"`
	HealthyThreshold   int64  `json:"healthy_threshold"`
	Matcher            string `json:"matcher"`
	Path               string `json:"path"`
	Protocol           string `json:"protocol"`
}

type Stickiness struct {
	CookieDuration int64  `json:"cookie_duration"`
	Enabled        bool   `json:"enabled"`
	Type           string `json:"type"`
}

// A LbTargetGroupStatus defines the observed state of a LbTargetGroup
type LbTargetGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LbTargetGroupObservation `json:",inline"`
}

// A LbTargetGroupObservation records the observed state of a LbTargetGroup
type LbTargetGroupObservation struct {
	Arn       string `json:"arn"`
	ArnSuffix string `json:"arn_suffix"`
}