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
	r, ok := mr.(*LbSslNegotiationPolicy)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a LbSslNegotiationPolicy.")
	}
	return EncodeLbSslNegotiationPolicy(*r), nil
}

func EncodeLbSslNegotiationPolicy(r LbSslNegotiationPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLbSslNegotiationPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeLbSslNegotiationPolicy_LbPort(r.Spec.ForProvider, ctyVal)
	EncodeLbSslNegotiationPolicy_LoadBalancer(r.Spec.ForProvider, ctyVal)
	EncodeLbSslNegotiationPolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeLbSslNegotiationPolicy_Attribute(r.Spec.ForProvider.Attribute, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeLbSslNegotiationPolicy_Id(p LbSslNegotiationPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLbSslNegotiationPolicy_LbPort(p LbSslNegotiationPolicyParameters, vals map[string]cty.Value) {
	vals["lb_port"] = cty.NumberIntVal(p.LbPort)
}

func EncodeLbSslNegotiationPolicy_LoadBalancer(p LbSslNegotiationPolicyParameters, vals map[string]cty.Value) {
	vals["load_balancer"] = cty.StringVal(p.LoadBalancer)
}

func EncodeLbSslNegotiationPolicy_Name(p LbSslNegotiationPolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLbSslNegotiationPolicy_Attribute(p Attribute, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbSslNegotiationPolicy_Attribute_Name(p, ctyVal)
	EncodeLbSslNegotiationPolicy_Attribute_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["attribute"] = cty.SetVal(valsForCollection)
}

func EncodeLbSslNegotiationPolicy_Attribute_Name(p Attribute, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLbSslNegotiationPolicy_Attribute_Value(p Attribute, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}