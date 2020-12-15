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
	r, ok := mr.(*RdsClusterInstance)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a RdsClusterInstance.")
	}
	return EncodeRdsClusterInstance(*r), nil
}

func EncodeRdsClusterInstance(r RdsClusterInstance) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRdsClusterInstance_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_PerformanceInsightsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_AutoMinorVersionUpgrade(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_DbParameterGroupName(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_EngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_MonitoringInterval(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_PubliclyAccessible(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_PromotionTier(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_CaCertIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_CopyTagsToSnapshot(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_DbSubnetGroupName(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_MonitoringRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_PreferredBackupWindow(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_IdentifierPrefix(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_ClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_Identifier(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_ApplyImmediately(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_Engine(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_Id(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_InstanceClass(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_PerformanceInsightsKmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_PreferredMaintenanceWindow(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterInstance_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeRdsClusterInstance_Arn(r.Status.AtProvider, ctyVal)
	EncodeRdsClusterInstance_Endpoint(r.Status.AtProvider, ctyVal)
	EncodeRdsClusterInstance_KmsKeyId(r.Status.AtProvider, ctyVal)
	EncodeRdsClusterInstance_Port(r.Status.AtProvider, ctyVal)
	EncodeRdsClusterInstance_Writer(r.Status.AtProvider, ctyVal)
	EncodeRdsClusterInstance_DbiResourceId(r.Status.AtProvider, ctyVal)
	EncodeRdsClusterInstance_StorageEncrypted(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeRdsClusterInstance_AvailabilityZone(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeRdsClusterInstance_PerformanceInsightsEnabled(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["performance_insights_enabled"] = cty.BoolVal(p.PerformanceInsightsEnabled)
}

func EncodeRdsClusterInstance_AutoMinorVersionUpgrade(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["auto_minor_version_upgrade"] = cty.BoolVal(p.AutoMinorVersionUpgrade)
}

func EncodeRdsClusterInstance_DbParameterGroupName(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["db_parameter_group_name"] = cty.StringVal(p.DbParameterGroupName)
}

func EncodeRdsClusterInstance_EngineVersion(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeRdsClusterInstance_MonitoringInterval(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["monitoring_interval"] = cty.NumberIntVal(p.MonitoringInterval)
}

func EncodeRdsClusterInstance_PubliclyAccessible(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["publicly_accessible"] = cty.BoolVal(p.PubliclyAccessible)
}

func EncodeRdsClusterInstance_PromotionTier(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["promotion_tier"] = cty.NumberIntVal(p.PromotionTier)
}

func EncodeRdsClusterInstance_CaCertIdentifier(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["ca_cert_identifier"] = cty.StringVal(p.CaCertIdentifier)
}

func EncodeRdsClusterInstance_CopyTagsToSnapshot(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["copy_tags_to_snapshot"] = cty.BoolVal(p.CopyTagsToSnapshot)
}

func EncodeRdsClusterInstance_DbSubnetGroupName(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["db_subnet_group_name"] = cty.StringVal(p.DbSubnetGroupName)
}

func EncodeRdsClusterInstance_MonitoringRoleArn(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["monitoring_role_arn"] = cty.StringVal(p.MonitoringRoleArn)
}

func EncodeRdsClusterInstance_PreferredBackupWindow(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["preferred_backup_window"] = cty.StringVal(p.PreferredBackupWindow)
}

func EncodeRdsClusterInstance_IdentifierPrefix(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["identifier_prefix"] = cty.StringVal(p.IdentifierPrefix)
}

func EncodeRdsClusterInstance_ClusterIdentifier(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["cluster_identifier"] = cty.StringVal(p.ClusterIdentifier)
}

func EncodeRdsClusterInstance_Identifier(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["identifier"] = cty.StringVal(p.Identifier)
}

func EncodeRdsClusterInstance_ApplyImmediately(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["apply_immediately"] = cty.BoolVal(p.ApplyImmediately)
}

func EncodeRdsClusterInstance_Engine(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeRdsClusterInstance_Id(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRdsClusterInstance_InstanceClass(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["instance_class"] = cty.StringVal(p.InstanceClass)
}

func EncodeRdsClusterInstance_PerformanceInsightsKmsKeyId(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["performance_insights_kms_key_id"] = cty.StringVal(p.PerformanceInsightsKmsKeyId)
}

func EncodeRdsClusterInstance_PreferredMaintenanceWindow(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
	vals["preferred_maintenance_window"] = cty.StringVal(p.PreferredMaintenanceWindow)
}

func EncodeRdsClusterInstance_Tags(p RdsClusterInstanceParameters, vals map[string]cty.Value) {
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

func EncodeRdsClusterInstance_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeRdsClusterInstance_Timeouts_Create(p, ctyVal)
	EncodeRdsClusterInstance_Timeouts_Delete(p, ctyVal)
	EncodeRdsClusterInstance_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeRdsClusterInstance_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeRdsClusterInstance_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeRdsClusterInstance_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeRdsClusterInstance_Arn(p RdsClusterInstanceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeRdsClusterInstance_Endpoint(p RdsClusterInstanceObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeRdsClusterInstance_KmsKeyId(p RdsClusterInstanceObservation, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeRdsClusterInstance_Port(p RdsClusterInstanceObservation, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeRdsClusterInstance_Writer(p RdsClusterInstanceObservation, vals map[string]cty.Value) {
	vals["writer"] = cty.BoolVal(p.Writer)
}

func EncodeRdsClusterInstance_DbiResourceId(p RdsClusterInstanceObservation, vals map[string]cty.Value) {
	vals["dbi_resource_id"] = cty.StringVal(p.DbiResourceId)
}

func EncodeRdsClusterInstance_StorageEncrypted(p RdsClusterInstanceObservation, vals map[string]cty.Value) {
	vals["storage_encrypted"] = cty.BoolVal(p.StorageEncrypted)
}