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
	r, ok := mr.(*IotThingType)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a IotThingType.")
	}
	return EncodeIotThingType(*r), nil
}

func EncodeIotThingType(r IotThingType) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIotThingType_Id(r.Spec.ForProvider, ctyVal)
	EncodeIotThingType_Name(r.Spec.ForProvider, ctyVal)
	EncodeIotThingType_Deprecated(r.Spec.ForProvider, ctyVal)
	EncodeIotThingType_Properties(r.Spec.ForProvider.Properties, ctyVal)
	EncodeIotThingType_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeIotThingType_Id(p IotThingTypeParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIotThingType_Name(p IotThingTypeParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIotThingType_Deprecated(p IotThingTypeParameters, vals map[string]cty.Value) {
	vals["deprecated"] = cty.BoolVal(p.Deprecated)
}

func EncodeIotThingType_Properties(p Properties, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotThingType_Properties_SearchableAttributes(p, ctyVal)
	EncodeIotThingType_Properties_Description(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["properties"] = cty.ListVal(valsForCollection)
}

func EncodeIotThingType_Properties_SearchableAttributes(p Properties, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SearchableAttributes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["searchable_attributes"] = cty.SetVal(colVals)
}

func EncodeIotThingType_Properties_Description(p Properties, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeIotThingType_Arn(p IotThingTypeObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}