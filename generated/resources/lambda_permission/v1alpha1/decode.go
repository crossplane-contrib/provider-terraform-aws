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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*LambdaPermission)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeLambdaPermission(r, ctyValue)
}

func DecodeLambdaPermission(prev *LambdaPermission, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeLambdaPermission_SourceAccount(&new.Spec.ForProvider, valMap)
	DecodeLambdaPermission_StatementId(&new.Spec.ForProvider, valMap)
	DecodeLambdaPermission_Action(&new.Spec.ForProvider, valMap)
	DecodeLambdaPermission_EventSourceToken(&new.Spec.ForProvider, valMap)
	DecodeLambdaPermission_FunctionName(&new.Spec.ForProvider, valMap)
	DecodeLambdaPermission_Id(&new.Spec.ForProvider, valMap)
	DecodeLambdaPermission_Principal(&new.Spec.ForProvider, valMap)
	DecodeLambdaPermission_Qualifier(&new.Spec.ForProvider, valMap)
	DecodeLambdaPermission_SourceArn(&new.Spec.ForProvider, valMap)
	DecodeLambdaPermission_StatementIdPrefix(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeLambdaPermission_SourceAccount(p *LambdaPermissionParameters, vals map[string]cty.Value) {
	p.SourceAccount = ctwhy.ValueAsString(vals["source_account"])
}

func DecodeLambdaPermission_StatementId(p *LambdaPermissionParameters, vals map[string]cty.Value) {
	p.StatementId = ctwhy.ValueAsString(vals["statement_id"])
}

func DecodeLambdaPermission_Action(p *LambdaPermissionParameters, vals map[string]cty.Value) {
	p.Action = ctwhy.ValueAsString(vals["action"])
}

func DecodeLambdaPermission_EventSourceToken(p *LambdaPermissionParameters, vals map[string]cty.Value) {
	p.EventSourceToken = ctwhy.ValueAsString(vals["event_source_token"])
}

func DecodeLambdaPermission_FunctionName(p *LambdaPermissionParameters, vals map[string]cty.Value) {
	p.FunctionName = ctwhy.ValueAsString(vals["function_name"])
}

func DecodeLambdaPermission_Id(p *LambdaPermissionParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeLambdaPermission_Principal(p *LambdaPermissionParameters, vals map[string]cty.Value) {
	p.Principal = ctwhy.ValueAsString(vals["principal"])
}

func DecodeLambdaPermission_Qualifier(p *LambdaPermissionParameters, vals map[string]cty.Value) {
	p.Qualifier = ctwhy.ValueAsString(vals["qualifier"])
}

func DecodeLambdaPermission_SourceArn(p *LambdaPermissionParameters, vals map[string]cty.Value) {
	p.SourceArn = ctwhy.ValueAsString(vals["source_arn"])
}

func DecodeLambdaPermission_StatementIdPrefix(p *LambdaPermissionParameters, vals map[string]cty.Value) {
	p.StatementIdPrefix = ctwhy.ValueAsString(vals["statement_id_prefix"])
}