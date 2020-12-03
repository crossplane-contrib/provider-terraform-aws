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

package v1alpha1func EncodeKmsGrant(r KmsGrant) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeKmsGrant_Operations(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_RetiringPrincipal(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_GranteePrincipal(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_Id(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_KeyId(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_Name(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_RetireOnDelete(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_GrantCreationTokens(r.Spec.ForProvider, ctyVal)
	EncodeKmsGrant_Constraints(r.Spec.ForProvider.Constraints, ctyVal)
	EncodeKmsGrant_GrantId(r.Status.AtProvider, ctyVal)
	EncodeKmsGrant_GrantToken(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeKmsGrant_Operations(p *KmsGrantParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Operations {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["operations"] = cty.SetVal(colVals)
}

func EncodeKmsGrant_RetiringPrincipal(p *KmsGrantParameters, vals map[string]cty.Value) {
	vals["retiring_principal"] = cty.StringVal(p.RetiringPrincipal)
}

func EncodeKmsGrant_GranteePrincipal(p *KmsGrantParameters, vals map[string]cty.Value) {
	vals["grantee_principal"] = cty.StringVal(p.GranteePrincipal)
}

func EncodeKmsGrant_Id(p *KmsGrantParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKmsGrant_KeyId(p *KmsGrantParameters, vals map[string]cty.Value) {
	vals["key_id"] = cty.StringVal(p.KeyId)
}

func EncodeKmsGrant_Name(p *KmsGrantParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeKmsGrant_RetireOnDelete(p *KmsGrantParameters, vals map[string]cty.Value) {
	vals["retire_on_delete"] = cty.BoolVal(p.RetireOnDelete)
}

func EncodeKmsGrant_GrantCreationTokens(p *KmsGrantParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.GrantCreationTokens {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["grant_creation_tokens"] = cty.SetVal(colVals)
}

func EncodeKmsGrant_Constraints(p *Constraints, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Constraints {
		ctyVal = make(map[string]cty.Value)
		EncodeKmsGrant_Constraints_EncryptionContextEquals(v, ctyVal)
		EncodeKmsGrant_Constraints_EncryptionContextSubset(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["constraints"] = cty.SetVal(valsForCollection)
}

func EncodeKmsGrant_Constraints_EncryptionContextEquals(p *Constraints, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.EncryptionContextEquals {
		mVals[key] = cty.StringVal(value)
	}
	vals["encryption_context_equals"] = cty.MapVal(mVals)
}

func EncodeKmsGrant_Constraints_EncryptionContextSubset(p *Constraints, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.EncryptionContextSubset {
		mVals[key] = cty.StringVal(value)
	}
	vals["encryption_context_subset"] = cty.MapVal(mVals)
}

func EncodeKmsGrant_GrantId(p *KmsGrantObservation, vals map[string]cty.Value) {
	vals["grant_id"] = cty.StringVal(p.GrantId)
}

func EncodeKmsGrant_GrantToken(p *KmsGrantObservation, vals map[string]cty.Value) {
	vals["grant_token"] = cty.StringVal(p.GrantToken)
}