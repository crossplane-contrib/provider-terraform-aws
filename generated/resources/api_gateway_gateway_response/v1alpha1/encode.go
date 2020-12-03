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

package v1alpha1func EncodeApiGatewayGatewayResponse(r ApiGatewayGatewayResponse) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeApiGatewayGatewayResponse_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayGatewayResponse_ResponseParameters(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayGatewayResponse_ResponseTemplates(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayGatewayResponse_ResponseType(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayGatewayResponse_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayGatewayResponse_StatusCode(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeApiGatewayGatewayResponse_Id(p *ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayGatewayResponse_ResponseParameters(p *ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseParameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["response_parameters"] = cty.MapVal(mVals)
}

func EncodeApiGatewayGatewayResponse_ResponseTemplates(p *ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseTemplates {
		mVals[key] = cty.StringVal(value)
	}
	vals["response_templates"] = cty.MapVal(mVals)
}

func EncodeApiGatewayGatewayResponse_ResponseType(p *ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	vals["response_type"] = cty.StringVal(p.ResponseType)
}

func EncodeApiGatewayGatewayResponse_RestApiId(p *ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayGatewayResponse_StatusCode(p *ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	vals["status_code"] = cty.StringVal(p.StatusCode)
}