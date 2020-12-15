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
	r, ok := mr.(*Wafv2IpSet)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Wafv2IpSet.")
	}
	return EncodeWafv2IpSet(*r), nil
}

func EncodeWafv2IpSet(r Wafv2IpSet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafv2IpSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafv2IpSet_IpAddressVersion(r.Spec.ForProvider, ctyVal)
	EncodeWafv2IpSet_Scope(r.Spec.ForProvider, ctyVal)
	EncodeWafv2IpSet_Addresses(r.Spec.ForProvider, ctyVal)
	EncodeWafv2IpSet_Description(r.Spec.ForProvider, ctyVal)
	EncodeWafv2IpSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafv2IpSet_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWafv2IpSet_LockToken(r.Status.AtProvider, ctyVal)
	EncodeWafv2IpSet_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeWafv2IpSet_Id(p Wafv2IpSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafv2IpSet_IpAddressVersion(p Wafv2IpSetParameters, vals map[string]cty.Value) {
	vals["ip_address_version"] = cty.StringVal(p.IpAddressVersion)
}

func EncodeWafv2IpSet_Scope(p Wafv2IpSetParameters, vals map[string]cty.Value) {
	vals["scope"] = cty.StringVal(p.Scope)
}

func EncodeWafv2IpSet_Addresses(p Wafv2IpSetParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Addresses {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["addresses"] = cty.SetVal(colVals)
}

func EncodeWafv2IpSet_Description(p Wafv2IpSetParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeWafv2IpSet_Name(p Wafv2IpSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2IpSet_Tags(p Wafv2IpSetParameters, vals map[string]cty.Value) {
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

func EncodeWafv2IpSet_LockToken(p Wafv2IpSetObservation, vals map[string]cty.Value) {
	vals["lock_token"] = cty.StringVal(p.LockToken)
}

func EncodeWafv2IpSet_Arn(p Wafv2IpSetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}