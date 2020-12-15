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
	r, ok := mr.(*DbSnapshot)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDbSnapshot(r, ctyValue)
}

func DecodeDbSnapshot(prev *DbSnapshot, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDbSnapshot_DbInstanceIdentifier(&new.Spec.ForProvider, valMap)
	DecodeDbSnapshot_DbSnapshotIdentifier(&new.Spec.ForProvider, valMap)
	DecodeDbSnapshot_Id(&new.Spec.ForProvider, valMap)
	DecodeDbSnapshot_Tags(&new.Spec.ForProvider, valMap)
	DecodeDbSnapshot_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDbSnapshot_LicenseModel(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_OptionGroupName(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_Port(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_SnapshotType(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_SourceDbSnapshotIdentifier(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_Engine(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_EngineVersion(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_Iops(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_KmsKeyId(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_Status(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_StorageType(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_AllocatedStorage(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_AvailabilityZone(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_DbSnapshotArn(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_Encrypted(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_SourceRegion(&new.Status.AtProvider, valMap)
	DecodeDbSnapshot_VpcId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_DbInstanceIdentifier(p *DbSnapshotParameters, vals map[string]cty.Value) {
	p.DbInstanceIdentifier = ctwhy.ValueAsString(vals["db_instance_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_DbSnapshotIdentifier(p *DbSnapshotParameters, vals map[string]cty.Value) {
	p.DbSnapshotIdentifier = ctwhy.ValueAsString(vals["db_snapshot_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_Id(p *DbSnapshotParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeDbSnapshot_Tags(p *DbSnapshotParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//containerTypeDecodeTemplate
func DecodeDbSnapshot_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDbSnapshot_Timeouts_Read(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_Timeouts_Read(p *Timeouts, vals map[string]cty.Value) {
	p.Read = ctwhy.ValueAsString(vals["read"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_LicenseModel(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.LicenseModel = ctwhy.ValueAsString(vals["license_model"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_OptionGroupName(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.OptionGroupName = ctwhy.ValueAsString(vals["option_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_Port(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.Port = ctwhy.ValueAsInt64(vals["port"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_SnapshotType(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.SnapshotType = ctwhy.ValueAsString(vals["snapshot_type"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_SourceDbSnapshotIdentifier(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.SourceDbSnapshotIdentifier = ctwhy.ValueAsString(vals["source_db_snapshot_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_Engine(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.Engine = ctwhy.ValueAsString(vals["engine"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_EngineVersion(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.EngineVersion = ctwhy.ValueAsString(vals["engine_version"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_Iops(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.Iops = ctwhy.ValueAsInt64(vals["iops"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_KmsKeyId(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_Status(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.Status = ctwhy.ValueAsString(vals["status"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_StorageType(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.StorageType = ctwhy.ValueAsString(vals["storage_type"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_AllocatedStorage(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.AllocatedStorage = ctwhy.ValueAsInt64(vals["allocated_storage"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_AvailabilityZone(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.AvailabilityZone = ctwhy.ValueAsString(vals["availability_zone"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_DbSnapshotArn(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.DbSnapshotArn = ctwhy.ValueAsString(vals["db_snapshot_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_Encrypted(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.Encrypted = ctwhy.ValueAsBool(vals["encrypted"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_SourceRegion(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.SourceRegion = ctwhy.ValueAsString(vals["source_region"])
}

//primitiveTypeDecodeTemplate
func DecodeDbSnapshot_VpcId(p *DbSnapshotObservation, vals map[string]cty.Value) {
	p.VpcId = ctwhy.ValueAsString(vals["vpc_id"])
}