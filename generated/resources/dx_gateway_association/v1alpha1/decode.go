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
	r, ok := mr.(*DxGatewayAssociation)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDxGatewayAssociation(r, ctyValue)
}

func DecodeDxGatewayAssociation(prev *DxGatewayAssociation, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDxGatewayAssociation_Id(&new.Spec.ForProvider, valMap)
	DecodeDxGatewayAssociation_ProposalId(&new.Spec.ForProvider, valMap)
	DecodeDxGatewayAssociation_AllowedPrefixes(&new.Spec.ForProvider, valMap)
	DecodeDxGatewayAssociation_AssociatedGatewayId(&new.Spec.ForProvider, valMap)
	DecodeDxGatewayAssociation_AssociatedGatewayOwnerAccountId(&new.Spec.ForProvider, valMap)
	DecodeDxGatewayAssociation_DxGatewayId(&new.Spec.ForProvider, valMap)
	DecodeDxGatewayAssociation_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDxGatewayAssociation_DxGatewayOwnerAccountId(&new.Status.AtProvider, valMap)
	DecodeDxGatewayAssociation_AssociatedGatewayType(&new.Status.AtProvider, valMap)
	DecodeDxGatewayAssociation_DxGatewayAssociationId(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeDxGatewayAssociation_Id(p *DxGatewayAssociationParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeDxGatewayAssociation_ProposalId(p *DxGatewayAssociationParameters, vals map[string]cty.Value) {
	p.ProposalId = ctwhy.ValueAsString(vals["proposal_id"])
}

func DecodeDxGatewayAssociation_AllowedPrefixes(p *DxGatewayAssociationParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["allowed_prefixes"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.AllowedPrefixes = goVals
}

func DecodeDxGatewayAssociation_AssociatedGatewayId(p *DxGatewayAssociationParameters, vals map[string]cty.Value) {
	p.AssociatedGatewayId = ctwhy.ValueAsString(vals["associated_gateway_id"])
}

func DecodeDxGatewayAssociation_AssociatedGatewayOwnerAccountId(p *DxGatewayAssociationParameters, vals map[string]cty.Value) {
	p.AssociatedGatewayOwnerAccountId = ctwhy.ValueAsString(vals["associated_gateway_owner_account_id"])
}

func DecodeDxGatewayAssociation_DxGatewayId(p *DxGatewayAssociationParameters, vals map[string]cty.Value) {
	p.DxGatewayId = ctwhy.ValueAsString(vals["dx_gateway_id"])
}

func DecodeDxGatewayAssociation_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDxGatewayAssociation_Timeouts_Create(p, valMap)
	DecodeDxGatewayAssociation_Timeouts_Delete(p, valMap)
	DecodeDxGatewayAssociation_Timeouts_Update(p, valMap)
}

func DecodeDxGatewayAssociation_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeDxGatewayAssociation_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeDxGatewayAssociation_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

func DecodeDxGatewayAssociation_DxGatewayOwnerAccountId(p *DxGatewayAssociationObservation, vals map[string]cty.Value) {
	p.DxGatewayOwnerAccountId = ctwhy.ValueAsString(vals["dx_gateway_owner_account_id"])
}

func DecodeDxGatewayAssociation_AssociatedGatewayType(p *DxGatewayAssociationObservation, vals map[string]cty.Value) {
	p.AssociatedGatewayType = ctwhy.ValueAsString(vals["associated_gateway_type"])
}

func DecodeDxGatewayAssociation_DxGatewayAssociationId(p *DxGatewayAssociationObservation, vals map[string]cty.Value) {
	p.DxGatewayAssociationId = ctwhy.ValueAsString(vals["dx_gateway_association_id"])
}