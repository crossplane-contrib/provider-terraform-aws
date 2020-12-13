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
	ForProvider                  LbTargetGroupParameters `json:"forProvider"`
}

// A LbTargetGroupParameters defines the desired state of a LbTargetGroup
type LbTargetGroupParameters struct {
	Id                             string            `json:"id"`
	Tags                           map[string]string `json:"tags"`
	NamePrefix                     string            `json:"name_prefix"`
	Port                           int64             `json:"port"`
	SlowStart                      int64             `json:"slow_start"`
	ProxyProtocolV2                bool              `json:"proxy_protocol_v2"`
	DeregistrationDelay            int64             `json:"deregistration_delay"`
	LambdaMultiValueHeadersEnabled bool              `json:"lambda_multi_value_headers_enabled"`
	LoadBalancingAlgorithmType     string            `json:"load_balancing_algorithm_type"`
	Name                           string            `json:"name"`
	Protocol                       string            `json:"protocol"`
	TargetType                     string            `json:"target_type"`
	VpcId                          string            `json:"vpc_id"`
	Stickiness                     Stickiness        `json:"stickiness"`
	HealthCheck                    HealthCheck       `json:"health_check"`
}

type Stickiness struct {
	CookieDuration int64  `json:"cookie_duration"`
	Enabled        bool   `json:"enabled"`
	Type           string `json:"type"`
}

type HealthCheck struct {
	Path               string `json:"path"`
	Protocol           string `json:"protocol"`
	Timeout            int64  `json:"timeout"`
	HealthyThreshold   int64  `json:"healthy_threshold"`
	Matcher            string `json:"matcher"`
	Port               string `json:"port"`
	UnhealthyThreshold int64  `json:"unhealthy_threshold"`
	Enabled            bool   `json:"enabled"`
	Interval           int64  `json:"interval"`
}

// A LbTargetGroupStatus defines the observed state of a LbTargetGroup
type LbTargetGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LbTargetGroupObservation `json:"atProvider"`
}

// A LbTargetGroupObservation records the observed state of a LbTargetGroup
type LbTargetGroupObservation struct {
	ArnSuffix string `json:"arn_suffix"`
	Arn       string `json:"arn"`
}