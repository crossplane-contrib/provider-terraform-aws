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

// WorklinkFleet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WorklinkFleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorklinkFleetSpec   `json:"spec"`
	Status WorklinkFleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorklinkFleet contains a list of WorklinkFleetList
type WorklinkFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorklinkFleet `json:"items"`
}

// A WorklinkFleetSpec defines the desired state of a WorklinkFleet
type WorklinkFleetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WorklinkFleetParameters `json:",inline"`
}

// A WorklinkFleetParameters defines the desired state of a WorklinkFleet
type WorklinkFleetParameters struct {
	DeviceCaCertificate        string           `json:"device_ca_certificate"`
	Id                         string           `json:"id"`
	Name                       string           `json:"name"`
	OptimizeForEndUserLocation bool             `json:"optimize_for_end_user_location"`
	AuditStreamArn             string           `json:"audit_stream_arn"`
	DisplayName                string           `json:"display_name"`
	Network                    Network          `json:"network"`
	IdentityProvider           IdentityProvider `json:"identity_provider"`
}

type Network struct {
	SecurityGroupIds []string `json:"security_group_ids"`
	SubnetIds        []string `json:"subnet_ids"`
	VpcId            string   `json:"vpc_id"`
}

type IdentityProvider struct {
	SamlMetadata string `json:"saml_metadata"`
	Type         string `json:"type"`
}

// A WorklinkFleetStatus defines the observed state of a WorklinkFleet
type WorklinkFleetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WorklinkFleetObservation `json:",inline"`
}

// A WorklinkFleetObservation records the observed state of a WorklinkFleet
type WorklinkFleetObservation struct {
	Arn             string `json:"arn"`
	CompanyCode     string `json:"company_code"`
	CreatedTime     string `json:"created_time"`
	LastUpdatedTime string `json:"last_updated_time"`
}