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
	"fmt"
	
	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*RdsCluster)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a RdsCluster.")
	}
	return EncodeRdsCluster(*r), nil
}

func EncodeRdsCluster(r RdsCluster) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRdsCluster_ClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_ClusterIdentifierPrefix(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_EnabledCloudwatchLogsExports(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_GlobalClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_IamDatabaseAuthenticationEnabled(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_AvailabilityZones(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_CopyTagsToSnapshot(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_DatabaseName(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_DeletionProtection(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_SkipFinalSnapshot(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_BackupRetentionPeriod(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_DbSubnetGroupName(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_EnableHttpEndpoint(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_EngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_PreferredBackupWindow(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_PreferredMaintenanceWindow(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_StorageEncrypted(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_ApplyImmediately(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_IamRoles(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_Id(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_ReplicationSourceIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_SnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_Engine(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_DbClusterParameterGroupName(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_Port(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_ClusterMembers(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_VpcSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_EngineMode(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_MasterUsername(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_SourceRegion(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_BacktrackWindow(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_FinalSnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_MasterPassword(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_AllowMajorVersionUpgrade(r.Spec.ForProvider, ctyVal)
	EncodeRdsCluster_ScalingConfiguration(r.Spec.ForProvider.ScalingConfiguration, ctyVal)
	EncodeRdsCluster_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeRdsCluster_S3Import(r.Spec.ForProvider.S3Import, ctyVal)
	EncodeRdsCluster_ReaderEndpoint(r.Status.AtProvider, ctyVal)
	EncodeRdsCluster_Endpoint(r.Status.AtProvider, ctyVal)
	EncodeRdsCluster_ClusterResourceId(r.Status.AtProvider, ctyVal)
	EncodeRdsCluster_HostedZoneId(r.Status.AtProvider, ctyVal)
	EncodeRdsCluster_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeRdsCluster_ClusterIdentifier(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["cluster_identifier"] = cty.StringVal(p.ClusterIdentifier)
}

func EncodeRdsCluster_ClusterIdentifierPrefix(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["cluster_identifier_prefix"] = cty.StringVal(p.ClusterIdentifierPrefix)
}

func EncodeRdsCluster_EnabledCloudwatchLogsExports(p RdsClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.EnabledCloudwatchLogsExports {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["enabled_cloudwatch_logs_exports"] = cty.SetVal(colVals)
}

func EncodeRdsCluster_GlobalClusterIdentifier(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["global_cluster_identifier"] = cty.StringVal(p.GlobalClusterIdentifier)
}

func EncodeRdsCluster_IamDatabaseAuthenticationEnabled(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["iam_database_authentication_enabled"] = cty.BoolVal(p.IamDatabaseAuthenticationEnabled)
}

func EncodeRdsCluster_AvailabilityZones(p RdsClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.SetVal(colVals)
}

func EncodeRdsCluster_CopyTagsToSnapshot(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["copy_tags_to_snapshot"] = cty.BoolVal(p.CopyTagsToSnapshot)
}

func EncodeRdsCluster_DatabaseName(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["database_name"] = cty.StringVal(p.DatabaseName)
}

func EncodeRdsCluster_DeletionProtection(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["deletion_protection"] = cty.BoolVal(p.DeletionProtection)
}

func EncodeRdsCluster_SkipFinalSnapshot(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["skip_final_snapshot"] = cty.BoolVal(p.SkipFinalSnapshot)
}

func EncodeRdsCluster_BackupRetentionPeriod(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["backup_retention_period"] = cty.NumberIntVal(p.BackupRetentionPeriod)
}

func EncodeRdsCluster_DbSubnetGroupName(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["db_subnet_group_name"] = cty.StringVal(p.DbSubnetGroupName)
}

func EncodeRdsCluster_EnableHttpEndpoint(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["enable_http_endpoint"] = cty.BoolVal(p.EnableHttpEndpoint)
}

func EncodeRdsCluster_EngineVersion(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeRdsCluster_PreferredBackupWindow(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["preferred_backup_window"] = cty.StringVal(p.PreferredBackupWindow)
}

func EncodeRdsCluster_PreferredMaintenanceWindow(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["preferred_maintenance_window"] = cty.StringVal(p.PreferredMaintenanceWindow)
}

func EncodeRdsCluster_StorageEncrypted(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["storage_encrypted"] = cty.BoolVal(p.StorageEncrypted)
}

func EncodeRdsCluster_ApplyImmediately(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["apply_immediately"] = cty.BoolVal(p.ApplyImmediately)
}

func EncodeRdsCluster_IamRoles(p RdsClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.IamRoles {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["iam_roles"] = cty.SetVal(colVals)
}

func EncodeRdsCluster_Id(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRdsCluster_ReplicationSourceIdentifier(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["replication_source_identifier"] = cty.StringVal(p.ReplicationSourceIdentifier)
}

func EncodeRdsCluster_SnapshotIdentifier(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["snapshot_identifier"] = cty.StringVal(p.SnapshotIdentifier)
}

func EncodeRdsCluster_Engine(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeRdsCluster_DbClusterParameterGroupName(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["db_cluster_parameter_group_name"] = cty.StringVal(p.DbClusterParameterGroupName)
}

func EncodeRdsCluster_Port(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeRdsCluster_ClusterMembers(p RdsClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ClusterMembers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["cluster_members"] = cty.SetVal(colVals)
}

func EncodeRdsCluster_KmsKeyId(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeRdsCluster_VpcSecurityGroupIds(p RdsClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeRdsCluster_EngineMode(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["engine_mode"] = cty.StringVal(p.EngineMode)
}

func EncodeRdsCluster_MasterUsername(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["master_username"] = cty.StringVal(p.MasterUsername)
}

func EncodeRdsCluster_SourceRegion(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["source_region"] = cty.StringVal(p.SourceRegion)
}

func EncodeRdsCluster_BacktrackWindow(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["backtrack_window"] = cty.NumberIntVal(p.BacktrackWindow)
}

func EncodeRdsCluster_FinalSnapshotIdentifier(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["final_snapshot_identifier"] = cty.StringVal(p.FinalSnapshotIdentifier)
}

func EncodeRdsCluster_MasterPassword(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["master_password"] = cty.StringVal(p.MasterPassword)
}

func EncodeRdsCluster_Tags(p RdsClusterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRdsCluster_AllowMajorVersionUpgrade(p RdsClusterParameters, vals map[string]cty.Value) {
	vals["allow_major_version_upgrade"] = cty.BoolVal(p.AllowMajorVersionUpgrade)
}

func EncodeRdsCluster_ScalingConfiguration(p ScalingConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeRdsCluster_ScalingConfiguration_AutoPause(p, ctyVal)
	EncodeRdsCluster_ScalingConfiguration_MaxCapacity(p, ctyVal)
	EncodeRdsCluster_ScalingConfiguration_MinCapacity(p, ctyVal)
	EncodeRdsCluster_ScalingConfiguration_SecondsUntilAutoPause(p, ctyVal)
	EncodeRdsCluster_ScalingConfiguration_TimeoutAction(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["scaling_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeRdsCluster_ScalingConfiguration_AutoPause(p ScalingConfiguration, vals map[string]cty.Value) {
	vals["auto_pause"] = cty.BoolVal(p.AutoPause)
}

func EncodeRdsCluster_ScalingConfiguration_MaxCapacity(p ScalingConfiguration, vals map[string]cty.Value) {
	vals["max_capacity"] = cty.NumberIntVal(p.MaxCapacity)
}

func EncodeRdsCluster_ScalingConfiguration_MinCapacity(p ScalingConfiguration, vals map[string]cty.Value) {
	vals["min_capacity"] = cty.NumberIntVal(p.MinCapacity)
}

func EncodeRdsCluster_ScalingConfiguration_SecondsUntilAutoPause(p ScalingConfiguration, vals map[string]cty.Value) {
	vals["seconds_until_auto_pause"] = cty.NumberIntVal(p.SecondsUntilAutoPause)
}

func EncodeRdsCluster_ScalingConfiguration_TimeoutAction(p ScalingConfiguration, vals map[string]cty.Value) {
	vals["timeout_action"] = cty.StringVal(p.TimeoutAction)
}

func EncodeRdsCluster_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeRdsCluster_Timeouts_Create(p, ctyVal)
	EncodeRdsCluster_Timeouts_Delete(p, ctyVal)
	EncodeRdsCluster_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeRdsCluster_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeRdsCluster_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeRdsCluster_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeRdsCluster_S3Import(p S3Import, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeRdsCluster_S3Import_IngestionRole(p, ctyVal)
	EncodeRdsCluster_S3Import_SourceEngine(p, ctyVal)
	EncodeRdsCluster_S3Import_SourceEngineVersion(p, ctyVal)
	EncodeRdsCluster_S3Import_BucketName(p, ctyVal)
	EncodeRdsCluster_S3Import_BucketPrefix(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["s3_import"] = cty.ListVal(valsForCollection)
}

func EncodeRdsCluster_S3Import_IngestionRole(p S3Import, vals map[string]cty.Value) {
	vals["ingestion_role"] = cty.StringVal(p.IngestionRole)
}

func EncodeRdsCluster_S3Import_SourceEngine(p S3Import, vals map[string]cty.Value) {
	vals["source_engine"] = cty.StringVal(p.SourceEngine)
}

func EncodeRdsCluster_S3Import_SourceEngineVersion(p S3Import, vals map[string]cty.Value) {
	vals["source_engine_version"] = cty.StringVal(p.SourceEngineVersion)
}

func EncodeRdsCluster_S3Import_BucketName(p S3Import, vals map[string]cty.Value) {
	vals["bucket_name"] = cty.StringVal(p.BucketName)
}

func EncodeRdsCluster_S3Import_BucketPrefix(p S3Import, vals map[string]cty.Value) {
	vals["bucket_prefix"] = cty.StringVal(p.BucketPrefix)
}

func EncodeRdsCluster_ReaderEndpoint(p RdsClusterObservation, vals map[string]cty.Value) {
	vals["reader_endpoint"] = cty.StringVal(p.ReaderEndpoint)
}

func EncodeRdsCluster_Endpoint(p RdsClusterObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeRdsCluster_ClusterResourceId(p RdsClusterObservation, vals map[string]cty.Value) {
	vals["cluster_resource_id"] = cty.StringVal(p.ClusterResourceId)
}

func EncodeRdsCluster_HostedZoneId(p RdsClusterObservation, vals map[string]cty.Value) {
	vals["hosted_zone_id"] = cty.StringVal(p.HostedZoneId)
}

func EncodeRdsCluster_Arn(p RdsClusterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}