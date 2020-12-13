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
	r, ok := mr.(*RamResourceAssociation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a RamResourceAssociation.")
	}
	return EncodeRamResourceAssociation(*r), nil
}

func EncodeRamResourceAssociation(r RamResourceAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRamResourceAssociation_ResourceShareArn(r.Spec.ForProvider, ctyVal)
	EncodeRamResourceAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeRamResourceAssociation_ResourceArn(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeRamResourceAssociation_ResourceShareArn(p RamResourceAssociationParameters, vals map[string]cty.Value) {
	vals["resource_share_arn"] = cty.StringVal(p.ResourceShareArn)
}

func EncodeRamResourceAssociation_Id(p RamResourceAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRamResourceAssociation_ResourceArn(p RamResourceAssociationParameters, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}