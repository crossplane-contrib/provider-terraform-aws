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
	r, ok := mr.(*ApiGatewayGatewayResponse)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ApiGatewayGatewayResponse.")
	}
	return EncodeApiGatewayGatewayResponse(*r), nil
}

func EncodeApiGatewayGatewayResponse(r ApiGatewayGatewayResponse) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayGatewayResponse_ResponseTemplates(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayGatewayResponse_ResponseType(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayGatewayResponse_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayGatewayResponse_StatusCode(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayGatewayResponse_ResponseParameters(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayGatewayResponse_ResponseTemplates(p ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	if len(p.ResponseTemplates) == 0 {
		vals["response_templates"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseTemplates {
		mVals[key] = cty.StringVal(value)
	}
	vals["response_templates"] = cty.MapVal(mVals)
}

func EncodeApiGatewayGatewayResponse_ResponseType(p ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	vals["response_type"] = cty.StringVal(p.ResponseType)
}

func EncodeApiGatewayGatewayResponse_RestApiId(p ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayGatewayResponse_StatusCode(p ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	vals["status_code"] = cty.StringVal(p.StatusCode)
}

func EncodeApiGatewayGatewayResponse_ResponseParameters(p ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	if len(p.ResponseParameters) == 0 {
		vals["response_parameters"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseParameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["response_parameters"] = cty.MapVal(mVals)
}