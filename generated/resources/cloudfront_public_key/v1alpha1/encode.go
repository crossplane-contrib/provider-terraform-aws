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
	r, ok := mr.(*CloudfrontPublicKey)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CloudfrontPublicKey.")
	}
	return EncodeCloudfrontPublicKey(*r), nil
}

func EncodeCloudfrontPublicKey(r CloudfrontPublicKey) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudfrontPublicKey_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontPublicKey_Name(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontPublicKey_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontPublicKey_Comment(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontPublicKey_EncodedKey(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontPublicKey_Etag(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontPublicKey_CallerReference(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudfrontPublicKey_Id(p CloudfrontPublicKeyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudfrontPublicKey_Name(p CloudfrontPublicKeyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloudfrontPublicKey_NamePrefix(p CloudfrontPublicKeyParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeCloudfrontPublicKey_Comment(p CloudfrontPublicKeyParameters, vals map[string]cty.Value) {
	vals["comment"] = cty.StringVal(p.Comment)
}

func EncodeCloudfrontPublicKey_EncodedKey(p CloudfrontPublicKeyParameters, vals map[string]cty.Value) {
	vals["encoded_key"] = cty.StringVal(p.EncodedKey)
}

func EncodeCloudfrontPublicKey_Etag(p CloudfrontPublicKeyObservation, vals map[string]cty.Value) {
	vals["etag"] = cty.StringVal(p.Etag)
}

func EncodeCloudfrontPublicKey_CallerReference(p CloudfrontPublicKeyObservation, vals map[string]cty.Value) {
	vals["caller_reference"] = cty.StringVal(p.CallerReference)
}