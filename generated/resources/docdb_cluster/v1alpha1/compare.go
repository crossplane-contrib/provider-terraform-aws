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
	k := kube.(*DocdbCluster)
	p := prov.(*DocdbCluster)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDocdbCluster_FinalSnapshotIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_MasterUsername(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_SnapshotIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_Port(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_AvailabilityZones(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_ClusterMembers(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_DbClusterParameterGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_DbSubnetGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_EngineVersion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_VpcSecurityGroupIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_ClusterIdentifierPrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_KmsKeyId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_PreferredBackupWindow(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_PreferredMaintenanceWindow(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_StorageEncrypted(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_EnabledCloudwatchLogsExports(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_MasterPassword(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_SkipFinalSnapshot(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_ApplyImmediately(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_DeletionProtection(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_ClusterIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_Engine(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_BackupRetentionPeriod(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_ClusterResourceId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_Endpoint(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_HostedZoneId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_ReaderEndpoint(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDocdbCluster_FinalSnapshotIdentifier(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.FinalSnapshotIdentifier != p.FinalSnapshotIdentifier {
		p.FinalSnapshotIdentifier = k.FinalSnapshotIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_MasterUsername(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.MasterUsername != p.MasterUsername {
		p.MasterUsername = k.MasterUsername
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_SnapshotIdentifier(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.SnapshotIdentifier != p.SnapshotIdentifier {
		p.SnapshotIdentifier = k.SnapshotIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_Port(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.Port != p.Port {
		p.Port = k.Port
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDocdbCluster_AvailabilityZones(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.AvailabilityZones, p.AvailabilityZones) {
		p.AvailabilityZones = k.AvailabilityZones
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDocdbCluster_ClusterMembers(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.ClusterMembers, p.ClusterMembers) {
		p.ClusterMembers = k.ClusterMembers
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_DbClusterParameterGroupName(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.DbClusterParameterGroupName != p.DbClusterParameterGroupName {
		p.DbClusterParameterGroupName = k.DbClusterParameterGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_DbSubnetGroupName(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.DbSubnetGroupName != p.DbSubnetGroupName {
		p.DbSubnetGroupName = k.DbSubnetGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_EngineVersion(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.EngineVersion != p.EngineVersion {
		p.EngineVersion = k.EngineVersion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDocdbCluster_Tags(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDocdbCluster_VpcSecurityGroupIds(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.VpcSecurityGroupIds, p.VpcSecurityGroupIds) {
		p.VpcSecurityGroupIds = k.VpcSecurityGroupIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_ClusterIdentifierPrefix(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.ClusterIdentifierPrefix != p.ClusterIdentifierPrefix {
		p.ClusterIdentifierPrefix = k.ClusterIdentifierPrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_KmsKeyId(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.KmsKeyId != p.KmsKeyId {
		p.KmsKeyId = k.KmsKeyId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_PreferredBackupWindow(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.PreferredBackupWindow != p.PreferredBackupWindow {
		p.PreferredBackupWindow = k.PreferredBackupWindow
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_PreferredMaintenanceWindow(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.PreferredMaintenanceWindow != p.PreferredMaintenanceWindow {
		p.PreferredMaintenanceWindow = k.PreferredMaintenanceWindow
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_StorageEncrypted(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.StorageEncrypted != p.StorageEncrypted {
		p.StorageEncrypted = k.StorageEncrypted
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDocdbCluster_EnabledCloudwatchLogsExports(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.EnabledCloudwatchLogsExports, p.EnabledCloudwatchLogsExports) {
		p.EnabledCloudwatchLogsExports = k.EnabledCloudwatchLogsExports
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_MasterPassword(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.MasterPassword != p.MasterPassword {
		p.MasterPassword = k.MasterPassword
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_SkipFinalSnapshot(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.SkipFinalSnapshot != p.SkipFinalSnapshot {
		p.SkipFinalSnapshot = k.SkipFinalSnapshot
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_ApplyImmediately(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.ApplyImmediately != p.ApplyImmediately {
		p.ApplyImmediately = k.ApplyImmediately
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_DeletionProtection(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.DeletionProtection != p.DeletionProtection {
		p.DeletionProtection = k.DeletionProtection
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_ClusterIdentifier(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.ClusterIdentifier != p.ClusterIdentifier {
		p.ClusterIdentifier = k.ClusterIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_Engine(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.Engine != p.Engine {
		p.Engine = k.Engine
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_BackupRetentionPeriod(k *DocdbClusterParameters, p *DocdbClusterParameters, md *plugin.MergeDescription) bool {
	if k.BackupRetentionPeriod != p.BackupRetentionPeriod {
		p.BackupRetentionPeriod = k.BackupRetentionPeriod
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDocdbCluster_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDocdbCluster_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbCluster_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbCluster_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbCluster_ClusterResourceId(k *DocdbClusterObservation, p *DocdbClusterObservation, md *plugin.MergeDescription) bool {
	if k.ClusterResourceId != p.ClusterResourceId {
		k.ClusterResourceId = p.ClusterResourceId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbCluster_Endpoint(k *DocdbClusterObservation, p *DocdbClusterObservation, md *plugin.MergeDescription) bool {
	if k.Endpoint != p.Endpoint {
		k.Endpoint = p.Endpoint
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbCluster_Arn(k *DocdbClusterObservation, p *DocdbClusterObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbCluster_HostedZoneId(k *DocdbClusterObservation, p *DocdbClusterObservation, md *plugin.MergeDescription) bool {
	if k.HostedZoneId != p.HostedZoneId {
		k.HostedZoneId = p.HostedZoneId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbCluster_ReaderEndpoint(k *DocdbClusterObservation, p *DocdbClusterObservation, md *plugin.MergeDescription) bool {
	if k.ReaderEndpoint != p.ReaderEndpoint {
		k.ReaderEndpoint = p.ReaderEndpoint
		md.StatusUpdated = true
		return true
	}
	return false
}