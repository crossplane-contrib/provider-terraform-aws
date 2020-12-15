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
	r, ok := mr.(*Apigatewayv2Model)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApigatewayv2Model(r, ctyValue)
}

func DecodeApigatewayv2Model(prev *Apigatewayv2Model, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApigatewayv2Model_Name(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Model_Schema(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Model_ApiId(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Model_ContentType(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Model_Description(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Model_Id(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Model_Name(p *Apigatewayv2ModelParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Model_Schema(p *Apigatewayv2ModelParameters, vals map[string]cty.Value) {
	p.Schema = ctwhy.ValueAsString(vals["schema"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Model_ApiId(p *Apigatewayv2ModelParameters, vals map[string]cty.Value) {
	p.ApiId = ctwhy.ValueAsString(vals["api_id"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Model_ContentType(p *Apigatewayv2ModelParameters, vals map[string]cty.Value) {
	p.ContentType = ctwhy.ValueAsString(vals["content_type"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Model_Description(p *Apigatewayv2ModelParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Model_Id(p *Apigatewayv2ModelParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}