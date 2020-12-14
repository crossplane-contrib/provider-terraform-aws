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
	r, ok := mr.(*NeptuneClusterSnapshot)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a NeptuneClusterSnapshot.")
	}
	return EncodeNeptuneClusterSnapshot(*r), nil
}

func EncodeNeptuneClusterSnapshot(r NeptuneClusterSnapshot) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeNeptuneClusterSnapshot_DbClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_Id(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_DbClusterSnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeNeptuneClusterSnapshot_AllocatedStorage(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_Port(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_StorageEncrypted(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_Engine(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_EngineVersion(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_LicenseModel(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_SourceDbClusterSnapshotArn(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_AvailabilityZones(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_Status(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_DbClusterSnapshotArn(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_KmsKeyId(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_SnapshotType(r.Status.AtProvider, ctyVal)
	EncodeNeptuneClusterSnapshot_VpcId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeNeptuneClusterSnapshot_DbClusterIdentifier(p NeptuneClusterSnapshotParameters, vals map[string]cty.Value) {
	vals["db_cluster_identifier"] = cty.StringVal(p.DbClusterIdentifier)
}

func EncodeNeptuneClusterSnapshot_Id(p NeptuneClusterSnapshotParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeNeptuneClusterSnapshot_DbClusterSnapshotIdentifier(p NeptuneClusterSnapshotParameters, vals map[string]cty.Value) {
	vals["db_cluster_snapshot_identifier"] = cty.StringVal(p.DbClusterSnapshotIdentifier)
}

func EncodeNeptuneClusterSnapshot_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeNeptuneClusterSnapshot_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeNeptuneClusterSnapshot_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeNeptuneClusterSnapshot_AllocatedStorage(p NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["allocated_storage"] = cty.NumberIntVal(p.AllocatedStorage)
}

func EncodeNeptuneClusterSnapshot_Port(p NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeNeptuneClusterSnapshot_StorageEncrypted(p NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["storage_encrypted"] = cty.BoolVal(p.StorageEncrypted)
}

func EncodeNeptuneClusterSnapshot_Engine(p NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeNeptuneClusterSnapshot_EngineVersion(p NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeNeptuneClusterSnapshot_LicenseModel(p NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["license_model"] = cty.StringVal(p.LicenseModel)
}

func EncodeNeptuneClusterSnapshot_SourceDbClusterSnapshotArn(p NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["source_db_cluster_snapshot_arn"] = cty.StringVal(p.SourceDbClusterSnapshotArn)
}

func EncodeNeptuneClusterSnapshot_AvailabilityZones(p NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.ListVal(colVals)
}

func EncodeNeptuneClusterSnapshot_Status(p NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeNeptuneClusterSnapshot_DbClusterSnapshotArn(p NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["db_cluster_snapshot_arn"] = cty.StringVal(p.DbClusterSnapshotArn)
}

func EncodeNeptuneClusterSnapshot_KmsKeyId(p NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeNeptuneClusterSnapshot_SnapshotType(p NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["snapshot_type"] = cty.StringVal(p.SnapshotType)
}

func EncodeNeptuneClusterSnapshot_VpcId(p NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}