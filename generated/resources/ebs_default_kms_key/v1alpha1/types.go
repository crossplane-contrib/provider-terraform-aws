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

// EbsDefaultKmsKey is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EbsDefaultKmsKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EbsDefaultKmsKeySpec   `json:"spec"`
	Status EbsDefaultKmsKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EbsDefaultKmsKey contains a list of EbsDefaultKmsKeyList
type EbsDefaultKmsKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EbsDefaultKmsKey `json:"items"`
}

// A EbsDefaultKmsKeySpec defines the desired state of a EbsDefaultKmsKey
type EbsDefaultKmsKeySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EbsDefaultKmsKeyParameters `json:"forProvider"`
}

// A EbsDefaultKmsKeyParameters defines the desired state of a EbsDefaultKmsKey
type EbsDefaultKmsKeyParameters struct {
	Id     string `json:"id"`
	KeyArn string `json:"key_arn"`
}

// A EbsDefaultKmsKeyStatus defines the observed state of a EbsDefaultKmsKey
type EbsDefaultKmsKeyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EbsDefaultKmsKeyObservation `json:"atProvider"`
}

// A EbsDefaultKmsKeyObservation records the observed state of a EbsDefaultKmsKey
type EbsDefaultKmsKeyObservation struct{}