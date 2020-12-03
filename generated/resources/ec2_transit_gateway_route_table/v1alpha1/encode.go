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

package v1alpha1func EncodeEc2TransitGatewayRouteTable(r Ec2TransitGatewayRouteTable) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEc2TransitGatewayRouteTable_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayRouteTable_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayRouteTable_TransitGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayRouteTable_DefaultPropagationRouteTable(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGatewayRouteTable_DefaultAssociationRouteTable(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeEc2TransitGatewayRouteTable_Id(p *Ec2TransitGatewayRouteTableParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2TransitGatewayRouteTable_Tags(p *Ec2TransitGatewayRouteTableParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEc2TransitGatewayRouteTable_TransitGatewayId(p *Ec2TransitGatewayRouteTableParameters, vals map[string]cty.Value) {
	vals["transit_gateway_id"] = cty.StringVal(p.TransitGatewayId)
}

func EncodeEc2TransitGatewayRouteTable_DefaultPropagationRouteTable(p *Ec2TransitGatewayRouteTableObservation, vals map[string]cty.Value) {
	vals["default_propagation_route_table"] = cty.BoolVal(p.DefaultPropagationRouteTable)
}

func EncodeEc2TransitGatewayRouteTable_DefaultAssociationRouteTable(p *Ec2TransitGatewayRouteTableObservation, vals map[string]cty.Value) {
	vals["default_association_route_table"] = cty.BoolVal(p.DefaultAssociationRouteTable)
}