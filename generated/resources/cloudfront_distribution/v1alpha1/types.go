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

// CloudfrontDistribution is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudfrontDistribution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudfrontDistributionSpec   `json:"spec"`
	Status CloudfrontDistributionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontDistribution contains a list of CloudfrontDistributionList
type CloudfrontDistributionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudfrontDistribution `json:"items"`
}

// A CloudfrontDistributionSpec defines the desired state of a CloudfrontDistribution
type CloudfrontDistributionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudfrontDistributionParameters `json:",inline"`
}

// A CloudfrontDistributionParameters defines the desired state of a CloudfrontDistribution
type CloudfrontDistributionParameters struct {
	Enabled           bool     `json:"enabled"`
	Aliases           []string `json:"aliases"`
	HttpVersion       string   `json:"http_version"`
	WaitForDeployment bool     `json:"wait_for_deployment"`
	IsIpv6Enabled     bool     `json:"is_ipv6_enabled"`
	RetainOnDelete    bool     `json:"retain_on_delete"`
	Comment           string   `json:"comment"`
	PriceClass        string   `json:"price_class"`
	WebAclId          string   `json:"web_acl_id"`
	DefaultRootObject string   `json:"default_root_object"`
}

// A CloudfrontDistributionStatus defines the observed state of a CloudfrontDistribution
type CloudfrontDistributionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudfrontDistributionObservation `json:",inline"`
}

// A CloudfrontDistributionObservation records the observed state of a CloudfrontDistribution
type CloudfrontDistributionObservation struct {
	Arn                         string `json:"arn"`
	HostedZoneId                string `json:"hosted_zone_id"`
	Id                          string `json:"id"`
	InProgressValidationBatches int    `json:"in_progress_validation_batches"`
	DomainName                  string `json:"domain_name"`
	Etag                        string `json:"etag"`
	Status                      string `json:"status"`
	CallerReference             string `json:"caller_reference"`
	LastModifiedTime            string `json:"last_modified_time"`
}