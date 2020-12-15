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
	r, ok := mr.(*Ec2LocalGatewayRouteTableVpcAssociation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Ec2LocalGatewayRouteTableVpcAssociation.")
	}
	return EncodeEc2LocalGatewayRouteTableVpcAssociation(*r), nil
}

func EncodeEc2LocalGatewayRouteTableVpcAssociation(r Ec2LocalGatewayRouteTableVpcAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2LocalGatewayRouteTableVpcAssociation_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeEc2LocalGatewayRouteTableVpcAssociation_LocalGatewayRouteTableId(r.Spec.ForProvider, ctyVal)
	EncodeEc2LocalGatewayRouteTableVpcAssociation_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEc2LocalGatewayRouteTableVpcAssociation_LocalGatewayId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2LocalGatewayRouteTableVpcAssociation_VpcId(p Ec2LocalGatewayRouteTableVpcAssociationParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeEc2LocalGatewayRouteTableVpcAssociation_LocalGatewayRouteTableId(p Ec2LocalGatewayRouteTableVpcAssociationParameters, vals map[string]cty.Value) {
	vals["local_gateway_route_table_id"] = cty.StringVal(p.LocalGatewayRouteTableId)
}

func EncodeEc2LocalGatewayRouteTableVpcAssociation_Tags(p Ec2LocalGatewayRouteTableVpcAssociationParameters, vals map[string]cty.Value) {
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

func EncodeEc2LocalGatewayRouteTableVpcAssociation_LocalGatewayId(p Ec2LocalGatewayRouteTableVpcAssociationObservation, vals map[string]cty.Value) {
	vals["local_gateway_id"] = cty.StringVal(p.LocalGatewayId)
}