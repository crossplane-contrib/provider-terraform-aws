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

func EncodeElasticacheSubnetGroup(r ElasticacheSubnetGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeElasticacheSubnetGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheSubnetGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheSubnetGroup_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheSubnetGroup_Description(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeElasticacheSubnetGroup_Id(p ElasticacheSubnetGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElasticacheSubnetGroup_Name(p ElasticacheSubnetGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElasticacheSubnetGroup_SubnetIds(p ElasticacheSubnetGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeElasticacheSubnetGroup_Description(p ElasticacheSubnetGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}