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

package v1alpha1func EncodeEc2ClientVpnNetworkAssociation(r Ec2ClientVpnNetworkAssociation) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEc2ClientVpnNetworkAssociation_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnNetworkAssociation_ClientVpnEndpointId(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnNetworkAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnNetworkAssociation_SecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnNetworkAssociation_VpcId(r.Status.AtProvider, ctyVal)
	EncodeEc2ClientVpnNetworkAssociation_AssociationId(r.Status.AtProvider, ctyVal)
	EncodeEc2ClientVpnNetworkAssociation_Status(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeEc2ClientVpnNetworkAssociation_SubnetId(p *Ec2ClientVpnNetworkAssociationParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeEc2ClientVpnNetworkAssociation_ClientVpnEndpointId(p *Ec2ClientVpnNetworkAssociationParameters, vals map[string]cty.Value) {
	vals["client_vpn_endpoint_id"] = cty.StringVal(p.ClientVpnEndpointId)
}

func EncodeEc2ClientVpnNetworkAssociation_Id(p *Ec2ClientVpnNetworkAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2ClientVpnNetworkAssociation_SecurityGroups(p *Ec2ClientVpnNetworkAssociationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeEc2ClientVpnNetworkAssociation_VpcId(p *Ec2ClientVpnNetworkAssociationObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeEc2ClientVpnNetworkAssociation_AssociationId(p *Ec2ClientVpnNetworkAssociationObservation, vals map[string]cty.Value) {
	vals["association_id"] = cty.StringVal(p.AssociationId)
}

func EncodeEc2ClientVpnNetworkAssociation_Status(p *Ec2ClientVpnNetworkAssociationObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}