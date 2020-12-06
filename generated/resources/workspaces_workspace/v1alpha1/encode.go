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

func EncodeWorkspacesWorkspace(r WorkspacesWorkspace) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWorkspacesWorkspace_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesWorkspace_UserName(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesWorkspace_UserVolumeEncryptionEnabled(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesWorkspace_DirectoryId(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesWorkspace_VolumeEncryptionKey(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesWorkspace_BundleId(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesWorkspace_Id(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesWorkspace_RootVolumeEncryptionEnabled(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesWorkspace_WorkspaceProperties(r.Spec.ForProvider.WorkspaceProperties, ctyVal)
	EncodeWorkspacesWorkspace_ComputerName(r.Status.AtProvider, ctyVal)
	EncodeWorkspacesWorkspace_IpAddress(r.Status.AtProvider, ctyVal)
	EncodeWorkspacesWorkspace_State(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeWorkspacesWorkspace_Tags(p WorkspacesWorkspaceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeWorkspacesWorkspace_UserName(p WorkspacesWorkspaceParameters, vals map[string]cty.Value) {
	vals["user_name"] = cty.StringVal(p.UserName)
}

func EncodeWorkspacesWorkspace_UserVolumeEncryptionEnabled(p WorkspacesWorkspaceParameters, vals map[string]cty.Value) {
	vals["user_volume_encryption_enabled"] = cty.BoolVal(p.UserVolumeEncryptionEnabled)
}

func EncodeWorkspacesWorkspace_DirectoryId(p WorkspacesWorkspaceParameters, vals map[string]cty.Value) {
	vals["directory_id"] = cty.StringVal(p.DirectoryId)
}

func EncodeWorkspacesWorkspace_VolumeEncryptionKey(p WorkspacesWorkspaceParameters, vals map[string]cty.Value) {
	vals["volume_encryption_key"] = cty.StringVal(p.VolumeEncryptionKey)
}

func EncodeWorkspacesWorkspace_BundleId(p WorkspacesWorkspaceParameters, vals map[string]cty.Value) {
	vals["bundle_id"] = cty.StringVal(p.BundleId)
}

func EncodeWorkspacesWorkspace_Id(p WorkspacesWorkspaceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWorkspacesWorkspace_RootVolumeEncryptionEnabled(p WorkspacesWorkspaceParameters, vals map[string]cty.Value) {
	vals["root_volume_encryption_enabled"] = cty.BoolVal(p.RootVolumeEncryptionEnabled)
}

func EncodeWorkspacesWorkspace_WorkspaceProperties(p WorkspaceProperties, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWorkspacesWorkspace_WorkspaceProperties_ComputeTypeName(p, ctyVal)
	EncodeWorkspacesWorkspace_WorkspaceProperties_RootVolumeSizeGib(p, ctyVal)
	EncodeWorkspacesWorkspace_WorkspaceProperties_RunningMode(p, ctyVal)
	EncodeWorkspacesWorkspace_WorkspaceProperties_RunningModeAutoStopTimeoutInMinutes(p, ctyVal)
	EncodeWorkspacesWorkspace_WorkspaceProperties_UserVolumeSizeGib(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["workspace_properties"] = cty.ListVal(valsForCollection)
}

func EncodeWorkspacesWorkspace_WorkspaceProperties_ComputeTypeName(p WorkspaceProperties, vals map[string]cty.Value) {
	vals["compute_type_name"] = cty.StringVal(p.ComputeTypeName)
}

func EncodeWorkspacesWorkspace_WorkspaceProperties_RootVolumeSizeGib(p WorkspaceProperties, vals map[string]cty.Value) {
	vals["root_volume_size_gib"] = cty.NumberIntVal(p.RootVolumeSizeGib)
}

func EncodeWorkspacesWorkspace_WorkspaceProperties_RunningMode(p WorkspaceProperties, vals map[string]cty.Value) {
	vals["running_mode"] = cty.StringVal(p.RunningMode)
}

func EncodeWorkspacesWorkspace_WorkspaceProperties_RunningModeAutoStopTimeoutInMinutes(p WorkspaceProperties, vals map[string]cty.Value) {
	vals["running_mode_auto_stop_timeout_in_minutes"] = cty.NumberIntVal(p.RunningModeAutoStopTimeoutInMinutes)
}

func EncodeWorkspacesWorkspace_WorkspaceProperties_UserVolumeSizeGib(p WorkspaceProperties, vals map[string]cty.Value) {
	vals["user_volume_size_gib"] = cty.NumberIntVal(p.UserVolumeSizeGib)
}

func EncodeWorkspacesWorkspace_ComputerName(p WorkspacesWorkspaceObservation, vals map[string]cty.Value) {
	vals["computer_name"] = cty.StringVal(p.ComputerName)
}

func EncodeWorkspacesWorkspace_IpAddress(p WorkspacesWorkspaceObservation, vals map[string]cty.Value) {
	vals["ip_address"] = cty.StringVal(p.IpAddress)
}

func EncodeWorkspacesWorkspace_State(p WorkspacesWorkspaceObservation, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}