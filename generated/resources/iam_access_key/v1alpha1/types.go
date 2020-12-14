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

// IamAccessKey is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamAccessKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamAccessKeySpec   `json:"spec"`
	Status IamAccessKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamAccessKey contains a list of IamAccessKeyList
type IamAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamAccessKey `json:"items"`
}

// A IamAccessKeySpec defines the desired state of a IamAccessKey
type IamAccessKeySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamAccessKeyParameters `json:"forProvider"`
}

// A IamAccessKeyParameters defines the desired state of a IamAccessKey
type IamAccessKeyParameters struct {
	PgpKey string `json:"pgp_key"`
	Status string `json:"status"`
	User   string `json:"user"`
	Id     string `json:"id"`
}

// A IamAccessKeyStatus defines the observed state of a IamAccessKey
type IamAccessKeyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamAccessKeyObservation `json:"atProvider"`
}

// A IamAccessKeyObservation records the observed state of a IamAccessKey
type IamAccessKeyObservation struct {
	KeyFingerprint    string `json:"key_fingerprint"`
	Secret            string `json:"secret"`
	SesSmtpPasswordV4 string `json:"ses_smtp_password_v4"`
	EncryptedSecret   string `json:"encrypted_secret"`
}