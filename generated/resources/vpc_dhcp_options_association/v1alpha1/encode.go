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

package v1alpha1func EncodeVpcDhcpOptionsAssociation(r VpcDhcpOptionsAssociation) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeVpcDhcpOptionsAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpcDhcpOptionsAssociation_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeVpcDhcpOptionsAssociation_DhcpOptionsId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeVpcDhcpOptionsAssociation_Id(p *VpcDhcpOptionsAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpcDhcpOptionsAssociation_VpcId(p *VpcDhcpOptionsAssociationParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeVpcDhcpOptionsAssociation_DhcpOptionsId(p *VpcDhcpOptionsAssociationParameters, vals map[string]cty.Value) {
	vals["dhcp_options_id"] = cty.StringVal(p.DhcpOptionsId)
}