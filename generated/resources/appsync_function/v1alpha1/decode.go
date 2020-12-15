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
	r, ok := mr.(*AppsyncFunction)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeAppsyncFunction(r, ctyValue)
}

func DecodeAppsyncFunction(prev *AppsyncFunction, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeAppsyncFunction_Name(&new.Spec.ForProvider, valMap)
	DecodeAppsyncFunction_RequestMappingTemplate(&new.Spec.ForProvider, valMap)
	DecodeAppsyncFunction_ApiId(&new.Spec.ForProvider, valMap)
	DecodeAppsyncFunction_DataSource(&new.Spec.ForProvider, valMap)
	DecodeAppsyncFunction_FunctionVersion(&new.Spec.ForProvider, valMap)
	DecodeAppsyncFunction_Description(&new.Spec.ForProvider, valMap)
	DecodeAppsyncFunction_Id(&new.Spec.ForProvider, valMap)
	DecodeAppsyncFunction_ResponseMappingTemplate(&new.Spec.ForProvider, valMap)
	DecodeAppsyncFunction_Arn(&new.Status.AtProvider, valMap)
	DecodeAppsyncFunction_FunctionId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeAppsyncFunction_Name(p *AppsyncFunctionParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeAppsyncFunction_RequestMappingTemplate(p *AppsyncFunctionParameters, vals map[string]cty.Value) {
	p.RequestMappingTemplate = ctwhy.ValueAsString(vals["request_mapping_template"])
}

//primitiveTypeDecodeTemplate
func DecodeAppsyncFunction_ApiId(p *AppsyncFunctionParameters, vals map[string]cty.Value) {
	p.ApiId = ctwhy.ValueAsString(vals["api_id"])
}

//primitiveTypeDecodeTemplate
func DecodeAppsyncFunction_DataSource(p *AppsyncFunctionParameters, vals map[string]cty.Value) {
	p.DataSource = ctwhy.ValueAsString(vals["data_source"])
}

//primitiveTypeDecodeTemplate
func DecodeAppsyncFunction_FunctionVersion(p *AppsyncFunctionParameters, vals map[string]cty.Value) {
	p.FunctionVersion = ctwhy.ValueAsString(vals["function_version"])
}

//primitiveTypeDecodeTemplate
func DecodeAppsyncFunction_Description(p *AppsyncFunctionParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeAppsyncFunction_Id(p *AppsyncFunctionParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeAppsyncFunction_ResponseMappingTemplate(p *AppsyncFunctionParameters, vals map[string]cty.Value) {
	p.ResponseMappingTemplate = ctwhy.ValueAsString(vals["response_mapping_template"])
}

//primitiveTypeDecodeTemplate
func DecodeAppsyncFunction_Arn(p *AppsyncFunctionObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeAppsyncFunction_FunctionId(p *AppsyncFunctionObservation, vals map[string]cty.Value) {
	p.FunctionId = ctwhy.ValueAsString(vals["function_id"])
}