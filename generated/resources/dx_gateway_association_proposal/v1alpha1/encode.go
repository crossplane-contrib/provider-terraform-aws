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
	r, ok := mr.(*DxGatewayAssociationProposal)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DxGatewayAssociationProposal.")
	}
	return EncodeDxGatewayAssociationProposal(*r), nil
}

func EncodeDxGatewayAssociationProposal(r DxGatewayAssociationProposal) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxGatewayAssociationProposal_Id(r.Spec.ForProvider, ctyVal)
	EncodeDxGatewayAssociationProposal_AllowedPrefixes(r.Spec.ForProvider, ctyVal)
	EncodeDxGatewayAssociationProposal_AssociatedGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeDxGatewayAssociationProposal_DxGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeDxGatewayAssociationProposal_DxGatewayOwnerAccountId(r.Spec.ForProvider, ctyVal)
	EncodeDxGatewayAssociationProposal_AssociatedGatewayOwnerAccountId(r.Status.AtProvider, ctyVal)
	EncodeDxGatewayAssociationProposal_AssociatedGatewayType(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDxGatewayAssociationProposal_Id(p DxGatewayAssociationProposalParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDxGatewayAssociationProposal_AllowedPrefixes(p DxGatewayAssociationProposalParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowedPrefixes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allowed_prefixes"] = cty.SetVal(colVals)
}

func EncodeDxGatewayAssociationProposal_AssociatedGatewayId(p DxGatewayAssociationProposalParameters, vals map[string]cty.Value) {
	vals["associated_gateway_id"] = cty.StringVal(p.AssociatedGatewayId)
}

func EncodeDxGatewayAssociationProposal_DxGatewayId(p DxGatewayAssociationProposalParameters, vals map[string]cty.Value) {
	vals["dx_gateway_id"] = cty.StringVal(p.DxGatewayId)
}

func EncodeDxGatewayAssociationProposal_DxGatewayOwnerAccountId(p DxGatewayAssociationProposalParameters, vals map[string]cty.Value) {
	vals["dx_gateway_owner_account_id"] = cty.StringVal(p.DxGatewayOwnerAccountId)
}

func EncodeDxGatewayAssociationProposal_AssociatedGatewayOwnerAccountId(p DxGatewayAssociationProposalObservation, vals map[string]cty.Value) {
	vals["associated_gateway_owner_account_id"] = cty.StringVal(p.AssociatedGatewayOwnerAccountId)
}

func EncodeDxGatewayAssociationProposal_AssociatedGatewayType(p DxGatewayAssociationProposalObservation, vals map[string]cty.Value) {
	vals["associated_gateway_type"] = cty.StringVal(p.AssociatedGatewayType)
}