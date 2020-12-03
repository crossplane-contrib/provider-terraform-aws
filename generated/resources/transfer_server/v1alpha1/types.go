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

// TransferServer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type TransferServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TransferServerSpec   `json:"spec"`
	Status TransferServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransferServer contains a list of TransferServerList
type TransferServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransferServer `json:"items"`
}

// A TransferServerSpec defines the desired state of a TransferServer
type TransferServerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  TransferServerParameters `json:",inline"`
}

// A TransferServerParameters defines the desired state of a TransferServer
type TransferServerParameters struct {
	IdentityProviderType string            `json:"identity_provider_type"`
	Url                  string            `json:"url"`
	EndpointType         string            `json:"endpoint_type"`
	HostKey              string            `json:"host_key"`
	Id                   string            `json:"id"`
	InvocationRole       string            `json:"invocation_role"`
	LoggingRole          string            `json:"logging_role"`
	Tags                 map[string]string `json:"tags"`
	ForceDestroy         bool              `json:"force_destroy"`
	EndpointDetails      EndpointDetails   `json:"endpoint_details"`
}

type EndpointDetails struct {
	AddressAllocationIds []string `json:"address_allocation_ids"`
	SubnetIds            []string `json:"subnet_ids"`
	VpcEndpointId        string   `json:"vpc_endpoint_id"`
	VpcId                string   `json:"vpc_id"`
}

// A TransferServerStatus defines the observed state of a TransferServer
type TransferServerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     TransferServerObservation `json:",inline"`
}

// A TransferServerObservation records the observed state of a TransferServer
type TransferServerObservation struct {
	Arn                string `json:"arn"`
	HostKeyFingerprint string `json:"host_key_fingerprint"`
	Endpoint           string `json:"endpoint"`
}