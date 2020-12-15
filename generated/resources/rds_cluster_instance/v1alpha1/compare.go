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
	k := kube.(*RdsClusterInstance)
	p := prov.(*RdsClusterInstance)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeRdsClusterInstance_ApplyImmediately(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_AvailabilityZone(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_CopyTagsToSnapshot(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_PreferredMaintenanceWindow(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_PubliclyAccessible(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_ClusterIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_IdentifierPrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_PreferredBackupWindow(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_DbParameterGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_Engine(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_PerformanceInsightsEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_CaCertIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_EngineVersion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_DbSubnetGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_MonitoringRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_AutoMinorVersionUpgrade(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_Identifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_MonitoringInterval(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_PerformanceInsightsKmsKeyId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_PromotionTier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_InstanceClass(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_KmsKeyId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_StorageEncrypted(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_Writer(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_Endpoint(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_Port(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_DbiResourceId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeRdsClusterInstance_ApplyImmediately(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.ApplyImmediately != p.ApplyImmediately {
		p.ApplyImmediately = k.ApplyImmediately
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_AvailabilityZone(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.AvailabilityZone != p.AvailabilityZone {
		p.AvailabilityZone = k.AvailabilityZone
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_CopyTagsToSnapshot(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.CopyTagsToSnapshot != p.CopyTagsToSnapshot {
		p.CopyTagsToSnapshot = k.CopyTagsToSnapshot
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_PreferredMaintenanceWindow(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PreferredMaintenanceWindow != p.PreferredMaintenanceWindow {
		p.PreferredMaintenanceWindow = k.PreferredMaintenanceWindow
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_PubliclyAccessible(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PubliclyAccessible != p.PubliclyAccessible {
		p.PubliclyAccessible = k.PubliclyAccessible
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_ClusterIdentifier(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.ClusterIdentifier != p.ClusterIdentifier {
		p.ClusterIdentifier = k.ClusterIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_IdentifierPrefix(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.IdentifierPrefix != p.IdentifierPrefix {
		p.IdentifierPrefix = k.IdentifierPrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_PreferredBackupWindow(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PreferredBackupWindow != p.PreferredBackupWindow {
		p.PreferredBackupWindow = k.PreferredBackupWindow
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_DbParameterGroupName(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.DbParameterGroupName != p.DbParameterGroupName {
		p.DbParameterGroupName = k.DbParameterGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_Engine(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.Engine != p.Engine {
		p.Engine = k.Engine
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_PerformanceInsightsEnabled(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PerformanceInsightsEnabled != p.PerformanceInsightsEnabled {
		p.PerformanceInsightsEnabled = k.PerformanceInsightsEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_CaCertIdentifier(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.CaCertIdentifier != p.CaCertIdentifier {
		p.CaCertIdentifier = k.CaCertIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_EngineVersion(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.EngineVersion != p.EngineVersion {
		p.EngineVersion = k.EngineVersion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_DbSubnetGroupName(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.DbSubnetGroupName != p.DbSubnetGroupName {
		p.DbSubnetGroupName = k.DbSubnetGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_MonitoringRoleArn(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.MonitoringRoleArn != p.MonitoringRoleArn {
		p.MonitoringRoleArn = k.MonitoringRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeRdsClusterInstance_Tags(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_AutoMinorVersionUpgrade(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.AutoMinorVersionUpgrade != p.AutoMinorVersionUpgrade {
		p.AutoMinorVersionUpgrade = k.AutoMinorVersionUpgrade
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_Identifier(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.Identifier != p.Identifier {
		p.Identifier = k.Identifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_MonitoringInterval(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.MonitoringInterval != p.MonitoringInterval {
		p.MonitoringInterval = k.MonitoringInterval
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_PerformanceInsightsKmsKeyId(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PerformanceInsightsKmsKeyId != p.PerformanceInsightsKmsKeyId {
		p.PerformanceInsightsKmsKeyId = k.PerformanceInsightsKmsKeyId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_PromotionTier(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.PromotionTier != p.PromotionTier {
		p.PromotionTier = k.PromotionTier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_InstanceClass(k *RdsClusterInstanceParameters, p *RdsClusterInstanceParameters, md *plugin.MergeDescription) bool {
	if k.InstanceClass != p.InstanceClass {
		p.InstanceClass = k.InstanceClass
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeRdsClusterInstance_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeRdsClusterInstance_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterInstance_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterInstance_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRdsClusterInstance_KmsKeyId(k *RdsClusterInstanceObservation, p *RdsClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.KmsKeyId != p.KmsKeyId {
		k.KmsKeyId = p.KmsKeyId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRdsClusterInstance_StorageEncrypted(k *RdsClusterInstanceObservation, p *RdsClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.StorageEncrypted != p.StorageEncrypted {
		k.StorageEncrypted = p.StorageEncrypted
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRdsClusterInstance_Writer(k *RdsClusterInstanceObservation, p *RdsClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Writer != p.Writer {
		k.Writer = p.Writer
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRdsClusterInstance_Endpoint(k *RdsClusterInstanceObservation, p *RdsClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Endpoint != p.Endpoint {
		k.Endpoint = p.Endpoint
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRdsClusterInstance_Port(k *RdsClusterInstanceObservation, p *RdsClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Port != p.Port {
		k.Port = p.Port
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRdsClusterInstance_DbiResourceId(k *RdsClusterInstanceObservation, p *RdsClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.DbiResourceId != p.DbiResourceId {
		k.DbiResourceId = p.DbiResourceId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRdsClusterInstance_Arn(k *RdsClusterInstanceObservation, p *RdsClusterInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}