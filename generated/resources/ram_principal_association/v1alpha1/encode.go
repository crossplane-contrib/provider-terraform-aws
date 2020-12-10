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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*RamPrincipalAssociation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a RamPrincipalAssociation.")
	}
	return EncodeRamPrincipalAssociation(*r), nil
}

func EncodeRamPrincipalAssociation(r RamPrincipalAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRamPrincipalAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeRamPrincipalAssociation_Principal(r.Spec.ForProvider, ctyVal)
	EncodeRamPrincipalAssociation_ResourceShareArn(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeRamPrincipalAssociation_Id(p RamPrincipalAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRamPrincipalAssociation_Principal(p RamPrincipalAssociationParameters, vals map[string]cty.Value) {
	vals["principal"] = cty.StringVal(p.Principal)
}

func EncodeRamPrincipalAssociation_ResourceShareArn(p RamPrincipalAssociationParameters, vals map[string]cty.Value) {
	vals["resource_share_arn"] = cty.StringVal(p.ResourceShareArn)
}