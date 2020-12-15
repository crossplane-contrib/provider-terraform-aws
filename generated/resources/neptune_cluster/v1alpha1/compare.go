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
	k := kube.(*NeptuneCluster)
	p := prov.(*NeptuneCluster)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeNeptuneCluster_KmsKeyArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_EngineVersion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_PreferredBackupWindow(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_ReplicationSourceIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_IamRoles(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_NeptuneSubnetGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_PreferredMaintenanceWindow(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_SnapshotIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_VpcSecurityGroupIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_DeletionProtection(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_Engine(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_IamDatabaseAuthenticationEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_StorageEncrypted(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_ApplyImmediately(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_SkipFinalSnapshot(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_AvailabilityZones(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_BackupRetentionPeriod(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_ClusterIdentifierPrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_EnableCloudwatchLogsExports(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_NeptuneClusterParameterGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_ClusterIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_FinalSnapshotIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_Port(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_HostedZoneId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_ReaderEndpoint(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_ClusterMembers(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_ClusterResourceId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_Endpoint(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeNeptuneCluster_KmsKeyArn(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.KmsKeyArn != p.KmsKeyArn {
		p.KmsKeyArn = k.KmsKeyArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_EngineVersion(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.EngineVersion != p.EngineVersion {
		p.EngineVersion = k.EngineVersion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_PreferredBackupWindow(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.PreferredBackupWindow != p.PreferredBackupWindow {
		p.PreferredBackupWindow = k.PreferredBackupWindow
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_ReplicationSourceIdentifier(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.ReplicationSourceIdentifier != p.ReplicationSourceIdentifier {
		p.ReplicationSourceIdentifier = k.ReplicationSourceIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeNeptuneCluster_IamRoles(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.IamRoles, p.IamRoles) {
		p.IamRoles = k.IamRoles
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_NeptuneSubnetGroupName(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.NeptuneSubnetGroupName != p.NeptuneSubnetGroupName {
		p.NeptuneSubnetGroupName = k.NeptuneSubnetGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_PreferredMaintenanceWindow(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.PreferredMaintenanceWindow != p.PreferredMaintenanceWindow {
		p.PreferredMaintenanceWindow = k.PreferredMaintenanceWindow
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_SnapshotIdentifier(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.SnapshotIdentifier != p.SnapshotIdentifier {
		p.SnapshotIdentifier = k.SnapshotIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeNeptuneCluster_VpcSecurityGroupIds(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.VpcSecurityGroupIds, p.VpcSecurityGroupIds) {
		p.VpcSecurityGroupIds = k.VpcSecurityGroupIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_DeletionProtection(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.DeletionProtection != p.DeletionProtection {
		p.DeletionProtection = k.DeletionProtection
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_Engine(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.Engine != p.Engine {
		p.Engine = k.Engine
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_IamDatabaseAuthenticationEnabled(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.IamDatabaseAuthenticationEnabled != p.IamDatabaseAuthenticationEnabled {
		p.IamDatabaseAuthenticationEnabled = k.IamDatabaseAuthenticationEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_StorageEncrypted(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.StorageEncrypted != p.StorageEncrypted {
		p.StorageEncrypted = k.StorageEncrypted
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_ApplyImmediately(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.ApplyImmediately != p.ApplyImmediately {
		p.ApplyImmediately = k.ApplyImmediately
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_SkipFinalSnapshot(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.SkipFinalSnapshot != p.SkipFinalSnapshot {
		p.SkipFinalSnapshot = k.SkipFinalSnapshot
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeNeptuneCluster_Tags(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeNeptuneCluster_AvailabilityZones(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.AvailabilityZones, p.AvailabilityZones) {
		p.AvailabilityZones = k.AvailabilityZones
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_BackupRetentionPeriod(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.BackupRetentionPeriod != p.BackupRetentionPeriod {
		p.BackupRetentionPeriod = k.BackupRetentionPeriod
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_ClusterIdentifierPrefix(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.ClusterIdentifierPrefix != p.ClusterIdentifierPrefix {
		p.ClusterIdentifierPrefix = k.ClusterIdentifierPrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeNeptuneCluster_EnableCloudwatchLogsExports(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.EnableCloudwatchLogsExports, p.EnableCloudwatchLogsExports) {
		p.EnableCloudwatchLogsExports = k.EnableCloudwatchLogsExports
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_NeptuneClusterParameterGroupName(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.NeptuneClusterParameterGroupName != p.NeptuneClusterParameterGroupName {
		p.NeptuneClusterParameterGroupName = k.NeptuneClusterParameterGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_ClusterIdentifier(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.ClusterIdentifier != p.ClusterIdentifier {
		p.ClusterIdentifier = k.ClusterIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_FinalSnapshotIdentifier(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.FinalSnapshotIdentifier != p.FinalSnapshotIdentifier {
		p.FinalSnapshotIdentifier = k.FinalSnapshotIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_Port(k *NeptuneClusterParameters, p *NeptuneClusterParameters, md *plugin.MergeDescription) bool {
	if k.Port != p.Port {
		p.Port = k.Port
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeNeptuneCluster_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeNeptuneCluster_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneCluster_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneCluster_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneCluster_Arn(k *NeptuneClusterObservation, p *NeptuneClusterObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneCluster_HostedZoneId(k *NeptuneClusterObservation, p *NeptuneClusterObservation, md *plugin.MergeDescription) bool {
	if k.HostedZoneId != p.HostedZoneId {
		k.HostedZoneId = p.HostedZoneId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneCluster_ReaderEndpoint(k *NeptuneClusterObservation, p *NeptuneClusterObservation, md *plugin.MergeDescription) bool {
	if k.ReaderEndpoint != p.ReaderEndpoint {
		k.ReaderEndpoint = p.ReaderEndpoint
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeNeptuneCluster_ClusterMembers(k *NeptuneClusterObservation, p *NeptuneClusterObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.ClusterMembers, p.ClusterMembers) {
		k.ClusterMembers = p.ClusterMembers
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneCluster_ClusterResourceId(k *NeptuneClusterObservation, p *NeptuneClusterObservation, md *plugin.MergeDescription) bool {
	if k.ClusterResourceId != p.ClusterResourceId {
		k.ClusterResourceId = p.ClusterResourceId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneCluster_Endpoint(k *NeptuneClusterObservation, p *NeptuneClusterObservation, md *plugin.MergeDescription) bool {
	if k.Endpoint != p.Endpoint {
		k.Endpoint = p.Endpoint
		md.StatusUpdated = true
		return true
	}
	return false
}