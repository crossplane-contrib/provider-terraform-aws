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
	r, ok := mr.(*KmsGrant)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a KmsGrant.")
	}
	return EncodeKmsGrant(*r), nil
}

func EncodeKmsGrant(r KmsGrant) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeKmsGrant_RetiringPrincipal(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_GranteePrincipal(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_Operations(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_RetireOnDelete(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_Name(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_GrantCreationTokens(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_Id(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_KeyId(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_Constraints(r.Spec.ForProvider.Constraints, ctyVal)
	EncodeKmsGrant_GrantId(r.Status.AtProvider, ctyVal)
	EncodeKmsGrant_GrantToken(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeKmsGrant_RetiringPrincipal(p KmsGrantParameters, vals map[string]cty.Value) {
	vals["retiring_principal"] = cty.StringVal(p.RetiringPrincipal)
}

func EncodeKmsGrant_GranteePrincipal(p KmsGrantParameters, vals map[string]cty.Value) {
	vals["grantee_principal"] = cty.StringVal(p.GranteePrincipal)
}

func EncodeKmsGrant_Operations(p KmsGrantParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Operations {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["operations"] = cty.SetVal(colVals)
}

func EncodeKmsGrant_RetireOnDelete(p KmsGrantParameters, vals map[string]cty.Value) {
	vals["retire_on_delete"] = cty.BoolVal(p.RetireOnDelete)
}

func EncodeKmsGrant_Name(p KmsGrantParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeKmsGrant_GrantCreationTokens(p KmsGrantParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.GrantCreationTokens {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["grant_creation_tokens"] = cty.SetVal(colVals)
}

func EncodeKmsGrant_Id(p KmsGrantParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKmsGrant_KeyId(p KmsGrantParameters, vals map[string]cty.Value) {
	vals["key_id"] = cty.StringVal(p.KeyId)
}

func EncodeKmsGrant_Constraints(p Constraints, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeKmsGrant_Constraints_EncryptionContextSubset(p, ctyVal)
	EncodeKmsGrant_Constraints_EncryptionContextEquals(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["constraints"] = cty.SetVal(valsForCollection)
}

func EncodeKmsGrant_Constraints_EncryptionContextSubset(p Constraints, vals map[string]cty.Value) {
	if len(p.EncryptionContextSubset) == 0 {
		vals["encryption_context_subset"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.EncryptionContextSubset {
		mVals[key] = cty.StringVal(value)
	}
	vals["encryption_context_subset"] = cty.MapVal(mVals)
}

func EncodeKmsGrant_Constraints_EncryptionContextEquals(p Constraints, vals map[string]cty.Value) {
	if len(p.EncryptionContextEquals) == 0 {
		vals["encryption_context_equals"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.EncryptionContextEquals {
		mVals[key] = cty.StringVal(value)
	}
	vals["encryption_context_equals"] = cty.MapVal(mVals)
}

func EncodeKmsGrant_GrantId(p KmsGrantObservation, vals map[string]cty.Value) {
	vals["grant_id"] = cty.StringVal(p.GrantId)
}

func EncodeKmsGrant_GrantToken(p KmsGrantObservation, vals map[string]cty.Value) {
	vals["grant_token"] = cty.StringVal(p.GrantToken)
}