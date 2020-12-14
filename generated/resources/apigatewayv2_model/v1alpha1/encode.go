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
	r, ok := mr.(*Apigatewayv2Model)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Apigatewayv2Model.")
	}
	return EncodeApigatewayv2Model(*r), nil
}

func EncodeApigatewayv2Model(r Apigatewayv2Model) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2Model_Schema(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Model_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Model_ContentType(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Model_Description(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Model_Id(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Model_Name(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeApigatewayv2Model_Schema(p Apigatewayv2ModelParameters, vals map[string]cty.Value) {
	vals["schema"] = cty.StringVal(p.Schema)
}

func EncodeApigatewayv2Model_ApiId(p Apigatewayv2ModelParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeApigatewayv2Model_ContentType(p Apigatewayv2ModelParameters, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeApigatewayv2Model_Description(p Apigatewayv2ModelParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeApigatewayv2Model_Id(p Apigatewayv2ModelParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApigatewayv2Model_Name(p Apigatewayv2ModelParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}