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
	r, ok := mr.(*Ec2TransitGatewayRouteTable)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEc2TransitGatewayRouteTable(r, ctyValue)
}

func DecodeEc2TransitGatewayRouteTable(prev *Ec2TransitGatewayRouteTable, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEc2TransitGatewayRouteTable_Id(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGatewayRouteTable_Tags(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGatewayRouteTable_TransitGatewayId(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGatewayRouteTable_DefaultPropagationRouteTable(&new.Status.AtProvider, valMap)
	DecodeEc2TransitGatewayRouteTable_DefaultAssociationRouteTable(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayRouteTable_Id(p *Ec2TransitGatewayRouteTableParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeEc2TransitGatewayRouteTable_Tags(p *Ec2TransitGatewayRouteTableParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayRouteTable_TransitGatewayId(p *Ec2TransitGatewayRouteTableParameters, vals map[string]cty.Value) {
	p.TransitGatewayId = ctwhy.ValueAsString(vals["transit_gateway_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayRouteTable_DefaultPropagationRouteTable(p *Ec2TransitGatewayRouteTableObservation, vals map[string]cty.Value) {
	p.DefaultPropagationRouteTable = ctwhy.ValueAsBool(vals["default_propagation_route_table"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayRouteTable_DefaultAssociationRouteTable(p *Ec2TransitGatewayRouteTableObservation, vals map[string]cty.Value) {
	p.DefaultAssociationRouteTable = ctwhy.ValueAsBool(vals["default_association_route_table"])
}