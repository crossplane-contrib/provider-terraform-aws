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

func EncodeAppsyncResolver(r AppsyncResolver) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncResolver_Kind(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncResolver_ResponseTemplate(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncResolver_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncResolver_RequestTemplate(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncResolver_Type(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncResolver_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncResolver_DataSource(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncResolver_Field(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncResolver_CachingConfig(r.Spec.ForProvider.CachingConfig, ctyVal)
	EncodeAppsyncResolver_PipelineConfig(r.Spec.ForProvider.PipelineConfig, ctyVal)
	EncodeAppsyncResolver_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAppsyncResolver_Kind(p AppsyncResolverParameters, vals map[string]cty.Value) {
	vals["kind"] = cty.StringVal(p.Kind)
}

func EncodeAppsyncResolver_ResponseTemplate(p AppsyncResolverParameters, vals map[string]cty.Value) {
	vals["response_template"] = cty.StringVal(p.ResponseTemplate)
}

func EncodeAppsyncResolver_Id(p AppsyncResolverParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppsyncResolver_RequestTemplate(p AppsyncResolverParameters, vals map[string]cty.Value) {
	vals["request_template"] = cty.StringVal(p.RequestTemplate)
}

func EncodeAppsyncResolver_Type(p AppsyncResolverParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeAppsyncResolver_ApiId(p AppsyncResolverParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeAppsyncResolver_DataSource(p AppsyncResolverParameters, vals map[string]cty.Value) {
	vals["data_source"] = cty.StringVal(p.DataSource)
}

func EncodeAppsyncResolver_Field(p AppsyncResolverParameters, vals map[string]cty.Value) {
	vals["field"] = cty.StringVal(p.Field)
}

func EncodeAppsyncResolver_CachingConfig(p CachingConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncResolver_CachingConfig_CachingKeys(p, ctyVal)
	EncodeAppsyncResolver_CachingConfig_Ttl(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["caching_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncResolver_CachingConfig_CachingKeys(p CachingConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CachingKeys {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["caching_keys"] = cty.SetVal(colVals)
}

func EncodeAppsyncResolver_CachingConfig_Ttl(p CachingConfig, vals map[string]cty.Value) {
	vals["ttl"] = cty.NumberIntVal(p.Ttl)
}

func EncodeAppsyncResolver_PipelineConfig(p PipelineConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncResolver_PipelineConfig_Functions(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["pipeline_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncResolver_PipelineConfig_Functions(p PipelineConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Functions {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["functions"] = cty.ListVal(colVals)
}

func EncodeAppsyncResolver_Arn(p AppsyncResolverObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}