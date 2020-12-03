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

package v1alpha1func EncodeBudgetsBudget(r BudgetsBudget) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeBudgetsBudget_BudgetType(r.Spec.ForProvider, ctyVal)
	EncodeBudgetsBudget_LimitUnit(r.Spec.ForProvider, ctyVal)
	EncodeBudgetsBudget_Name(r.Spec.ForProvider, ctyVal)
	EncodeBudgetsBudget_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeBudgetsBudget_TimePeriodEnd(r.Spec.ForProvider, ctyVal)
	EncodeBudgetsBudget_TimeUnit(r.Spec.ForProvider, ctyVal)
	EncodeBudgetsBudget_AccountId(r.Spec.ForProvider, ctyVal)
	EncodeBudgetsBudget_CostFilters(r.Spec.ForProvider, ctyVal)
	EncodeBudgetsBudget_Id(r.Spec.ForProvider, ctyVal)
	EncodeBudgetsBudget_LimitAmount(r.Spec.ForProvider, ctyVal)
	EncodeBudgetsBudget_TimePeriodStart(r.Spec.ForProvider, ctyVal)
	EncodeBudgetsBudget_CostTypes(r.Spec.ForProvider.CostTypes, ctyVal)
	EncodeBudgetsBudget_Notification(r.Spec.ForProvider.Notification, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeBudgetsBudget_BudgetType(p *BudgetsBudgetParameters, vals map[string]cty.Value) {
	vals["budget_type"] = cty.StringVal(p.BudgetType)
}

func EncodeBudgetsBudget_LimitUnit(p *BudgetsBudgetParameters, vals map[string]cty.Value) {
	vals["limit_unit"] = cty.StringVal(p.LimitUnit)
}

func EncodeBudgetsBudget_Name(p *BudgetsBudgetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeBudgetsBudget_NamePrefix(p *BudgetsBudgetParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeBudgetsBudget_TimePeriodEnd(p *BudgetsBudgetParameters, vals map[string]cty.Value) {
	vals["time_period_end"] = cty.StringVal(p.TimePeriodEnd)
}

func EncodeBudgetsBudget_TimeUnit(p *BudgetsBudgetParameters, vals map[string]cty.Value) {
	vals["time_unit"] = cty.StringVal(p.TimeUnit)
}

func EncodeBudgetsBudget_AccountId(p *BudgetsBudgetParameters, vals map[string]cty.Value) {
	vals["account_id"] = cty.StringVal(p.AccountId)
}

func EncodeBudgetsBudget_CostFilters(p *BudgetsBudgetParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.CostFilters {
		mVals[key] = cty.StringVal(value)
	}
	vals["cost_filters"] = cty.MapVal(mVals)
}

func EncodeBudgetsBudget_Id(p *BudgetsBudgetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeBudgetsBudget_LimitAmount(p *BudgetsBudgetParameters, vals map[string]cty.Value) {
	vals["limit_amount"] = cty.StringVal(p.LimitAmount)
}

func EncodeBudgetsBudget_TimePeriodStart(p *BudgetsBudgetParameters, vals map[string]cty.Value) {
	vals["time_period_start"] = cty.StringVal(p.TimePeriodStart)
}

func EncodeBudgetsBudget_CostTypes(p *CostTypes, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CostTypes {
		ctyVal = make(map[string]cty.Value)
		EncodeBudgetsBudget_CostTypes_IncludeDiscount(v, ctyVal)
		EncodeBudgetsBudget_CostTypes_IncludeRecurring(v, ctyVal)
		EncodeBudgetsBudget_CostTypes_IncludeRefund(v, ctyVal)
		EncodeBudgetsBudget_CostTypes_IncludeSupport(v, ctyVal)
		EncodeBudgetsBudget_CostTypes_IncludeTax(v, ctyVal)
		EncodeBudgetsBudget_CostTypes_IncludeUpfront(v, ctyVal)
		EncodeBudgetsBudget_CostTypes_IncludeCredit(v, ctyVal)
		EncodeBudgetsBudget_CostTypes_IncludeOtherSubscription(v, ctyVal)
		EncodeBudgetsBudget_CostTypes_IncludeSubscription(v, ctyVal)
		EncodeBudgetsBudget_CostTypes_UseAmortized(v, ctyVal)
		EncodeBudgetsBudget_CostTypes_UseBlended(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cost_types"] = cty.ListVal(valsForCollection)
}

func EncodeBudgetsBudget_CostTypes_IncludeDiscount(p *CostTypes, vals map[string]cty.Value) {
	vals["include_discount"] = cty.BoolVal(p.IncludeDiscount)
}

func EncodeBudgetsBudget_CostTypes_IncludeRecurring(p *CostTypes, vals map[string]cty.Value) {
	vals["include_recurring"] = cty.BoolVal(p.IncludeRecurring)
}

func EncodeBudgetsBudget_CostTypes_IncludeRefund(p *CostTypes, vals map[string]cty.Value) {
	vals["include_refund"] = cty.BoolVal(p.IncludeRefund)
}

func EncodeBudgetsBudget_CostTypes_IncludeSupport(p *CostTypes, vals map[string]cty.Value) {
	vals["include_support"] = cty.BoolVal(p.IncludeSupport)
}

func EncodeBudgetsBudget_CostTypes_IncludeTax(p *CostTypes, vals map[string]cty.Value) {
	vals["include_tax"] = cty.BoolVal(p.IncludeTax)
}

func EncodeBudgetsBudget_CostTypes_IncludeUpfront(p *CostTypes, vals map[string]cty.Value) {
	vals["include_upfront"] = cty.BoolVal(p.IncludeUpfront)
}

func EncodeBudgetsBudget_CostTypes_IncludeCredit(p *CostTypes, vals map[string]cty.Value) {
	vals["include_credit"] = cty.BoolVal(p.IncludeCredit)
}

func EncodeBudgetsBudget_CostTypes_IncludeOtherSubscription(p *CostTypes, vals map[string]cty.Value) {
	vals["include_other_subscription"] = cty.BoolVal(p.IncludeOtherSubscription)
}

func EncodeBudgetsBudget_CostTypes_IncludeSubscription(p *CostTypes, vals map[string]cty.Value) {
	vals["include_subscription"] = cty.BoolVal(p.IncludeSubscription)
}

func EncodeBudgetsBudget_CostTypes_UseAmortized(p *CostTypes, vals map[string]cty.Value) {
	vals["use_amortized"] = cty.BoolVal(p.UseAmortized)
}

func EncodeBudgetsBudget_CostTypes_UseBlended(p *CostTypes, vals map[string]cty.Value) {
	vals["use_blended"] = cty.BoolVal(p.UseBlended)
}

func EncodeBudgetsBudget_Notification(p *Notification, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Notification {
		ctyVal = make(map[string]cty.Value)
		EncodeBudgetsBudget_Notification_NotificationType(v, ctyVal)
		EncodeBudgetsBudget_Notification_SubscriberEmailAddresses(v, ctyVal)
		EncodeBudgetsBudget_Notification_SubscriberSnsTopicArns(v, ctyVal)
		EncodeBudgetsBudget_Notification_Threshold(v, ctyVal)
		EncodeBudgetsBudget_Notification_ThresholdType(v, ctyVal)
		EncodeBudgetsBudget_Notification_ComparisonOperator(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["notification"] = cty.SetVal(valsForCollection)
}

func EncodeBudgetsBudget_Notification_NotificationType(p *Notification, vals map[string]cty.Value) {
	vals["notification_type"] = cty.StringVal(p.NotificationType)
}

func EncodeBudgetsBudget_Notification_SubscriberEmailAddresses(p *Notification, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubscriberEmailAddresses {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subscriber_email_addresses"] = cty.SetVal(colVals)
}

func EncodeBudgetsBudget_Notification_SubscriberSnsTopicArns(p *Notification, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubscriberSnsTopicArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subscriber_sns_topic_arns"] = cty.SetVal(colVals)
}

func EncodeBudgetsBudget_Notification_Threshold(p *Notification, vals map[string]cty.Value) {
	vals["threshold"] = cty.IntVal(p.Threshold)
}

func EncodeBudgetsBudget_Notification_ThresholdType(p *Notification, vals map[string]cty.Value) {
	vals["threshold_type"] = cty.StringVal(p.ThresholdType)
}

func EncodeBudgetsBudget_Notification_ComparisonOperator(p *Notification, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}