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

package v1alpha1func EncodeVpcEndpointRouteTableAssociation(r VpcEndpointRouteTableAssociation) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeVpcEndpointRouteTableAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointRouteTableAssociation_RouteTableId(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointRouteTableAssociation_VpcEndpointId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeVpcEndpointRouteTableAssociation_Id(p *VpcEndpointRouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpcEndpointRouteTableAssociation_RouteTableId(p *VpcEndpointRouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["route_table_id"] = cty.StringVal(p.RouteTableId)
}

func EncodeVpcEndpointRouteTableAssociation_VpcEndpointId(p *VpcEndpointRouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["vpc_endpoint_id"] = cty.StringVal(p.VpcEndpointId)
}