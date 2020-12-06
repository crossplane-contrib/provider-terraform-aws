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

func EncodeLambdaPermission(r LambdaPermission) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaPermission_StatementIdPrefix(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_FunctionName(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_Id(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_SourceArn(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_Qualifier(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_SourceAccount(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_StatementId(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_Action(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_EventSourceToken(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_Principal(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeLambdaPermission_StatementIdPrefix(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["statement_id_prefix"] = cty.StringVal(p.StatementIdPrefix)
}

func EncodeLambdaPermission_FunctionName(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["function_name"] = cty.StringVal(p.FunctionName)
}

func EncodeLambdaPermission_Id(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLambdaPermission_SourceArn(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["source_arn"] = cty.StringVal(p.SourceArn)
}

func EncodeLambdaPermission_Qualifier(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["qualifier"] = cty.StringVal(p.Qualifier)
}

func EncodeLambdaPermission_SourceAccount(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["source_account"] = cty.StringVal(p.SourceAccount)
}

func EncodeLambdaPermission_StatementId(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["statement_id"] = cty.StringVal(p.StatementId)
}

func EncodeLambdaPermission_Action(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["action"] = cty.StringVal(p.Action)
}

func EncodeLambdaPermission_EventSourceToken(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["event_source_token"] = cty.StringVal(p.EventSourceToken)
}

func EncodeLambdaPermission_Principal(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["principal"] = cty.StringVal(p.Principal)
}