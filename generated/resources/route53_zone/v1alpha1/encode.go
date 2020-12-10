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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*Route53Zone)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Route53Zone.")
	}
	return EncodeRoute53Zone(*r), nil
}

func EncodeRoute53Zone(r Route53Zone) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53Zone_DelegationSetId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Zone_ForceDestroy(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Zone_Id(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Zone_Name(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Zone_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Zone_Comment(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Zone_Vpc(r.Spec.ForProvider.Vpc, ctyVal)
	EncodeRoute53Zone_NameServers(r.Status.AtProvider, ctyVal)
	EncodeRoute53Zone_ZoneId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeRoute53Zone_DelegationSetId(p Route53ZoneParameters, vals map[string]cty.Value) {
	vals["delegation_set_id"] = cty.StringVal(p.DelegationSetId)
}

func EncodeRoute53Zone_ForceDestroy(p Route53ZoneParameters, vals map[string]cty.Value) {
	vals["force_destroy"] = cty.BoolVal(p.ForceDestroy)
}

func EncodeRoute53Zone_Id(p Route53ZoneParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRoute53Zone_Name(p Route53ZoneParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRoute53Zone_Tags(p Route53ZoneParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRoute53Zone_Comment(p Route53ZoneParameters, vals map[string]cty.Value) {
	vals["comment"] = cty.StringVal(p.Comment)
}

func EncodeRoute53Zone_Vpc(p Vpc, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53Zone_Vpc_VpcId(p, ctyVal)
	EncodeRoute53Zone_Vpc_VpcRegion(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["vpc"] = cty.SetVal(valsForCollection)
}

func EncodeRoute53Zone_Vpc_VpcId(p Vpc, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeRoute53Zone_Vpc_VpcRegion(p Vpc, vals map[string]cty.Value) {
	vals["vpc_region"] = cty.StringVal(p.VpcRegion)
}

func EncodeRoute53Zone_NameServers(p Route53ZoneObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NameServers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["name_servers"] = cty.ListVal(colVals)
}

func EncodeRoute53Zone_ZoneId(p Route53ZoneObservation, vals map[string]cty.Value) {
	vals["zone_id"] = cty.StringVal(p.ZoneId)
}