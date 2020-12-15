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
	r, ok := mr.(*Ec2ClientVpnNetworkAssociation)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEc2ClientVpnNetworkAssociation(r, ctyValue)
}

func DecodeEc2ClientVpnNetworkAssociation(prev *Ec2ClientVpnNetworkAssociation, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEc2ClientVpnNetworkAssociation_SecurityGroups(&new.Spec.ForProvider, valMap)
	DecodeEc2ClientVpnNetworkAssociation_SubnetId(&new.Spec.ForProvider, valMap)
	DecodeEc2ClientVpnNetworkAssociation_ClientVpnEndpointId(&new.Spec.ForProvider, valMap)
	DecodeEc2ClientVpnNetworkAssociation_Status(&new.Status.AtProvider, valMap)
	DecodeEc2ClientVpnNetworkAssociation_VpcId(&new.Status.AtProvider, valMap)
	DecodeEc2ClientVpnNetworkAssociation_AssociationId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveCollectionTypeDecodeTemplate
func DecodeEc2ClientVpnNetworkAssociation_SecurityGroups(p *Ec2ClientVpnNetworkAssociationParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["security_groups"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.SecurityGroups = goVals
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnNetworkAssociation_SubnetId(p *Ec2ClientVpnNetworkAssociationParameters, vals map[string]cty.Value) {
	p.SubnetId = ctwhy.ValueAsString(vals["subnet_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnNetworkAssociation_ClientVpnEndpointId(p *Ec2ClientVpnNetworkAssociationParameters, vals map[string]cty.Value) {
	p.ClientVpnEndpointId = ctwhy.ValueAsString(vals["client_vpn_endpoint_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnNetworkAssociation_Status(p *Ec2ClientVpnNetworkAssociationObservation, vals map[string]cty.Value) {
	p.Status = ctwhy.ValueAsString(vals["status"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnNetworkAssociation_VpcId(p *Ec2ClientVpnNetworkAssociationObservation, vals map[string]cty.Value) {
	p.VpcId = ctwhy.ValueAsString(vals["vpc_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnNetworkAssociation_AssociationId(p *Ec2ClientVpnNetworkAssociationObservation, vals map[string]cty.Value) {
	p.AssociationId = ctwhy.ValueAsString(vals["association_id"])
}