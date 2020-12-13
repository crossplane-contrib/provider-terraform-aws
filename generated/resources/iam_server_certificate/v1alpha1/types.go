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

// IamServerCertificate is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamServerCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamServerCertificateSpec   `json:"spec"`
	Status IamServerCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamServerCertificate contains a list of IamServerCertificateList
type IamServerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamServerCertificate `json:"items"`
}

// A IamServerCertificateSpec defines the desired state of a IamServerCertificate
type IamServerCertificateSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamServerCertificateParameters `json:"forProvider"`
}

// A IamServerCertificateParameters defines the desired state of a IamServerCertificate
type IamServerCertificateParameters struct {
	PrivateKey       string `json:"private_key"`
	Arn              string `json:"arn"`
	CertificateBody  string `json:"certificate_body"`
	CertificateChain string `json:"certificate_chain"`
	Id               string `json:"id"`
	Name             string `json:"name"`
	NamePrefix       string `json:"name_prefix"`
	Path             string `json:"path"`
}

// A IamServerCertificateStatus defines the observed state of a IamServerCertificate
type IamServerCertificateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamServerCertificateObservation `json:"atProvider"`
}

// A IamServerCertificateObservation records the observed state of a IamServerCertificate
type IamServerCertificateObservation struct{}