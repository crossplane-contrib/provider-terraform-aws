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

// GuarddutyDetector is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GuarddutyDetector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GuarddutyDetectorSpec   `json:"spec"`
	Status GuarddutyDetectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuarddutyDetector contains a list of GuarddutyDetectorList
type GuarddutyDetectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GuarddutyDetector `json:"items"`
}

// A GuarddutyDetectorSpec defines the desired state of a GuarddutyDetector
type GuarddutyDetectorSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GuarddutyDetectorParameters `json:",inline"`
}

// A GuarddutyDetectorParameters defines the desired state of a GuarddutyDetector
type GuarddutyDetectorParameters struct {
	Enable bool `json:"enable"`
}

// A GuarddutyDetectorStatus defines the observed state of a GuarddutyDetector
type GuarddutyDetectorStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GuarddutyDetectorObservation `json:",inline"`
}

// A GuarddutyDetectorObservation records the observed state of a GuarddutyDetector
type GuarddutyDetectorObservation struct {
	Arn                        string `json:"arn"`
	FindingPublishingFrequency string `json:"finding_publishing_frequency"`
	Id                         string `json:"id"`
	AccountId                  string `json:"account_id"`
}