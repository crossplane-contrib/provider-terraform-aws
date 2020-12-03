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

package v1alpha1func EncodeCognitoIdentityPoolRolesAttachment(r CognitoIdentityPoolRolesAttachment) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCognitoIdentityPoolRolesAttachment_IdentityPoolId(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPoolRolesAttachment_Roles(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPoolRolesAttachment_Id(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPoolRolesAttachment_RoleMapping(r.Spec.ForProvider.RoleMapping, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeCognitoIdentityPoolRolesAttachment_IdentityPoolId(p *CognitoIdentityPoolRolesAttachmentParameters, vals map[string]cty.Value) {
	vals["identity_pool_id"] = cty.StringVal(p.IdentityPoolId)
}

func EncodeCognitoIdentityPoolRolesAttachment_Roles(p *CognitoIdentityPoolRolesAttachmentParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Roles {
		mVals[key] = cty.StringVal(value)
	}
	vals["roles"] = cty.MapVal(mVals)
}

func EncodeCognitoIdentityPoolRolesAttachment_Id(p *CognitoIdentityPoolRolesAttachmentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCognitoIdentityPoolRolesAttachment_RoleMapping(p *RoleMapping, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RoleMapping {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_AmbiguousRoleResolution(v, ctyVal)
		EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_IdentityProvider(v, ctyVal)
		EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_Type(v, ctyVal)
		EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_MappingRule(v.MappingRule, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["role_mapping"] = cty.SetVal(valsForCollection)
}

func EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_AmbiguousRoleResolution(p *RoleMapping, vals map[string]cty.Value) {
	vals["ambiguous_role_resolution"] = cty.StringVal(p.AmbiguousRoleResolution)
}

func EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_IdentityProvider(p *RoleMapping, vals map[string]cty.Value) {
	vals["identity_provider"] = cty.StringVal(p.IdentityProvider)
}

func EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_Type(p *RoleMapping, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_MappingRule(p *MappingRule, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.MappingRule {
		ctyVal = make(map[string]cty.Value)
		EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_MappingRule_RoleArn(v, ctyVal)
		EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_MappingRule_Value(v, ctyVal)
		EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_MappingRule_Claim(v, ctyVal)
		EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_MappingRule_MatchType(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["mapping_rule"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_MappingRule_RoleArn(p *MappingRule, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_MappingRule_Value(p *MappingRule, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_MappingRule_Claim(p *MappingRule, vals map[string]cty.Value) {
	vals["claim"] = cty.StringVal(p.Claim)
}

func EncodeCognitoIdentityPoolRolesAttachment_RoleMapping_MappingRule_MatchType(p *MappingRule, vals map[string]cty.Value) {
	vals["match_type"] = cty.StringVal(p.MatchType)
}