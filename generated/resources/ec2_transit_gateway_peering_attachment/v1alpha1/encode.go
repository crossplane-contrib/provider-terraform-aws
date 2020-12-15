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
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*Ec2TransitGatewayPeeringAttachment)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Ec2TransitGatewayPeeringAttachment.")
	}
	return EncodeEc2TransitGatewayPeeringAttachment(*r), nil
}

func EncodeEc2TransitGatewayPeeringAttachment(r Ec2TransitGatewayPeeringAttachment) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2TransitGatewayPeeringAttachment_PeerAccountId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachment_PeerRegion(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachment_PeerTransitGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachment_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayPeeringAttachment_TransitGatewayId(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2TransitGatewayPeeringAttachment_PeerAccountId(p Ec2TransitGatewayPeeringAttachmentParameters, vals map[string]cty.Value) {
	vals["peer_account_id"] = cty.StringVal(p.PeerAccountId)
}

func EncodeEc2TransitGatewayPeeringAttachment_PeerRegion(p Ec2TransitGatewayPeeringAttachmentParameters, vals map[string]cty.Value) {
	vals["peer_region"] = cty.StringVal(p.PeerRegion)
}

func EncodeEc2TransitGatewayPeeringAttachment_PeerTransitGatewayId(p Ec2TransitGatewayPeeringAttachmentParameters, vals map[string]cty.Value) {
	vals["peer_transit_gateway_id"] = cty.StringVal(p.PeerTransitGatewayId)
}

func EncodeEc2TransitGatewayPeeringAttachment_Tags(p Ec2TransitGatewayPeeringAttachmentParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEc2TransitGatewayPeeringAttachment_TransitGatewayId(p Ec2TransitGatewayPeeringAttachmentParameters, vals map[string]cty.Value) {
	vals["transit_gateway_id"] = cty.StringVal(p.TransitGatewayId)
}