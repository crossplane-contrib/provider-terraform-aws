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
	r, ok := mr.(*Apigatewayv2Route)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApigatewayv2Route(r, ctyValue)
}

func DecodeApigatewayv2Route(prev *Apigatewayv2Route, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApigatewayv2Route_RequestModels(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Route_RouteResponseSelectionExpression(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Route_ApiId(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Route_ApiKeyRequired(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Route_AuthorizationScopes(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Route_AuthorizationType(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Route_OperationName(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Route_AuthorizerId(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Route_ModelSelectionExpression(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Route_RouteKey(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Route_Target(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeApigatewayv2Route_RequestModels(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["request_models"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.RequestModels = vMap
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Route_RouteResponseSelectionExpression(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	p.RouteResponseSelectionExpression = ctwhy.ValueAsString(vals["route_response_selection_expression"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Route_ApiId(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	p.ApiId = ctwhy.ValueAsString(vals["api_id"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Route_ApiKeyRequired(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	p.ApiKeyRequired = ctwhy.ValueAsBool(vals["api_key_required"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeApigatewayv2Route_AuthorizationScopes(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["authorization_scopes"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.AuthorizationScopes = goVals
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Route_AuthorizationType(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	p.AuthorizationType = ctwhy.ValueAsString(vals["authorization_type"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Route_OperationName(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	p.OperationName = ctwhy.ValueAsString(vals["operation_name"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Route_AuthorizerId(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	p.AuthorizerId = ctwhy.ValueAsString(vals["authorizer_id"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Route_ModelSelectionExpression(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	p.ModelSelectionExpression = ctwhy.ValueAsString(vals["model_selection_expression"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Route_RouteKey(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	p.RouteKey = ctwhy.ValueAsString(vals["route_key"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2Route_Target(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	p.Target = ctwhy.ValueAsString(vals["target"])
}