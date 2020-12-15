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
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*DocdbCluster)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DocdbCluster.")
	}
	return EncodeDocdbCluster(*r), nil
}

func EncodeDocdbCluster(r DocdbCluster) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDocdbCluster_DeletionProtection(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_ClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_SkipFinalSnapshot(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_FinalSnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_DbClusterParameterGroupName(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_SnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_VpcSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_ApplyImmediately(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_ClusterMembers(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_DbSubnetGroupName(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_MasterUsername(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_ClusterIdentifierPrefix(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_PreferredBackupWindow(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_Port(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_Id(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_EngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_StorageEncrypted(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_PreferredMaintenanceWindow(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_AvailabilityZones(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_BackupRetentionPeriod(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_EnabledCloudwatchLogsExports(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_Engine(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_MasterPassword(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDocdbCluster_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDocdbCluster_ClusterResourceId(r.Status.AtProvider, ctyVal)
	EncodeDocdbCluster_Endpoint(r.Status.AtProvider, ctyVal)
	EncodeDocdbCluster_HostedZoneId(r.Status.AtProvider, ctyVal)
	EncodeDocdbCluster_ReaderEndpoint(r.Status.AtProvider, ctyVal)
	EncodeDocdbCluster_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDocdbCluster_DeletionProtection(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["deletion_protection"] = cty.BoolVal(p.DeletionProtection)
}

func EncodeDocdbCluster_ClusterIdentifier(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["cluster_identifier"] = cty.StringVal(p.ClusterIdentifier)
}

func EncodeDocdbCluster_KmsKeyId(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeDocdbCluster_SkipFinalSnapshot(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["skip_final_snapshot"] = cty.BoolVal(p.SkipFinalSnapshot)
}

func EncodeDocdbCluster_FinalSnapshotIdentifier(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["final_snapshot_identifier"] = cty.StringVal(p.FinalSnapshotIdentifier)
}

func EncodeDocdbCluster_DbClusterParameterGroupName(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["db_cluster_parameter_group_name"] = cty.StringVal(p.DbClusterParameterGroupName)
}

func EncodeDocdbCluster_SnapshotIdentifier(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["snapshot_identifier"] = cty.StringVal(p.SnapshotIdentifier)
}

func EncodeDocdbCluster_VpcSecurityGroupIds(p DocdbClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeDocdbCluster_ApplyImmediately(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["apply_immediately"] = cty.BoolVal(p.ApplyImmediately)
}

func EncodeDocdbCluster_ClusterMembers(p DocdbClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ClusterMembers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["cluster_members"] = cty.SetVal(colVals)
}

func EncodeDocdbCluster_DbSubnetGroupName(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["db_subnet_group_name"] = cty.StringVal(p.DbSubnetGroupName)
}

func EncodeDocdbCluster_MasterUsername(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["master_username"] = cty.StringVal(p.MasterUsername)
}

func EncodeDocdbCluster_ClusterIdentifierPrefix(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["cluster_identifier_prefix"] = cty.StringVal(p.ClusterIdentifierPrefix)
}

func EncodeDocdbCluster_PreferredBackupWindow(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["preferred_backup_window"] = cty.StringVal(p.PreferredBackupWindow)
}

func EncodeDocdbCluster_Port(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeDocdbCluster_Id(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDocdbCluster_EngineVersion(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeDocdbCluster_StorageEncrypted(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["storage_encrypted"] = cty.BoolVal(p.StorageEncrypted)
}

func EncodeDocdbCluster_PreferredMaintenanceWindow(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["preferred_maintenance_window"] = cty.StringVal(p.PreferredMaintenanceWindow)
}

func EncodeDocdbCluster_AvailabilityZones(p DocdbClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.SetVal(colVals)
}

func EncodeDocdbCluster_BackupRetentionPeriod(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["backup_retention_period"] = cty.NumberIntVal(p.BackupRetentionPeriod)
}

func EncodeDocdbCluster_EnabledCloudwatchLogsExports(p DocdbClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.EnabledCloudwatchLogsExports {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["enabled_cloudwatch_logs_exports"] = cty.ListVal(colVals)
}

func EncodeDocdbCluster_Engine(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeDocdbCluster_MasterPassword(p DocdbClusterParameters, vals map[string]cty.Value) {
	vals["master_password"] = cty.StringVal(p.MasterPassword)
}

func EncodeDocdbCluster_Tags(p DocdbClusterParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDocdbCluster_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDocdbCluster_Timeouts_Create(p, ctyVal)
	EncodeDocdbCluster_Timeouts_Delete(p, ctyVal)
	EncodeDocdbCluster_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDocdbCluster_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDocdbCluster_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDocdbCluster_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDocdbCluster_ClusterResourceId(p DocdbClusterObservation, vals map[string]cty.Value) {
	vals["cluster_resource_id"] = cty.StringVal(p.ClusterResourceId)
}

func EncodeDocdbCluster_Endpoint(p DocdbClusterObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeDocdbCluster_HostedZoneId(p DocdbClusterObservation, vals map[string]cty.Value) {
	vals["hosted_zone_id"] = cty.StringVal(p.HostedZoneId)
}

func EncodeDocdbCluster_ReaderEndpoint(p DocdbClusterObservation, vals map[string]cty.Value) {
	vals["reader_endpoint"] = cty.StringVal(p.ReaderEndpoint)
}

func EncodeDocdbCluster_Arn(p DocdbClusterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}