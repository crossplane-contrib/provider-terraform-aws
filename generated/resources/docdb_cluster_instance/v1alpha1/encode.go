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
	r, ok := mr.(*DocdbClusterInstance)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DocdbClusterInstance.")
	}
	return EncodeDocdbClusterInstance(*r), nil
}

func EncodeDocdbClusterInstance(r DocdbClusterInstance) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDocdbClusterInstance_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterInstance_AutoMinorVersionUpgrade(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterInstance_ApplyImmediately(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterInstance_Engine(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterInstance_Identifier(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterInstance_IdentifierPrefix(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterInstance_PreferredMaintenanceWindow(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterInstance_Id(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterInstance_InstanceClass(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterInstance_PromotionTier(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterInstance_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterInstance_CaCertIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterInstance_ClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterInstance_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDocdbClusterInstance_Writer(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterInstance_Arn(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterInstance_Endpoint(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterInstance_KmsKeyId(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterInstance_PubliclyAccessible(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterInstance_DbSubnetGroupName(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterInstance_DbiResourceId(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterInstance_Port(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterInstance_PreferredBackupWindow(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterInstance_EngineVersion(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterInstance_StorageEncrypted(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDocdbClusterInstance_Tags(p DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDocdbClusterInstance_AutoMinorVersionUpgrade(p DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	vals["auto_minor_version_upgrade"] = cty.BoolVal(p.AutoMinorVersionUpgrade)
}

func EncodeDocdbClusterInstance_ApplyImmediately(p DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	vals["apply_immediately"] = cty.BoolVal(p.ApplyImmediately)
}

func EncodeDocdbClusterInstance_Engine(p DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeDocdbClusterInstance_Identifier(p DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	vals["identifier"] = cty.StringVal(p.Identifier)
}

func EncodeDocdbClusterInstance_IdentifierPrefix(p DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	vals["identifier_prefix"] = cty.StringVal(p.IdentifierPrefix)
}

func EncodeDocdbClusterInstance_PreferredMaintenanceWindow(p DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	vals["preferred_maintenance_window"] = cty.StringVal(p.PreferredMaintenanceWindow)
}

func EncodeDocdbClusterInstance_Id(p DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDocdbClusterInstance_InstanceClass(p DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	vals["instance_class"] = cty.StringVal(p.InstanceClass)
}

func EncodeDocdbClusterInstance_PromotionTier(p DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	vals["promotion_tier"] = cty.NumberIntVal(p.PromotionTier)
}

func EncodeDocdbClusterInstance_AvailabilityZone(p DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeDocdbClusterInstance_CaCertIdentifier(p DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	vals["ca_cert_identifier"] = cty.StringVal(p.CaCertIdentifier)
}

func EncodeDocdbClusterInstance_ClusterIdentifier(p DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	vals["cluster_identifier"] = cty.StringVal(p.ClusterIdentifier)
}

func EncodeDocdbClusterInstance_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDocdbClusterInstance_Timeouts_Update(p, ctyVal)
	EncodeDocdbClusterInstance_Timeouts_Create(p, ctyVal)
	EncodeDocdbClusterInstance_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDocdbClusterInstance_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDocdbClusterInstance_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDocdbClusterInstance_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDocdbClusterInstance_Writer(p DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	vals["writer"] = cty.BoolVal(p.Writer)
}

func EncodeDocdbClusterInstance_Arn(p DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDocdbClusterInstance_Endpoint(p DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeDocdbClusterInstance_KmsKeyId(p DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeDocdbClusterInstance_PubliclyAccessible(p DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	vals["publicly_accessible"] = cty.BoolVal(p.PubliclyAccessible)
}

func EncodeDocdbClusterInstance_DbSubnetGroupName(p DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	vals["db_subnet_group_name"] = cty.StringVal(p.DbSubnetGroupName)
}

func EncodeDocdbClusterInstance_DbiResourceId(p DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	vals["dbi_resource_id"] = cty.StringVal(p.DbiResourceId)
}

func EncodeDocdbClusterInstance_Port(p DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeDocdbClusterInstance_PreferredBackupWindow(p DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	vals["preferred_backup_window"] = cty.StringVal(p.PreferredBackupWindow)
}

func EncodeDocdbClusterInstance_EngineVersion(p DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeDocdbClusterInstance_StorageEncrypted(p DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	vals["storage_encrypted"] = cty.BoolVal(p.StorageEncrypted)
}