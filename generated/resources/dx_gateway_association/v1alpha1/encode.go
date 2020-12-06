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
	"github.com/zclconf/go-cty/cty"
)

func EncodeDxGatewayAssociation(r DxGatewayAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxGatewayAssociation_AllowedPrefixes(r.Spec.ForProvider, ctyVal)
	EncodeDxGatewayAssociation_ProposalId(r.Spec.ForProvider, ctyVal)
	EncodeDxGatewayAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeDxGatewayAssociation_AssociatedGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeDxGatewayAssociation_AssociatedGatewayOwnerAccountId(r.Spec.ForProvider, ctyVal)
	EncodeDxGatewayAssociation_DxGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeDxGatewayAssociation_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDxGatewayAssociation_DxGatewayAssociationId(r.Status.AtProvider, ctyVal)
	EncodeDxGatewayAssociation_AssociatedGatewayType(r.Status.AtProvider, ctyVal)
	EncodeDxGatewayAssociation_DxGatewayOwnerAccountId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDxGatewayAssociation_AllowedPrefixes(p DxGatewayAssociationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowedPrefixes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allowed_prefixes"] = cty.SetVal(colVals)
}

func EncodeDxGatewayAssociation_ProposalId(p DxGatewayAssociationParameters, vals map[string]cty.Value) {
	vals["proposal_id"] = cty.StringVal(p.ProposalId)
}

func EncodeDxGatewayAssociation_Id(p DxGatewayAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDxGatewayAssociation_AssociatedGatewayId(p DxGatewayAssociationParameters, vals map[string]cty.Value) {
	vals["associated_gateway_id"] = cty.StringVal(p.AssociatedGatewayId)
}

func EncodeDxGatewayAssociation_AssociatedGatewayOwnerAccountId(p DxGatewayAssociationParameters, vals map[string]cty.Value) {
	vals["associated_gateway_owner_account_id"] = cty.StringVal(p.AssociatedGatewayOwnerAccountId)
}

func EncodeDxGatewayAssociation_DxGatewayId(p DxGatewayAssociationParameters, vals map[string]cty.Value) {
	vals["dx_gateway_id"] = cty.StringVal(p.DxGatewayId)
}

func EncodeDxGatewayAssociation_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDxGatewayAssociation_Timeouts_Delete(p, ctyVal)
	EncodeDxGatewayAssociation_Timeouts_Update(p, ctyVal)
	EncodeDxGatewayAssociation_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDxGatewayAssociation_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDxGatewayAssociation_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDxGatewayAssociation_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDxGatewayAssociation_DxGatewayAssociationId(p DxGatewayAssociationObservation, vals map[string]cty.Value) {
	vals["dx_gateway_association_id"] = cty.StringVal(p.DxGatewayAssociationId)
}

func EncodeDxGatewayAssociation_AssociatedGatewayType(p DxGatewayAssociationObservation, vals map[string]cty.Value) {
	vals["associated_gateway_type"] = cty.StringVal(p.AssociatedGatewayType)
}

func EncodeDxGatewayAssociation_DxGatewayOwnerAccountId(p DxGatewayAssociationObservation, vals map[string]cty.Value) {
	vals["dx_gateway_owner_account_id"] = cty.StringVal(p.DxGatewayOwnerAccountId)
}