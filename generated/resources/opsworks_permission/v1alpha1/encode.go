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

func EncodeOpsworksPermission(r OpsworksPermission) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksPermission_AllowSudo(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPermission_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPermission_Level(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPermission_StackId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPermission_UserArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPermission_AllowSsh(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksPermission_AllowSudo(p OpsworksPermissionParameters, vals map[string]cty.Value) {
	vals["allow_sudo"] = cty.BoolVal(p.AllowSudo)
}

func EncodeOpsworksPermission_Id(p OpsworksPermissionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksPermission_Level(p OpsworksPermissionParameters, vals map[string]cty.Value) {
	vals["level"] = cty.StringVal(p.Level)
}

func EncodeOpsworksPermission_StackId(p OpsworksPermissionParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}

func EncodeOpsworksPermission_UserArn(p OpsworksPermissionParameters, vals map[string]cty.Value) {
	vals["user_arn"] = cty.StringVal(p.UserArn)
}

func EncodeOpsworksPermission_AllowSsh(p OpsworksPermissionParameters, vals map[string]cty.Value) {
	vals["allow_ssh"] = cty.BoolVal(p.AllowSsh)
}