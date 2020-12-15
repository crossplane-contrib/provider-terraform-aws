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
	r, ok := mr.(*FsxLustreFileSystem)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a FsxLustreFileSystem.")
	}
	return EncodeFsxLustreFileSystem(*r), nil
}

func EncodeFsxLustreFileSystem(r FsxLustreFileSystem) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeFsxLustreFileSystem_DeploymentType(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_ImportPath(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_WeeklyMaintenanceStartTime(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_AutoImportPolicy(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_DailyAutomaticBackupStartTime(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_PerUnitStorageThroughput(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_SecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_Tags(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_StorageType(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_AutomaticBackupRetentionDays(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_ExportPath(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_ImportedFileChunkSize(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_StorageCapacity(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_DriveCacheType(r.Spec.ForProvider, ctyVal)
	EncodeFsxLustreFileSystem_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeFsxLustreFileSystem_Arn(r.Status.AtProvider, ctyVal)
	EncodeFsxLustreFileSystem_DnsName(r.Status.AtProvider, ctyVal)
	EncodeFsxLustreFileSystem_MountName(r.Status.AtProvider, ctyVal)
	EncodeFsxLustreFileSystem_NetworkInterfaceIds(r.Status.AtProvider, ctyVal)
	EncodeFsxLustreFileSystem_VpcId(r.Status.AtProvider, ctyVal)
	EncodeFsxLustreFileSystem_OwnerId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeFsxLustreFileSystem_DeploymentType(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	vals["deployment_type"] = cty.StringVal(p.DeploymentType)
}

func EncodeFsxLustreFileSystem_ImportPath(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	vals["import_path"] = cty.StringVal(p.ImportPath)
}

func EncodeFsxLustreFileSystem_KmsKeyId(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeFsxLustreFileSystem_WeeklyMaintenanceStartTime(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	vals["weekly_maintenance_start_time"] = cty.StringVal(p.WeeklyMaintenanceStartTime)
}

func EncodeFsxLustreFileSystem_AutoImportPolicy(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	vals["auto_import_policy"] = cty.StringVal(p.AutoImportPolicy)
}

func EncodeFsxLustreFileSystem_DailyAutomaticBackupStartTime(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	vals["daily_automatic_backup_start_time"] = cty.StringVal(p.DailyAutomaticBackupStartTime)
}

func EncodeFsxLustreFileSystem_PerUnitStorageThroughput(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	vals["per_unit_storage_throughput"] = cty.NumberIntVal(p.PerUnitStorageThroughput)
}

func EncodeFsxLustreFileSystem_SecurityGroupIds(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeFsxLustreFileSystem_Tags(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
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

func EncodeFsxLustreFileSystem_StorageType(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	vals["storage_type"] = cty.StringVal(p.StorageType)
}

func EncodeFsxLustreFileSystem_AutomaticBackupRetentionDays(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	vals["automatic_backup_retention_days"] = cty.NumberIntVal(p.AutomaticBackupRetentionDays)
}

func EncodeFsxLustreFileSystem_ExportPath(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	vals["export_path"] = cty.StringVal(p.ExportPath)
}

func EncodeFsxLustreFileSystem_ImportedFileChunkSize(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	vals["imported_file_chunk_size"] = cty.NumberIntVal(p.ImportedFileChunkSize)
}

func EncodeFsxLustreFileSystem_StorageCapacity(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	vals["storage_capacity"] = cty.NumberIntVal(p.StorageCapacity)
}

func EncodeFsxLustreFileSystem_SubnetIds(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.ListVal(colVals)
}

func EncodeFsxLustreFileSystem_DriveCacheType(p FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	vals["drive_cache_type"] = cty.StringVal(p.DriveCacheType)
}

func EncodeFsxLustreFileSystem_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeFsxLustreFileSystem_Timeouts_Create(p, ctyVal)
	EncodeFsxLustreFileSystem_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeFsxLustreFileSystem_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeFsxLustreFileSystem_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeFsxLustreFileSystem_Arn(p FsxLustreFileSystemObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeFsxLustreFileSystem_DnsName(p FsxLustreFileSystemObservation, vals map[string]cty.Value) {
	vals["dns_name"] = cty.StringVal(p.DnsName)
}

func EncodeFsxLustreFileSystem_MountName(p FsxLustreFileSystemObservation, vals map[string]cty.Value) {
	vals["mount_name"] = cty.StringVal(p.MountName)
}

func EncodeFsxLustreFileSystem_NetworkInterfaceIds(p FsxLustreFileSystemObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NetworkInterfaceIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["network_interface_ids"] = cty.ListVal(colVals)
}

func EncodeFsxLustreFileSystem_VpcId(p FsxLustreFileSystemObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeFsxLustreFileSystem_OwnerId(p FsxLustreFileSystemObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}