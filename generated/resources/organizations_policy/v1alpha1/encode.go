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

package v1alpha1func EncodeOrganizationsPolicy(r OrganizationsPolicy) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeOrganizationsPolicy_Type(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicy_Content(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicy_Description(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicy_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicy_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeOrganizationsPolicy_Type(p *OrganizationsPolicyParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOrganizationsPolicy_Content(p *OrganizationsPolicyParameters, vals map[string]cty.Value) {
	vals["content"] = cty.StringVal(p.Content)
}

func EncodeOrganizationsPolicy_Description(p *OrganizationsPolicyParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeOrganizationsPolicy_Id(p *OrganizationsPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOrganizationsPolicy_Name(p *OrganizationsPolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOrganizationsPolicy_Tags(p *OrganizationsPolicyParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOrganizationsPolicy_Arn(p *OrganizationsPolicyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}