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
	r, ok := mr.(*WorklinkWebsiteCertificateAuthorityAssociation)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeWorklinkWebsiteCertificateAuthorityAssociation(r, ctyValue)
}

func DecodeWorklinkWebsiteCertificateAuthorityAssociation(prev *WorklinkWebsiteCertificateAuthorityAssociation, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeWorklinkWebsiteCertificateAuthorityAssociation_Certificate(&new.Spec.ForProvider, valMap)
	DecodeWorklinkWebsiteCertificateAuthorityAssociation_DisplayName(&new.Spec.ForProvider, valMap)
	DecodeWorklinkWebsiteCertificateAuthorityAssociation_FleetArn(&new.Spec.ForProvider, valMap)
	DecodeWorklinkWebsiteCertificateAuthorityAssociation_WebsiteCaId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeWorklinkWebsiteCertificateAuthorityAssociation_Certificate(p *WorklinkWebsiteCertificateAuthorityAssociationParameters, vals map[string]cty.Value) {
	p.Certificate = ctwhy.ValueAsString(vals["certificate"])
}

//primitiveTypeDecodeTemplate
func DecodeWorklinkWebsiteCertificateAuthorityAssociation_DisplayName(p *WorklinkWebsiteCertificateAuthorityAssociationParameters, vals map[string]cty.Value) {
	p.DisplayName = ctwhy.ValueAsString(vals["display_name"])
}

//primitiveTypeDecodeTemplate
func DecodeWorklinkWebsiteCertificateAuthorityAssociation_FleetArn(p *WorklinkWebsiteCertificateAuthorityAssociationParameters, vals map[string]cty.Value) {
	p.FleetArn = ctwhy.ValueAsString(vals["fleet_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeWorklinkWebsiteCertificateAuthorityAssociation_WebsiteCaId(p *WorklinkWebsiteCertificateAuthorityAssociationObservation, vals map[string]cty.Value) {
	p.WebsiteCaId = ctwhy.ValueAsString(vals["website_ca_id"])
}