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

func EncodeNeptuneSubnetGroup(r NeptuneSubnetGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeNeptuneSubnetGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneSubnetGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneSubnetGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneSubnetGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneSubnetGroup_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneSubnetGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneSubnetGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeNeptuneSubnetGroup_Description(p NeptuneSubnetGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeNeptuneSubnetGroup_Id(p NeptuneSubnetGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeNeptuneSubnetGroup_Name(p NeptuneSubnetGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeNeptuneSubnetGroup_NamePrefix(p NeptuneSubnetGroupParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeNeptuneSubnetGroup_SubnetIds(p NeptuneSubnetGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeNeptuneSubnetGroup_Tags(p NeptuneSubnetGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeNeptuneSubnetGroup_Arn(p NeptuneSubnetGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}