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

// ServicecatalogPortfolio is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ServicecatalogPortfolio struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServicecatalogPortfolioSpec   `json:"spec"`
	Status ServicecatalogPortfolioStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogPortfolio contains a list of ServicecatalogPortfolioList
type ServicecatalogPortfolioList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicecatalogPortfolio `json:"items"`
}

// A ServicecatalogPortfolioSpec defines the desired state of a ServicecatalogPortfolio
type ServicecatalogPortfolioSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ServicecatalogPortfolioParameters `json:",inline"`
}

// A ServicecatalogPortfolioParameters defines the desired state of a ServicecatalogPortfolio
type ServicecatalogPortfolioParameters struct {
	Name         string            `json:"name"`
	ProviderName string            `json:"provider_name"`
	Tags         map[string]string `json:"tags"`
	Timeouts     []Timeouts        `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A ServicecatalogPortfolioStatus defines the observed state of a ServicecatalogPortfolio
type ServicecatalogPortfolioStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ServicecatalogPortfolioObservation `json:",inline"`
}

// A ServicecatalogPortfolioObservation records the observed state of a ServicecatalogPortfolio
type ServicecatalogPortfolioObservation struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Arn         string `json:"arn"`
	CreatedTime string `json:"created_time"`
}