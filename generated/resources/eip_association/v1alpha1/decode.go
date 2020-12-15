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
	r, ok := mr.(*EipAssociation)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEipAssociation(r, ctyValue)
}

func DecodeEipAssociation(prev *EipAssociation, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEipAssociation_AllocationId(&new.Spec.ForProvider, valMap)
	DecodeEipAssociation_AllowReassociation(&new.Spec.ForProvider, valMap)
	DecodeEipAssociation_InstanceId(&new.Spec.ForProvider, valMap)
	DecodeEipAssociation_NetworkInterfaceId(&new.Spec.ForProvider, valMap)
	DecodeEipAssociation_PrivateIpAddress(&new.Spec.ForProvider, valMap)
	DecodeEipAssociation_PublicIp(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeEipAssociation_AllocationId(p *EipAssociationParameters, vals map[string]cty.Value) {
	p.AllocationId = ctwhy.ValueAsString(vals["allocation_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEipAssociation_AllowReassociation(p *EipAssociationParameters, vals map[string]cty.Value) {
	p.AllowReassociation = ctwhy.ValueAsBool(vals["allow_reassociation"])
}

//primitiveTypeDecodeTemplate
func DecodeEipAssociation_InstanceId(p *EipAssociationParameters, vals map[string]cty.Value) {
	p.InstanceId = ctwhy.ValueAsString(vals["instance_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEipAssociation_NetworkInterfaceId(p *EipAssociationParameters, vals map[string]cty.Value) {
	p.NetworkInterfaceId = ctwhy.ValueAsString(vals["network_interface_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEipAssociation_PrivateIpAddress(p *EipAssociationParameters, vals map[string]cty.Value) {
	p.PrivateIpAddress = ctwhy.ValueAsString(vals["private_ip_address"])
}

//primitiveTypeDecodeTemplate
func DecodeEipAssociation_PublicIp(p *EipAssociationParameters, vals map[string]cty.Value) {
	p.PublicIp = ctwhy.ValueAsString(vals["public_ip"])
}