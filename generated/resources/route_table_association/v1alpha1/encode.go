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
	r, ok := mr.(*RouteTableAssociation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a RouteTableAssociation.")
	}
	return EncodeRouteTableAssociation(*r), nil
}

func EncodeRouteTableAssociation(r RouteTableAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRouteTableAssociation_GatewayId(r.Spec.ForProvider, ctyVal)
	EncodeRouteTableAssociation_RouteTableId(r.Spec.ForProvider, ctyVal)
	EncodeRouteTableAssociation_SubnetId(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeRouteTableAssociation_GatewayId(p RouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["gateway_id"] = cty.StringVal(p.GatewayId)
}

func EncodeRouteTableAssociation_RouteTableId(p RouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["route_table_id"] = cty.StringVal(p.RouteTableId)
}

func EncodeRouteTableAssociation_SubnetId(p RouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}