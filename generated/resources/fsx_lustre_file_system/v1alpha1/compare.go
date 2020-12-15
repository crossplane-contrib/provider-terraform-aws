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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
)

//mergeManagedResourceEntrypointTemplate
type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube resource.Managed, prov resource.Managed) plugin.MergeDescription {
	k := kube.(*FsxLustreFileSystem)
	p := prov.(*FsxLustreFileSystem)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeFsxLustreFileSystem_DeploymentType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_ImportPath(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_KmsKeyId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_WeeklyMaintenanceStartTime(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_AutoImportPolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_DailyAutomaticBackupStartTime(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_PerUnitStorageThroughput(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_SecurityGroupIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_StorageType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_AutomaticBackupRetentionDays(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_ExportPath(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_ImportedFileChunkSize(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_StorageCapacity(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_SubnetIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_DriveCacheType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_DnsName(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_MountName(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_NetworkInterfaceIds(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_VpcId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_OwnerId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	for key, v := range p.Annotations {
		if k.Annotations[key] != v {
			k.Annotations[key] = v
			md.AnnotationsUpdated = true
		}
	}
	md.AnyFieldUpdated = anyChildUpdated
	return *md
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_DeploymentType(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.DeploymentType != p.DeploymentType {
		p.DeploymentType = k.DeploymentType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_ImportPath(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.ImportPath != p.ImportPath {
		p.ImportPath = k.ImportPath
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_KmsKeyId(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.KmsKeyId != p.KmsKeyId {
		p.KmsKeyId = k.KmsKeyId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_WeeklyMaintenanceStartTime(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.WeeklyMaintenanceStartTime != p.WeeklyMaintenanceStartTime {
		p.WeeklyMaintenanceStartTime = k.WeeklyMaintenanceStartTime
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_AutoImportPolicy(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.AutoImportPolicy != p.AutoImportPolicy {
		p.AutoImportPolicy = k.AutoImportPolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_DailyAutomaticBackupStartTime(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.DailyAutomaticBackupStartTime != p.DailyAutomaticBackupStartTime {
		p.DailyAutomaticBackupStartTime = k.DailyAutomaticBackupStartTime
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_PerUnitStorageThroughput(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.PerUnitStorageThroughput != p.PerUnitStorageThroughput {
		p.PerUnitStorageThroughput = k.PerUnitStorageThroughput
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeFsxLustreFileSystem_SecurityGroupIds(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.SecurityGroupIds, p.SecurityGroupIds) {
		p.SecurityGroupIds = k.SecurityGroupIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeFsxLustreFileSystem_Tags(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_StorageType(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.StorageType != p.StorageType {
		p.StorageType = k.StorageType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_AutomaticBackupRetentionDays(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.AutomaticBackupRetentionDays != p.AutomaticBackupRetentionDays {
		p.AutomaticBackupRetentionDays = k.AutomaticBackupRetentionDays
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_ExportPath(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.ExportPath != p.ExportPath {
		p.ExportPath = k.ExportPath
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_ImportedFileChunkSize(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.ImportedFileChunkSize != p.ImportedFileChunkSize {
		p.ImportedFileChunkSize = k.ImportedFileChunkSize
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_StorageCapacity(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.StorageCapacity != p.StorageCapacity {
		p.StorageCapacity = k.StorageCapacity
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeFsxLustreFileSystem_SubnetIds(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.SubnetIds, p.SubnetIds) {
		p.SubnetIds = k.SubnetIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_DriveCacheType(k *FsxLustreFileSystemParameters, p *FsxLustreFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.DriveCacheType != p.DriveCacheType {
		p.DriveCacheType = k.DriveCacheType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeFsxLustreFileSystem_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeFsxLustreFileSystem_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFsxLustreFileSystem_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFsxLustreFileSystem_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeFsxLustreFileSystem_Arn(k *FsxLustreFileSystemObservation, p *FsxLustreFileSystemObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeFsxLustreFileSystem_DnsName(k *FsxLustreFileSystemObservation, p *FsxLustreFileSystemObservation, md *plugin.MergeDescription) bool {
	if k.DnsName != p.DnsName {
		k.DnsName = p.DnsName
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeFsxLustreFileSystem_MountName(k *FsxLustreFileSystemObservation, p *FsxLustreFileSystemObservation, md *plugin.MergeDescription) bool {
	if k.MountName != p.MountName {
		k.MountName = p.MountName
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeFsxLustreFileSystem_NetworkInterfaceIds(k *FsxLustreFileSystemObservation, p *FsxLustreFileSystemObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.NetworkInterfaceIds, p.NetworkInterfaceIds) {
		k.NetworkInterfaceIds = p.NetworkInterfaceIds
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeFsxLustreFileSystem_VpcId(k *FsxLustreFileSystemObservation, p *FsxLustreFileSystemObservation, md *plugin.MergeDescription) bool {
	if k.VpcId != p.VpcId {
		k.VpcId = p.VpcId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeFsxLustreFileSystem_OwnerId(k *FsxLustreFileSystemObservation, p *FsxLustreFileSystemObservation, md *plugin.MergeDescription) bool {
	if k.OwnerId != p.OwnerId {
		k.OwnerId = p.OwnerId
		md.StatusUpdated = true
		return true
	}
	return false
}