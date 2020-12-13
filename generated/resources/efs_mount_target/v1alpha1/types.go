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

// EfsMountTarget is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EfsMountTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EfsMountTargetSpec   `json:"spec"`
	Status EfsMountTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EfsMountTarget contains a list of EfsMountTargetList
type EfsMountTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EfsMountTarget `json:"items"`
}

// A EfsMountTargetSpec defines the desired state of a EfsMountTarget
type EfsMountTargetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EfsMountTargetParameters `json:"forProvider"`
}

// A EfsMountTargetParameters defines the desired state of a EfsMountTarget
type EfsMountTargetParameters struct {
	FileSystemId   string   `json:"file_system_id"`
	Id             string   `json:"id"`
	SubnetId       string   `json:"subnet_id"`
	IpAddress      string   `json:"ip_address"`
	SecurityGroups []string `json:"security_groups"`
}

// A EfsMountTargetStatus defines the observed state of a EfsMountTarget
type EfsMountTargetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EfsMountTargetObservation `json:"atProvider"`
}

// A EfsMountTargetObservation records the observed state of a EfsMountTarget
type EfsMountTargetObservation struct {
	DnsName              string `json:"dns_name"`
	FileSystemArn        string `json:"file_system_arn"`
	NetworkInterfaceId   string `json:"network_interface_id"`
	AvailabilityZoneId   string `json:"availability_zone_id"`
	MountTargetDnsName   string `json:"mount_target_dns_name"`
	OwnerId              string `json:"owner_id"`
	AvailabilityZoneName string `json:"availability_zone_name"`
}