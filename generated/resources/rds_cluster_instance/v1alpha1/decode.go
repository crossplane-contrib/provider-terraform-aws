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
	r, ok := mr.(*RdsClusterInstance)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeRdsClusterInstance(r, ctyValue)
}

func DecodeRdsClusterInstance(prev *RdsClusterInstance, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeRdsClusterInstance_IdentifierPrefix(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_PerformanceInsightsEnabled(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_ClusterIdentifier(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_PromotionTier(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_AvailabilityZone(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_CopyTagsToSnapshot(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_DbParameterGroupName(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_PreferredMaintenanceWindow(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_PubliclyAccessible(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_ApplyImmediately(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_DbSubnetGroupName(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_PerformanceInsightsKmsKeyId(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_AutoMinorVersionUpgrade(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_MonitoringRoleArn(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_InstanceClass(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_PreferredBackupWindow(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_Tags(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_EngineVersion(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_Identifier(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_CaCertIdentifier(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_Engine(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_MonitoringInterval(&new.Spec.ForProvider, valMap)
	DecodeRdsClusterInstance_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeRdsClusterInstance_Port(&new.Status.AtProvider, valMap)
	DecodeRdsClusterInstance_Endpoint(&new.Status.AtProvider, valMap)
	DecodeRdsClusterInstance_StorageEncrypted(&new.Status.AtProvider, valMap)
	DecodeRdsClusterInstance_KmsKeyId(&new.Status.AtProvider, valMap)
	DecodeRdsClusterInstance_Writer(&new.Status.AtProvider, valMap)
	DecodeRdsClusterInstance_DbiResourceId(&new.Status.AtProvider, valMap)
	DecodeRdsClusterInstance_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_IdentifierPrefix(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.IdentifierPrefix = ctwhy.ValueAsString(vals["identifier_prefix"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_PerformanceInsightsEnabled(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.PerformanceInsightsEnabled = ctwhy.ValueAsBool(vals["performance_insights_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_ClusterIdentifier(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.ClusterIdentifier = ctwhy.ValueAsString(vals["cluster_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_PromotionTier(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.PromotionTier = ctwhy.ValueAsInt64(vals["promotion_tier"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_AvailabilityZone(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.AvailabilityZone = ctwhy.ValueAsString(vals["availability_zone"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_CopyTagsToSnapshot(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.CopyTagsToSnapshot = ctwhy.ValueAsBool(vals["copy_tags_to_snapshot"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_DbParameterGroupName(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.DbParameterGroupName = ctwhy.ValueAsString(vals["db_parameter_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_PreferredMaintenanceWindow(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.PreferredMaintenanceWindow = ctwhy.ValueAsString(vals["preferred_maintenance_window"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_PubliclyAccessible(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.PubliclyAccessible = ctwhy.ValueAsBool(vals["publicly_accessible"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_ApplyImmediately(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.ApplyImmediately = ctwhy.ValueAsBool(vals["apply_immediately"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_DbSubnetGroupName(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.DbSubnetGroupName = ctwhy.ValueAsString(vals["db_subnet_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_PerformanceInsightsKmsKeyId(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.PerformanceInsightsKmsKeyId = ctwhy.ValueAsString(vals["performance_insights_kms_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_AutoMinorVersionUpgrade(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.AutoMinorVersionUpgrade = ctwhy.ValueAsBool(vals["auto_minor_version_upgrade"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_MonitoringRoleArn(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.MonitoringRoleArn = ctwhy.ValueAsString(vals["monitoring_role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_InstanceClass(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.InstanceClass = ctwhy.ValueAsString(vals["instance_class"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_PreferredBackupWindow(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.PreferredBackupWindow = ctwhy.ValueAsString(vals["preferred_backup_window"])
}

//primitiveMapTypeDecodeTemplate
func DecodeRdsClusterInstance_Tags(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_EngineVersion(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.EngineVersion = ctwhy.ValueAsString(vals["engine_version"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_Identifier(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.Identifier = ctwhy.ValueAsString(vals["identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_CaCertIdentifier(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.CaCertIdentifier = ctwhy.ValueAsString(vals["ca_cert_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_Engine(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.Engine = ctwhy.ValueAsString(vals["engine"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_MonitoringInterval(p *RdsClusterInstanceParameters, vals map[string]cty.Value) {
	p.MonitoringInterval = ctwhy.ValueAsInt64(vals["monitoring_interval"])
}

//containerTypeDecodeTemplate
func DecodeRdsClusterInstance_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeRdsClusterInstance_Timeouts_Create(p, valMap)
	DecodeRdsClusterInstance_Timeouts_Delete(p, valMap)
	DecodeRdsClusterInstance_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_Port(p *RdsClusterInstanceObservation, vals map[string]cty.Value) {
	p.Port = ctwhy.ValueAsInt64(vals["port"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_Endpoint(p *RdsClusterInstanceObservation, vals map[string]cty.Value) {
	p.Endpoint = ctwhy.ValueAsString(vals["endpoint"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_StorageEncrypted(p *RdsClusterInstanceObservation, vals map[string]cty.Value) {
	p.StorageEncrypted = ctwhy.ValueAsBool(vals["storage_encrypted"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_KmsKeyId(p *RdsClusterInstanceObservation, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_Writer(p *RdsClusterInstanceObservation, vals map[string]cty.Value) {
	p.Writer = ctwhy.ValueAsBool(vals["writer"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_DbiResourceId(p *RdsClusterInstanceObservation, vals map[string]cty.Value) {
	p.DbiResourceId = ctwhy.ValueAsString(vals["dbi_resource_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRdsClusterInstance_Arn(p *RdsClusterInstanceObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}