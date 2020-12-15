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
	r, ok := mr.(*Route53ResolverQueryLogConfig)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Route53ResolverQueryLogConfig.")
	}
	return EncodeRoute53ResolverQueryLogConfig(*r), nil
}

func EncodeRoute53ResolverQueryLogConfig(r Route53ResolverQueryLogConfig) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53ResolverQueryLogConfig_DestinationArn(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverQueryLogConfig_Name(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverQueryLogConfig_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverQueryLogConfig_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeRoute53ResolverQueryLogConfig_ShareStatus(r.Status.AtProvider, ctyVal)
	EncodeRoute53ResolverQueryLogConfig_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeRoute53ResolverQueryLogConfig_DestinationArn(p Route53ResolverQueryLogConfigParameters, vals map[string]cty.Value) {
	vals["destination_arn"] = cty.StringVal(p.DestinationArn)
}

func EncodeRoute53ResolverQueryLogConfig_Name(p Route53ResolverQueryLogConfigParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRoute53ResolverQueryLogConfig_Tags(p Route53ResolverQueryLogConfigParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRoute53ResolverQueryLogConfig_OwnerId(p Route53ResolverQueryLogConfigObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeRoute53ResolverQueryLogConfig_ShareStatus(p Route53ResolverQueryLogConfigObservation, vals map[string]cty.Value) {
	vals["share_status"] = cty.StringVal(p.ShareStatus)
}

func EncodeRoute53ResolverQueryLogConfig_Arn(p Route53ResolverQueryLogConfigObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}