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
	r, ok := mr.(*Ec2LocalGatewayRouteTableVpcAssociation)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEc2LocalGatewayRouteTableVpcAssociation(r, ctyValue)
}

func DecodeEc2LocalGatewayRouteTableVpcAssociation(prev *Ec2LocalGatewayRouteTableVpcAssociation, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEc2LocalGatewayRouteTableVpcAssociation_Id(&new.Spec.ForProvider, valMap)
	DecodeEc2LocalGatewayRouteTableVpcAssociation_LocalGatewayRouteTableId(&new.Spec.ForProvider, valMap)
	DecodeEc2LocalGatewayRouteTableVpcAssociation_Tags(&new.Spec.ForProvider, valMap)
	DecodeEc2LocalGatewayRouteTableVpcAssociation_VpcId(&new.Spec.ForProvider, valMap)
	DecodeEc2LocalGatewayRouteTableVpcAssociation_LocalGatewayId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeEc2LocalGatewayRouteTableVpcAssociation_Id(p *Ec2LocalGatewayRouteTableVpcAssociationParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2LocalGatewayRouteTableVpcAssociation_LocalGatewayRouteTableId(p *Ec2LocalGatewayRouteTableVpcAssociationParameters, vals map[string]cty.Value) {
	p.LocalGatewayRouteTableId = ctwhy.ValueAsString(vals["local_gateway_route_table_id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeEc2LocalGatewayRouteTableVpcAssociation_Tags(p *Ec2LocalGatewayRouteTableVpcAssociationParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeEc2LocalGatewayRouteTableVpcAssociation_VpcId(p *Ec2LocalGatewayRouteTableVpcAssociationParameters, vals map[string]cty.Value) {
	p.VpcId = ctwhy.ValueAsString(vals["vpc_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2LocalGatewayRouteTableVpcAssociation_LocalGatewayId(p *Ec2LocalGatewayRouteTableVpcAssociationObservation, vals map[string]cty.Value) {
	p.LocalGatewayId = ctwhy.ValueAsString(vals["local_gateway_id"])
}