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
	r, ok := mr.(*KeyPair)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a KeyPair.")
	}
	return EncodeKeyPair(*r), nil
}

func EncodeKeyPair(r KeyPair) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeKeyPair_Id(r.Spec.ForProvider, ctyVal)
	EncodeKeyPair_KeyName(r.Spec.ForProvider, ctyVal)
	EncodeKeyPair_KeyNamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeKeyPair_PublicKey(r.Spec.ForProvider, ctyVal)
	EncodeKeyPair_Tags(r.Spec.ForProvider, ctyVal)
	EncodeKeyPair_Arn(r.Status.AtProvider, ctyVal)
	EncodeKeyPair_Fingerprint(r.Status.AtProvider, ctyVal)
	EncodeKeyPair_KeyPairId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeKeyPair_Id(p KeyPairParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKeyPair_KeyName(p KeyPairParameters, vals map[string]cty.Value) {
	vals["key_name"] = cty.StringVal(p.KeyName)
}

func EncodeKeyPair_KeyNamePrefix(p KeyPairParameters, vals map[string]cty.Value) {
	vals["key_name_prefix"] = cty.StringVal(p.KeyNamePrefix)
}

func EncodeKeyPair_PublicKey(p KeyPairParameters, vals map[string]cty.Value) {
	vals["public_key"] = cty.StringVal(p.PublicKey)
}

func EncodeKeyPair_Tags(p KeyPairParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeKeyPair_Arn(p KeyPairObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeKeyPair_Fingerprint(p KeyPairObservation, vals map[string]cty.Value) {
	vals["fingerprint"] = cty.StringVal(p.Fingerprint)
}

func EncodeKeyPair_KeyPairId(p KeyPairObservation, vals map[string]cty.Value) {
	vals["key_pair_id"] = cty.StringVal(p.KeyPairId)
}