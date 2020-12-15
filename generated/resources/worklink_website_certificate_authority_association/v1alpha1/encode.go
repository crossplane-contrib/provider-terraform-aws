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
	r, ok := mr.(*WorklinkWebsiteCertificateAuthorityAssociation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a WorklinkWebsiteCertificateAuthorityAssociation.")
	}
	return EncodeWorklinkWebsiteCertificateAuthorityAssociation(*r), nil
}

func EncodeWorklinkWebsiteCertificateAuthorityAssociation(r WorklinkWebsiteCertificateAuthorityAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWorklinkWebsiteCertificateAuthorityAssociation_FleetArn(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkWebsiteCertificateAuthorityAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkWebsiteCertificateAuthorityAssociation_Certificate(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkWebsiteCertificateAuthorityAssociation_DisplayName(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkWebsiteCertificateAuthorityAssociation_WebsiteCaId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeWorklinkWebsiteCertificateAuthorityAssociation_FleetArn(p WorklinkWebsiteCertificateAuthorityAssociationParameters, vals map[string]cty.Value) {
	vals["fleet_arn"] = cty.StringVal(p.FleetArn)
}

func EncodeWorklinkWebsiteCertificateAuthorityAssociation_Id(p WorklinkWebsiteCertificateAuthorityAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWorklinkWebsiteCertificateAuthorityAssociation_Certificate(p WorklinkWebsiteCertificateAuthorityAssociationParameters, vals map[string]cty.Value) {
	vals["certificate"] = cty.StringVal(p.Certificate)
}

func EncodeWorklinkWebsiteCertificateAuthorityAssociation_DisplayName(p WorklinkWebsiteCertificateAuthorityAssociationParameters, vals map[string]cty.Value) {
	vals["display_name"] = cty.StringVal(p.DisplayName)
}

func EncodeWorklinkWebsiteCertificateAuthorityAssociation_WebsiteCaId(p WorklinkWebsiteCertificateAuthorityAssociationObservation, vals map[string]cty.Value) {
	vals["website_ca_id"] = cty.StringVal(p.WebsiteCaId)
}