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

func EncodeApiGatewayMethodResponse(r ApiGatewayMethodResponse) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayMethodResponse_ResourceId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodResponse_ResponseModels(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodResponse_ResponseParameters(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodResponse_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodResponse_StatusCode(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodResponse_HttpMethod(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodResponse_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayMethodResponse_ResourceId(p ApiGatewayMethodResponseParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeApiGatewayMethodResponse_ResponseModels(p ApiGatewayMethodResponseParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseModels {
		mVals[key] = cty.StringVal(value)
	}
	vals["response_models"] = cty.MapVal(mVals)
}

func EncodeApiGatewayMethodResponse_ResponseParameters(p ApiGatewayMethodResponseParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseParameters {
		mVals[key] = cty.BoolVal(value)
	}
	vals["response_parameters"] = cty.MapVal(mVals)
}

func EncodeApiGatewayMethodResponse_RestApiId(p ApiGatewayMethodResponseParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayMethodResponse_StatusCode(p ApiGatewayMethodResponseParameters, vals map[string]cty.Value) {
	vals["status_code"] = cty.StringVal(p.StatusCode)
}

func EncodeApiGatewayMethodResponse_HttpMethod(p ApiGatewayMethodResponseParameters, vals map[string]cty.Value) {
	vals["http_method"] = cty.StringVal(p.HttpMethod)
}

func EncodeApiGatewayMethodResponse_Id(p ApiGatewayMethodResponseParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}