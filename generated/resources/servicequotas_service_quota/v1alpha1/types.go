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

// ServicequotasServiceQuota is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ServicequotasServiceQuota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServicequotasServiceQuotaSpec   `json:"spec"`
	Status ServicequotasServiceQuotaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicequotasServiceQuota contains a list of ServicequotasServiceQuotaList
type ServicequotasServiceQuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicequotasServiceQuota `json:"items"`
}

// A ServicequotasServiceQuotaSpec defines the desired state of a ServicequotasServiceQuota
type ServicequotasServiceQuotaSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ServicequotasServiceQuotaParameters `json:",inline"`
}

// A ServicequotasServiceQuotaParameters defines the desired state of a ServicequotasServiceQuota
type ServicequotasServiceQuotaParameters struct {
	QuotaCode   string `json:"quota_code"`
	ServiceCode string `json:"service_code"`
	Value       int    `json:"value"`
}

// A ServicequotasServiceQuotaStatus defines the observed state of a ServicequotasServiceQuota
type ServicequotasServiceQuotaStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ServicequotasServiceQuotaObservation `json:",inline"`
}

// A ServicequotasServiceQuotaObservation records the observed state of a ServicequotasServiceQuota
type ServicequotasServiceQuotaObservation struct {
	QuotaName     string `json:"quota_name"`
	RequestStatus string `json:"request_status"`
	ServiceName   string `json:"service_name"`
	DefaultValue  int    `json:"default_value"`
	Arn           string `json:"arn"`
	Id            string `json:"id"`
	RequestId     string `json:"request_id"`
	Adjustable    bool   `json:"adjustable"`
}