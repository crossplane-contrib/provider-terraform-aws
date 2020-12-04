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

func EncodeIamUser(r IamUser) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamUser_Tags(r.Spec.ForProvider, ctyVal)
	EncodeIamUser_ForceDestroy(r.Spec.ForProvider, ctyVal)
	EncodeIamUser_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamUser_Name(r.Spec.ForProvider, ctyVal)
	EncodeIamUser_Path(r.Spec.ForProvider, ctyVal)
	EncodeIamUser_PermissionsBoundary(r.Spec.ForProvider, ctyVal)
	EncodeIamUser_UniqueId(r.Status.AtProvider, ctyVal)
	EncodeIamUser_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeIamUser_Tags(p IamUserParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeIamUser_ForceDestroy(p IamUserParameters, vals map[string]cty.Value) {
	vals["force_destroy"] = cty.BoolVal(p.ForceDestroy)
}

func EncodeIamUser_Id(p IamUserParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamUser_Name(p IamUserParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamUser_Path(p IamUserParameters, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeIamUser_PermissionsBoundary(p IamUserParameters, vals map[string]cty.Value) {
	vals["permissions_boundary"] = cty.StringVal(p.PermissionsBoundary)
}

func EncodeIamUser_UniqueId(p IamUserObservation, vals map[string]cty.Value) {
	vals["unique_id"] = cty.StringVal(p.UniqueId)
}

func EncodeIamUser_Arn(p IamUserObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}