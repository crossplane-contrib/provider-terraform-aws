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

func EncodeCodestarnotificationsNotificationRule(r CodestarnotificationsNotificationRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodestarnotificationsNotificationRule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCodestarnotificationsNotificationRule_DetailType(r.Spec.ForProvider, ctyVal)
	EncodeCodestarnotificationsNotificationRule_EventTypeIds(r.Spec.ForProvider, ctyVal)
	EncodeCodestarnotificationsNotificationRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodestarnotificationsNotificationRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeCodestarnotificationsNotificationRule_Resource(r.Spec.ForProvider, ctyVal)
	EncodeCodestarnotificationsNotificationRule_Status(r.Spec.ForProvider, ctyVal)
	EncodeCodestarnotificationsNotificationRule_Target(r.Spec.ForProvider.Target, ctyVal)
	EncodeCodestarnotificationsNotificationRule_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCodestarnotificationsNotificationRule_Tags(p CodestarnotificationsNotificationRuleParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCodestarnotificationsNotificationRule_DetailType(p CodestarnotificationsNotificationRuleParameters, vals map[string]cty.Value) {
	vals["detail_type"] = cty.StringVal(p.DetailType)
}

func EncodeCodestarnotificationsNotificationRule_EventTypeIds(p CodestarnotificationsNotificationRuleParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.EventTypeIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["event_type_ids"] = cty.SetVal(colVals)
}

func EncodeCodestarnotificationsNotificationRule_Id(p CodestarnotificationsNotificationRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodestarnotificationsNotificationRule_Name(p CodestarnotificationsNotificationRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodestarnotificationsNotificationRule_Resource(p CodestarnotificationsNotificationRuleParameters, vals map[string]cty.Value) {
	vals["resource"] = cty.StringVal(p.Resource)
}

func EncodeCodestarnotificationsNotificationRule_Status(p CodestarnotificationsNotificationRuleParameters, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeCodestarnotificationsNotificationRule_Target(p []Target, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeCodestarnotificationsNotificationRule_Target_Address(v, ctyVal)
		EncodeCodestarnotificationsNotificationRule_Target_Status(v, ctyVal)
		EncodeCodestarnotificationsNotificationRule_Target_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["target"] = cty.SetVal(valsForCollection)
}

func EncodeCodestarnotificationsNotificationRule_Target_Address(p Target, vals map[string]cty.Value) {
	vals["address"] = cty.StringVal(p.Address)
}

func EncodeCodestarnotificationsNotificationRule_Target_Status(p Target, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeCodestarnotificationsNotificationRule_Target_Type(p Target, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodestarnotificationsNotificationRule_Arn(p CodestarnotificationsNotificationRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}