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

package v1alpha1func EncodeCognitoResourceServer(r CognitoResourceServer) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCognitoResourceServer_Id(r.Spec.ForProvider, ctyVal)
	EncodeCognitoResourceServer_Identifier(r.Spec.ForProvider, ctyVal)
	EncodeCognitoResourceServer_Name(r.Spec.ForProvider, ctyVal)
	EncodeCognitoResourceServer_UserPoolId(r.Spec.ForProvider, ctyVal)
	EncodeCognitoResourceServer_Scope(r.Spec.ForProvider.Scope, ctyVal)
	EncodeCognitoResourceServer_ScopeIdentifiers(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeCognitoResourceServer_Id(p *CognitoResourceServerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCognitoResourceServer_Identifier(p *CognitoResourceServerParameters, vals map[string]cty.Value) {
	vals["identifier"] = cty.StringVal(p.Identifier)
}

func EncodeCognitoResourceServer_Name(p *CognitoResourceServerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCognitoResourceServer_UserPoolId(p *CognitoResourceServerParameters, vals map[string]cty.Value) {
	vals["user_pool_id"] = cty.StringVal(p.UserPoolId)
}

func EncodeCognitoResourceServer_Scope(p *Scope, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Scope {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoResourceServer_Scope_ScopeDescription(v, ctyVal)
		EncodeCognitoResourceServer_Scope_ScopeName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["scope"] = cty.SetVal(valsForCollection)
}

func EncodeCognitoResourceServer_Scope_ScopeDescription(p *Scope, vals map[string]cty.Value) {
	vals["scope_description"] = cty.StringVal(p.ScopeDescription)
}

func EncodeCognitoResourceServer_Scope_ScopeName(p *Scope, vals map[string]cty.Value) {
	vals["scope_name"] = cty.StringVal(p.ScopeName)
}

func EncodeCognitoResourceServer_ScopeIdentifiers(p *CognitoResourceServerObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ScopeIdentifiers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["scope_identifiers"] = cty.ListVal(colVals)
}