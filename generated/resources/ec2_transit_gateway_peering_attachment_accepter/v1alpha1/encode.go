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
	r, ok := mr.(*Ec2TransitGatewayPeeringAttachmentAccepter)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Ec2TransitGatewayPeeringAttachmentAccepter.")
	}
	return EncodeEc2TransitGatewayPeeringAttachmentAccepter(*r), nil
}

func EncodeEc2TransitGatewayPeeringAttachmentAccepter(r Ec2TransitGatewayPeeringAttachmentAccepter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2TransitGatewayPeeringAttachmentAccepter_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachmentAccepter_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachmentAccepter_TransitGatewayAttachmentId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachmentAccepter_PeerAccountId(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachmentAccepter_PeerRegion(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachmentAccepter_PeerTransitGatewayId(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachmentAccepter_TransitGatewayId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2TransitGatewayPeeringAttachmentAccepter_Id(p Ec2TransitGatewayPeeringAttachmentAccepterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2TransitGatewayPeeringAttachmentAccepter_Tags(p Ec2TransitGatewayPeeringAttachmentAccepterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEc2TransitGatewayPeeringAttachmentAccepter_TransitGatewayAttachmentId(p Ec2TransitGatewayPeeringAttachmentAccepterParameters, vals map[string]cty.Value) {
	vals["transit_gateway_attachment_id"] = cty.StringVal(p.TransitGatewayAttachmentId)
}

func EncodeEc2TransitGatewayPeeringAttachmentAccepter_PeerAccountId(p Ec2TransitGatewayPeeringAttachmentAccepterObservation, vals map[string]cty.Value) {
	vals["peer_account_id"] = cty.StringVal(p.PeerAccountId)
}

func EncodeEc2TransitGatewayPeeringAttachmentAccepter_PeerRegion(p Ec2TransitGatewayPeeringAttachmentAccepterObservation, vals map[string]cty.Value) {
	vals["peer_region"] = cty.StringVal(p.PeerRegion)
}

func EncodeEc2TransitGatewayPeeringAttachmentAccepter_PeerTransitGatewayId(p Ec2TransitGatewayPeeringAttachmentAccepterObservation, vals map[string]cty.Value) {
	vals["peer_transit_gateway_id"] = cty.StringVal(p.PeerTransitGatewayId)
}

func EncodeEc2TransitGatewayPeeringAttachmentAccepter_TransitGatewayId(p Ec2TransitGatewayPeeringAttachmentAccepterObservation, vals map[string]cty.Value) {
	vals["transit_gateway_id"] = cty.StringVal(p.TransitGatewayId)
}