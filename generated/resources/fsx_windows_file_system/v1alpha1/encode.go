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
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*FsxWindowsFileSystem)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a FsxWindowsFileSystem.")
	}
	return EncodeFsxWindowsFileSystem(*r), nil
}

func EncodeFsxWindowsFileSystem(r FsxWindowsFileSystem) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeFsxWindowsFileSystem_DailyAutomaticBackupStartTime(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_Id(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_SecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_StorageCapacity(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_Tags(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_ActiveDirectoryId(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_DeploymentType(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_CopyTagsToBackups(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_PreferredSubnetId(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_SkipFinalBackup(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_StorageType(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_ThroughputCapacity(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_WeeklyMaintenanceStartTime(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_AutomaticBackupRetentionDays(r.Spec.ForProvider, ctyVal)
	EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory(r.Spec.ForProvider.SelfManagedActiveDirectory, ctyVal)
	EncodeFsxWindowsFileSystem_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeFsxWindowsFileSystem_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeFsxWindowsFileSystem_Arn(r.Status.AtProvider, ctyVal)
	EncodeFsxWindowsFileSystem_VpcId(r.Status.AtProvider, ctyVal)
	EncodeFsxWindowsFileSystem_DnsName(r.Status.AtProvider, ctyVal)
	EncodeFsxWindowsFileSystem_NetworkInterfaceIds(r.Status.AtProvider, ctyVal)
	EncodeFsxWindowsFileSystem_PreferredFileServerIp(r.Status.AtProvider, ctyVal)
	EncodeFsxWindowsFileSystem_RemoteAdministrationEndpoint(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeFsxWindowsFileSystem_DailyAutomaticBackupStartTime(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["daily_automatic_backup_start_time"] = cty.StringVal(p.DailyAutomaticBackupStartTime)
}

func EncodeFsxWindowsFileSystem_Id(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeFsxWindowsFileSystem_KmsKeyId(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeFsxWindowsFileSystem_SecurityGroupIds(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeFsxWindowsFileSystem_StorageCapacity(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["storage_capacity"] = cty.NumberIntVal(p.StorageCapacity)
}

func EncodeFsxWindowsFileSystem_Tags(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeFsxWindowsFileSystem_ActiveDirectoryId(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["active_directory_id"] = cty.StringVal(p.ActiveDirectoryId)
}

func EncodeFsxWindowsFileSystem_DeploymentType(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["deployment_type"] = cty.StringVal(p.DeploymentType)
}

func EncodeFsxWindowsFileSystem_SubnetIds(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.ListVal(colVals)
}

func EncodeFsxWindowsFileSystem_CopyTagsToBackups(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["copy_tags_to_backups"] = cty.BoolVal(p.CopyTagsToBackups)
}

func EncodeFsxWindowsFileSystem_PreferredSubnetId(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["preferred_subnet_id"] = cty.StringVal(p.PreferredSubnetId)
}

func EncodeFsxWindowsFileSystem_SkipFinalBackup(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["skip_final_backup"] = cty.BoolVal(p.SkipFinalBackup)
}

func EncodeFsxWindowsFileSystem_StorageType(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["storage_type"] = cty.StringVal(p.StorageType)
}

func EncodeFsxWindowsFileSystem_ThroughputCapacity(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["throughput_capacity"] = cty.NumberIntVal(p.ThroughputCapacity)
}

func EncodeFsxWindowsFileSystem_WeeklyMaintenanceStartTime(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["weekly_maintenance_start_time"] = cty.StringVal(p.WeeklyMaintenanceStartTime)
}

func EncodeFsxWindowsFileSystem_AutomaticBackupRetentionDays(p FsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["automatic_backup_retention_days"] = cty.NumberIntVal(p.AutomaticBackupRetentionDays)
}

func EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory(p SelfManagedActiveDirectory, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory_DnsIps(p, ctyVal)
	EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory_DomainName(p, ctyVal)
	EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory_FileSystemAdministratorsGroup(p, ctyVal)
	EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory_OrganizationalUnitDistinguishedName(p, ctyVal)
	EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory_Password(p, ctyVal)
	EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory_Username(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["self_managed_active_directory"] = cty.ListVal(valsForCollection)
}

func EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory_DnsIps(p SelfManagedActiveDirectory, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.DnsIps {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["dns_ips"] = cty.SetVal(colVals)
}

func EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory_DomainName(p SelfManagedActiveDirectory, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory_FileSystemAdministratorsGroup(p SelfManagedActiveDirectory, vals map[string]cty.Value) {
	vals["file_system_administrators_group"] = cty.StringVal(p.FileSystemAdministratorsGroup)
}

func EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory_OrganizationalUnitDistinguishedName(p SelfManagedActiveDirectory, vals map[string]cty.Value) {
	vals["organizational_unit_distinguished_name"] = cty.StringVal(p.OrganizationalUnitDistinguishedName)
}

func EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory_Password(p SelfManagedActiveDirectory, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeFsxWindowsFileSystem_SelfManagedActiveDirectory_Username(p SelfManagedActiveDirectory, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeFsxWindowsFileSystem_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeFsxWindowsFileSystem_Timeouts_Create(p, ctyVal)
	EncodeFsxWindowsFileSystem_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeFsxWindowsFileSystem_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeFsxWindowsFileSystem_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeFsxWindowsFileSystem_OwnerId(p FsxWindowsFileSystemObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeFsxWindowsFileSystem_Arn(p FsxWindowsFileSystemObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeFsxWindowsFileSystem_VpcId(p FsxWindowsFileSystemObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeFsxWindowsFileSystem_DnsName(p FsxWindowsFileSystemObservation, vals map[string]cty.Value) {
	vals["dns_name"] = cty.StringVal(p.DnsName)
}

func EncodeFsxWindowsFileSystem_NetworkInterfaceIds(p FsxWindowsFileSystemObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NetworkInterfaceIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["network_interface_ids"] = cty.SetVal(colVals)
}

func EncodeFsxWindowsFileSystem_PreferredFileServerIp(p FsxWindowsFileSystemObservation, vals map[string]cty.Value) {
	vals["preferred_file_server_ip"] = cty.StringVal(p.PreferredFileServerIp)
}

func EncodeFsxWindowsFileSystem_RemoteAdministrationEndpoint(p FsxWindowsFileSystemObservation, vals map[string]cty.Value) {
	vals["remote_administration_endpoint"] = cty.StringVal(p.RemoteAdministrationEndpoint)
}