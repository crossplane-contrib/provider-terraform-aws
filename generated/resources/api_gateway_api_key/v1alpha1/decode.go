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
	r, ok := mr.(*ApiGatewayApiKey)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApiGatewayApiKey(r, ctyValue)
}

func DecodeApiGatewayApiKey(prev *ApiGatewayApiKey, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApiGatewayApiKey_Tags(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayApiKey_Value(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayApiKey_Enabled(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayApiKey_Description(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayApiKey_Id(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayApiKey_Name(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayApiKey_Arn(&new.Status.AtProvider, valMap)
	DecodeApiGatewayApiKey_CreatedDate(&new.Status.AtProvider, valMap)
	DecodeApiGatewayApiKey_LastUpdatedDate(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeApiGatewayApiKey_Tags(p *ApiGatewayApiKeyParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeApiGatewayApiKey_Value(p *ApiGatewayApiKeyParameters, vals map[string]cty.Value) {
	p.Value = ctwhy.ValueAsString(vals["value"])
}

func DecodeApiGatewayApiKey_Enabled(p *ApiGatewayApiKeyParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

func DecodeApiGatewayApiKey_Description(p *ApiGatewayApiKeyParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeApiGatewayApiKey_Id(p *ApiGatewayApiKeyParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeApiGatewayApiKey_Name(p *ApiGatewayApiKeyParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeApiGatewayApiKey_Arn(p *ApiGatewayApiKeyObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeApiGatewayApiKey_CreatedDate(p *ApiGatewayApiKeyObservation, vals map[string]cty.Value) {
	p.CreatedDate = ctwhy.ValueAsString(vals["created_date"])
}

func DecodeApiGatewayApiKey_LastUpdatedDate(p *ApiGatewayApiKeyObservation, vals map[string]cty.Value) {
	p.LastUpdatedDate = ctwhy.ValueAsString(vals["last_updated_date"])
}