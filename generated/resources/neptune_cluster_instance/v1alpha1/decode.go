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
	r, ok := mr.(*NeptuneClusterInstance)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeNeptuneClusterInstance(r, ctyValue)
}

func DecodeNeptuneClusterInstance(prev *NeptuneClusterInstance, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeNeptuneClusterInstance_NeptuneSubnetGroupName(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_AvailabilityZone(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_Engine(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_EngineVersion(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_Identifier(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_InstanceClass(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_NeptuneParameterGroupName(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_ApplyImmediately(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_ClusterIdentifier(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_Port(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_PubliclyAccessible(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_PromotionTier(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_AutoMinorVersionUpgrade(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_Id(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_IdentifierPrefix(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_PreferredBackupWindow(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_PreferredMaintenanceWindow(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_Tags(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterInstance_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeNeptuneClusterInstance_Endpoint(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterInstance_Arn(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterInstance_DbiResourceId(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterInstance_StorageEncrypted(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterInstance_Writer(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterInstance_Address(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterInstance_KmsKeyArn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_NeptuneSubnetGroupName(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.NeptuneSubnetGroupName = ctwhy.ValueAsString(vals["neptune_subnet_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_AvailabilityZone(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.AvailabilityZone = ctwhy.ValueAsString(vals["availability_zone"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_Engine(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.Engine = ctwhy.ValueAsString(vals["engine"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_EngineVersion(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.EngineVersion = ctwhy.ValueAsString(vals["engine_version"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_Identifier(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.Identifier = ctwhy.ValueAsString(vals["identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_InstanceClass(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.InstanceClass = ctwhy.ValueAsString(vals["instance_class"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_NeptuneParameterGroupName(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.NeptuneParameterGroupName = ctwhy.ValueAsString(vals["neptune_parameter_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_ApplyImmediately(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.ApplyImmediately = ctwhy.ValueAsBool(vals["apply_immediately"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_ClusterIdentifier(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.ClusterIdentifier = ctwhy.ValueAsString(vals["cluster_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_Port(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.Port = ctwhy.ValueAsInt64(vals["port"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_PubliclyAccessible(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.PubliclyAccessible = ctwhy.ValueAsBool(vals["publicly_accessible"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_PromotionTier(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.PromotionTier = ctwhy.ValueAsInt64(vals["promotion_tier"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_AutoMinorVersionUpgrade(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.AutoMinorVersionUpgrade = ctwhy.ValueAsBool(vals["auto_minor_version_upgrade"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_Id(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_IdentifierPrefix(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.IdentifierPrefix = ctwhy.ValueAsString(vals["identifier_prefix"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_PreferredBackupWindow(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.PreferredBackupWindow = ctwhy.ValueAsString(vals["preferred_backup_window"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_PreferredMaintenanceWindow(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	p.PreferredMaintenanceWindow = ctwhy.ValueAsString(vals["preferred_maintenance_window"])
}

//primitiveMapTypeDecodeTemplate
func DecodeNeptuneClusterInstance_Tags(p *NeptuneClusterInstanceParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//containerTypeDecodeTemplate
func DecodeNeptuneClusterInstance_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeNeptuneClusterInstance_Timeouts_Create(p, valMap)
	DecodeNeptuneClusterInstance_Timeouts_Delete(p, valMap)
	DecodeNeptuneClusterInstance_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_Endpoint(p *NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	p.Endpoint = ctwhy.ValueAsString(vals["endpoint"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_Arn(p *NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_DbiResourceId(p *NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	p.DbiResourceId = ctwhy.ValueAsString(vals["dbi_resource_id"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_StorageEncrypted(p *NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	p.StorageEncrypted = ctwhy.ValueAsBool(vals["storage_encrypted"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_Writer(p *NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	p.Writer = ctwhy.ValueAsBool(vals["writer"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_Address(p *NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	p.Address = ctwhy.ValueAsString(vals["address"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterInstance_KmsKeyArn(p *NeptuneClusterInstanceObservation, vals map[string]cty.Value) {
	p.KmsKeyArn = ctwhy.ValueAsString(vals["kms_key_arn"])
}