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
	r, ok := mr.(*GlueUserDefinedFunction)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GlueUserDefinedFunction.")
	}
	return EncodeGlueUserDefinedFunction(*r), nil
}

func EncodeGlueUserDefinedFunction(r GlueUserDefinedFunction) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlueUserDefinedFunction_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlueUserDefinedFunction_OwnerName(r.Spec.ForProvider, ctyVal)
	EncodeGlueUserDefinedFunction_CatalogId(r.Spec.ForProvider, ctyVal)
	EncodeGlueUserDefinedFunction_ClassName(r.Spec.ForProvider, ctyVal)
	EncodeGlueUserDefinedFunction_DatabaseName(r.Spec.ForProvider, ctyVal)
	EncodeGlueUserDefinedFunction_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlueUserDefinedFunction_OwnerType(r.Spec.ForProvider, ctyVal)
	EncodeGlueUserDefinedFunction_ResourceUris(r.Spec.ForProvider.ResourceUris, ctyVal)
	EncodeGlueUserDefinedFunction_CreateTime(r.Status.AtProvider, ctyVal)
	EncodeGlueUserDefinedFunction_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeGlueUserDefinedFunction_Id(p GlueUserDefinedFunctionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlueUserDefinedFunction_OwnerName(p GlueUserDefinedFunctionParameters, vals map[string]cty.Value) {
	vals["owner_name"] = cty.StringVal(p.OwnerName)
}

func EncodeGlueUserDefinedFunction_CatalogId(p GlueUserDefinedFunctionParameters, vals map[string]cty.Value) {
	vals["catalog_id"] = cty.StringVal(p.CatalogId)
}

func EncodeGlueUserDefinedFunction_ClassName(p GlueUserDefinedFunctionParameters, vals map[string]cty.Value) {
	vals["class_name"] = cty.StringVal(p.ClassName)
}

func EncodeGlueUserDefinedFunction_DatabaseName(p GlueUserDefinedFunctionParameters, vals map[string]cty.Value) {
	vals["database_name"] = cty.StringVal(p.DatabaseName)
}

func EncodeGlueUserDefinedFunction_Name(p GlueUserDefinedFunctionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueUserDefinedFunction_OwnerType(p GlueUserDefinedFunctionParameters, vals map[string]cty.Value) {
	vals["owner_type"] = cty.StringVal(p.OwnerType)
}

func EncodeGlueUserDefinedFunction_ResourceUris(p []ResourceUris, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeGlueUserDefinedFunction_ResourceUris_ResourceType(v, ctyVal)
		EncodeGlueUserDefinedFunction_ResourceUris_Uri(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["resource_uris"] = cty.SetVal(valsForCollection)
}

func EncodeGlueUserDefinedFunction_ResourceUris_ResourceType(p ResourceUris, vals map[string]cty.Value) {
	vals["resource_type"] = cty.StringVal(p.ResourceType)
}

func EncodeGlueUserDefinedFunction_ResourceUris_Uri(p ResourceUris, vals map[string]cty.Value) {
	vals["uri"] = cty.StringVal(p.Uri)
}

func EncodeGlueUserDefinedFunction_CreateTime(p GlueUserDefinedFunctionObservation, vals map[string]cty.Value) {
	vals["create_time"] = cty.StringVal(p.CreateTime)
}

func EncodeGlueUserDefinedFunction_Arn(p GlueUserDefinedFunctionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}