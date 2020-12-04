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

// AppmeshRoute is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppmeshRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppmeshRouteSpec   `json:"spec"`
	Status AppmeshRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshRoute contains a list of AppmeshRouteList
type AppmeshRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppmeshRoute `json:"items"`
}

// A AppmeshRouteSpec defines the desired state of a AppmeshRoute
type AppmeshRouteSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppmeshRouteParameters `json:",inline"`
}

// A AppmeshRouteParameters defines the desired state of a AppmeshRoute
type AppmeshRouteParameters struct {
	MeshOwner         string            `json:"mesh_owner"`
	Tags              map[string]string `json:"tags"`
	VirtualRouterName string            `json:"virtual_router_name"`
	Id                string            `json:"id"`
	MeshName          string            `json:"mesh_name"`
	Name              string            `json:"name"`
	Spec              Spec              `json:"spec"`
}

type Spec struct {
	Priority   int64      `json:"priority"`
	GrpcRoute  GrpcRoute  `json:"grpc_route"`
	Http2Route Http2Route `json:"http2_route"`
	HttpRoute  HttpRoute  `json:"http_route"`
	TcpRoute   TcpRoute   `json:"tcp_route"`
}

type GrpcRoute struct {
	RetryPolicy RetryPolicy `json:"retry_policy"`
	Action      Action      `json:"action"`
	Match       Match       `json:"match"`
}

type RetryPolicy struct {
	GrpcRetryEvents []string        `json:"grpc_retry_events"`
	HttpRetryEvents []string        `json:"http_retry_events"`
	MaxRetries      int64           `json:"max_retries"`
	TcpRetryEvents  []string        `json:"tcp_retry_events"`
	PerRetryTimeout PerRetryTimeout `json:"per_retry_timeout"`
}

type PerRetryTimeout struct {
	Unit  string `json:"unit"`
	Value int64  `json:"value"`
}

type Action struct {
	WeightedTarget []WeightedTarget `json:"weighted_target"`
}

type WeightedTarget struct {
	VirtualNode string `json:"virtual_node"`
	Weight      int64  `json:"weight"`
}

type Match struct {
	MethodName  string     `json:"method_name"`
	ServiceName string     `json:"service_name"`
	Metadata    []Metadata `json:"metadata"`
}

type Metadata struct {
	Invert bool   `json:"invert"`
	Name   string `json:"name"`
	Match  Match  `json:"match"`
}

type Match struct {
	Prefix string `json:"prefix"`
	Regex  string `json:"regex"`
	Suffix string `json:"suffix"`
	Exact  string `json:"exact"`
	Range  Range  `json:"range"`
}

type Range struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
}

type Http2Route struct {
	Action      Action      `json:"action"`
	Match       Match       `json:"match"`
	RetryPolicy RetryPolicy `json:"retry_policy"`
}

type Action struct {
	WeightedTarget []WeightedTarget `json:"weighted_target"`
}

type WeightedTarget struct {
	VirtualNode string `json:"virtual_node"`
	Weight      int64  `json:"weight"`
}

type Match struct {
	Scheme string   `json:"scheme"`
	Method string   `json:"method"`
	Prefix string   `json:"prefix"`
	Header []Header `json:"header"`
}

type Header struct {
	Name   string `json:"name"`
	Invert bool   `json:"invert"`
	Match  Match  `json:"match"`
}

type Match struct {
	Regex  string `json:"regex"`
	Suffix string `json:"suffix"`
	Exact  string `json:"exact"`
	Prefix string `json:"prefix"`
	Range  Range  `json:"range"`
}

type Range struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
}

type RetryPolicy struct {
	MaxRetries      int64           `json:"max_retries"`
	TcpRetryEvents  []string        `json:"tcp_retry_events"`
	HttpRetryEvents []string        `json:"http_retry_events"`
	PerRetryTimeout PerRetryTimeout `json:"per_retry_timeout"`
}

type PerRetryTimeout struct {
	Unit  string `json:"unit"`
	Value int64  `json:"value"`
}

type HttpRoute struct {
	Action      Action      `json:"action"`
	Match       Match       `json:"match"`
	RetryPolicy RetryPolicy `json:"retry_policy"`
}

type Action struct {
	WeightedTarget []WeightedTarget `json:"weighted_target"`
}

type WeightedTarget struct {
	VirtualNode string `json:"virtual_node"`
	Weight      int64  `json:"weight"`
}

type Match struct {
	Prefix string   `json:"prefix"`
	Scheme string   `json:"scheme"`
	Method string   `json:"method"`
	Header []Header `json:"header"`
}

type Header struct {
	Invert bool   `json:"invert"`
	Name   string `json:"name"`
	Match  Match  `json:"match"`
}

type Match struct {
	Regex  string `json:"regex"`
	Suffix string `json:"suffix"`
	Exact  string `json:"exact"`
	Prefix string `json:"prefix"`
	Range  Range  `json:"range"`
}

type Range struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
}

type RetryPolicy struct {
	HttpRetryEvents []string        `json:"http_retry_events"`
	MaxRetries      int64           `json:"max_retries"`
	TcpRetryEvents  []string        `json:"tcp_retry_events"`
	PerRetryTimeout PerRetryTimeout `json:"per_retry_timeout"`
}

type PerRetryTimeout struct {
	Unit  string `json:"unit"`
	Value int64  `json:"value"`
}

type TcpRoute struct {
	Action Action `json:"action"`
}

type Action struct {
	WeightedTarget []WeightedTarget `json:"weighted_target"`
}

type WeightedTarget struct {
	VirtualNode string `json:"virtual_node"`
	Weight      int64  `json:"weight"`
}

// A AppmeshRouteStatus defines the observed state of a AppmeshRoute
type AppmeshRouteStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppmeshRouteObservation `json:",inline"`
}

// A AppmeshRouteObservation records the observed state of a AppmeshRoute
type AppmeshRouteObservation struct {
	ResourceOwner   string `json:"resource_owner"`
	Arn             string `json:"arn"`
	CreatedDate     string `json:"created_date"`
	LastUpdatedDate string `json:"last_updated_date"`
}