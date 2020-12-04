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

func EncodeApiGatewayDocumentationPart(r ApiGatewayDocumentationPart) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayDocumentationPart_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDocumentationPart_Properties(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDocumentationPart_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDocumentationPart_Location(r.Spec.ForProvider.Location, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayDocumentationPart_Id(p ApiGatewayDocumentationPartParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayDocumentationPart_Properties(p ApiGatewayDocumentationPartParameters, vals map[string]cty.Value) {
	vals["properties"] = cty.StringVal(p.Properties)
}

func EncodeApiGatewayDocumentationPart_RestApiId(p ApiGatewayDocumentationPartParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayDocumentationPart_Location(p Location, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayDocumentationPart_Location_Method(p, ctyVal)
	EncodeApiGatewayDocumentationPart_Location_Name(p, ctyVal)
	EncodeApiGatewayDocumentationPart_Location_Path(p, ctyVal)
	EncodeApiGatewayDocumentationPart_Location_StatusCode(p, ctyVal)
	EncodeApiGatewayDocumentationPart_Location_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["location"] = cty.ListVal(valsForCollection)
}

func EncodeApiGatewayDocumentationPart_Location_Method(p Location, vals map[string]cty.Value) {
	vals["method"] = cty.StringVal(p.Method)
}

func EncodeApiGatewayDocumentationPart_Location_Name(p Location, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeApiGatewayDocumentationPart_Location_Path(p Location, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeApiGatewayDocumentationPart_Location_StatusCode(p Location, vals map[string]cty.Value) {
	vals["status_code"] = cty.StringVal(p.StatusCode)
}

func EncodeApiGatewayDocumentationPart_Location_Type(p Location, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}