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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*FsxLustreFileSystem)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeFsxLustreFileSystem(r, ctyValue)
}

func DecodeFsxLustreFileSystem(prev *FsxLustreFileSystem, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeFsxLustreFileSystem_DeploymentType(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_KmsKeyId(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_StorageCapacity(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_Tags(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_DailyAutomaticBackupStartTime(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_DriveCacheType(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_ImportedFileChunkSize(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_StorageType(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_SubnetIds(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_AutoImportPolicy(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_ExportPath(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_Id(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_SecurityGroupIds(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_ImportPath(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_PerUnitStorageThroughput(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_WeeklyMaintenanceStartTime(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_AutomaticBackupRetentionDays(&new.Spec.ForProvider, valMap)
	DecodeFsxLustreFileSystem_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeFsxLustreFileSystem_OwnerId(&new.Status.AtProvider, valMap)
	DecodeFsxLustreFileSystem_VpcId(&new.Status.AtProvider, valMap)
	DecodeFsxLustreFileSystem_NetworkInterfaceIds(&new.Status.AtProvider, valMap)
	DecodeFsxLustreFileSystem_DnsName(&new.Status.AtProvider, valMap)
	DecodeFsxLustreFileSystem_Arn(&new.Status.AtProvider, valMap)
	DecodeFsxLustreFileSystem_MountName(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeFsxLustreFileSystem_DeploymentType(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.DeploymentType = ctwhy.ValueAsString(vals["deployment_type"])
}

func DecodeFsxLustreFileSystem_KmsKeyId(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

func DecodeFsxLustreFileSystem_StorageCapacity(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.StorageCapacity = ctwhy.ValueAsInt64(vals["storage_capacity"])
}

func DecodeFsxLustreFileSystem_Tags(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeFsxLustreFileSystem_DailyAutomaticBackupStartTime(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.DailyAutomaticBackupStartTime = ctwhy.ValueAsString(vals["daily_automatic_backup_start_time"])
}

func DecodeFsxLustreFileSystem_DriveCacheType(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.DriveCacheType = ctwhy.ValueAsString(vals["drive_cache_type"])
}

func DecodeFsxLustreFileSystem_ImportedFileChunkSize(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.ImportedFileChunkSize = ctwhy.ValueAsInt64(vals["imported_file_chunk_size"])
}

func DecodeFsxLustreFileSystem_StorageType(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.StorageType = ctwhy.ValueAsString(vals["storage_type"])
}

func DecodeFsxLustreFileSystem_SubnetIds(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["subnet_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.SubnetIds = goVals
}

func DecodeFsxLustreFileSystem_AutoImportPolicy(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.AutoImportPolicy = ctwhy.ValueAsString(vals["auto_import_policy"])
}

func DecodeFsxLustreFileSystem_ExportPath(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.ExportPath = ctwhy.ValueAsString(vals["export_path"])
}

func DecodeFsxLustreFileSystem_Id(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeFsxLustreFileSystem_SecurityGroupIds(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["security_group_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.SecurityGroupIds = goVals
}

func DecodeFsxLustreFileSystem_ImportPath(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.ImportPath = ctwhy.ValueAsString(vals["import_path"])
}

func DecodeFsxLustreFileSystem_PerUnitStorageThroughput(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.PerUnitStorageThroughput = ctwhy.ValueAsInt64(vals["per_unit_storage_throughput"])
}

func DecodeFsxLustreFileSystem_WeeklyMaintenanceStartTime(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.WeeklyMaintenanceStartTime = ctwhy.ValueAsString(vals["weekly_maintenance_start_time"])
}

func DecodeFsxLustreFileSystem_AutomaticBackupRetentionDays(p *FsxLustreFileSystemParameters, vals map[string]cty.Value) {
	p.AutomaticBackupRetentionDays = ctwhy.ValueAsInt64(vals["automatic_backup_retention_days"])
}

func DecodeFsxLustreFileSystem_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeFsxLustreFileSystem_Timeouts_Delete(p, valMap)
	DecodeFsxLustreFileSystem_Timeouts_Create(p, valMap)
}

func DecodeFsxLustreFileSystem_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeFsxLustreFileSystem_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeFsxLustreFileSystem_OwnerId(p *FsxLustreFileSystemObservation, vals map[string]cty.Value) {
	p.OwnerId = ctwhy.ValueAsString(vals["owner_id"])
}

func DecodeFsxLustreFileSystem_VpcId(p *FsxLustreFileSystemObservation, vals map[string]cty.Value) {
	p.VpcId = ctwhy.ValueAsString(vals["vpc_id"])
}

func DecodeFsxLustreFileSystem_NetworkInterfaceIds(p *FsxLustreFileSystemObservation, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["network_interface_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.NetworkInterfaceIds = goVals
}

func DecodeFsxLustreFileSystem_DnsName(p *FsxLustreFileSystemObservation, vals map[string]cty.Value) {
	p.DnsName = ctwhy.ValueAsString(vals["dns_name"])
}

func DecodeFsxLustreFileSystem_Arn(p *FsxLustreFileSystemObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeFsxLustreFileSystem_MountName(p *FsxLustreFileSystemObservation, vals map[string]cty.Value) {
	p.MountName = ctwhy.ValueAsString(vals["mount_name"])
}