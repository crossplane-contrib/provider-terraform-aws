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

package v1alpha1func EncodeEc2TransitGatewayPeeringAttachment(r Ec2TransitGatewayPeeringAttachment) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEc2TransitGatewayPeeringAttachment_PeerRegion(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachment_PeerTransitGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachment_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachment_TransitGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachment_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachment_PeerAccountId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeEc2TransitGatewayPeeringAttachment_PeerRegion(p *Ec2TransitGatewayPeeringAttachmentParameters, vals map[string]cty.Value) {
	vals["peer_region"] = cty.StringVal(p.PeerRegion)
}

func EncodeEc2TransitGatewayPeeringAttachment_PeerTransitGatewayId(p *Ec2TransitGatewayPeeringAttachmentParameters, vals map[string]cty.Value) {
	vals["peer_transit_gateway_id"] = cty.StringVal(p.PeerTransitGatewayId)
}

func EncodeEc2TransitGatewayPeeringAttachment_Tags(p *Ec2TransitGatewayPeeringAttachmentParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEc2TransitGatewayPeeringAttachment_TransitGatewayId(p *Ec2TransitGatewayPeeringAttachmentParameters, vals map[string]cty.Value) {
	vals["transit_gateway_id"] = cty.StringVal(p.TransitGatewayId)
}

func EncodeEc2TransitGatewayPeeringAttachment_Id(p *Ec2TransitGatewayPeeringAttachmentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2TransitGatewayPeeringAttachment_PeerAccountId(p *Ec2TransitGatewayPeeringAttachmentParameters, vals map[string]cty.Value) {
	vals["peer_account_id"] = cty.StringVal(p.PeerAccountId)
}