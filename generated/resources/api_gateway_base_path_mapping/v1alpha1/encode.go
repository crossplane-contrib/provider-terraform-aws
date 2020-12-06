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

func EncodeApiGatewayBasePathMapping(r ApiGatewayBasePathMapping) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayBasePathMapping_StageName(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayBasePathMapping_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayBasePathMapping_BasePath(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayBasePathMapping_DomainName(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayBasePathMapping_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayBasePathMapping_StageName(p ApiGatewayBasePathMappingParameters, vals map[string]cty.Value) {
	vals["stage_name"] = cty.StringVal(p.StageName)
}

func EncodeApiGatewayBasePathMapping_ApiId(p ApiGatewayBasePathMappingParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeApiGatewayBasePathMapping_BasePath(p ApiGatewayBasePathMappingParameters, vals map[string]cty.Value) {
	vals["base_path"] = cty.StringVal(p.BasePath)
}

func EncodeApiGatewayBasePathMapping_DomainName(p ApiGatewayBasePathMappingParameters, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeApiGatewayBasePathMapping_Id(p ApiGatewayBasePathMappingParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}