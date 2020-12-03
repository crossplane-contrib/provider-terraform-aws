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

package v1alpha1func EncodeWorkspacesIpGroup(r WorkspacesIpGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeWorkspacesIpGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesIpGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesIpGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesIpGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesIpGroup_Rules(r.Spec.ForProvider.Rules, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeWorkspacesIpGroup_Tags(p *WorkspacesIpGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeWorkspacesIpGroup_Description(p *WorkspacesIpGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeWorkspacesIpGroup_Id(p *WorkspacesIpGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWorkspacesIpGroup_Name(p *WorkspacesIpGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWorkspacesIpGroup_Rules(p *Rules, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Rules {
		ctyVal = make(map[string]cty.Value)
		EncodeWorkspacesIpGroup_Rules_Description(v, ctyVal)
		EncodeWorkspacesIpGroup_Rules_Source(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["rules"] = cty.SetVal(valsForCollection)
}

func EncodeWorkspacesIpGroup_Rules_Description(p *Rules, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeWorkspacesIpGroup_Rules_Source(p *Rules, vals map[string]cty.Value) {
	vals["source"] = cty.StringVal(p.Source)
}