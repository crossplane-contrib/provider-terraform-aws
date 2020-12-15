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

// AppCookieStickinessPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppCookieStickinessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppCookieStickinessPolicySpec   `json:"spec"`
	Status AppCookieStickinessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppCookieStickinessPolicy contains a list of AppCookieStickinessPolicyList
type AppCookieStickinessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppCookieStickinessPolicy `json:"items"`
}

// A AppCookieStickinessPolicySpec defines the desired state of a AppCookieStickinessPolicy
type AppCookieStickinessPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppCookieStickinessPolicyParameters `json:"forProvider"`
}

// A AppCookieStickinessPolicyParameters defines the desired state of a AppCookieStickinessPolicy
type AppCookieStickinessPolicyParameters struct {
	LoadBalancer string `json:"load_balancer"`
	Name         string `json:"name"`
	CookieName   string `json:"cookie_name"`
	LbPort       int64  `json:"lb_port"`
}

// A AppCookieStickinessPolicyStatus defines the observed state of a AppCookieStickinessPolicy
type AppCookieStickinessPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppCookieStickinessPolicyObservation `json:"atProvider"`
}

// A AppCookieStickinessPolicyObservation records the observed state of a AppCookieStickinessPolicy
type AppCookieStickinessPolicyObservation struct{}