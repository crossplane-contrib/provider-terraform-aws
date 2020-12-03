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

package v1alpha1func EncodeTransferSshKey(r TransferSshKey) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeTransferSshKey_Body(r.Spec.ForProvider, ctyVal)
	EncodeTransferSshKey_Id(r.Spec.ForProvider, ctyVal)
	EncodeTransferSshKey_ServerId(r.Spec.ForProvider, ctyVal)
	EncodeTransferSshKey_UserName(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeTransferSshKey_Body(p *TransferSshKeyParameters, vals map[string]cty.Value) {
	vals["body"] = cty.StringVal(p.Body)
}

func EncodeTransferSshKey_Id(p *TransferSshKeyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeTransferSshKey_ServerId(p *TransferSshKeyParameters, vals map[string]cty.Value) {
	vals["server_id"] = cty.StringVal(p.ServerId)
}

func EncodeTransferSshKey_UserName(p *TransferSshKeyParameters, vals map[string]cty.Value) {
	vals["user_name"] = cty.StringVal(p.UserName)
}