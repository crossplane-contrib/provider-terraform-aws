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

func EncodeDmsEventSubscription(r DmsEventSubscription) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDmsEventSubscription_Name(r.Spec.ForProvider, ctyVal)
	EncodeDmsEventSubscription_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeDmsEventSubscription_EventCategories(r.Spec.ForProvider, ctyVal)
	EncodeDmsEventSubscription_Id(r.Spec.ForProvider, ctyVal)
	EncodeDmsEventSubscription_SnsTopicArn(r.Spec.ForProvider, ctyVal)
	EncodeDmsEventSubscription_SourceIds(r.Spec.ForProvider, ctyVal)
	EncodeDmsEventSubscription_SourceType(r.Spec.ForProvider, ctyVal)
	EncodeDmsEventSubscription_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDmsEventSubscription_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDmsEventSubscription_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDmsEventSubscription_Name(p DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDmsEventSubscription_Enabled(p DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeDmsEventSubscription_EventCategories(p DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.EventCategories {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["event_categories"] = cty.SetVal(colVals)
}

func EncodeDmsEventSubscription_Id(p DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDmsEventSubscription_SnsTopicArn(p DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["sns_topic_arn"] = cty.StringVal(p.SnsTopicArn)
}

func EncodeDmsEventSubscription_SourceIds(p DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SourceIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["source_ids"] = cty.SetVal(colVals)
}

func EncodeDmsEventSubscription_SourceType(p DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["source_type"] = cty.StringVal(p.SourceType)
}

func EncodeDmsEventSubscription_Tags(p DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDmsEventSubscription_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDmsEventSubscription_Timeouts_Create(p, ctyVal)
	EncodeDmsEventSubscription_Timeouts_Delete(p, ctyVal)
	EncodeDmsEventSubscription_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDmsEventSubscription_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDmsEventSubscription_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDmsEventSubscription_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDmsEventSubscription_Arn(p DmsEventSubscriptionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}