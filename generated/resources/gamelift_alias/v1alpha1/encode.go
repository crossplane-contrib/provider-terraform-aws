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
	r, ok := mr.(*GameliftAlias)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GameliftAlias.")
	}
	return EncodeGameliftAlias(*r), nil
}

func EncodeGameliftAlias(r GameliftAlias) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGameliftAlias_Description(r.Spec.ForProvider, ctyVal)
	EncodeGameliftAlias_Id(r.Spec.ForProvider, ctyVal)
	EncodeGameliftAlias_Name(r.Spec.ForProvider, ctyVal)
	EncodeGameliftAlias_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGameliftAlias_RoutingStrategy(r.Spec.ForProvider.RoutingStrategy, ctyVal)
	EncodeGameliftAlias_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeGameliftAlias_Description(p GameliftAliasParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeGameliftAlias_Id(p GameliftAliasParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGameliftAlias_Name(p GameliftAliasParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGameliftAlias_Tags(p GameliftAliasParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGameliftAlias_RoutingStrategy(p RoutingStrategy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGameliftAlias_RoutingStrategy_FleetId(p, ctyVal)
	EncodeGameliftAlias_RoutingStrategy_Message(p, ctyVal)
	EncodeGameliftAlias_RoutingStrategy_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["routing_strategy"] = cty.ListVal(valsForCollection)
}

func EncodeGameliftAlias_RoutingStrategy_FleetId(p RoutingStrategy, vals map[string]cty.Value) {
	vals["fleet_id"] = cty.StringVal(p.FleetId)
}

func EncodeGameliftAlias_RoutingStrategy_Message(p RoutingStrategy, vals map[string]cty.Value) {
	vals["message"] = cty.StringVal(p.Message)
}

func EncodeGameliftAlias_RoutingStrategy_Type(p RoutingStrategy, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeGameliftAlias_Arn(p GameliftAliasObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}