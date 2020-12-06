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

func EncodeTransferUser(r TransferUser) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeTransferUser_HomeDirectory(r.Spec.ForProvider, ctyVal)
	EncodeTransferUser_Policy(r.Spec.ForProvider, ctyVal)
	EncodeTransferUser_Role(r.Spec.ForProvider, ctyVal)
	EncodeTransferUser_ServerId(r.Spec.ForProvider, ctyVal)
	EncodeTransferUser_Tags(r.Spec.ForProvider, ctyVal)
	EncodeTransferUser_UserName(r.Spec.ForProvider, ctyVal)
	EncodeTransferUser_HomeDirectoryType(r.Spec.ForProvider, ctyVal)
	EncodeTransferUser_Id(r.Spec.ForProvider, ctyVal)
	EncodeTransferUser_HomeDirectoryMappings(r.Spec.ForProvider.HomeDirectoryMappings, ctyVal)
	EncodeTransferUser_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeTransferUser_HomeDirectory(p TransferUserParameters, vals map[string]cty.Value) {
	vals["home_directory"] = cty.StringVal(p.HomeDirectory)
}

func EncodeTransferUser_Policy(p TransferUserParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeTransferUser_Role(p TransferUserParameters, vals map[string]cty.Value) {
	vals["role"] = cty.StringVal(p.Role)
}

func EncodeTransferUser_ServerId(p TransferUserParameters, vals map[string]cty.Value) {
	vals["server_id"] = cty.StringVal(p.ServerId)
}

func EncodeTransferUser_Tags(p TransferUserParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeTransferUser_UserName(p TransferUserParameters, vals map[string]cty.Value) {
	vals["user_name"] = cty.StringVal(p.UserName)
}

func EncodeTransferUser_HomeDirectoryType(p TransferUserParameters, vals map[string]cty.Value) {
	vals["home_directory_type"] = cty.StringVal(p.HomeDirectoryType)
}

func EncodeTransferUser_Id(p TransferUserParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeTransferUser_HomeDirectoryMappings(p HomeDirectoryMappings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeTransferUser_HomeDirectoryMappings_Entry(p, ctyVal)
	EncodeTransferUser_HomeDirectoryMappings_Target(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["home_directory_mappings"] = cty.ListVal(valsForCollection)
}

func EncodeTransferUser_HomeDirectoryMappings_Entry(p HomeDirectoryMappings, vals map[string]cty.Value) {
	vals["entry"] = cty.StringVal(p.Entry)
}

func EncodeTransferUser_HomeDirectoryMappings_Target(p HomeDirectoryMappings, vals map[string]cty.Value) {
	vals["target"] = cty.StringVal(p.Target)
}

func EncodeTransferUser_Arn(p TransferUserObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}