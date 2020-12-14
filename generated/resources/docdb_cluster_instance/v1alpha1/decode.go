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
	r, ok := mr.(*DocdbClusterInstance)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDocdbClusterInstance(r, ctyValue)
}

func DecodeDocdbClusterInstance(prev *DocdbClusterInstance, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDocdbClusterInstance_InstanceClass(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterInstance_AvailabilityZone(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterInstance_Engine(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterInstance_CaCertIdentifier(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterInstance_Identifier(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterInstance_IdentifierPrefix(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterInstance_AutoMinorVersionUpgrade(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterInstance_PreferredMaintenanceWindow(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterInstance_ApplyImmediately(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterInstance_Id(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterInstance_PromotionTier(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterInstance_Tags(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterInstance_ClusterIdentifier(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterInstance_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDocdbClusterInstance_EngineVersion(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterInstance_StorageEncrypted(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterInstance_DbSubnetGroupName(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterInstance_Port(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterInstance_Arn(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterInstance_KmsKeyId(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterInstance_PreferredBackupWindow(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterInstance_DbiResourceId(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterInstance_PubliclyAccessible(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterInstance_Writer(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterInstance_Endpoint(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeDocdbClusterInstance_InstanceClass(p *DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	p.InstanceClass = ctwhy.ValueAsString(vals["instance_class"])
}

func DecodeDocdbClusterInstance_AvailabilityZone(p *DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	p.AvailabilityZone = ctwhy.ValueAsString(vals["availability_zone"])
}

func DecodeDocdbClusterInstance_Engine(p *DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	p.Engine = ctwhy.ValueAsString(vals["engine"])
}

func DecodeDocdbClusterInstance_CaCertIdentifier(p *DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	p.CaCertIdentifier = ctwhy.ValueAsString(vals["ca_cert_identifier"])
}

func DecodeDocdbClusterInstance_Identifier(p *DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	p.Identifier = ctwhy.ValueAsString(vals["identifier"])
}

func DecodeDocdbClusterInstance_IdentifierPrefix(p *DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	p.IdentifierPrefix = ctwhy.ValueAsString(vals["identifier_prefix"])
}

func DecodeDocdbClusterInstance_AutoMinorVersionUpgrade(p *DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	p.AutoMinorVersionUpgrade = ctwhy.ValueAsBool(vals["auto_minor_version_upgrade"])
}

func DecodeDocdbClusterInstance_PreferredMaintenanceWindow(p *DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	p.PreferredMaintenanceWindow = ctwhy.ValueAsString(vals["preferred_maintenance_window"])
}

func DecodeDocdbClusterInstance_ApplyImmediately(p *DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	p.ApplyImmediately = ctwhy.ValueAsBool(vals["apply_immediately"])
}

func DecodeDocdbClusterInstance_Id(p *DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeDocdbClusterInstance_PromotionTier(p *DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	p.PromotionTier = ctwhy.ValueAsInt64(vals["promotion_tier"])
}

func DecodeDocdbClusterInstance_Tags(p *DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeDocdbClusterInstance_ClusterIdentifier(p *DocdbClusterInstanceParameters, vals map[string]cty.Value) {
	p.ClusterIdentifier = ctwhy.ValueAsString(vals["cluster_identifier"])
}

func DecodeDocdbClusterInstance_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDocdbClusterInstance_Timeouts_Create(p, valMap)
	DecodeDocdbClusterInstance_Timeouts_Delete(p, valMap)
	DecodeDocdbClusterInstance_Timeouts_Update(p, valMap)
}

func DecodeDocdbClusterInstance_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeDocdbClusterInstance_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeDocdbClusterInstance_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

func DecodeDocdbClusterInstance_EngineVersion(p *DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	p.EngineVersion = ctwhy.ValueAsString(vals["engine_version"])
}

func DecodeDocdbClusterInstance_StorageEncrypted(p *DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	p.StorageEncrypted = ctwhy.ValueAsBool(vals["storage_encrypted"])
}

func DecodeDocdbClusterInstance_DbSubnetGroupName(p *DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	p.DbSubnetGroupName = ctwhy.ValueAsString(vals["db_subnet_group_name"])
}

func DecodeDocdbClusterInstance_Port(p *DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	p.Port = ctwhy.ValueAsInt64(vals["port"])
}

func DecodeDocdbClusterInstance_Arn(p *DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeDocdbClusterInstance_KmsKeyId(p *DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

func DecodeDocdbClusterInstance_PreferredBackupWindow(p *DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	p.PreferredBackupWindow = ctwhy.ValueAsString(vals["preferred_backup_window"])
}

func DecodeDocdbClusterInstance_DbiResourceId(p *DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	p.DbiResourceId = ctwhy.ValueAsString(vals["dbi_resource_id"])
}

func DecodeDocdbClusterInstance_PubliclyAccessible(p *DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	p.PubliclyAccessible = ctwhy.ValueAsBool(vals["publicly_accessible"])
}

func DecodeDocdbClusterInstance_Writer(p *DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	p.Writer = ctwhy.ValueAsBool(vals["writer"])
}

func DecodeDocdbClusterInstance_Endpoint(p *DocdbClusterInstanceObservation, vals map[string]cty.Value) {
	p.Endpoint = ctwhy.ValueAsString(vals["endpoint"])
}