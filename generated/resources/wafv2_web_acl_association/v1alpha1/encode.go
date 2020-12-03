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

package v1alpha1func EncodeWafv2WebAclAssociation(r Wafv2WebAclAssociation) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeWafv2WebAclAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafv2WebAclAssociation_ResourceArn(r.Spec.ForProvider, ctyVal)
	EncodeWafv2WebAclAssociation_WebAclArn(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeWafv2WebAclAssociation_Id(p *Wafv2WebAclAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafv2WebAclAssociation_ResourceArn(p *Wafv2WebAclAssociationParameters, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}

func EncodeWafv2WebAclAssociation_WebAclArn(p *Wafv2WebAclAssociationParameters, vals map[string]cty.Value) {
	vals["web_acl_arn"] = cty.StringVal(p.WebAclArn)
}