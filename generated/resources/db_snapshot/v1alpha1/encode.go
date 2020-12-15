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
	r, ok := mr.(*DbSnapshot)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DbSnapshot.")
	}
	return EncodeDbSnapshot(*r), nil
}

func EncodeDbSnapshot(r DbSnapshot) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDbSnapshot_DbInstanceIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDbSnapshot_DbSnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDbSnapshot_Id(r.Spec.ForProvider, ctyVal)
	EncodeDbSnapshot_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDbSnapshot_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDbSnapshot_LicenseModel(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_OptionGroupName(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_Port(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_SnapshotType(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_SourceDbSnapshotIdentifier(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_Engine(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_EngineVersion(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_Iops(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_KmsKeyId(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_Status(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_StorageType(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_AllocatedStorage(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_AvailabilityZone(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_DbSnapshotArn(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_Encrypted(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_SourceRegion(r.Status.AtProvider, ctyVal)
	EncodeDbSnapshot_VpcId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDbSnapshot_DbInstanceIdentifier(p DbSnapshotParameters, vals map[string]cty.Value) {
	vals["db_instance_identifier"] = cty.StringVal(p.DbInstanceIdentifier)
}

func EncodeDbSnapshot_DbSnapshotIdentifier(p DbSnapshotParameters, vals map[string]cty.Value) {
	vals["db_snapshot_identifier"] = cty.StringVal(p.DbSnapshotIdentifier)
}

func EncodeDbSnapshot_Id(p DbSnapshotParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDbSnapshot_Tags(p DbSnapshotParameters, vals map[string]cty.Value) {
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

func EncodeDbSnapshot_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDbSnapshot_Timeouts_Read(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDbSnapshot_Timeouts_Read(p Timeouts, vals map[string]cty.Value) {
	vals["read"] = cty.StringVal(p.Read)
}

func EncodeDbSnapshot_LicenseModel(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["license_model"] = cty.StringVal(p.LicenseModel)
}

func EncodeDbSnapshot_OptionGroupName(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["option_group_name"] = cty.StringVal(p.OptionGroupName)
}

func EncodeDbSnapshot_Port(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeDbSnapshot_SnapshotType(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["snapshot_type"] = cty.StringVal(p.SnapshotType)
}

func EncodeDbSnapshot_SourceDbSnapshotIdentifier(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["source_db_snapshot_identifier"] = cty.StringVal(p.SourceDbSnapshotIdentifier)
}

func EncodeDbSnapshot_Engine(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeDbSnapshot_EngineVersion(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeDbSnapshot_Iops(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeDbSnapshot_KmsKeyId(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeDbSnapshot_Status(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeDbSnapshot_StorageType(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["storage_type"] = cty.StringVal(p.StorageType)
}

func EncodeDbSnapshot_AllocatedStorage(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["allocated_storage"] = cty.NumberIntVal(p.AllocatedStorage)
}

func EncodeDbSnapshot_AvailabilityZone(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeDbSnapshot_DbSnapshotArn(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["db_snapshot_arn"] = cty.StringVal(p.DbSnapshotArn)
}

func EncodeDbSnapshot_Encrypted(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeDbSnapshot_SourceRegion(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["source_region"] = cty.StringVal(p.SourceRegion)
}

func EncodeDbSnapshot_VpcId(p DbSnapshotObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}