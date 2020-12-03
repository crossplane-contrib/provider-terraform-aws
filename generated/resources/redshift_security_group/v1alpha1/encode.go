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

package v1alpha1func EncodeRedshiftSecurityGroup(r RedshiftSecurityGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeRedshiftSecurityGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSecurityGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSecurityGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSecurityGroup_Ingress(r.Spec.ForProvider.Ingress, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeRedshiftSecurityGroup_Name(p *RedshiftSecurityGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRedshiftSecurityGroup_Description(p *RedshiftSecurityGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeRedshiftSecurityGroup_Id(p *RedshiftSecurityGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRedshiftSecurityGroup_Ingress(p *Ingress, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Ingress {
		ctyVal = make(map[string]cty.Value)
		EncodeRedshiftSecurityGroup_Ingress_SecurityGroupOwnerId(v, ctyVal)
		EncodeRedshiftSecurityGroup_Ingress_Cidr(v, ctyVal)
		EncodeRedshiftSecurityGroup_Ingress_SecurityGroupName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ingress"] = cty.SetVal(valsForCollection)
}

func EncodeRedshiftSecurityGroup_Ingress_SecurityGroupOwnerId(p *Ingress, vals map[string]cty.Value) {
	vals["security_group_owner_id"] = cty.StringVal(p.SecurityGroupOwnerId)
}

func EncodeRedshiftSecurityGroup_Ingress_Cidr(p *Ingress, vals map[string]cty.Value) {
	vals["cidr"] = cty.StringVal(p.Cidr)
}

func EncodeRedshiftSecurityGroup_Ingress_SecurityGroupName(p *Ingress, vals map[string]cty.Value) {
	vals["security_group_name"] = cty.StringVal(p.SecurityGroupName)
}