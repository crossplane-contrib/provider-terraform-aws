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

// LightsailInstance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LightsailInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LightsailInstanceSpec   `json:"spec"`
	Status LightsailInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LightsailInstance contains a list of LightsailInstanceList
type LightsailInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LightsailInstance `json:"items"`
}

// A LightsailInstanceSpec defines the desired state of a LightsailInstance
type LightsailInstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LightsailInstanceParameters `json:",inline"`
}

// A LightsailInstanceParameters defines the desired state of a LightsailInstance
type LightsailInstanceParameters struct {
	AvailabilityZone string `json:"availability_zone"`
	KeyPairName      string `json:"key_pair_name"`
	Name             string `json:"name"`
	UserData         string `json:"user_data"`
	BlueprintId      string `json:"blueprint_id"`
	BundleId         string `json:"bundle_id"`
}

// A LightsailInstanceStatus defines the observed state of a LightsailInstance
type LightsailInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LightsailInstanceObservation `json:",inline"`
}

// A LightsailInstanceObservation records the observed state of a LightsailInstance
type LightsailInstanceObservation struct {
	Username         string `json:"username"`
	Id               string `json:"id"`
	Ipv6Address      string `json:"ipv6_address"`
	IsStaticIp       bool   `json:"is_static_ip"`
	CpuCount         int    `json:"cpu_count"`
	PrivateIpAddress string `json:"private_ip_address"`
	Arn              string `json:"arn"`
	RamSize          int    `json:"ram_size"`
	CreatedAt        string `json:"created_at"`
	PublicIpAddress  string `json:"public_ip_address"`
}