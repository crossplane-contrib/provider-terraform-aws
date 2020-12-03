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

package v1alpha1func EncodeApiGatewayModel(r ApiGatewayModel) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeApiGatewayModel_ContentType(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayModel_Description(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayModel_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayModel_Name(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayModel_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayModel_Schema(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeApiGatewayModel_ContentType(p *ApiGatewayModelParameters, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeApiGatewayModel_Description(p *ApiGatewayModelParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeApiGatewayModel_Id(p *ApiGatewayModelParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayModel_Name(p *ApiGatewayModelParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeApiGatewayModel_RestApiId(p *ApiGatewayModelParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayModel_Schema(p *ApiGatewayModelParameters, vals map[string]cty.Value) {
	vals["schema"] = cty.StringVal(p.Schema)
}