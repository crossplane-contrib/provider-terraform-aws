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

// LbListenerRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LbListenerRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LbListenerRuleSpec   `json:"spec"`
	Status LbListenerRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbListenerRule contains a list of LbListenerRuleList
type LbListenerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbListenerRule `json:"items"`
}

// A LbListenerRuleSpec defines the desired state of a LbListenerRule
type LbListenerRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LbListenerRuleParameters `json:",inline"`
}

// A LbListenerRuleParameters defines the desired state of a LbListenerRule
type LbListenerRuleParameters struct {
	ListenerArn string      `json:"listener_arn"`
	Priority    int         `json:"priority"`
	Id          string      `json:"id"`
	Action      []Action    `json:"action"`
	Condition   []Condition `json:"condition"`
}

type Action struct {
	TargetGroupArn      string              `json:"target_group_arn"`
	Type                string              `json:"type"`
	Order               int                 `json:"order"`
	AuthenticateCognito AuthenticateCognito `json:"authenticate_cognito"`
	AuthenticateOidc    AuthenticateOidc    `json:"authenticate_oidc"`
	FixedResponse       FixedResponse       `json:"fixed_response"`
	Forward             Forward             `json:"forward"`
	Redirect            Redirect            `json:"redirect"`
}

type AuthenticateCognito struct {
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   int               `json:"session_timeout"`
	UserPoolArn                      string            `json:"user_pool_arn"`
}

type AuthenticateOidc struct {
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	AuthorizationEndpoint            string            `json:"authorization_endpoint"`
	Issuer                           string            `json:"issuer"`
	Scope                            string            `json:"scope"`
	TokenEndpoint                    string            `json:"token_endpoint"`
	UserInfoEndpoint                 string            `json:"user_info_endpoint"`
	ClientId                         string            `json:"client_id"`
	ClientSecret                     string            `json:"client_secret"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   int               `json:"session_timeout"`
}

type FixedResponse struct {
	ContentType string `json:"content_type"`
	MessageBody string `json:"message_body"`
	StatusCode  string `json:"status_code"`
}

type Forward struct {
	Stickiness  Stickiness    `json:"stickiness"`
	TargetGroup []TargetGroup `json:"target_group"`
}

type Stickiness struct {
	Duration int  `json:"duration"`
	Enabled  bool `json:"enabled"`
}

type TargetGroup struct {
	Arn    string `json:"arn"`
	Weight int    `json:"weight"`
}

type Redirect struct {
	Protocol   string `json:"protocol"`
	Query      string `json:"query"`
	StatusCode string `json:"status_code"`
	Host       string `json:"host"`
	Path       string `json:"path"`
	Port       string `json:"port"`
}

type Condition struct {
	HostHeader        HostHeader        `json:"host_header"`
	HttpHeader        HttpHeader        `json:"http_header"`
	HttpRequestMethod HttpRequestMethod `json:"http_request_method"`
	PathPattern       PathPattern       `json:"path_pattern"`
	QueryString       []QueryString     `json:"query_string"`
	SourceIp          SourceIp          `json:"source_ip"`
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

type PathPattern struct {
	Values []string `json:"values"`
}

type QueryString struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SourceIp struct {
	Values []string `json:"values"`
}

// A LbListenerRuleStatus defines the observed state of a LbListenerRule
type LbListenerRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LbListenerRuleObservation `json:",inline"`
}

// A LbListenerRuleObservation records the observed state of a LbListenerRule
type LbListenerRuleObservation struct {
	Arn string `json:"arn"`
}