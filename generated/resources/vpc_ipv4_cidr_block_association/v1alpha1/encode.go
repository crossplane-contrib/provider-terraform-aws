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
	r, ok := mr.(*VpcIpv4CidrBlockAssociation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a VpcIpv4CidrBlockAssociation.")
	}
	return EncodeVpcIpv4CidrBlockAssociation(*r), nil
}

func EncodeVpcIpv4CidrBlockAssociation(r VpcIpv4CidrBlockAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVpcIpv4CidrBlockAssociation_CidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeVpcIpv4CidrBlockAssociation_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeVpcIpv4CidrBlockAssociation_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeVpcIpv4CidrBlockAssociation_CidrBlock(p VpcIpv4CidrBlockAssociationParameters, vals map[string]cty.Value) {
	vals["cidr_block"] = cty.StringVal(p.CidrBlock)
}

func EncodeVpcIpv4CidrBlockAssociation_VpcId(p VpcIpv4CidrBlockAssociationParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeVpcIpv4CidrBlockAssociation_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeVpcIpv4CidrBlockAssociation_Timeouts_Create(p, ctyVal)
	EncodeVpcIpv4CidrBlockAssociation_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeVpcIpv4CidrBlockAssociation_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeVpcIpv4CidrBlockAssociation_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}