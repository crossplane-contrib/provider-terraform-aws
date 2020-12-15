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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*Ec2TransitGatewayRoute)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEc2TransitGatewayRoute(r, ctyValue)
}

func DecodeEc2TransitGatewayRoute(prev *Ec2TransitGatewayRoute, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEc2TransitGatewayRoute_DestinationCidrBlock(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGatewayRoute_TransitGatewayAttachmentId(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGatewayRoute_TransitGatewayRouteTableId(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGatewayRoute_Blackhole(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayRoute_DestinationCidrBlock(p *Ec2TransitGatewayRouteParameters, vals map[string]cty.Value) {
	p.DestinationCidrBlock = ctwhy.ValueAsString(vals["destination_cidr_block"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayRoute_TransitGatewayAttachmentId(p *Ec2TransitGatewayRouteParameters, vals map[string]cty.Value) {
	p.TransitGatewayAttachmentId = ctwhy.ValueAsString(vals["transit_gateway_attachment_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayRoute_TransitGatewayRouteTableId(p *Ec2TransitGatewayRouteParameters, vals map[string]cty.Value) {
	p.TransitGatewayRouteTableId = ctwhy.ValueAsString(vals["transit_gateway_route_table_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayRoute_Blackhole(p *Ec2TransitGatewayRouteParameters, vals map[string]cty.Value) {
	p.Blackhole = ctwhy.ValueAsBool(vals["blackhole"])
}