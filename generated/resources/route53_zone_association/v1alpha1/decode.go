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
	r, ok := mr.(*Route53ZoneAssociation)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeRoute53ZoneAssociation(r, ctyValue)
}

func DecodeRoute53ZoneAssociation(prev *Route53ZoneAssociation, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeRoute53ZoneAssociation_VpcId(&new.Spec.ForProvider, valMap)
	DecodeRoute53ZoneAssociation_VpcRegion(&new.Spec.ForProvider, valMap)
	DecodeRoute53ZoneAssociation_ZoneId(&new.Spec.ForProvider, valMap)
	DecodeRoute53ZoneAssociation_OwningAccount(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeRoute53ZoneAssociation_VpcId(p *Route53ZoneAssociationParameters, vals map[string]cty.Value) {
	p.VpcId = ctwhy.ValueAsString(vals["vpc_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute53ZoneAssociation_VpcRegion(p *Route53ZoneAssociationParameters, vals map[string]cty.Value) {
	p.VpcRegion = ctwhy.ValueAsString(vals["vpc_region"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute53ZoneAssociation_ZoneId(p *Route53ZoneAssociationParameters, vals map[string]cty.Value) {
	p.ZoneId = ctwhy.ValueAsString(vals["zone_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute53ZoneAssociation_OwningAccount(p *Route53ZoneAssociationObservation, vals map[string]cty.Value) {
	p.OwningAccount = ctwhy.ValueAsString(vals["owning_account"])
}