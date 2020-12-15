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
	r, ok := mr.(*NatGateway)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeNatGateway(r, ctyValue)
}

func DecodeNatGateway(prev *NatGateway, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeNatGateway_SubnetId(&new.Spec.ForProvider, valMap)
	DecodeNatGateway_Tags(&new.Spec.ForProvider, valMap)
	DecodeNatGateway_AllocationId(&new.Spec.ForProvider, valMap)
	DecodeNatGateway_Id(&new.Spec.ForProvider, valMap)
	DecodeNatGateway_NetworkInterfaceId(&new.Status.AtProvider, valMap)
	DecodeNatGateway_PrivateIp(&new.Status.AtProvider, valMap)
	DecodeNatGateway_PublicIp(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeNatGateway_SubnetId(p *NatGatewayParameters, vals map[string]cty.Value) {
	p.SubnetId = ctwhy.ValueAsString(vals["subnet_id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeNatGateway_Tags(p *NatGatewayParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeNatGateway_AllocationId(p *NatGatewayParameters, vals map[string]cty.Value) {
	p.AllocationId = ctwhy.ValueAsString(vals["allocation_id"])
}

//primitiveTypeDecodeTemplate
func DecodeNatGateway_Id(p *NatGatewayParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeNatGateway_NetworkInterfaceId(p *NatGatewayObservation, vals map[string]cty.Value) {
	p.NetworkInterfaceId = ctwhy.ValueAsString(vals["network_interface_id"])
}

//primitiveTypeDecodeTemplate
func DecodeNatGateway_PrivateIp(p *NatGatewayObservation, vals map[string]cty.Value) {
	p.PrivateIp = ctwhy.ValueAsString(vals["private_ip"])
}

//primitiveTypeDecodeTemplate
func DecodeNatGateway_PublicIp(p *NatGatewayObservation, vals map[string]cty.Value) {
	p.PublicIp = ctwhy.ValueAsString(vals["public_ip"])
}