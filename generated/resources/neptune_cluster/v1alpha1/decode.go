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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*NeptuneCluster)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeNeptuneCluster(r, ctyValue)
}

func DecodeNeptuneCluster(prev *NeptuneCluster, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeNeptuneCluster_IamDatabaseAuthenticationEnabled(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_IamRoles(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_KmsKeyArn(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_NeptuneClusterParameterGroupName(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_BackupRetentionPeriod(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_DeletionProtection(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_Id(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_EnableCloudwatchLogsExports(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_Engine(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_FinalSnapshotIdentifier(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_ApplyImmediately(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_AvailabilityZones(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_ClusterIdentifierPrefix(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_PreferredMaintenanceWindow(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_StorageEncrypted(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_NeptuneSubnetGroupName(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_ReplicationSourceIdentifier(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_Tags(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_ClusterIdentifier(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_EngineVersion(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_Port(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_SnapshotIdentifier(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_PreferredBackupWindow(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_SkipFinalSnapshot(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_VpcSecurityGroupIds(&new.Spec.ForProvider, valMap)
	DecodeNeptuneCluster_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeNeptuneCluster_ClusterMembers(&new.Status.AtProvider, valMap)
	DecodeNeptuneCluster_Endpoint(&new.Status.AtProvider, valMap)
	DecodeNeptuneCluster_ClusterResourceId(&new.Status.AtProvider, valMap)
	DecodeNeptuneCluster_Arn(&new.Status.AtProvider, valMap)
	DecodeNeptuneCluster_HostedZoneId(&new.Status.AtProvider, valMap)
	DecodeNeptuneCluster_ReaderEndpoint(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_IamDatabaseAuthenticationEnabled(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.IamDatabaseAuthenticationEnabled = ctwhy.ValueAsBool(vals["iam_database_authentication_enabled"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeNeptuneCluster_IamRoles(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["iam_roles"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.IamRoles = goVals
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_KmsKeyArn(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.KmsKeyArn = ctwhy.ValueAsString(vals["kms_key_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_NeptuneClusterParameterGroupName(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.NeptuneClusterParameterGroupName = ctwhy.ValueAsString(vals["neptune_cluster_parameter_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_BackupRetentionPeriod(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.BackupRetentionPeriod = ctwhy.ValueAsInt64(vals["backup_retention_period"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_DeletionProtection(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.DeletionProtection = ctwhy.ValueAsBool(vals["deletion_protection"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_Id(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeNeptuneCluster_EnableCloudwatchLogsExports(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["enable_cloudwatch_logs_exports"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.EnableCloudwatchLogsExports = goVals
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_Engine(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.Engine = ctwhy.ValueAsString(vals["engine"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_FinalSnapshotIdentifier(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.FinalSnapshotIdentifier = ctwhy.ValueAsString(vals["final_snapshot_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_ApplyImmediately(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.ApplyImmediately = ctwhy.ValueAsBool(vals["apply_immediately"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeNeptuneCluster_AvailabilityZones(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["availability_zones"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.AvailabilityZones = goVals
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_ClusterIdentifierPrefix(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.ClusterIdentifierPrefix = ctwhy.ValueAsString(vals["cluster_identifier_prefix"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_PreferredMaintenanceWindow(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.PreferredMaintenanceWindow = ctwhy.ValueAsString(vals["preferred_maintenance_window"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_StorageEncrypted(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.StorageEncrypted = ctwhy.ValueAsBool(vals["storage_encrypted"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_NeptuneSubnetGroupName(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.NeptuneSubnetGroupName = ctwhy.ValueAsString(vals["neptune_subnet_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_ReplicationSourceIdentifier(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.ReplicationSourceIdentifier = ctwhy.ValueAsString(vals["replication_source_identifier"])
}

//primitiveMapTypeDecodeTemplate
func DecodeNeptuneCluster_Tags(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_ClusterIdentifier(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.ClusterIdentifier = ctwhy.ValueAsString(vals["cluster_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_EngineVersion(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.EngineVersion = ctwhy.ValueAsString(vals["engine_version"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_Port(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.Port = ctwhy.ValueAsInt64(vals["port"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_SnapshotIdentifier(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.SnapshotIdentifier = ctwhy.ValueAsString(vals["snapshot_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_PreferredBackupWindow(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.PreferredBackupWindow = ctwhy.ValueAsString(vals["preferred_backup_window"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_SkipFinalSnapshot(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	p.SkipFinalSnapshot = ctwhy.ValueAsBool(vals["skip_final_snapshot"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeNeptuneCluster_VpcSecurityGroupIds(p *NeptuneClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["vpc_security_group_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.VpcSecurityGroupIds = goVals
}

//containerTypeDecodeTemplate
func DecodeNeptuneCluster_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeNeptuneCluster_Timeouts_Create(p, valMap)
	DecodeNeptuneCluster_Timeouts_Delete(p, valMap)
	DecodeNeptuneCluster_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeNeptuneCluster_ClusterMembers(p *NeptuneClusterObservation, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["cluster_members"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ClusterMembers = goVals
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_Endpoint(p *NeptuneClusterObservation, vals map[string]cty.Value) {
	p.Endpoint = ctwhy.ValueAsString(vals["endpoint"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_ClusterResourceId(p *NeptuneClusterObservation, vals map[string]cty.Value) {
	p.ClusterResourceId = ctwhy.ValueAsString(vals["cluster_resource_id"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_Arn(p *NeptuneClusterObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_HostedZoneId(p *NeptuneClusterObservation, vals map[string]cty.Value) {
	p.HostedZoneId = ctwhy.ValueAsString(vals["hosted_zone_id"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneCluster_ReaderEndpoint(p *NeptuneClusterObservation, vals map[string]cty.Value) {
	p.ReaderEndpoint = ctwhy.ValueAsString(vals["reader_endpoint"])
}