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

// CognitoIdentityPoolRolesAttachment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CognitoIdentityPoolRolesAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CognitoIdentityPoolRolesAttachmentSpec   `json:"spec"`
	Status CognitoIdentityPoolRolesAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoIdentityPoolRolesAttachment contains a list of CognitoIdentityPoolRolesAttachmentList
type CognitoIdentityPoolRolesAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoIdentityPoolRolesAttachment `json:"items"`
}

// A CognitoIdentityPoolRolesAttachmentSpec defines the desired state of a CognitoIdentityPoolRolesAttachment
type CognitoIdentityPoolRolesAttachmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CognitoIdentityPoolRolesAttachmentParameters `json:"forProvider"`
}

// A CognitoIdentityPoolRolesAttachmentParameters defines the desired state of a CognitoIdentityPoolRolesAttachment
type CognitoIdentityPoolRolesAttachmentParameters struct {
	Id             string            `json:"id"`
	IdentityPoolId string            `json:"identity_pool_id"`
	Roles          map[string]string `json:"roles"`
	RoleMapping    RoleMapping       `json:"role_mapping"`
}

type RoleMapping struct {
	Type                    string        `json:"type"`
	AmbiguousRoleResolution string        `json:"ambiguous_role_resolution"`
	IdentityProvider        string        `json:"identity_provider"`
	MappingRule             []MappingRule `json:"mapping_rule"`
}

type MappingRule struct {
	Claim     string `json:"claim"`
	MatchType string `json:"match_type"`
	RoleArn   string `json:"role_arn"`
	Value     string `json:"value"`
}

// A CognitoIdentityPoolRolesAttachmentStatus defines the observed state of a CognitoIdentityPoolRolesAttachment
type CognitoIdentityPoolRolesAttachmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CognitoIdentityPoolRolesAttachmentObservation `json:"atProvider"`
}

// A CognitoIdentityPoolRolesAttachmentObservation records the observed state of a CognitoIdentityPoolRolesAttachment
type CognitoIdentityPoolRolesAttachmentObservation struct{}