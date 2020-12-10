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

// BudgetsBudget is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type BudgetsBudget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BudgetsBudgetSpec   `json:"spec"`
	Status BudgetsBudgetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetsBudget contains a list of BudgetsBudgetList
type BudgetsBudgetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BudgetsBudget `json:"items"`
}

// A BudgetsBudgetSpec defines the desired state of a BudgetsBudget
type BudgetsBudgetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  BudgetsBudgetParameters `json:",inline"`
}

// A BudgetsBudgetParameters defines the desired state of a BudgetsBudget
type BudgetsBudgetParameters struct {
	Id              string            `json:"id"`
	LimitUnit       string            `json:"limit_unit"`
	Name            string            `json:"name"`
	NamePrefix      string            `json:"name_prefix"`
	TimePeriodEnd   string            `json:"time_period_end"`
	BudgetType      string            `json:"budget_type"`
	CostFilters     map[string]string `json:"cost_filters"`
	LimitAmount     string            `json:"limit_amount"`
	TimePeriodStart string            `json:"time_period_start"`
	TimeUnit        string            `json:"time_unit"`
	AccountId       string            `json:"account_id"`
	CostTypes       CostTypes         `json:"cost_types"`
	Notification    Notification      `json:"notification"`
}

type CostTypes struct {
	IncludeRefund            bool `json:"include_refund"`
	IncludeSupport           bool `json:"include_support"`
	IncludeTax               bool `json:"include_tax"`
	UseBlended               bool `json:"use_blended"`
	UseAmortized             bool `json:"use_amortized"`
	IncludeCredit            bool `json:"include_credit"`
	IncludeDiscount          bool `json:"include_discount"`
	IncludeOtherSubscription bool `json:"include_other_subscription"`
	IncludeRecurring         bool `json:"include_recurring"`
	IncludeSubscription      bool `json:"include_subscription"`
	IncludeUpfront           bool `json:"include_upfront"`
}

type Notification struct {
	Threshold                int64    `json:"threshold"`
	ThresholdType            string   `json:"threshold_type"`
	ComparisonOperator       string   `json:"comparison_operator"`
	NotificationType         string   `json:"notification_type"`
	SubscriberEmailAddresses []string `json:"subscriber_email_addresses"`
	SubscriberSnsTopicArns   []string `json:"subscriber_sns_topic_arns"`
}

// A BudgetsBudgetStatus defines the observed state of a BudgetsBudget
type BudgetsBudgetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     BudgetsBudgetObservation `json:",inline"`
}

// A BudgetsBudgetObservation records the observed state of a BudgetsBudget
type BudgetsBudgetObservation struct{}