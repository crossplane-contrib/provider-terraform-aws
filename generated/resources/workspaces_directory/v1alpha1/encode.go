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

func EncodeWorkspacesDirectory(r WorkspacesDirectory) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWorkspacesDirectory_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesDirectory_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesDirectory_DirectoryId(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesDirectory_Id(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesDirectory_SelfServicePermissions(r.Spec.ForProvider.SelfServicePermissions, ctyVal)
	EncodeWorkspacesDirectory_DirectoryType(r.Status.AtProvider, ctyVal)
	EncodeWorkspacesDirectory_DnsIpAddresses(r.Status.AtProvider, ctyVal)
	EncodeWorkspacesDirectory_Alias(r.Status.AtProvider, ctyVal)
	EncodeWorkspacesDirectory_CustomerUserName(r.Status.AtProvider, ctyVal)
	EncodeWorkspacesDirectory_DirectoryName(r.Status.AtProvider, ctyVal)
	EncodeWorkspacesDirectory_IamRoleId(r.Status.AtProvider, ctyVal)
	EncodeWorkspacesDirectory_IpGroupIds(r.Status.AtProvider, ctyVal)
	EncodeWorkspacesDirectory_RegistrationCode(r.Status.AtProvider, ctyVal)
	EncodeWorkspacesDirectory_WorkspaceSecurityGroupId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeWorkspacesDirectory_SubnetIds(p WorkspacesDirectoryParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeWorkspacesDirectory_Tags(p WorkspacesDirectoryParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeWorkspacesDirectory_DirectoryId(p WorkspacesDirectoryParameters, vals map[string]cty.Value) {
	vals["directory_id"] = cty.StringVal(p.DirectoryId)
}

func EncodeWorkspacesDirectory_Id(p WorkspacesDirectoryParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWorkspacesDirectory_SelfServicePermissions(p SelfServicePermissions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWorkspacesDirectory_SelfServicePermissions_RestartWorkspace(p, ctyVal)
	EncodeWorkspacesDirectory_SelfServicePermissions_SwitchRunningMode(p, ctyVal)
	EncodeWorkspacesDirectory_SelfServicePermissions_ChangeComputeType(p, ctyVal)
	EncodeWorkspacesDirectory_SelfServicePermissions_IncreaseVolumeSize(p, ctyVal)
	EncodeWorkspacesDirectory_SelfServicePermissions_RebuildWorkspace(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["self_service_permissions"] = cty.ListVal(valsForCollection)
}

func EncodeWorkspacesDirectory_SelfServicePermissions_RestartWorkspace(p SelfServicePermissions, vals map[string]cty.Value) {
	vals["restart_workspace"] = cty.BoolVal(p.RestartWorkspace)
}

func EncodeWorkspacesDirectory_SelfServicePermissions_SwitchRunningMode(p SelfServicePermissions, vals map[string]cty.Value) {
	vals["switch_running_mode"] = cty.BoolVal(p.SwitchRunningMode)
}

func EncodeWorkspacesDirectory_SelfServicePermissions_ChangeComputeType(p SelfServicePermissions, vals map[string]cty.Value) {
	vals["change_compute_type"] = cty.BoolVal(p.ChangeComputeType)
}

func EncodeWorkspacesDirectory_SelfServicePermissions_IncreaseVolumeSize(p SelfServicePermissions, vals map[string]cty.Value) {
	vals["increase_volume_size"] = cty.BoolVal(p.IncreaseVolumeSize)
}

func EncodeWorkspacesDirectory_SelfServicePermissions_RebuildWorkspace(p SelfServicePermissions, vals map[string]cty.Value) {
	vals["rebuild_workspace"] = cty.BoolVal(p.RebuildWorkspace)
}

func EncodeWorkspacesDirectory_DirectoryType(p WorkspacesDirectoryObservation, vals map[string]cty.Value) {
	vals["directory_type"] = cty.StringVal(p.DirectoryType)
}

func EncodeWorkspacesDirectory_DnsIpAddresses(p WorkspacesDirectoryObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.DnsIpAddresses {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["dns_ip_addresses"] = cty.SetVal(colVals)
}

func EncodeWorkspacesDirectory_Alias(p WorkspacesDirectoryObservation, vals map[string]cty.Value) {
	vals["alias"] = cty.StringVal(p.Alias)
}

func EncodeWorkspacesDirectory_CustomerUserName(p WorkspacesDirectoryObservation, vals map[string]cty.Value) {
	vals["customer_user_name"] = cty.StringVal(p.CustomerUserName)
}

func EncodeWorkspacesDirectory_DirectoryName(p WorkspacesDirectoryObservation, vals map[string]cty.Value) {
	vals["directory_name"] = cty.StringVal(p.DirectoryName)
}

func EncodeWorkspacesDirectory_IamRoleId(p WorkspacesDirectoryObservation, vals map[string]cty.Value) {
	vals["iam_role_id"] = cty.StringVal(p.IamRoleId)
}

func EncodeWorkspacesDirectory_IpGroupIds(p WorkspacesDirectoryObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.IpGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["ip_group_ids"] = cty.SetVal(colVals)
}

func EncodeWorkspacesDirectory_RegistrationCode(p WorkspacesDirectoryObservation, vals map[string]cty.Value) {
	vals["registration_code"] = cty.StringVal(p.RegistrationCode)
}

func EncodeWorkspacesDirectory_WorkspaceSecurityGroupId(p WorkspacesDirectoryObservation, vals map[string]cty.Value) {
	vals["workspace_security_group_id"] = cty.StringVal(p.WorkspaceSecurityGroupId)
}