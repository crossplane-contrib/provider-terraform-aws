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
	r, ok := mr.(*ApiGatewayModel)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApiGatewayModel(r, ctyValue)
}

func DecodeApiGatewayModel(prev *ApiGatewayModel, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApiGatewayModel_ContentType(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayModel_Description(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayModel_Id(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayModel_Name(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayModel_RestApiId(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayModel_Schema(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayModel_ContentType(p *ApiGatewayModelParameters, vals map[string]cty.Value) {
	p.ContentType = ctwhy.ValueAsString(vals["content_type"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayModel_Description(p *ApiGatewayModelParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayModel_Id(p *ApiGatewayModelParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayModel_Name(p *ApiGatewayModelParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayModel_RestApiId(p *ApiGatewayModelParameters, vals map[string]cty.Value) {
	p.RestApiId = ctwhy.ValueAsString(vals["rest_api_id"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayModel_Schema(p *ApiGatewayModelParameters, vals map[string]cty.Value) {
	p.Schema = ctwhy.ValueAsString(vals["schema"])
}