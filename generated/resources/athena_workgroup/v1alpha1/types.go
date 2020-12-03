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

// AthenaWorkgroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AthenaWorkgroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AthenaWorkgroupSpec   `json:"spec"`
	Status AthenaWorkgroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AthenaWorkgroup contains a list of AthenaWorkgroupList
type AthenaWorkgroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AthenaWorkgroup `json:"items"`
}

// A AthenaWorkgroupSpec defines the desired state of a AthenaWorkgroup
type AthenaWorkgroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AthenaWorkgroupParameters `json:",inline"`
}

// A AthenaWorkgroupParameters defines the desired state of a AthenaWorkgroup
type AthenaWorkgroupParameters struct {
	Id            string            `json:"id"`
	Name          string            `json:"name"`
	State         string            `json:"state"`
	Tags          map[string]string `json:"tags"`
	Description   string            `json:"description"`
	ForceDestroy  bool              `json:"force_destroy"`
	Configuration Configuration     `json:"configuration"`
}

type Configuration struct {
	BytesScannedCutoffPerQuery      int                 `json:"bytes_scanned_cutoff_per_query"`
	EnforceWorkgroupConfiguration   bool                `json:"enforce_workgroup_configuration"`
	PublishCloudwatchMetricsEnabled bool                `json:"publish_cloudwatch_metrics_enabled"`
	ResultConfiguration             ResultConfiguration `json:"result_configuration"`
}

type ResultConfiguration struct {
	OutputLocation          string                  `json:"output_location"`
	EncryptionConfiguration EncryptionConfiguration `json:"encryption_configuration"`
}

type EncryptionConfiguration struct {
	EncryptionOption string `json:"encryption_option"`
	KmsKeyArn        string `json:"kms_key_arn"`
}

// A AthenaWorkgroupStatus defines the observed state of a AthenaWorkgroup
type AthenaWorkgroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AthenaWorkgroupObservation `json:",inline"`
}

// A AthenaWorkgroupObservation records the observed state of a AthenaWorkgroup
type AthenaWorkgroupObservation struct {
	Arn string `json:"arn"`
}