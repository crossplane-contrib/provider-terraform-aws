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
	k := kube.(*DocdbClusterInstance)
	p := prov.(*DocdbClusterInstance)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDocdbClusterInstance_IdentifierPrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_PreferredMaintenanceWindow(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_CaCertIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_ClusterIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_PromotionTier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_ApplyImmediately(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_AvailabilityZone(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_Identifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_Engine(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_InstanceClass(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_AutoMinorVersionUpgrade(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_KmsKeyId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_Port(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_PubliclyAccessible(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_EngineVersion(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_Writer(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_Endpoint(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_PreferredBackupWindow(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_StorageEncrypted(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_DbSubnetGroupName(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_DbiResourceId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDocdbClusterInstance_IdentifierPrefix(k *DocdbClusterInstanceParameters, p *DocdbClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.IdentifierPrefix != p.IdentifierPrefix {
		p.IdentifierPrefix = k.IdentifierPrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterInstance_PreferredMaintenanceWindow(k *DocdbClusterInstanceParameters, p *DocdbClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PreferredMaintenanceWindow != p.PreferredMaintenanceWindow {
		p.PreferredMaintenanceWindow = k.PreferredMaintenanceWindow
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterInstance_CaCertIdentifier(k *DocdbClusterInstanceParameters, p *DocdbClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.CaCertIdentifier != p.CaCertIdentifier {
		p.CaCertIdentifier = k.CaCertIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterInstance_ClusterIdentifier(k *DocdbClusterInstanceParameters, p *DocdbClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.ClusterIdentifier != p.ClusterIdentifier {
		p.ClusterIdentifier = k.ClusterIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterInstance_PromotionTier(k *DocdbClusterInstanceParameters, p *DocdbClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PromotionTier != p.PromotionTier {
		p.PromotionTier = k.PromotionTier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterInstance_ApplyImmediately(k *DocdbClusterInstanceParameters, p *DocdbClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.ApplyImmediately != p.ApplyImmediately {
		p.ApplyImmediately = k.ApplyImmediately
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterInstance_AvailabilityZone(k *DocdbClusterInstanceParameters, p *DocdbClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.AvailabilityZone != p.AvailabilityZone {
		p.AvailabilityZone = k.AvailabilityZone
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterInstance_Identifier(k *DocdbClusterInstanceParameters, p *DocdbClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.Identifier != p.Identifier {
		p.Identifier = k.Identifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDocdbClusterInstance_Tags(k *DocdbClusterInstanceParameters, p *DocdbClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterInstance_Engine(k *DocdbClusterInstanceParameters, p *DocdbClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.Engine != p.Engine {
		p.Engine = k.Engine
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterInstance_InstanceClass(k *DocdbClusterInstanceParameters, p *DocdbClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.InstanceClass != p.InstanceClass {
		p.InstanceClass = k.InstanceClass
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterInstance_AutoMinorVersionUpgrade(k *DocdbClusterInstanceParameters, p *DocdbClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.AutoMinorVersionUpgrade != p.AutoMinorVersionUpgrade {
		p.AutoMinorVersionUpgrade = k.AutoMinorVersionUpgrade
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDocdbClusterInstance_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDocdbClusterInstance_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterInstance_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterInstance_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterInstance_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterInstance_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterInstance_KmsKeyId(k *DocdbClusterInstanceObservation, p *DocdbClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.KmsKeyId != p.KmsKeyId {
		k.KmsKeyId = p.KmsKeyId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterInstance_Port(k *DocdbClusterInstanceObservation, p *DocdbClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Port != p.Port {
		k.Port = p.Port
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterInstance_PubliclyAccessible(k *DocdbClusterInstanceObservation, p *DocdbClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.PubliclyAccessible != p.PubliclyAccessible {
		k.PubliclyAccessible = p.PubliclyAccessible
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterInstance_EngineVersion(k *DocdbClusterInstanceObservation, p *DocdbClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.EngineVersion != p.EngineVersion {
		k.EngineVersion = p.EngineVersion
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterInstance_Writer(k *DocdbClusterInstanceObservation, p *DocdbClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Writer != p.Writer {
		k.Writer = p.Writer
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterInstance_Arn(k *DocdbClusterInstanceObservation, p *DocdbClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterInstance_Endpoint(k *DocdbClusterInstanceObservation, p *DocdbClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Endpoint != p.Endpoint {
		k.Endpoint = p.Endpoint
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterInstance_PreferredBackupWindow(k *DocdbClusterInstanceObservation, p *DocdbClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.PreferredBackupWindow != p.PreferredBackupWindow {
		k.PreferredBackupWindow = p.PreferredBackupWindow
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterInstance_StorageEncrypted(k *DocdbClusterInstanceObservation, p *DocdbClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.StorageEncrypted != p.StorageEncrypted {
		k.StorageEncrypted = p.StorageEncrypted
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterInstance_DbSubnetGroupName(k *DocdbClusterInstanceObservation, p *DocdbClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.DbSubnetGroupName != p.DbSubnetGroupName {
		k.DbSubnetGroupName = p.DbSubnetGroupName
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterInstance_DbiResourceId(k *DocdbClusterInstanceObservation, p *DocdbClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.DbiResourceId != p.DbiResourceId {
		k.DbiResourceId = p.DbiResourceId
		md.StatusUpdated = true
		return true
	}
	return false
}