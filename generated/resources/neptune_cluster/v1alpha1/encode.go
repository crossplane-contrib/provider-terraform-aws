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
	"github.com/zclconf/go-cty/cty"
)

func EncodeNeptuneCluster(r NeptuneCluster) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeNeptuneCluster_SkipFinalSnapshot(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_EnableCloudwatchLogsExports(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_EngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_Port(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_PreferredBackupWindow(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_AvailabilityZones(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_ClusterIdentifierPrefix(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_DeletionProtection(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_PreferredMaintenanceWindow(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_StorageEncrypted(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_ApplyImmediately(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_FinalSnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_IamDatabaseAuthenticationEnabled(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_NeptuneSubnetGroupName(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_Id(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_ReplicationSourceIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_Tags(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_VpcSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_BackupRetentionPeriod(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_Engine(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_ClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_KmsKeyArn(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_NeptuneClusterParameterGroupName(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_IamRoles(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_SnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneCluster_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeNeptuneCluster_ClusterResourceId(r.Status.AtProvider, ctyVal)
	EncodeNeptuneCluster_HostedZoneId(r.Status.AtProvider, ctyVal)
	EncodeNeptuneCluster_ClusterMembers(r.Status.AtProvider, ctyVal)
	EncodeNeptuneCluster_Arn(r.Status.AtProvider, ctyVal)
	EncodeNeptuneCluster_ReaderEndpoint(r.Status.AtProvider, ctyVal)
	EncodeNeptuneCluster_Endpoint(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeNeptuneCluster_SkipFinalSnapshot(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["skip_final_snapshot"] = cty.BoolVal(p.SkipFinalSnapshot)
}

func EncodeNeptuneCluster_EnableCloudwatchLogsExports(p NeptuneClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.EnableCloudwatchLogsExports {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["enable_cloudwatch_logs_exports"] = cty.SetVal(colVals)
}

func EncodeNeptuneCluster_EngineVersion(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeNeptuneCluster_Port(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeNeptuneCluster_PreferredBackupWindow(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["preferred_backup_window"] = cty.StringVal(p.PreferredBackupWindow)
}

func EncodeNeptuneCluster_AvailabilityZones(p NeptuneClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.SetVal(colVals)
}

func EncodeNeptuneCluster_ClusterIdentifierPrefix(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["cluster_identifier_prefix"] = cty.StringVal(p.ClusterIdentifierPrefix)
}

func EncodeNeptuneCluster_DeletionProtection(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["deletion_protection"] = cty.BoolVal(p.DeletionProtection)
}

func EncodeNeptuneCluster_PreferredMaintenanceWindow(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["preferred_maintenance_window"] = cty.StringVal(p.PreferredMaintenanceWindow)
}

func EncodeNeptuneCluster_StorageEncrypted(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["storage_encrypted"] = cty.BoolVal(p.StorageEncrypted)
}

func EncodeNeptuneCluster_ApplyImmediately(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["apply_immediately"] = cty.BoolVal(p.ApplyImmediately)
}

func EncodeNeptuneCluster_FinalSnapshotIdentifier(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["final_snapshot_identifier"] = cty.StringVal(p.FinalSnapshotIdentifier)
}

func EncodeNeptuneCluster_IamDatabaseAuthenticationEnabled(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["iam_database_authentication_enabled"] = cty.BoolVal(p.IamDatabaseAuthenticationEnabled)
}

func EncodeNeptuneCluster_NeptuneSubnetGroupName(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["neptune_subnet_group_name"] = cty.StringVal(p.NeptuneSubnetGroupName)
}

func EncodeNeptuneCluster_Id(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeNeptuneCluster_ReplicationSourceIdentifier(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["replication_source_identifier"] = cty.StringVal(p.ReplicationSourceIdentifier)
}

func EncodeNeptuneCluster_Tags(p NeptuneClusterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeNeptuneCluster_VpcSecurityGroupIds(p NeptuneClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeNeptuneCluster_BackupRetentionPeriod(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["backup_retention_period"] = cty.NumberIntVal(p.BackupRetentionPeriod)
}

func EncodeNeptuneCluster_Engine(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeNeptuneCluster_ClusterIdentifier(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["cluster_identifier"] = cty.StringVal(p.ClusterIdentifier)
}

func EncodeNeptuneCluster_KmsKeyArn(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeNeptuneCluster_NeptuneClusterParameterGroupName(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["neptune_cluster_parameter_group_name"] = cty.StringVal(p.NeptuneClusterParameterGroupName)
}

func EncodeNeptuneCluster_IamRoles(p NeptuneClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.IamRoles {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["iam_roles"] = cty.SetVal(colVals)
}

func EncodeNeptuneCluster_SnapshotIdentifier(p NeptuneClusterParameters, vals map[string]cty.Value) {
	vals["snapshot_identifier"] = cty.StringVal(p.SnapshotIdentifier)
}

func EncodeNeptuneCluster_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeNeptuneCluster_Timeouts_Create(p, ctyVal)
	EncodeNeptuneCluster_Timeouts_Delete(p, ctyVal)
	EncodeNeptuneCluster_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeNeptuneCluster_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeNeptuneCluster_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeNeptuneCluster_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeNeptuneCluster_ClusterResourceId(p NeptuneClusterObservation, vals map[string]cty.Value) {
	vals["cluster_resource_id"] = cty.StringVal(p.ClusterResourceId)
}

func EncodeNeptuneCluster_HostedZoneId(p NeptuneClusterObservation, vals map[string]cty.Value) {
	vals["hosted_zone_id"] = cty.StringVal(p.HostedZoneId)
}

func EncodeNeptuneCluster_ClusterMembers(p NeptuneClusterObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ClusterMembers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["cluster_members"] = cty.SetVal(colVals)
}

func EncodeNeptuneCluster_Arn(p NeptuneClusterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeNeptuneCluster_ReaderEndpoint(p NeptuneClusterObservation, vals map[string]cty.Value) {
	vals["reader_endpoint"] = cty.StringVal(p.ReaderEndpoint)
}

func EncodeNeptuneCluster_Endpoint(p NeptuneClusterObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}