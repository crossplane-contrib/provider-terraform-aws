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
	k := kube.(*DbSnapshot)
	p := prov.(*DbSnapshot)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDbSnapshot_DbInstanceIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_DbSnapshotIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_AllocatedStorage(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_AvailabilityZone(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_KmsKeyId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_SnapshotType(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_Status(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_StorageType(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_VpcId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_Engine(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_EngineVersion(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_Iops(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_SourceDbSnapshotIdentifier(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_DbSnapshotArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_LicenseModel(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_OptionGroupName(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_Port(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_Encrypted(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSnapshot_SourceRegion(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDbSnapshot_DbInstanceIdentifier(k *DbSnapshotParameters, p *DbSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.DbInstanceIdentifier != p.DbInstanceIdentifier {
		p.DbInstanceIdentifier = k.DbInstanceIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDbSnapshot_DbSnapshotIdentifier(k *DbSnapshotParameters, p *DbSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.DbSnapshotIdentifier != p.DbSnapshotIdentifier {
		p.DbSnapshotIdentifier = k.DbSnapshotIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDbSnapshot_Tags(k *DbSnapshotParameters, p *DbSnapshotParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDbSnapshot_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDbSnapshot_Timeouts_Read(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDbSnapshot_Timeouts_Read(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Read != p.Read {
		p.Read = k.Read
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_AllocatedStorage(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.AllocatedStorage != p.AllocatedStorage {
		k.AllocatedStorage = p.AllocatedStorage
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_AvailabilityZone(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.AvailabilityZone != p.AvailabilityZone {
		k.AvailabilityZone = p.AvailabilityZone
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_KmsKeyId(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.KmsKeyId != p.KmsKeyId {
		k.KmsKeyId = p.KmsKeyId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_SnapshotType(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.SnapshotType != p.SnapshotType {
		k.SnapshotType = p.SnapshotType
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_Status(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.Status != p.Status {
		k.Status = p.Status
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_StorageType(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.StorageType != p.StorageType {
		k.StorageType = p.StorageType
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_VpcId(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.VpcId != p.VpcId {
		k.VpcId = p.VpcId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_Engine(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.Engine != p.Engine {
		k.Engine = p.Engine
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_EngineVersion(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.EngineVersion != p.EngineVersion {
		k.EngineVersion = p.EngineVersion
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_Iops(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.Iops != p.Iops {
		k.Iops = p.Iops
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_SourceDbSnapshotIdentifier(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.SourceDbSnapshotIdentifier != p.SourceDbSnapshotIdentifier {
		k.SourceDbSnapshotIdentifier = p.SourceDbSnapshotIdentifier
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_DbSnapshotArn(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.DbSnapshotArn != p.DbSnapshotArn {
		k.DbSnapshotArn = p.DbSnapshotArn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_LicenseModel(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.LicenseModel != p.LicenseModel {
		k.LicenseModel = p.LicenseModel
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_OptionGroupName(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.OptionGroupName != p.OptionGroupName {
		k.OptionGroupName = p.OptionGroupName
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_Port(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.Port != p.Port {
		k.Port = p.Port
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_Encrypted(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.Encrypted != p.Encrypted {
		k.Encrypted = p.Encrypted
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSnapshot_SourceRegion(k *DbSnapshotObservation, p *DbSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.SourceRegion != p.SourceRegion {
		k.SourceRegion = p.SourceRegion
		md.StatusUpdated = true
		return true
	}
	return false
}