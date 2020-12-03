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

// AlbListener is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AlbListener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlbListenerSpec   `json:"spec"`
	Status AlbListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlbListener contains a list of AlbListenerList
type AlbListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlbListener `json:"items"`
}

// A AlbListenerSpec defines the desired state of a AlbListener
type AlbListenerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AlbListenerParameters `json:",inline"`
}

// A AlbListenerParameters defines the desired state of a AlbListener
type AlbListenerParameters struct {
	Id              string          `json:"id"`
	LoadBalancerArn string          `json:"load_balancer_arn"`
	Port            int             `json:"port"`
	Protocol        string          `json:"protocol"`
	SslPolicy       string          `json:"ssl_policy"`
	CertificateArn  string          `json:"certificate_arn"`
	DefaultAction   []DefaultAction `json:"default_action"`
	Timeouts        []Timeouts      `json:"timeouts"`
}

type DefaultAction struct {
	Order               int                 `json:"order"`
	TargetGroupArn      string              `json:"target_group_arn"`
	Type                string              `json:"type"`
	FixedResponse       FixedResponse       `json:"fixed_response"`
	Forward             Forward             `json:"forward"`
	Redirect            Redirect            `json:"redirect"`
	AuthenticateCognito AuthenticateCognito `json:"authenticate_cognito"`
	AuthenticateOidc    AuthenticateOidc    `json:"authenticate_oidc"`
}

type FixedResponse struct {
	StatusCode  string `json:"status_code"`
	ContentType string `json:"content_type"`
	MessageBody string `json:"message_body"`
}

type Forward struct {
	TargetGroup []TargetGroup `json:"target_group"`
	Stickiness  Stickiness    `json:"stickiness"`
}

type TargetGroup struct {
	Arn    string `json:"arn"`
	Weight int    `json:"weight"`
}

type Stickiness struct {
	Duration int  `json:"duration"`
	Enabled  bool `json:"enabled"`
}

type Redirect struct {
	Host       string `json:"host"`
	Path       string `json:"path"`
	Port       string `json:"port"`
	Protocol   string `json:"protocol"`
	Query      string `json:"query"`
	StatusCode string `json:"status_code"`
}

type AuthenticateCognito struct {
	SessionTimeout                   int               `json:"session_timeout"`
	UserPoolArn                      string            `json:"user_pool_arn"`
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
}

type AuthenticateOidc struct {
	ClientSecret                     string            `json:"client_secret"`
	Issuer                           string            `json:"issuer"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	TokenEndpoint                    string            `json:"token_endpoint"`
	AuthorizationEndpoint            string            `json:"authorization_endpoint"`
	ClientId                         string            `json:"client_id"`
	UserInfoEndpoint                 string            `json:"user_info_endpoint"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	SessionTimeout                   int               `json:"session_timeout"`
}

type Timeouts struct {
	Read string `json:"read"`
}

// A AlbListenerStatus defines the observed state of a AlbListener
type AlbListenerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AlbListenerObservation `json:",inline"`
}

// A AlbListenerObservation records the observed state of a AlbListener
type AlbListenerObservation struct {
	Arn string `json:"arn"`
}