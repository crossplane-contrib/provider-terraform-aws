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
	k := kube.(*DmsReplicationInstance)
	p := prov.(*DmsReplicationInstance)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDmsReplicationInstance_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_VpcSecurityGroupIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_AutoMinorVersionUpgrade(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_EngineVersion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_MultiAz(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_AvailabilityZone(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_PreferredMaintenanceWindow(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_ReplicationSubnetGroupId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_AllowMajorVersionUpgrade(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_ApplyImmediately(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_KmsKeyArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_PubliclyAccessible(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_ReplicationInstanceId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_AllocatedStorage(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_ReplicationInstanceClass(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_ReplicationInstancePrivateIps(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_ReplicationInstanceArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_ReplicationInstancePublicIps(&k.Status.AtProvider, &p.Status.AtProvider, md)
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

//mergePrimitiveContainerTemplateSpec
func MergeDmsReplicationInstance_Tags(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDmsReplicationInstance_VpcSecurityGroupIds(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.VpcSecurityGroupIds, p.VpcSecurityGroupIds) {
		p.VpcSecurityGroupIds = k.VpcSecurityGroupIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_AutoMinorVersionUpgrade(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if k.AutoMinorVersionUpgrade != p.AutoMinorVersionUpgrade {
		p.AutoMinorVersionUpgrade = k.AutoMinorVersionUpgrade
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_EngineVersion(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if k.EngineVersion != p.EngineVersion {
		p.EngineVersion = k.EngineVersion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_MultiAz(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if k.MultiAz != p.MultiAz {
		p.MultiAz = k.MultiAz
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_AvailabilityZone(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if k.AvailabilityZone != p.AvailabilityZone {
		p.AvailabilityZone = k.AvailabilityZone
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_PreferredMaintenanceWindow(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PreferredMaintenanceWindow != p.PreferredMaintenanceWindow {
		p.PreferredMaintenanceWindow = k.PreferredMaintenanceWindow
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_ReplicationSubnetGroupId(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if k.ReplicationSubnetGroupId != p.ReplicationSubnetGroupId {
		p.ReplicationSubnetGroupId = k.ReplicationSubnetGroupId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_AllowMajorVersionUpgrade(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if k.AllowMajorVersionUpgrade != p.AllowMajorVersionUpgrade {
		p.AllowMajorVersionUpgrade = k.AllowMajorVersionUpgrade
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_ApplyImmediately(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if k.ApplyImmediately != p.ApplyImmediately {
		p.ApplyImmediately = k.ApplyImmediately
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_KmsKeyArn(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if k.KmsKeyArn != p.KmsKeyArn {
		p.KmsKeyArn = k.KmsKeyArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_PubliclyAccessible(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PubliclyAccessible != p.PubliclyAccessible {
		p.PubliclyAccessible = k.PubliclyAccessible
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_ReplicationInstanceId(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if k.ReplicationInstanceId != p.ReplicationInstanceId {
		p.ReplicationInstanceId = k.ReplicationInstanceId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_AllocatedStorage(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if k.AllocatedStorage != p.AllocatedStorage {
		p.AllocatedStorage = k.AllocatedStorage
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_ReplicationInstanceClass(k *DmsReplicationInstanceParameters, p *DmsReplicationInstanceParameters, md *plugin.MergeDescription) bool {
	if k.ReplicationInstanceClass != p.ReplicationInstanceClass {
		p.ReplicationInstanceClass = k.ReplicationInstanceClass
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDmsReplicationInstance_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDmsReplicationInstance_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationInstance_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationInstance_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeDmsReplicationInstance_ReplicationInstancePrivateIps(k *DmsReplicationInstanceObservation, p *DmsReplicationInstanceObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.ReplicationInstancePrivateIps, p.ReplicationInstancePrivateIps) {
		k.ReplicationInstancePrivateIps = p.ReplicationInstancePrivateIps
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDmsReplicationInstance_ReplicationInstanceArn(k *DmsReplicationInstanceObservation, p *DmsReplicationInstanceObservation, md *plugin.MergeDescription) bool {
	if k.ReplicationInstanceArn != p.ReplicationInstanceArn {
		k.ReplicationInstanceArn = p.ReplicationInstanceArn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeDmsReplicationInstance_ReplicationInstancePublicIps(k *DmsReplicationInstanceObservation, p *DmsReplicationInstanceObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.ReplicationInstancePublicIps, p.ReplicationInstancePublicIps) {
		k.ReplicationInstancePublicIps = p.ReplicationInstancePublicIps
		md.StatusUpdated = true
		return true
	}
	return false
}