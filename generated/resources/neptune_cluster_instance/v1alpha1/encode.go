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
	r, ok := mr.(*NeptuneClusterInstance)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a NeptuneClusterInstance.")
	}
	return EncodeNeptuneClusterInstance(*r), nil
}

func EncodeNeptuneClusterInstance(r NeptuneClusterInstance) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeNeptuneClusterInstance_ApplyImmediately(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_Engine(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_InstanceClass(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_NeptuneSubnetGroupName(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_Tags(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_AutoMinorVersionUpgrade(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_ClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_EngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_Id(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_PreferredBackupWindow(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_PromotionTier(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_Identifier(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_Port(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_IdentifierPrefix(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_NeptuneParameterGroupName(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_PreferredMaintenanceWindow(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_PubliclyAccessible(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterInstance_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeNeptuneClusterInstance_DbiResourceId(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterInstance_Arn(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterInstance_Endpoint(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterInstance_KmsKeyArn(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterInstance_Writer(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterInstance_Address(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterInstance_StorageEncrypted(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeNeptuneClusterInstance_ApplyImmediately(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["apply_immediately"] = cty.BoolVal(p.ApplyImmediately)
}

func EncodeNeptuneClusterInstance_AvailabilityZone(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeNeptuneClusterInstance_Engine(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeNeptuneClusterInstance_InstanceClass(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["instance_class"] = cty.StringVal(p.InstanceClass)
}

func EncodeNeptuneClusterInstance_NeptuneSubnetGroupName(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["neptune_subnet_group_name"] = cty.StringVal(p.NeptuneSubnetGroupName)
}

func EncodeNeptuneClusterInstance_Tags(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
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

func EncodeNeptuneClusterInstance_AutoMinorVersionUpgrade(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["auto_minor_version_upgrade"] = cty.BoolVal(p.AutoMinorVersionUpgrade)
}

func EncodeNeptuneClusterInstance_ClusterIdentifier(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["cluster_identifier"] = cty.StringVal(p.ClusterIdentifier)
}

func EncodeNeptuneClusterInstance_EngineVersion(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeNeptuneClusterInstance_Id(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeNeptuneClusterInstance_PreferredBackupWindow(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["preferred_backup_window"] = cty.StringVal(p.PreferredBackupWindow)
}

func EncodeNeptuneClusterInstance_PromotionTier(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["promotion_tier"] = cty.NumberIntVal(p.PromotionTier)
}

func EncodeNeptuneClusterInstance_Identifier(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["identifier"] = cty.StringVal(p.Identifier)
}

func EncodeNeptuneClusterInstance_Port(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeNeptuneClusterInstance_IdentifierPrefix(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["identifier_prefix"] = cty.StringVal(p.IdentifierPrefix)
}

func EncodeNeptuneClusterInstance_NeptuneParameterGroupName(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["neptune_parameter_group_name"] = cty.StringVal(p.NeptuneParameterGroupName)
}

func EncodeNeptuneClusterInstance_PreferredMaintenanceWindow(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["preferred_maintenance_window"] = cty.StringVal(p.PreferredMaintenanceWindow)
}

func EncodeNeptuneClusterInstance_PubliclyAccessible(p NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	vals["publicly_accessible"] = cty.BoolVal(p.PubliclyAccessible)
}

func EncodeNeptuneClusterInstance_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeNeptuneClusterInstance_Timeouts_Delete(p, ctyVal)
	EncodeNeptuneClusterInstance_Timeouts_Update(p, ctyVal)
	EncodeNeptuneClusterInstance_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeNeptuneClusterInstance_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeNeptuneClusterInstance_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeNeptuneClusterInstance_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeNeptuneClusterInstance_DbiResourceId(p NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	vals["dbi_resource_id"] = cty.StringVal(p.DbiResourceId)
}

func EncodeNeptuneClusterInstance_Arn(p NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeNeptuneClusterInstance_Endpoint(p NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeNeptuneClusterInstance_KmsKeyArn(p NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeNeptuneClusterInstance_Writer(p NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	vals["writer"] = cty.BoolVal(p.Writer)
}

func EncodeNeptuneClusterInstance_Address(p NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	vals["address"] = cty.StringVal(p.Address)
}

func EncodeNeptuneClusterInstance_StorageEncrypted(p NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	vals["storage_encrypted"] = cty.BoolVal(p.StorageEncrypted)
}