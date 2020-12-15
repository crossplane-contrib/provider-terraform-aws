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
	r, ok := mr.(*DocdbCluster)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDocdbCluster(r, ctyValue)
}

func DecodeDocdbCluster(prev *DocdbCluster, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDocdbCluster_FinalSnapshotIdentifier(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_MasterUsername(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_SnapshotIdentifier(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_Port(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_AvailabilityZones(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_ClusterMembers(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_DbClusterParameterGroupName(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_DbSubnetGroupName(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_EngineVersion(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_Tags(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_VpcSecurityGroupIds(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_ClusterIdentifierPrefix(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_KmsKeyId(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_PreferredBackupWindow(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_PreferredMaintenanceWindow(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_StorageEncrypted(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_EnabledCloudwatchLogsExports(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_MasterPassword(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_SkipFinalSnapshot(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_ApplyImmediately(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_DeletionProtection(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_ClusterIdentifier(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_Engine(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_BackupRetentionPeriod(&new.Spec.ForProvider, valMap)
	DecodeDocdbCluster_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDocdbCluster_ClusterResourceId(&new.Status.AtProvider, valMap)
	DecodeDocdbCluster_Endpoint(&new.Status.AtProvider, valMap)
	DecodeDocdbCluster_Arn(&new.Status.AtProvider, valMap)
	DecodeDocdbCluster_HostedZoneId(&new.Status.AtProvider, valMap)
	DecodeDocdbCluster_ReaderEndpoint(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_FinalSnapshotIdentifier(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.FinalSnapshotIdentifier = ctwhy.ValueAsString(vals["final_snapshot_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_MasterUsername(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.MasterUsername = ctwhy.ValueAsString(vals["master_username"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_SnapshotIdentifier(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.SnapshotIdentifier = ctwhy.ValueAsString(vals["snapshot_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_Port(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.Port = ctwhy.ValueAsInt64(vals["port"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDocdbCluster_AvailabilityZones(p *DocdbClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["availability_zones"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.AvailabilityZones = goVals
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDocdbCluster_ClusterMembers(p *DocdbClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["cluster_members"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ClusterMembers = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_DbClusterParameterGroupName(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.DbClusterParameterGroupName = ctwhy.ValueAsString(vals["db_cluster_parameter_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_DbSubnetGroupName(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.DbSubnetGroupName = ctwhy.ValueAsString(vals["db_subnet_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_EngineVersion(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.EngineVersion = ctwhy.ValueAsString(vals["engine_version"])
}

//primitiveMapTypeDecodeTemplate
func DecodeDocdbCluster_Tags(p *DocdbClusterParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDocdbCluster_VpcSecurityGroupIds(p *DocdbClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["vpc_security_group_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.VpcSecurityGroupIds = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_ClusterIdentifierPrefix(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.ClusterIdentifierPrefix = ctwhy.ValueAsString(vals["cluster_identifier_prefix"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_KmsKeyId(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_PreferredBackupWindow(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.PreferredBackupWindow = ctwhy.ValueAsString(vals["preferred_backup_window"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_PreferredMaintenanceWindow(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.PreferredMaintenanceWindow = ctwhy.ValueAsString(vals["preferred_maintenance_window"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_StorageEncrypted(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.StorageEncrypted = ctwhy.ValueAsBool(vals["storage_encrypted"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDocdbCluster_EnabledCloudwatchLogsExports(p *DocdbClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["enabled_cloudwatch_logs_exports"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.EnabledCloudwatchLogsExports = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_MasterPassword(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.MasterPassword = ctwhy.ValueAsString(vals["master_password"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_SkipFinalSnapshot(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.SkipFinalSnapshot = ctwhy.ValueAsBool(vals["skip_final_snapshot"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_ApplyImmediately(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.ApplyImmediately = ctwhy.ValueAsBool(vals["apply_immediately"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_DeletionProtection(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.DeletionProtection = ctwhy.ValueAsBool(vals["deletion_protection"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_ClusterIdentifier(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.ClusterIdentifier = ctwhy.ValueAsString(vals["cluster_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_Engine(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.Engine = ctwhy.ValueAsString(vals["engine"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_BackupRetentionPeriod(p *DocdbClusterParameters, vals map[string]cty.Value) {
	p.BackupRetentionPeriod = ctwhy.ValueAsInt64(vals["backup_retention_period"])
}

//containerTypeDecodeTemplate
func DecodeDocdbCluster_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDocdbCluster_Timeouts_Create(p, valMap)
	DecodeDocdbCluster_Timeouts_Delete(p, valMap)
	DecodeDocdbCluster_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_ClusterResourceId(p *DocdbClusterObservation, vals map[string]cty.Value) {
	p.ClusterResourceId = ctwhy.ValueAsString(vals["cluster_resource_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_Endpoint(p *DocdbClusterObservation, vals map[string]cty.Value) {
	p.Endpoint = ctwhy.ValueAsString(vals["endpoint"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_Arn(p *DocdbClusterObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_HostedZoneId(p *DocdbClusterObservation, vals map[string]cty.Value) {
	p.HostedZoneId = ctwhy.ValueAsString(vals["hosted_zone_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbCluster_ReaderEndpoint(p *DocdbClusterObservation, vals map[string]cty.Value) {
	p.ReaderEndpoint = ctwhy.ValueAsString(vals["reader_endpoint"])
}