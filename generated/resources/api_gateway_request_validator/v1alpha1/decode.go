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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*ApiGatewayRequestValidator)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApiGatewayRequestValidator(r, ctyValue)
}

func DecodeApiGatewayRequestValidator(prev *ApiGatewayRequestValidator, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApiGatewayRequestValidator_Name(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayRequestValidator_RestApiId(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayRequestValidator_ValidateRequestBody(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayRequestValidator_ValidateRequestParameters(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayRequestValidator_Id(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayRequestValidator_Name(p *ApiGatewayRequestValidatorParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayRequestValidator_RestApiId(p *ApiGatewayRequestValidatorParameters, vals map[string]cty.Value) {
	p.RestApiId = ctwhy.ValueAsString(vals["rest_api_id"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayRequestValidator_ValidateRequestBody(p *ApiGatewayRequestValidatorParameters, vals map[string]cty.Value) {
	p.ValidateRequestBody = ctwhy.ValueAsBool(vals["validate_request_body"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayRequestValidator_ValidateRequestParameters(p *ApiGatewayRequestValidatorParameters, vals map[string]cty.Value) {
	p.ValidateRequestParameters = ctwhy.ValueAsBool(vals["validate_request_parameters"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayRequestValidator_Id(p *ApiGatewayRequestValidatorParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}