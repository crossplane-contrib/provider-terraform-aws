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

package v1alpha1

import (
	"github.com/zclconf/go-cty/cty"
)

func EncodeDbSecurityGroup(r DbSecurityGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDbSecurityGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeDbSecurityGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeDbSecurityGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeDbSecurityGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDbSecurityGroup_Ingress(r.Spec.ForProvider.Ingress, ctyVal)
	EncodeDbSecurityGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDbSecurityGroup_Description(p DbSecurityGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDbSecurityGroup_Id(p DbSecurityGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDbSecurityGroup_Name(p DbSecurityGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDbSecurityGroup_Tags(p DbSecurityGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDbSecurityGroup_Ingress(p []Ingress, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeDbSecurityGroup_Ingress_SecurityGroupId(v, ctyVal)
		EncodeDbSecurityGroup_Ingress_SecurityGroupName(v, ctyVal)
		EncodeDbSecurityGroup_Ingress_SecurityGroupOwnerId(v, ctyVal)
		EncodeDbSecurityGroup_Ingress_Cidr(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ingress"] = cty.SetVal(valsForCollection)
}

func EncodeDbSecurityGroup_Ingress_SecurityGroupId(p Ingress, vals map[string]cty.Value) {
	vals["security_group_id"] = cty.StringVal(p.SecurityGroupId)
}

func EncodeDbSecurityGroup_Ingress_SecurityGroupName(p Ingress, vals map[string]cty.Value) {
	vals["security_group_name"] = cty.StringVal(p.SecurityGroupName)
}

func EncodeDbSecurityGroup_Ingress_SecurityGroupOwnerId(p Ingress, vals map[string]cty.Value) {
	vals["security_group_owner_id"] = cty.StringVal(p.SecurityGroupOwnerId)
}

func EncodeDbSecurityGroup_Ingress_Cidr(p Ingress, vals map[string]cty.Value) {
	vals["cidr"] = cty.StringVal(p.Cidr)
}

func EncodeDbSecurityGroup_Arn(p DbSecurityGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}