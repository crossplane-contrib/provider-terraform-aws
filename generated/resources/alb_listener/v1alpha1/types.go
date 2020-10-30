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
	Port            int    `json:"port"`
	Protocol        string `json:"protocol"`
	CertificateArn  string `json:"certificate_arn"`
	LoadBalancerArn string `json:"load_balancer_arn"`
}

// A AlbListenerStatus defines the observed state of a AlbListener
type AlbListenerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AlbListenerObservation `json:",inline"`
}

// A AlbListenerObservation records the observed state of a AlbListener
type AlbListenerObservation struct {
	SslPolicy string `json:"ssl_policy"`
	Arn       string `json:"arn"`
	Id        string `json:"id"`
}