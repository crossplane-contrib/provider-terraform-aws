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
	r, ok := mr.(*Ec2ClientVpnRoute)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEc2ClientVpnRoute(r, ctyValue)
}

func DecodeEc2ClientVpnRoute(prev *Ec2ClientVpnRoute, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEc2ClientVpnRoute_ClientVpnEndpointId(&new.Spec.ForProvider, valMap)
	DecodeEc2ClientVpnRoute_Description(&new.Spec.ForProvider, valMap)
	DecodeEc2ClientVpnRoute_DestinationCidrBlock(&new.Spec.ForProvider, valMap)
	DecodeEc2ClientVpnRoute_TargetVpcSubnetId(&new.Spec.ForProvider, valMap)
	DecodeEc2ClientVpnRoute_Origin(&new.Status.AtProvider, valMap)
	DecodeEc2ClientVpnRoute_Type(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnRoute_ClientVpnEndpointId(p *Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	p.ClientVpnEndpointId = ctwhy.ValueAsString(vals["client_vpn_endpoint_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnRoute_Description(p *Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnRoute_DestinationCidrBlock(p *Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	p.DestinationCidrBlock = ctwhy.ValueAsString(vals["destination_cidr_block"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnRoute_TargetVpcSubnetId(p *Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	p.TargetVpcSubnetId = ctwhy.ValueAsString(vals["target_vpc_subnet_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnRoute_Origin(p *Ec2ClientVpnRouteObservation, vals map[string]cty.Value) {
	p.Origin = ctwhy.ValueAsString(vals["origin"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnRoute_Type(p *Ec2ClientVpnRouteObservation, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}