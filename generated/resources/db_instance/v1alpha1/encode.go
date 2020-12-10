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
	r, ok := mr.(*DbInstance)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DbInstance.")
	}
	return EncodeDbInstance(*r), nil
}

func EncodeDbInstance(r DbInstance) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDbInstance_IamDatabaseAuthenticationEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_ParameterGroupName(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_Identifier(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_InstanceClass(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_PubliclyAccessible(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_StorageEncrypted(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_Port(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_SecurityGroupNames(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_Id(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_ReplicateSourceDb(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_Timezone(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_AutoMinorVersionUpgrade(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_CaCertIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_SnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_CharacterSetName(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_DomainIamRoleName(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_MaxAllocatedStorage(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_AllowMajorVersionUpgrade(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_EngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_LicenseModel(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_DeletionProtection(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_FinalSnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_MaintenanceWindow(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_Name(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_Password(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_Username(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_CopyTagsToSnapshot(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_Iops(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_PerformanceInsightsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_PerformanceInsightsKmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_VpcSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_StorageType(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_BackupRetentionPeriod(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_Domain(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_IdentifierPrefix(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_OptionGroupName(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_SkipFinalSnapshot(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_MonitoringRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_AllocatedStorage(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_BackupWindow(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_DbSubnetGroupName(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_DeleteAutomatedBackups(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_Engine(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_MonitoringInterval(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_ApplyImmediately(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_EnabledCloudwatchLogsExports(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_MultiAz(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_PerformanceInsightsRetentionPeriod(r.Spec.ForProvider, ctyVal)
	EncodeDbInstance_S3Import(r.Spec.ForProvider.S3Import, ctyVal)
	EncodeDbInstance_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDbInstance_Address(r.Status.AtProvider, ctyVal)
	EncodeDbInstance_ResourceId(r.Status.AtProvider, ctyVal)
	EncodeDbInstance_Status(r.Status.AtProvider, ctyVal)
	EncodeDbInstance_HostedZoneId(r.Status.AtProvider, ctyVal)
	EncodeDbInstance_Arn(r.Status.AtProvider, ctyVal)
	EncodeDbInstance_Endpoint(r.Status.AtProvider, ctyVal)
	EncodeDbInstance_Replicas(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDbInstance_IamDatabaseAuthenticationEnabled(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["iam_database_authentication_enabled"] = cty.BoolVal(p.IamDatabaseAuthenticationEnabled)
}

func EncodeDbInstance_ParameterGroupName(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["parameter_group_name"] = cty.StringVal(p.ParameterGroupName)
}

func EncodeDbInstance_Identifier(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["identifier"] = cty.StringVal(p.Identifier)
}

func EncodeDbInstance_InstanceClass(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["instance_class"] = cty.StringVal(p.InstanceClass)
}

func EncodeDbInstance_PubliclyAccessible(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["publicly_accessible"] = cty.BoolVal(p.PubliclyAccessible)
}

func EncodeDbInstance_StorageEncrypted(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["storage_encrypted"] = cty.BoolVal(p.StorageEncrypted)
}

func EncodeDbInstance_Tags(p DbInstanceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDbInstance_Port(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeDbInstance_KmsKeyId(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeDbInstance_SecurityGroupNames(p DbInstanceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_names"] = cty.SetVal(colVals)
}

func EncodeDbInstance_Id(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDbInstance_ReplicateSourceDb(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["replicate_source_db"] = cty.StringVal(p.ReplicateSourceDb)
}

func EncodeDbInstance_Timezone(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["timezone"] = cty.StringVal(p.Timezone)
}

func EncodeDbInstance_AutoMinorVersionUpgrade(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["auto_minor_version_upgrade"] = cty.BoolVal(p.AutoMinorVersionUpgrade)
}

func EncodeDbInstance_CaCertIdentifier(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["ca_cert_identifier"] = cty.StringVal(p.CaCertIdentifier)
}

func EncodeDbInstance_SnapshotIdentifier(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["snapshot_identifier"] = cty.StringVal(p.SnapshotIdentifier)
}

func EncodeDbInstance_CharacterSetName(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["character_set_name"] = cty.StringVal(p.CharacterSetName)
}

func EncodeDbInstance_DomainIamRoleName(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["domain_iam_role_name"] = cty.StringVal(p.DomainIamRoleName)
}

func EncodeDbInstance_MaxAllocatedStorage(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["max_allocated_storage"] = cty.NumberIntVal(p.MaxAllocatedStorage)
}

func EncodeDbInstance_AllowMajorVersionUpgrade(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["allow_major_version_upgrade"] = cty.BoolVal(p.AllowMajorVersionUpgrade)
}

func EncodeDbInstance_EngineVersion(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeDbInstance_LicenseModel(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["license_model"] = cty.StringVal(p.LicenseModel)
}

func EncodeDbInstance_DeletionProtection(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["deletion_protection"] = cty.BoolVal(p.DeletionProtection)
}

func EncodeDbInstance_FinalSnapshotIdentifier(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["final_snapshot_identifier"] = cty.StringVal(p.FinalSnapshotIdentifier)
}

func EncodeDbInstance_MaintenanceWindow(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["maintenance_window"] = cty.StringVal(p.MaintenanceWindow)
}

func EncodeDbInstance_Name(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDbInstance_Password(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeDbInstance_Username(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeDbInstance_CopyTagsToSnapshot(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["copy_tags_to_snapshot"] = cty.BoolVal(p.CopyTagsToSnapshot)
}

func EncodeDbInstance_Iops(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeDbInstance_PerformanceInsightsEnabled(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["performance_insights_enabled"] = cty.BoolVal(p.PerformanceInsightsEnabled)
}

func EncodeDbInstance_PerformanceInsightsKmsKeyId(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["performance_insights_kms_key_id"] = cty.StringVal(p.PerformanceInsightsKmsKeyId)
}

func EncodeDbInstance_VpcSecurityGroupIds(p DbInstanceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeDbInstance_StorageType(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["storage_type"] = cty.StringVal(p.StorageType)
}

func EncodeDbInstance_BackupRetentionPeriod(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["backup_retention_period"] = cty.NumberIntVal(p.BackupRetentionPeriod)
}

func EncodeDbInstance_Domain(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["domain"] = cty.StringVal(p.Domain)
}

func EncodeDbInstance_IdentifierPrefix(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["identifier_prefix"] = cty.StringVal(p.IdentifierPrefix)
}

func EncodeDbInstance_OptionGroupName(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["option_group_name"] = cty.StringVal(p.OptionGroupName)
}

func EncodeDbInstance_SkipFinalSnapshot(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["skip_final_snapshot"] = cty.BoolVal(p.SkipFinalSnapshot)
}

func EncodeDbInstance_MonitoringRoleArn(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["monitoring_role_arn"] = cty.StringVal(p.MonitoringRoleArn)
}

func EncodeDbInstance_AllocatedStorage(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["allocated_storage"] = cty.NumberIntVal(p.AllocatedStorage)
}

func EncodeDbInstance_BackupWindow(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["backup_window"] = cty.StringVal(p.BackupWindow)
}

func EncodeDbInstance_DbSubnetGroupName(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["db_subnet_group_name"] = cty.StringVal(p.DbSubnetGroupName)
}

func EncodeDbInstance_DeleteAutomatedBackups(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["delete_automated_backups"] = cty.BoolVal(p.DeleteAutomatedBackups)
}

func EncodeDbInstance_Engine(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeDbInstance_MonitoringInterval(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["monitoring_interval"] = cty.NumberIntVal(p.MonitoringInterval)
}

func EncodeDbInstance_ApplyImmediately(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["apply_immediately"] = cty.BoolVal(p.ApplyImmediately)
}

func EncodeDbInstance_AvailabilityZone(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeDbInstance_EnabledCloudwatchLogsExports(p DbInstanceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.EnabledCloudwatchLogsExports {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["enabled_cloudwatch_logs_exports"] = cty.SetVal(colVals)
}

func EncodeDbInstance_MultiAz(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["multi_az"] = cty.BoolVal(p.MultiAz)
}

func EncodeDbInstance_PerformanceInsightsRetentionPeriod(p DbInstanceParameters, vals map[string]cty.Value) {
	vals["performance_insights_retention_period"] = cty.NumberIntVal(p.PerformanceInsightsRetentionPeriod)
}

func EncodeDbInstance_S3Import(p S3Import, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDbInstance_S3Import_BucketPrefix(p, ctyVal)
	EncodeDbInstance_S3Import_IngestionRole(p, ctyVal)
	EncodeDbInstance_S3Import_SourceEngine(p, ctyVal)
	EncodeDbInstance_S3Import_SourceEngineVersion(p, ctyVal)
	EncodeDbInstance_S3Import_BucketName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["s3_import"] = cty.ListVal(valsForCollection)
}

func EncodeDbInstance_S3Import_BucketPrefix(p S3Import, vals map[string]cty.Value) {
	vals["bucket_prefix"] = cty.StringVal(p.BucketPrefix)
}

func EncodeDbInstance_S3Import_IngestionRole(p S3Import, vals map[string]cty.Value) {
	vals["ingestion_role"] = cty.StringVal(p.IngestionRole)
}

func EncodeDbInstance_S3Import_SourceEngine(p S3Import, vals map[string]cty.Value) {
	vals["source_engine"] = cty.StringVal(p.SourceEngine)
}

func EncodeDbInstance_S3Import_SourceEngineVersion(p S3Import, vals map[string]cty.Value) {
	vals["source_engine_version"] = cty.StringVal(p.SourceEngineVersion)
}

func EncodeDbInstance_S3Import_BucketName(p S3Import, vals map[string]cty.Value) {
	vals["bucket_name"] = cty.StringVal(p.BucketName)
}

func EncodeDbInstance_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDbInstance_Timeouts_Create(p, ctyVal)
	EncodeDbInstance_Timeouts_Delete(p, ctyVal)
	EncodeDbInstance_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDbInstance_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDbInstance_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDbInstance_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDbInstance_Address(p DbInstanceObservation, vals map[string]cty.Value) {
	vals["address"] = cty.StringVal(p.Address)
}

func EncodeDbInstance_ResourceId(p DbInstanceObservation, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeDbInstance_Status(p DbInstanceObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeDbInstance_HostedZoneId(p DbInstanceObservation, vals map[string]cty.Value) {
	vals["hosted_zone_id"] = cty.StringVal(p.HostedZoneId)
}

func EncodeDbInstance_Arn(p DbInstanceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDbInstance_Endpoint(p DbInstanceObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeDbInstance_Replicas(p DbInstanceObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Replicas {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["replicas"] = cty.ListVal(colVals)
}