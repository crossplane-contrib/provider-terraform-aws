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

package v1alpha1func EncodeXrayGroup(r XrayGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeXrayGroup_FilterExpression(r.Spec.ForProvider, ctyVal)
	EncodeXrayGroup_GroupName(r.Spec.ForProvider, ctyVal)
	EncodeXrayGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeXrayGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeXrayGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeXrayGroup_FilterExpression(p *XrayGroupParameters, vals map[string]cty.Value) {
	vals["filter_expression"] = cty.StringVal(p.FilterExpression)
}

func EncodeXrayGroup_GroupName(p *XrayGroupParameters, vals map[string]cty.Value) {
	vals["group_name"] = cty.StringVal(p.GroupName)
}

func EncodeXrayGroup_Id(p *XrayGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeXrayGroup_Tags(p *XrayGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeXrayGroup_Arn(p *XrayGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}