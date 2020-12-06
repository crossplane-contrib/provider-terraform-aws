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

func EncodeOpsworksUserProfile(r OpsworksUserProfile) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksUserProfile_SshUsername(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksUserProfile_UserArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksUserProfile_AllowSelfManagement(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksUserProfile_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksUserProfile_SshPublicKey(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksUserProfile_SshUsername(p OpsworksUserProfileParameters, vals map[string]cty.Value) {
	vals["ssh_username"] = cty.StringVal(p.SshUsername)
}

func EncodeOpsworksUserProfile_UserArn(p OpsworksUserProfileParameters, vals map[string]cty.Value) {
	vals["user_arn"] = cty.StringVal(p.UserArn)
}

func EncodeOpsworksUserProfile_AllowSelfManagement(p OpsworksUserProfileParameters, vals map[string]cty.Value) {
	vals["allow_self_management"] = cty.BoolVal(p.AllowSelfManagement)
}

func EncodeOpsworksUserProfile_Id(p OpsworksUserProfileParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksUserProfile_SshPublicKey(p OpsworksUserProfileParameters, vals map[string]cty.Value) {
	vals["ssh_public_key"] = cty.StringVal(p.SshPublicKey)
}