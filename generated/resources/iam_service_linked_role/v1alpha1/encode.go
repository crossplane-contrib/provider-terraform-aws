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

package v1alpha1func EncodeIamServiceLinkedRole(r IamServiceLinkedRole) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeIamServiceLinkedRole_AwsServiceName(r.Spec.ForProvider, ctyVal)
	EncodeIamServiceLinkedRole_Description(r.Spec.ForProvider, ctyVal)
	EncodeIamServiceLinkedRole_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamServiceLinkedRole_CustomSuffix(r.Spec.ForProvider, ctyVal)
	EncodeIamServiceLinkedRole_Arn(r.Status.AtProvider, ctyVal)
	EncodeIamServiceLinkedRole_CreateDate(r.Status.AtProvider, ctyVal)
	EncodeIamServiceLinkedRole_Name(r.Status.AtProvider, ctyVal)
	EncodeIamServiceLinkedRole_Path(r.Status.AtProvider, ctyVal)
	EncodeIamServiceLinkedRole_UniqueId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeIamServiceLinkedRole_AwsServiceName(p *IamServiceLinkedRoleParameters, vals map[string]cty.Value) {
	vals["aws_service_name"] = cty.StringVal(p.AwsServiceName)
}

func EncodeIamServiceLinkedRole_Description(p *IamServiceLinkedRoleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeIamServiceLinkedRole_Id(p *IamServiceLinkedRoleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamServiceLinkedRole_CustomSuffix(p *IamServiceLinkedRoleParameters, vals map[string]cty.Value) {
	vals["custom_suffix"] = cty.StringVal(p.CustomSuffix)
}

func EncodeIamServiceLinkedRole_Arn(p *IamServiceLinkedRoleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeIamServiceLinkedRole_CreateDate(p *IamServiceLinkedRoleObservation, vals map[string]cty.Value) {
	vals["create_date"] = cty.StringVal(p.CreateDate)
}

func EncodeIamServiceLinkedRole_Name(p *IamServiceLinkedRoleObservation, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamServiceLinkedRole_Path(p *IamServiceLinkedRoleObservation, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeIamServiceLinkedRole_UniqueId(p *IamServiceLinkedRoleObservation, vals map[string]cty.Value) {
	vals["unique_id"] = cty.StringVal(p.UniqueId)
}