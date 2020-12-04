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
	"github.com/zclconf/go-cty/cty"
)

func EncodeRedshiftEventSubscription(r RedshiftEventSubscription) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRedshiftEventSubscription_SourceIds(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftEventSubscription_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftEventSubscription_Name(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftEventSubscription_Severity(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftEventSubscription_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftEventSubscription_EventCategories(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftEventSubscription_Id(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftEventSubscription_SnsTopicArn(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftEventSubscription_SourceType(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftEventSubscription_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeRedshiftEventSubscription_Status(r.Status.AtProvider, ctyVal)
	EncodeRedshiftEventSubscription_Arn(r.Status.AtProvider, ctyVal)
	EncodeRedshiftEventSubscription_CustomerAwsId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeRedshiftEventSubscription_SourceIds(p RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SourceIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["source_ids"] = cty.SetVal(colVals)
}

func EncodeRedshiftEventSubscription_Tags(p RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRedshiftEventSubscription_Name(p RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRedshiftEventSubscription_Severity(p RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["severity"] = cty.StringVal(p.Severity)
}

func EncodeRedshiftEventSubscription_Enabled(p RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeRedshiftEventSubscription_EventCategories(p RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.EventCategories {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["event_categories"] = cty.SetVal(colVals)
}

func EncodeRedshiftEventSubscription_Id(p RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRedshiftEventSubscription_SnsTopicArn(p RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["sns_topic_arn"] = cty.StringVal(p.SnsTopicArn)
}

func EncodeRedshiftEventSubscription_SourceType(p RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["source_type"] = cty.StringVal(p.SourceType)
}

func EncodeRedshiftEventSubscription_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeRedshiftEventSubscription_Timeouts_Delete(p, ctyVal)
	EncodeRedshiftEventSubscription_Timeouts_Update(p, ctyVal)
	EncodeRedshiftEventSubscription_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeRedshiftEventSubscription_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeRedshiftEventSubscription_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeRedshiftEventSubscription_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeRedshiftEventSubscription_Status(p RedshiftEventSubscriptionObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeRedshiftEventSubscription_Arn(p RedshiftEventSubscriptionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeRedshiftEventSubscription_CustomerAwsId(p RedshiftEventSubscriptionObservation, vals map[string]cty.Value) {
	vals["customer_aws_id"] = cty.StringVal(p.CustomerAwsId)
}