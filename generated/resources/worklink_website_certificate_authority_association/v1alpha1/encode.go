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

package v1alpha1func EncodeWorklinkWebsiteCertificateAuthorityAssociation(r WorklinkWebsiteCertificateAuthorityAssociation) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeWorklinkWebsiteCertificateAuthorityAssociation_Certificate(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkWebsiteCertificateAuthorityAssociation_DisplayName(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkWebsiteCertificateAuthorityAssociation_FleetArn(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkWebsiteCertificateAuthorityAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkWebsiteCertificateAuthorityAssociation_WebsiteCaId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeWorklinkWebsiteCertificateAuthorityAssociation_Certificate(p *WorklinkWebsiteCertificateAuthorityAssociationParameters, vals map[string]cty.Value) {
	vals["certificate"] = cty.StringVal(p.Certificate)
}

func EncodeWorklinkWebsiteCertificateAuthorityAssociation_DisplayName(p *WorklinkWebsiteCertificateAuthorityAssociationParameters, vals map[string]cty.Value) {
	vals["display_name"] = cty.StringVal(p.DisplayName)
}

func EncodeWorklinkWebsiteCertificateAuthorityAssociation_FleetArn(p *WorklinkWebsiteCertificateAuthorityAssociationParameters, vals map[string]cty.Value) {
	vals["fleet_arn"] = cty.StringVal(p.FleetArn)
}

func EncodeWorklinkWebsiteCertificateAuthorityAssociation_Id(p *WorklinkWebsiteCertificateAuthorityAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWorklinkWebsiteCertificateAuthorityAssociation_WebsiteCaId(p *WorklinkWebsiteCertificateAuthorityAssociationObservation, vals map[string]cty.Value) {
	vals["website_ca_id"] = cty.StringVal(p.WebsiteCaId)
}