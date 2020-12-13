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
	r, ok := mr.(*ApiGatewayMethod)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ApiGatewayMethod.")
	}
	return EncodeApiGatewayMethod(*r), nil
}

func EncodeApiGatewayMethod(r ApiGatewayMethod) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayMethod_ApiKeyRequired(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_Authorization(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_AuthorizationScopes(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_AuthorizerId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_RequestValidatorId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_HttpMethod(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_RequestModels(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_RequestParameters(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_ResourceId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_RestApiId(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayMethod_ApiKeyRequired(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["api_key_required"] = cty.BoolVal(p.ApiKeyRequired)
}

func EncodeApiGatewayMethod_Authorization(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["authorization"] = cty.StringVal(p.Authorization)
}

func EncodeApiGatewayMethod_AuthorizationScopes(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AuthorizationScopes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["authorization_scopes"] = cty.SetVal(colVals)
}

func EncodeApiGatewayMethod_AuthorizerId(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["authorizer_id"] = cty.StringVal(p.AuthorizerId)
}

func EncodeApiGatewayMethod_Id(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayMethod_RequestValidatorId(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["request_validator_id"] = cty.StringVal(p.RequestValidatorId)
}

func EncodeApiGatewayMethod_HttpMethod(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["http_method"] = cty.StringVal(p.HttpMethod)
}

func EncodeApiGatewayMethod_RequestModels(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	if len(p.RequestModels) == 0 {
		vals["request_models"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.RequestModels {
		mVals[key] = cty.StringVal(value)
	}
	vals["request_models"] = cty.MapVal(mVals)
}

func EncodeApiGatewayMethod_RequestParameters(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	if len(p.RequestParameters) == 0 {
		vals["request_parameters"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.RequestParameters {
		mVals[key] = cty.BoolVal(value)
	}
	vals["request_parameters"] = cty.MapVal(mVals)
}

func EncodeApiGatewayMethod_ResourceId(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeApiGatewayMethod_RestApiId(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}