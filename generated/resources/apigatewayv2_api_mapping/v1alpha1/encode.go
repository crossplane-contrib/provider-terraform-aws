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
	r, ok := mr.(*Apigatewayv2ApiMapping)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Apigatewayv2ApiMapping.")
	}
	return EncodeApigatewayv2ApiMapping(*r), nil
}

func EncodeApigatewayv2ApiMapping(r Apigatewayv2ApiMapping) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2ApiMapping_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2ApiMapping_ApiMappingKey(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2ApiMapping_DomainName(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2ApiMapping_Id(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2ApiMapping_Stage(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeApigatewayv2ApiMapping_ApiId(p Apigatewayv2ApiMappingParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeApigatewayv2ApiMapping_ApiMappingKey(p Apigatewayv2ApiMappingParameters, vals map[string]cty.Value) {
	vals["api_mapping_key"] = cty.StringVal(p.ApiMappingKey)
}

func EncodeApigatewayv2ApiMapping_DomainName(p Apigatewayv2ApiMappingParameters, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeApigatewayv2ApiMapping_Id(p Apigatewayv2ApiMappingParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApigatewayv2ApiMapping_Stage(p Apigatewayv2ApiMappingParameters, vals map[string]cty.Value) {
	vals["stage"] = cty.StringVal(p.Stage)
}