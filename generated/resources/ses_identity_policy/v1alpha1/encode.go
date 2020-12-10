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
	r, ok := mr.(*SesIdentityPolicy)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SesIdentityPolicy.")
	}
	return EncodeSesIdentityPolicy(*r), nil
}

func EncodeSesIdentityPolicy(r SesIdentityPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSesIdentityPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeSesIdentityPolicy_Identity(r.Spec.ForProvider, ctyVal)
	EncodeSesIdentityPolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeSesIdentityPolicy_Policy(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeSesIdentityPolicy_Id(p SesIdentityPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSesIdentityPolicy_Identity(p SesIdentityPolicyParameters, vals map[string]cty.Value) {
	vals["identity"] = cty.StringVal(p.Identity)
}

func EncodeSesIdentityPolicy_Name(p SesIdentityPolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSesIdentityPolicy_Policy(p SesIdentityPolicyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}