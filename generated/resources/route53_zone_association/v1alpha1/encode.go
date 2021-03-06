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
	r, ok := mr.(*Route53ZoneAssociation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Route53ZoneAssociation.")
	}
	return EncodeRoute53ZoneAssociation(*r), nil
}

func EncodeRoute53ZoneAssociation(r Route53ZoneAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53ZoneAssociation_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ZoneAssociation_VpcRegion(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ZoneAssociation_ZoneId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ZoneAssociation_OwningAccount(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeRoute53ZoneAssociation_VpcId(p Route53ZoneAssociationParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeRoute53ZoneAssociation_VpcRegion(p Route53ZoneAssociationParameters, vals map[string]cty.Value) {
	vals["vpc_region"] = cty.StringVal(p.VpcRegion)
}

func EncodeRoute53ZoneAssociation_ZoneId(p Route53ZoneAssociationParameters, vals map[string]cty.Value) {
	vals["zone_id"] = cty.StringVal(p.ZoneId)
}

func EncodeRoute53ZoneAssociation_OwningAccount(p Route53ZoneAssociationObservation, vals map[string]cty.Value) {
	vals["owning_account"] = cty.StringVal(p.OwningAccount)
}