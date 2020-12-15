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
	r, ok := mr.(*ApiGatewayRequestValidator)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ApiGatewayRequestValidator.")
	}
	return EncodeApiGatewayRequestValidator(*r), nil
}

func EncodeApiGatewayRequestValidator(r ApiGatewayRequestValidator) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayRequestValidator_Name(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayRequestValidator_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayRequestValidator_ValidateRequestBody(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayRequestValidator_ValidateRequestParameters(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayRequestValidator_Name(p ApiGatewayRequestValidatorParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeApiGatewayRequestValidator_RestApiId(p ApiGatewayRequestValidatorParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayRequestValidator_ValidateRequestBody(p ApiGatewayRequestValidatorParameters, vals map[string]cty.Value) {
	vals["validate_request_body"] = cty.BoolVal(p.ValidateRequestBody)
}

func EncodeApiGatewayRequestValidator_ValidateRequestParameters(p ApiGatewayRequestValidatorParameters, vals map[string]cty.Value) {
	vals["validate_request_parameters"] = cty.BoolVal(p.ValidateRequestParameters)
}