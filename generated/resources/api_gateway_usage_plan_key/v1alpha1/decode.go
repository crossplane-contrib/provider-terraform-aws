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
	r, ok := mr.(*ApiGatewayUsagePlanKey)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApiGatewayUsagePlanKey(r, ctyValue)
}

func DecodeApiGatewayUsagePlanKey(prev *ApiGatewayUsagePlanKey, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApiGatewayUsagePlanKey_KeyType(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayUsagePlanKey_UsagePlanId(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayUsagePlanKey_Id(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayUsagePlanKey_KeyId(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayUsagePlanKey_Name(&new.Status.AtProvider, valMap)
	DecodeApiGatewayUsagePlanKey_Value(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeApiGatewayUsagePlanKey_KeyType(p *ApiGatewayUsagePlanKeyParameters, vals map[string]cty.Value) {
	p.KeyType = ctwhy.ValueAsString(vals["key_type"])
}

func DecodeApiGatewayUsagePlanKey_UsagePlanId(p *ApiGatewayUsagePlanKeyParameters, vals map[string]cty.Value) {
	p.UsagePlanId = ctwhy.ValueAsString(vals["usage_plan_id"])
}

func DecodeApiGatewayUsagePlanKey_Id(p *ApiGatewayUsagePlanKeyParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeApiGatewayUsagePlanKey_KeyId(p *ApiGatewayUsagePlanKeyParameters, vals map[string]cty.Value) {
	p.KeyId = ctwhy.ValueAsString(vals["key_id"])
}

func DecodeApiGatewayUsagePlanKey_Name(p *ApiGatewayUsagePlanKeyObservation, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeApiGatewayUsagePlanKey_Value(p *ApiGatewayUsagePlanKeyObservation, vals map[string]cty.Value) {
	p.Value = ctwhy.ValueAsString(vals["value"])
}