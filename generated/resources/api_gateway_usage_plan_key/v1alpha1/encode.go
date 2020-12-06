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

func EncodeApiGatewayUsagePlanKey(r ApiGatewayUsagePlanKey) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayUsagePlanKey_UsagePlanId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayUsagePlanKey_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayUsagePlanKey_KeyId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayUsagePlanKey_KeyType(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayUsagePlanKey_Name(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayUsagePlanKey_Value(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayUsagePlanKey_UsagePlanId(p ApiGatewayUsagePlanKeyParameters, vals map[string]cty.Value) {
	vals["usage_plan_id"] = cty.StringVal(p.UsagePlanId)
}

func EncodeApiGatewayUsagePlanKey_Id(p ApiGatewayUsagePlanKeyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayUsagePlanKey_KeyId(p ApiGatewayUsagePlanKeyParameters, vals map[string]cty.Value) {
	vals["key_id"] = cty.StringVal(p.KeyId)
}

func EncodeApiGatewayUsagePlanKey_KeyType(p ApiGatewayUsagePlanKeyParameters, vals map[string]cty.Value) {
	vals["key_type"] = cty.StringVal(p.KeyType)
}

func EncodeApiGatewayUsagePlanKey_Name(p ApiGatewayUsagePlanKeyObservation, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeApiGatewayUsagePlanKey_Value(p ApiGatewayUsagePlanKeyObservation, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}