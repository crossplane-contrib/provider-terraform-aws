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

func EncodeCloudwatchEventPermission(r CloudwatchEventPermission) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchEventPermission_Action(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventPermission_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventPermission_Principal(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventPermission_StatementId(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventPermission_Condition(r.Spec.ForProvider.Condition, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeCloudwatchEventPermission_Action(p CloudwatchEventPermissionParameters, vals map[string]cty.Value) {
	vals["action"] = cty.StringVal(p.Action)
}

func EncodeCloudwatchEventPermission_Id(p CloudwatchEventPermissionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudwatchEventPermission_Principal(p CloudwatchEventPermissionParameters, vals map[string]cty.Value) {
	vals["principal"] = cty.StringVal(p.Principal)
}

func EncodeCloudwatchEventPermission_StatementId(p CloudwatchEventPermissionParameters, vals map[string]cty.Value) {
	vals["statement_id"] = cty.StringVal(p.StatementId)
}

func EncodeCloudwatchEventPermission_Condition(p Condition, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchEventPermission_Condition_Key(p, ctyVal)
	EncodeCloudwatchEventPermission_Condition_Type(p, ctyVal)
	EncodeCloudwatchEventPermission_Condition_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["condition"] = cty.ListVal(valsForCollection)
}

func EncodeCloudwatchEventPermission_Condition_Key(p Condition, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeCloudwatchEventPermission_Condition_Type(p Condition, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCloudwatchEventPermission_Condition_Value(p Condition, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}