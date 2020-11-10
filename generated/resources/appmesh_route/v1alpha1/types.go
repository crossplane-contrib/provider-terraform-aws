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
	Tags              map[string]string `json:"tags"`
	Name              string            `json:"name"`
	VirtualRouterName string            `json:"virtual_router_name"`
	MeshName          string            `json:"mesh_name"`
	Spec              Spec              `json:"spec"`
}

type Spec struct {
	Priority   int        `json:"priority"`
	TcpRoute   TcpRoute   `json:"tcp_route"`
	GrpcRoute  GrpcRoute  `json:"grpc_route"`
	Http2Route Http2Route `json:"http2_route"`
	HttpRoute  HttpRoute  `json:"http_route"`
}

type TcpRoute struct {
	Action Action `json:"action"`
}

type Action struct {
	WeightedTarget []WeightedTarget `json:"weighted_target"`
}

type WeightedTarget struct {
	VirtualNode string `json:"virtual_node"`
	Weight      int    `json:"weight"`
}

type GrpcRoute struct {
	Action      Action      `json:"action"`
	Match       Match       `json:"match"`
	RetryPolicy RetryPolicy `json:"retry_policy"`
}

type Action struct {
	WeightedTarget []WeightedTarget `json:"weighted_target"`
}

type WeightedTarget struct {
	VirtualNode string `json:"virtual_node"`
	Weight      int    `json:"weight"`
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
	Exact  string `json:"exact"`
	Prefix string `json:"prefix"`
	Regex  string `json:"regex"`
	Suffix string `json:"suffix"`
	Range  Range  `json:"range"`
}

type Range struct {
	End   int `json:"end"`
	Start int `json:"start"`
}

type RetryPolicy struct {
	GrpcRetryEvents []string        `json:"grpc_retry_events"`
	HttpRetryEvents []string        `json:"http_retry_events"`
	MaxRetries      int             `json:"max_retries"`
	TcpRetryEvents  []string        `json:"tcp_retry_events"`
	PerRetryTimeout PerRetryTimeout `json:"per_retry_timeout"`
}

type PerRetryTimeout struct {
	Unit  string `json:"unit"`
	Value int    `json:"value"`
}

type Http2Route struct {
	Match       Match       `json:"match"`
	RetryPolicy RetryPolicy `json:"retry_policy"`
	Action      Action      `json:"action"`
}

type Match struct {
	Method string   `json:"method"`
	Prefix string   `json:"prefix"`
	Scheme string   `json:"scheme"`
	Header []Header `json:"header"`
}

type Header struct {
	Invert bool   `json:"invert"`
	Name   string `json:"name"`
	Match  Match  `json:"match"`
}

type Match struct {
	Exact  string `json:"exact"`
	Prefix string `json:"prefix"`
	Regex  string `json:"regex"`
	Suffix string `json:"suffix"`
	Range  Range  `json:"range"`
}

type Range struct {
	End   int `json:"end"`
	Start int `json:"start"`
}

type RetryPolicy struct {
	HttpRetryEvents []string        `json:"http_retry_events"`
	MaxRetries      int             `json:"max_retries"`
	TcpRetryEvents  []string        `json:"tcp_retry_events"`
	PerRetryTimeout PerRetryTimeout `json:"per_retry_timeout"`
}

type PerRetryTimeout struct {
	Unit  string `json:"unit"`
	Value int    `json:"value"`
}

type Action struct {
	WeightedTarget []WeightedTarget `json:"weighted_target"`
}

type WeightedTarget struct {
	VirtualNode string `json:"virtual_node"`
	Weight      int    `json:"weight"`
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
	Weight      int    `json:"weight"`
}

type Match struct {
	Method string   `json:"method"`
	Prefix string   `json:"prefix"`
	Scheme string   `json:"scheme"`
	Header []Header `json:"header"`
}

type Header struct {
	Invert bool   `json:"invert"`
	Name   string `json:"name"`
	Match  Match  `json:"match"`
}

type Match struct {
	Exact  string `json:"exact"`
	Prefix string `json:"prefix"`
	Regex  string `json:"regex"`
	Suffix string `json:"suffix"`
	Range  Range  `json:"range"`
}

type Range struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type RetryPolicy struct {
	HttpRetryEvents []string        `json:"http_retry_events"`
	MaxRetries      int             `json:"max_retries"`
	TcpRetryEvents  []string        `json:"tcp_retry_events"`
	PerRetryTimeout PerRetryTimeout `json:"per_retry_timeout"`
}

type PerRetryTimeout struct {
	Unit  string `json:"unit"`
	Value int    `json:"value"`
}

// A AppmeshRouteStatus defines the observed state of a AppmeshRoute
type AppmeshRouteStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppmeshRouteObservation `json:",inline"`
}

// A AppmeshRouteObservation records the observed state of a AppmeshRoute
type AppmeshRouteObservation struct {
	Arn             string `json:"arn"`
	MeshOwner       string `json:"mesh_owner"`
	ResourceOwner   string `json:"resource_owner"`
	CreatedDate     string `json:"created_date"`
	Id              string `json:"id"`
	LastUpdatedDate string `json:"last_updated_date"`
}