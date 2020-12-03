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

package v1alpha1func EncodeIamInstanceProfile(r IamInstanceProfile) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeIamInstanceProfile_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamInstanceProfile_Name(r.Spec.ForProvider, ctyVal)
	EncodeIamInstanceProfile_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeIamInstanceProfile_Path(r.Spec.ForProvider, ctyVal)
	EncodeIamInstanceProfile_Role(r.Spec.ForProvider, ctyVal)
	EncodeIamInstanceProfile_CreateDate(r.Status.AtProvider, ctyVal)
	EncodeIamInstanceProfile_UniqueId(r.Status.AtProvider, ctyVal)
	EncodeIamInstanceProfile_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeIamInstanceProfile_Id(p *IamInstanceProfileParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamInstanceProfile_Name(p *IamInstanceProfileParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamInstanceProfile_NamePrefix(p *IamInstanceProfileParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeIamInstanceProfile_Path(p *IamInstanceProfileParameters, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeIamInstanceProfile_Role(p *IamInstanceProfileParameters, vals map[string]cty.Value) {
	vals["role"] = cty.StringVal(p.Role)
}

func EncodeIamInstanceProfile_CreateDate(p *IamInstanceProfileObservation, vals map[string]cty.Value) {
	vals["create_date"] = cty.StringVal(p.CreateDate)
}

func EncodeIamInstanceProfile_UniqueId(p *IamInstanceProfileObservation, vals map[string]cty.Value) {
	vals["unique_id"] = cty.StringVal(p.UniqueId)
}

func EncodeIamInstanceProfile_Arn(p *IamInstanceProfileObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}