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

package v1alpha1func EncodeIamGroup(r IamGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeIamGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeIamGroup_Path(r.Spec.ForProvider, ctyVal)
	EncodeIamGroup_UniqueId(r.Status.AtProvider, ctyVal)
	EncodeIamGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeIamGroup_Id(p *IamGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamGroup_Name(p *IamGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamGroup_Path(p *IamGroupParameters, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeIamGroup_UniqueId(p *IamGroupObservation, vals map[string]cty.Value) {
	vals["unique_id"] = cty.StringVal(p.UniqueId)
}

func EncodeIamGroup_Arn(p *IamGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}