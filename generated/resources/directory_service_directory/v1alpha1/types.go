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

// DirectoryServiceDirectory is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DirectoryServiceDirectory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DirectoryServiceDirectorySpec   `json:"spec"`
	Status DirectoryServiceDirectoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DirectoryServiceDirectory contains a list of DirectoryServiceDirectoryList
type DirectoryServiceDirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DirectoryServiceDirectory `json:"items"`
}

// A DirectoryServiceDirectorySpec defines the desired state of a DirectoryServiceDirectory
type DirectoryServiceDirectorySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DirectoryServiceDirectoryParameters `json:",inline"`
}

// A DirectoryServiceDirectoryParameters defines the desired state of a DirectoryServiceDirectory
type DirectoryServiceDirectoryParameters struct {
	Description     string            `json:"description"`
	EnableSso       bool              `json:"enable_sso"`
	ShortName       string            `json:"short_name"`
	Tags            map[string]string `json:"tags"`
	Edition         string            `json:"edition"`
	Password        string            `json:"password"`
	Size            string            `json:"size"`
	Type            string            `json:"type"`
	Alias           string            `json:"alias"`
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	ConnectSettings ConnectSettings   `json:"connect_settings"`
	VpcSettings     VpcSettings       `json:"vpc_settings"`
}

type ConnectSettings struct {
	VpcId             string   `json:"vpc_id"`
	AvailabilityZones []string `json:"availability_zones"`
	ConnectIps        []string `json:"connect_ips"`
	CustomerDnsIps    []string `json:"customer_dns_ips"`
	CustomerUsername  string   `json:"customer_username"`
	SubnetIds         []string `json:"subnet_ids"`
}

type VpcSettings struct {
	SubnetIds         []string `json:"subnet_ids"`
	VpcId             string   `json:"vpc_id"`
	AvailabilityZones []string `json:"availability_zones"`
}

// A DirectoryServiceDirectoryStatus defines the observed state of a DirectoryServiceDirectory
type DirectoryServiceDirectoryStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DirectoryServiceDirectoryObservation `json:",inline"`
}

// A DirectoryServiceDirectoryObservation records the observed state of a DirectoryServiceDirectory
type DirectoryServiceDirectoryObservation struct {
	AccessUrl       string   `json:"access_url"`
	DnsIpAddresses  []string `json:"dns_ip_addresses"`
	SecurityGroupId string   `json:"security_group_id"`
}