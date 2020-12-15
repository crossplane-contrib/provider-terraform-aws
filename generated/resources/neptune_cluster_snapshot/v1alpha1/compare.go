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
	k := kube.(*NeptuneClusterSnapshot)
	p := prov.(*NeptuneClusterSnapshot)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeNeptuneClusterSnapshot_DbClusterIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_DbClusterSnapshotIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_VpcId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_AllocatedStorage(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_LicenseModel(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_StorageEncrypted(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_Port(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_Status(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_DbClusterSnapshotArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_EngineVersion(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_KmsKeyId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_Engine(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_AvailabilityZones(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_SnapshotType(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterSnapshot_SourceDbClusterSnapshotArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeNeptuneClusterSnapshot_DbClusterIdentifier(k *NeptuneClusterSnapshotParameters, p *NeptuneClusterSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.DbClusterIdentifier != p.DbClusterIdentifier {
		p.DbClusterIdentifier = k.DbClusterIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterSnapshot_DbClusterSnapshotIdentifier(k *NeptuneClusterSnapshotParameters, p *NeptuneClusterSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.DbClusterSnapshotIdentifier != p.DbClusterSnapshotIdentifier {
		p.DbClusterSnapshotIdentifier = k.DbClusterSnapshotIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeNeptuneClusterSnapshot_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeNeptuneClusterSnapshot_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterSnapshot_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterSnapshot_VpcId(k *NeptuneClusterSnapshotObservation, p *NeptuneClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.VpcId != p.VpcId {
		k.VpcId = p.VpcId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterSnapshot_AllocatedStorage(k *NeptuneClusterSnapshotObservation, p *NeptuneClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.AllocatedStorage != p.AllocatedStorage {
		k.AllocatedStorage = p.AllocatedStorage
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterSnapshot_LicenseModel(k *NeptuneClusterSnapshotObservation, p *NeptuneClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.LicenseModel != p.LicenseModel {
		k.LicenseModel = p.LicenseModel
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterSnapshot_StorageEncrypted(k *NeptuneClusterSnapshotObservation, p *NeptuneClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.StorageEncrypted != p.StorageEncrypted {
		k.StorageEncrypted = p.StorageEncrypted
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterSnapshot_Port(k *NeptuneClusterSnapshotObservation, p *NeptuneClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.Port != p.Port {
		k.Port = p.Port
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterSnapshot_Status(k *NeptuneClusterSnapshotObservation, p *NeptuneClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.Status != p.Status {
		k.Status = p.Status
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterSnapshot_DbClusterSnapshotArn(k *NeptuneClusterSnapshotObservation, p *NeptuneClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.DbClusterSnapshotArn != p.DbClusterSnapshotArn {
		k.DbClusterSnapshotArn = p.DbClusterSnapshotArn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterSnapshot_EngineVersion(k *NeptuneClusterSnapshotObservation, p *NeptuneClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.EngineVersion != p.EngineVersion {
		k.EngineVersion = p.EngineVersion
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterSnapshot_KmsKeyId(k *NeptuneClusterSnapshotObservation, p *NeptuneClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.KmsKeyId != p.KmsKeyId {
		k.KmsKeyId = p.KmsKeyId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterSnapshot_Engine(k *NeptuneClusterSnapshotObservation, p *NeptuneClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.Engine != p.Engine {
		k.Engine = p.Engine
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeNeptuneClusterSnapshot_AvailabilityZones(k *NeptuneClusterSnapshotObservation, p *NeptuneClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.AvailabilityZones, p.AvailabilityZones) {
		k.AvailabilityZones = p.AvailabilityZones
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterSnapshot_SnapshotType(k *NeptuneClusterSnapshotObservation, p *NeptuneClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.SnapshotType != p.SnapshotType {
		k.SnapshotType = p.SnapshotType
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterSnapshot_SourceDbClusterSnapshotArn(k *NeptuneClusterSnapshotObservation, p *NeptuneClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.SourceDbClusterSnapshotArn != p.SourceDbClusterSnapshotArn {
		k.SourceDbClusterSnapshotArn = p.SourceDbClusterSnapshotArn
		md.StatusUpdated = true
		return true
	}
	return false
}