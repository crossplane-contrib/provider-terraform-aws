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

func EncodeApigatewayv2Model(r Apigatewayv2Model) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2Model_Schema(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Model_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Model_ContentType(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Model_Description(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Model_Id(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Model_Name(r.Spec.ForProvider, ctyVal)

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