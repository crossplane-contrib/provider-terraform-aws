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
	r, ok := mr.(*Ec2TransitGatewayRouteTablePropagation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Ec2TransitGatewayRouteTablePropagation.")
	}
	return EncodeEc2TransitGatewayRouteTablePropagation(*r), nil
}

func EncodeEc2TransitGatewayRouteTablePropagation(r Ec2TransitGatewayRouteTablePropagation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2TransitGatewayRouteTablePropagation_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayRouteTablePropagation_TransitGatewayAttachmentId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayRouteTablePropagation_TransitGatewayRouteTableId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayRouteTablePropagation_ResourceId(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGatewayRouteTablePropagation_ResourceType(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2TransitGatewayRouteTablePropagation_Id(p Ec2TransitGatewayRouteTablePropagationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2TransitGatewayRouteTablePropagation_TransitGatewayAttachmentId(p Ec2TransitGatewayRouteTablePropagationParameters, vals map[string]cty.Value) {
	vals["transit_gateway_attachment_id"] = cty.StringVal(p.TransitGatewayAttachmentId)
}

func EncodeEc2TransitGatewayRouteTablePropagation_TransitGatewayRouteTableId(p Ec2TransitGatewayRouteTablePropagationParameters, vals map[string]cty.Value) {
	vals["transit_gateway_route_table_id"] = cty.StringVal(p.TransitGatewayRouteTableId)
}

func EncodeEc2TransitGatewayRouteTablePropagation_ResourceId(p Ec2TransitGatewayRouteTablePropagationObservation, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeEc2TransitGatewayRouteTablePropagation_ResourceType(p Ec2TransitGatewayRouteTablePropagationObservation, vals map[string]cty.Value) {
	vals["resource_type"] = cty.StringVal(p.ResourceType)
}