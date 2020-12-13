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
	r, ok := mr.(*Ec2ClientVpnRoute)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Ec2ClientVpnRoute.")
	}
	return EncodeEc2ClientVpnRoute(*r), nil
}

func EncodeEc2ClientVpnRoute(r Ec2ClientVpnRoute) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2ClientVpnRoute_DestinationCidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnRoute_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnRoute_TargetVpcSubnetId(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnRoute_ClientVpnEndpointId(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnRoute_Description(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnRoute_Origin(r.Status.AtProvider, ctyVal)
	EncodeEc2ClientVpnRoute_Type(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2ClientVpnRoute_DestinationCidrBlock(p Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	vals["destination_cidr_block"] = cty.StringVal(p.DestinationCidrBlock)
}

func EncodeEc2ClientVpnRoute_Id(p Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2ClientVpnRoute_TargetVpcSubnetId(p Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	vals["target_vpc_subnet_id"] = cty.StringVal(p.TargetVpcSubnetId)
}

func EncodeEc2ClientVpnRoute_ClientVpnEndpointId(p Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	vals["client_vpn_endpoint_id"] = cty.StringVal(p.ClientVpnEndpointId)
}

func EncodeEc2ClientVpnRoute_Description(p Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeEc2ClientVpnRoute_Origin(p Ec2ClientVpnRouteObservation, vals map[string]cty.Value) {
	vals["origin"] = cty.StringVal(p.Origin)
}

func EncodeEc2ClientVpnRoute_Type(p Ec2ClientVpnRouteObservation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}