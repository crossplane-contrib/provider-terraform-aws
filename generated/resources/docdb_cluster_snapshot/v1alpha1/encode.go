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
	r, ok := mr.(*DocdbClusterSnapshot)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DocdbClusterSnapshot.")
	}
	return EncodeDocdbClusterSnapshot(*r), nil
}

func EncodeDocdbClusterSnapshot(r DocdbClusterSnapshot) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDocdbClusterSnapshot_DbClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterSnapshot_DbClusterSnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterSnapshot_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDocdbClusterSnapshot_SnapshotType(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterSnapshot_SourceDbClusterSnapshotArn(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterSnapshot_StorageEncrypted(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterSnapshot_Engine(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterSnapshot_EngineVersion(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterSnapshot_Status(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterSnapshot_AvailabilityZones(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterSnapshot_DbClusterSnapshotArn(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterSnapshot_KmsKeyId(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterSnapshot_Port(r.Status.AtProvider, ctyVal)
	EncodeDocdbClusterSnapshot_VpcId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDocdbClusterSnapshot_DbClusterIdentifier(p DocdbClusterSnapshotParameters, vals map[string]cty.Value) {
	vals["db_cluster_identifier"] = cty.StringVal(p.DbClusterIdentifier)
}

func EncodeDocdbClusterSnapshot_DbClusterSnapshotIdentifier(p DocdbClusterSnapshotParameters, vals map[string]cty.Value) {
	vals["db_cluster_snapshot_identifier"] = cty.StringVal(p.DbClusterSnapshotIdentifier)
}

func EncodeDocdbClusterSnapshot_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDocdbClusterSnapshot_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDocdbClusterSnapshot_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDocdbClusterSnapshot_SnapshotType(p DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["snapshot_type"] = cty.StringVal(p.SnapshotType)
}

func EncodeDocdbClusterSnapshot_SourceDbClusterSnapshotArn(p DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["source_db_cluster_snapshot_arn"] = cty.StringVal(p.SourceDbClusterSnapshotArn)
}

func EncodeDocdbClusterSnapshot_StorageEncrypted(p DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["storage_encrypted"] = cty.BoolVal(p.StorageEncrypted)
}

func EncodeDocdbClusterSnapshot_Engine(p DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeDocdbClusterSnapshot_EngineVersion(p DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeDocdbClusterSnapshot_Status(p DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeDocdbClusterSnapshot_AvailabilityZones(p DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.ListVal(colVals)
}

func EncodeDocdbClusterSnapshot_DbClusterSnapshotArn(p DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["db_cluster_snapshot_arn"] = cty.StringVal(p.DbClusterSnapshotArn)
}

func EncodeDocdbClusterSnapshot_KmsKeyId(p DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeDocdbClusterSnapshot_Port(p DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeDocdbClusterSnapshot_VpcId(p DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}