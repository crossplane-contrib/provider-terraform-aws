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

func EncodeAppsyncFunction(r AppsyncFunction) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncFunction_DataSource(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncFunction_Description(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncFunction_FunctionVersion(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncFunction_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncFunction_Name(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncFunction_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncFunction_ResponseMappingTemplate(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncFunction_RequestMappingTemplate(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncFunction_Arn(r.Status.AtProvider, ctyVal)
	EncodeAppsyncFunction_FunctionId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAppsyncFunction_DataSource(p AppsyncFunctionParameters, vals map[string]cty.Value) {
	vals["data_source"] = cty.StringVal(p.DataSource)
}

func EncodeAppsyncFunction_Description(p AppsyncFunctionParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeAppsyncFunction_FunctionVersion(p AppsyncFunctionParameters, vals map[string]cty.Value) {
	vals["function_version"] = cty.StringVal(p.FunctionVersion)
}

func EncodeAppsyncFunction_Id(p AppsyncFunctionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppsyncFunction_Name(p AppsyncFunctionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppsyncFunction_ApiId(p AppsyncFunctionParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeAppsyncFunction_ResponseMappingTemplate(p AppsyncFunctionParameters, vals map[string]cty.Value) {
	vals["response_mapping_template"] = cty.StringVal(p.ResponseMappingTemplate)
}

func EncodeAppsyncFunction_RequestMappingTemplate(p AppsyncFunctionParameters, vals map[string]cty.Value) {
	vals["request_mapping_template"] = cty.StringVal(p.RequestMappingTemplate)
}

func EncodeAppsyncFunction_Arn(p AppsyncFunctionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeAppsyncFunction_FunctionId(p AppsyncFunctionObservation, vals map[string]cty.Value) {
	vals["function_id"] = cty.StringVal(p.FunctionId)
}