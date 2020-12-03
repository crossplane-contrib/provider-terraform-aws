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

package v1alpha1func EncodeDaxSubnetGroup(r DaxSubnetGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeDaxSubnetGroup_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeDaxSubnetGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeDaxSubnetGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeDaxSubnetGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeDaxSubnetGroup_VpcId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeDaxSubnetGroup_SubnetIds(p *DaxSubnetGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeDaxSubnetGroup_Description(p *DaxSubnetGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDaxSubnetGroup_Id(p *DaxSubnetGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDaxSubnetGroup_Name(p *DaxSubnetGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDaxSubnetGroup_VpcId(p *DaxSubnetGroupObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}