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
	r, ok := mr.(*Route53ResolverQueryLogConfigAssociation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Route53ResolverQueryLogConfigAssociation.")
	}
	return EncodeRoute53ResolverQueryLogConfigAssociation(*r), nil
}

func EncodeRoute53ResolverQueryLogConfigAssociation(r Route53ResolverQueryLogConfigAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53ResolverQueryLogConfigAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverQueryLogConfigAssociation_ResolverQueryLogConfigId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverQueryLogConfigAssociation_ResourceId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeRoute53ResolverQueryLogConfigAssociation_Id(p Route53ResolverQueryLogConfigAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRoute53ResolverQueryLogConfigAssociation_ResolverQueryLogConfigId(p Route53ResolverQueryLogConfigAssociationParameters, vals map[string]cty.Value) {
	vals["resolver_query_log_config_id"] = cty.StringVal(p.ResolverQueryLogConfigId)
}

func EncodeRoute53ResolverQueryLogConfigAssociation_ResourceId(p Route53ResolverQueryLogConfigAssociationParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}