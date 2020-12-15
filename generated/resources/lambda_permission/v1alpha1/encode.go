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
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*LambdaPermission)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a LambdaPermission.")
	}
	return EncodeLambdaPermission(*r), nil
}

func EncodeLambdaPermission(r LambdaPermission) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaPermission_Action(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_Principal(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_StatementId(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_Qualifier(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_SourceAccount(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_SourceArn(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_StatementIdPrefix(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_EventSourceToken(r.Spec.ForProvider, ctyVal)
	EncodeLambdaPermission_FunctionName(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeLambdaPermission_Action(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["action"] = cty.StringVal(p.Action)
}

func EncodeLambdaPermission_Principal(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["principal"] = cty.StringVal(p.Principal)
}

func EncodeLambdaPermission_StatementId(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["statement_id"] = cty.StringVal(p.StatementId)
}

func EncodeLambdaPermission_Qualifier(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["qualifier"] = cty.StringVal(p.Qualifier)
}

func EncodeLambdaPermission_SourceAccount(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["source_account"] = cty.StringVal(p.SourceAccount)
}

func EncodeLambdaPermission_SourceArn(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["source_arn"] = cty.StringVal(p.SourceArn)
}

func EncodeLambdaPermission_StatementIdPrefix(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["statement_id_prefix"] = cty.StringVal(p.StatementIdPrefix)
}

func EncodeLambdaPermission_EventSourceToken(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["event_source_token"] = cty.StringVal(p.EventSourceToken)
}

func EncodeLambdaPermission_FunctionName(p LambdaPermissionParameters, vals map[string]cty.Value) {
	vals["function_name"] = cty.StringVal(p.FunctionName)
}