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

func EncodeRedshiftSubnetGroup(r RedshiftSubnetGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRedshiftSubnetGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSubnetGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSubnetGroup_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSubnetGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSubnetGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSubnetGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeRedshiftSubnetGroup_Id(p RedshiftSubnetGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRedshiftSubnetGroup_Name(p RedshiftSubnetGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRedshiftSubnetGroup_SubnetIds(p RedshiftSubnetGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeRedshiftSubnetGroup_Tags(p RedshiftSubnetGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRedshiftSubnetGroup_Description(p RedshiftSubnetGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeRedshiftSubnetGroup_Arn(p RedshiftSubnetGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}