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
	r, ok := mr.(*ApiGatewayMethodResponse)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ApiGatewayMethodResponse.")
	}
	return EncodeApiGatewayMethodResponse(*r), nil
}

func EncodeApiGatewayMethodResponse(r ApiGatewayMethodResponse) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayMethodResponse_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodResponse_StatusCode(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodResponse_HttpMethod(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodResponse_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodResponse_ResourceId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodResponse_ResponseModels(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodResponse_ResponseParameters(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
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

func EncodeApiGatewayMethodResponse_ResourceId(p ApiGatewayMethodResponseParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeApiGatewayMethodResponse_ResponseModels(p ApiGatewayMethodResponseParameters, vals map[string]cty.Value) {
	if len(p.ResponseModels) == 0 {
		vals["response_models"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseModels {
		mVals[key] = cty.StringVal(value)
	}
	vals["response_models"] = cty.MapVal(mVals)
}

func EncodeApiGatewayMethodResponse_ResponseParameters(p ApiGatewayMethodResponseParameters, vals map[string]cty.Value) {
	if len(p.ResponseParameters) == 0 {
		vals["response_parameters"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseParameters {
		mVals[key] = cty.BoolVal(value)
	}
	vals["response_parameters"] = cty.MapVal(mVals)
}