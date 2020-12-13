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

// AlbListenerRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AlbListenerRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlbListenerRuleSpec   `json:"spec"`
	Status AlbListenerRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlbListenerRule contains a list of AlbListenerRuleList
type AlbListenerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlbListenerRule `json:"items"`
}

// A AlbListenerRuleSpec defines the desired state of a AlbListenerRule
type AlbListenerRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AlbListenerRuleParameters `json:"forProvider"`
}

// A AlbListenerRuleParameters defines the desired state of a AlbListenerRule
type AlbListenerRuleParameters struct {
	Id          string      `json:"id"`
	ListenerArn string      `json:"listener_arn"`
	Priority    int64       `json:"priority"`
	Condition   []Condition `json:"condition"`
	Action      []Action    `json:"action"`
}

type Condition struct {
	PathPattern       PathPattern       `json:"path_pattern"`
	QueryString       QueryString       `json:"query_string"`
	SourceIp          SourceIp          `json:"source_ip"`
	HostHeader        HostHeader        `json:"host_header"`
	HttpHeader        HttpHeader        `json:"http_header"`
	HttpRequestMethod HttpRequestMethod `json:"http_request_method"`
}

type PathPattern struct {
	Values []string `json:"values"`
}

type QueryString struct {
	Value string `json:"value"`
	Key   string `json:"key"`
}

type SourceIp struct {
	Values []string `json:"values"`
}

type HostHeader struct {
	Values []string `json:"values"`
}

type HttpHeader struct {
	HttpHeaderName string   `json:"http_header_name"`
	Values         []string `json:"values"`
}

type HttpRequestMethod struct {
	Values []string `json:"values"`
}

type Action struct {
	Order               int64               `json:"order"`
	TargetGroupArn      string              `json:"target_group_arn"`
	Type                string              `json:"type"`
	AuthenticateCognito AuthenticateCognito `json:"authenticate_cognito"`
	AuthenticateOidc    AuthenticateOidc    `json:"authenticate_oidc"`
	FixedResponse       FixedResponse       `json:"fixed_response"`
	Forward             Forward             `json:"forward"`
	Redirect            Redirect            `json:"redirect"`
}

type AuthenticateCognito struct {
	SessionTimeout                   int64             `json:"session_timeout"`
	UserPoolArn                      string            `json:"user_pool_arn"`
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
}

type AuthenticateOidc struct {
	AuthorizationEndpoint            string            `json:"authorization_endpoint"`
	ClientId                         string            `json:"client_id"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	UserInfoEndpoint                 string            `json:"user_info_endpoint"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	ClientSecret                     string            `json:"client_secret"`
	Issuer                           string            `json:"issuer"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   int64             `json:"session_timeout"`
	TokenEndpoint                    string            `json:"token_endpoint"`
}

type FixedResponse struct {
	MessageBody string `json:"message_body"`
	StatusCode  string `json:"status_code"`
	ContentType string `json:"content_type"`
}

type Forward struct {
	Stickiness  Stickiness    `json:"stickiness"`
	TargetGroup []TargetGroup `json:"target_group"`
}

type Stickiness struct {
	Duration int64 `json:"duration"`
	Enabled  bool  `json:"enabled"`
}

type TargetGroup struct {
	Arn    string `json:"arn"`
	Weight int64  `json:"weight"`
}

type Redirect struct {
	Path       string `json:"path"`
	Port       string `json:"port"`
	Protocol   string `json:"protocol"`
	Query      string `json:"query"`
	StatusCode string `json:"status_code"`
	Host       string `json:"host"`
}

// A AlbListenerRuleStatus defines the observed state of a AlbListenerRule
type AlbListenerRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AlbListenerRuleObservation `json:"atProvider"`
}

// A AlbListenerRuleObservation records the observed state of a AlbListenerRule
type AlbListenerRuleObservation struct {
	Arn string `json:"arn"`
}