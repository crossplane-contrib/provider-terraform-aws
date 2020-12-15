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
	r, ok := mr.(*VpnGatewayRoutePropagation)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVpnGatewayRoutePropagation(r, ctyValue)
}

func DecodeVpnGatewayRoutePropagation(prev *VpnGatewayRoutePropagation, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVpnGatewayRoutePropagation_Id(&new.Spec.ForProvider, valMap)
	DecodeVpnGatewayRoutePropagation_RouteTableId(&new.Spec.ForProvider, valMap)
	DecodeVpnGatewayRoutePropagation_VpnGatewayId(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeVpnGatewayRoutePropagation_Id(p *VpnGatewayRoutePropagationParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeVpnGatewayRoutePropagation_RouteTableId(p *VpnGatewayRoutePropagationParameters, vals map[string]cty.Value) {
	p.RouteTableId = ctwhy.ValueAsString(vals["route_table_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVpnGatewayRoutePropagation_VpnGatewayId(p *VpnGatewayRoutePropagationParameters, vals map[string]cty.Value) {
	p.VpnGatewayId = ctwhy.ValueAsString(vals["vpn_gateway_id"])
}