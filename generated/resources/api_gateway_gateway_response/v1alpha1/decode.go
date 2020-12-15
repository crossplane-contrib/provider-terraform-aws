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
	r, ok := mr.(*ApiGatewayGatewayResponse)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApiGatewayGatewayResponse(r, ctyValue)
}

func DecodeApiGatewayGatewayResponse(prev *ApiGatewayGatewayResponse, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApiGatewayGatewayResponse_ResponseTemplates(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayGatewayResponse_ResponseType(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayGatewayResponse_RestApiId(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayGatewayResponse_StatusCode(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayGatewayResponse_ResponseParameters(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeApiGatewayGatewayResponse_ResponseTemplates(p *ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["response_templates"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.ResponseTemplates = vMap
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayGatewayResponse_ResponseType(p *ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	p.ResponseType = ctwhy.ValueAsString(vals["response_type"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayGatewayResponse_RestApiId(p *ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	p.RestApiId = ctwhy.ValueAsString(vals["rest_api_id"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayGatewayResponse_StatusCode(p *ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	p.StatusCode = ctwhy.ValueAsString(vals["status_code"])
}

//primitiveMapTypeDecodeTemplate
func DecodeApiGatewayGatewayResponse_ResponseParameters(p *ApiGatewayGatewayResponseParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["response_parameters"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.ResponseParameters = vMap
}