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
	r, ok := mr.(*Ec2TransitGatewayRoute)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Ec2TransitGatewayRoute.")
	}
	return EncodeEc2TransitGatewayRoute(*r), nil
}

func EncodeEc2TransitGatewayRoute(r Ec2TransitGatewayRoute) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2TransitGatewayRoute_DestinationCidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayRoute_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayRoute_TransitGatewayAttachmentId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayRoute_TransitGatewayRouteTableId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayRoute_Blackhole(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeEc2TransitGatewayRoute_DestinationCidrBlock(p Ec2TransitGatewayRouteParameters, vals map[string]cty.Value) {
	vals["destination_cidr_block"] = cty.StringVal(p.DestinationCidrBlock)
}

func EncodeEc2TransitGatewayRoute_Id(p Ec2TransitGatewayRouteParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2TransitGatewayRoute_TransitGatewayAttachmentId(p Ec2TransitGatewayRouteParameters, vals map[string]cty.Value) {
	vals["transit_gateway_attachment_id"] = cty.StringVal(p.TransitGatewayAttachmentId)
}

func EncodeEc2TransitGatewayRoute_TransitGatewayRouteTableId(p Ec2TransitGatewayRouteParameters, vals map[string]cty.Value) {
	vals["transit_gateway_route_table_id"] = cty.StringVal(p.TransitGatewayRouteTableId)
}

func EncodeEc2TransitGatewayRoute_Blackhole(p Ec2TransitGatewayRouteParameters, vals map[string]cty.Value) {
	vals["blackhole"] = cty.BoolVal(p.Blackhole)
}