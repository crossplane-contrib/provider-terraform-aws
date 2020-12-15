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

// CloudfrontOriginAccessIdentity is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudfrontOriginAccessIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudfrontOriginAccessIdentitySpec   `json:"spec"`
	Status CloudfrontOriginAccessIdentityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontOriginAccessIdentity contains a list of CloudfrontOriginAccessIdentityList
type CloudfrontOriginAccessIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudfrontOriginAccessIdentity `json:"items"`
}

// A CloudfrontOriginAccessIdentitySpec defines the desired state of a CloudfrontOriginAccessIdentity
type CloudfrontOriginAccessIdentitySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudfrontOriginAccessIdentityParameters `json:"forProvider"`
}

// A CloudfrontOriginAccessIdentityParameters defines the desired state of a CloudfrontOriginAccessIdentity
type CloudfrontOriginAccessIdentityParameters struct {
	Comment string `json:"comment"`
}

// A CloudfrontOriginAccessIdentityStatus defines the observed state of a CloudfrontOriginAccessIdentity
type CloudfrontOriginAccessIdentityStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudfrontOriginAccessIdentityObservation `json:"atProvider"`
}

// A CloudfrontOriginAccessIdentityObservation records the observed state of a CloudfrontOriginAccessIdentity
type CloudfrontOriginAccessIdentityObservation struct {
	IamArn                       string `json:"iam_arn"`
	S3CanonicalUserId            string `json:"s3_canonical_user_id"`
	CallerReference              string `json:"caller_reference"`
	CloudfrontAccessIdentityPath string `json:"cloudfront_access_identity_path"`
	Etag                         string `json:"etag"`
}