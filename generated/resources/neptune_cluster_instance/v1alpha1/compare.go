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
	k := kube.(*NeptuneClusterInstance)
	p := prov.(*NeptuneClusterInstance)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeNeptuneClusterInstance_ApplyImmediately(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_AutoMinorVersionUpgrade(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_ClusterIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_Engine(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_EngineVersion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_NeptuneSubnetGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_PreferredBackupWindow(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_PromotionTier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_IdentifierPrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_InstanceClass(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_NeptuneParameterGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_Port(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_AvailabilityZone(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_Identifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_PreferredMaintenanceWindow(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_PubliclyAccessible(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_Address(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_DbiResourceId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_Endpoint(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_StorageEncrypted(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_Writer(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_KmsKeyArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeNeptuneClusterInstance_ApplyImmediately(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.ApplyImmediately != p.ApplyImmediately {
		p.ApplyImmediately = k.ApplyImmediately
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_AutoMinorVersionUpgrade(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.AutoMinorVersionUpgrade != p.AutoMinorVersionUpgrade {
		p.AutoMinorVersionUpgrade = k.AutoMinorVersionUpgrade
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_ClusterIdentifier(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.ClusterIdentifier != p.ClusterIdentifier {
		p.ClusterIdentifier = k.ClusterIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_Engine(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.Engine != p.Engine {
		p.Engine = k.Engine
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_EngineVersion(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.EngineVersion != p.EngineVersion {
		p.EngineVersion = k.EngineVersion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_NeptuneSubnetGroupName(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.NeptuneSubnetGroupName != p.NeptuneSubnetGroupName {
		p.NeptuneSubnetGroupName = k.NeptuneSubnetGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_PreferredBackupWindow(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PreferredBackupWindow != p.PreferredBackupWindow {
		p.PreferredBackupWindow = k.PreferredBackupWindow
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_PromotionTier(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PromotionTier != p.PromotionTier {
		p.PromotionTier = k.PromotionTier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_IdentifierPrefix(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.IdentifierPrefix != p.IdentifierPrefix {
		p.IdentifierPrefix = k.IdentifierPrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_InstanceClass(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.InstanceClass != p.InstanceClass {
		p.InstanceClass = k.InstanceClass
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_NeptuneParameterGroupName(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.NeptuneParameterGroupName != p.NeptuneParameterGroupName {
		p.NeptuneParameterGroupName = k.NeptuneParameterGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_Port(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.Port != p.Port {
		p.Port = k.Port
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_AvailabilityZone(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.AvailabilityZone != p.AvailabilityZone {
		p.AvailabilityZone = k.AvailabilityZone
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_Identifier(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.Identifier != p.Identifier {
		p.Identifier = k.Identifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_PreferredMaintenanceWindow(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PreferredMaintenanceWindow != p.PreferredMaintenanceWindow {
		p.PreferredMaintenanceWindow = k.PreferredMaintenanceWindow
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_PubliclyAccessible(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PubliclyAccessible != p.PubliclyAccessible {
		p.PubliclyAccessible = k.PubliclyAccessible
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeNeptuneClusterInstance_Tags(k *NeptuneClusterInstanceParameters, p *NeptuneClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeNeptuneClusterInstance_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeNeptuneClusterInstance_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneClusterInstance_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneClusterInstance_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterInstance_Address(k *NeptuneClusterInstanceObservation, p *NeptuneClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Address != p.Address {
		k.Address = p.Address
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterInstance_DbiResourceId(k *NeptuneClusterInstanceObservation, p *NeptuneClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.DbiResourceId != p.DbiResourceId {
		k.DbiResourceId = p.DbiResourceId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterInstance_Endpoint(k *NeptuneClusterInstanceObservation, p *NeptuneClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Endpoint != p.Endpoint {
		k.Endpoint = p.Endpoint
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterInstance_StorageEncrypted(k *NeptuneClusterInstanceObservation, p *NeptuneClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.StorageEncrypted != p.StorageEncrypted {
		k.StorageEncrypted = p.StorageEncrypted
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterInstance_Writer(k *NeptuneClusterInstanceObservation, p *NeptuneClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Writer != p.Writer {
		k.Writer = p.Writer
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterInstance_Arn(k *NeptuneClusterInstanceObservation, p *NeptuneClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneClusterInstance_KmsKeyArn(k *NeptuneClusterInstanceObservation, p *NeptuneClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.KmsKeyArn != p.KmsKeyArn {
		k.KmsKeyArn = p.KmsKeyArn
		md.StatusUpdated = true
		return true
	}
	return false
}