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

func EncodeConfigAggregateAuthorization(r ConfigAggregateAuthorization) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeConfigAggregateAuthorization_Id(r.Spec.ForProvider, ctyVal)
	EncodeConfigAggregateAuthorization_Region(r.Spec.ForProvider, ctyVal)
	EncodeConfigAggregateAuthorization_Tags(r.Spec.ForProvider, ctyVal)
	EncodeConfigAggregateAuthorization_AccountId(r.Spec.ForProvider, ctyVal)
	EncodeConfigAggregateAuthorization_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeConfigAggregateAuthorization_Id(p ConfigAggregateAuthorizationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeConfigAggregateAuthorization_Region(p ConfigAggregateAuthorizationParameters, vals map[string]cty.Value) {
	vals["region"] = cty.StringVal(p.Region)
}

func EncodeConfigAggregateAuthorization_Tags(p ConfigAggregateAuthorizationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeConfigAggregateAuthorization_AccountId(p ConfigAggregateAuthorizationParameters, vals map[string]cty.Value) {
	vals["account_id"] = cty.StringVal(p.AccountId)
}

func EncodeConfigAggregateAuthorization_Arn(p ConfigAggregateAuthorizationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}