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

// DmsEndpoint is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DmsEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DmsEndpointSpec   `json:"spec"`
	Status DmsEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DmsEndpoint contains a list of DmsEndpointList
type DmsEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DmsEndpoint `json:"items"`
}

// A DmsEndpointSpec defines the desired state of a DmsEndpoint
type DmsEndpointSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DmsEndpointParameters `json:",inline"`
}

// A DmsEndpointParameters defines the desired state of a DmsEndpoint
type DmsEndpointParameters struct {
	ServiceAccessRole string `json:"service_access_role"`
	EngineName        string `json:"engine_name"`
	Password          string `json:"password"`
	ServerName        string `json:"server_name"`
	Username          string `json:"username"`
	EndpointId        string `json:"endpoint_id"`
	Port              int    `json:"port"`
	DatabaseName      string `json:"database_name"`
	EndpointType      string `json:"endpoint_type"`
}

// A DmsEndpointStatus defines the observed state of a DmsEndpoint
type DmsEndpointStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DmsEndpointObservation `json:",inline"`
}

// A DmsEndpointObservation records the observed state of a DmsEndpoint
type DmsEndpointObservation struct {
	CertificateArn            string `json:"certificate_arn"`
	ExtraConnectionAttributes string `json:"extra_connection_attributes"`
	SslMode                   string `json:"ssl_mode"`
	EndpointArn               string `json:"endpoint_arn"`
	Id                        string `json:"id"`
	KmsKeyArn                 string `json:"kms_key_arn"`
}