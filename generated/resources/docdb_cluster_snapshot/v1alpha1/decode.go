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
	r, ok := mr.(*DocdbClusterSnapshot)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDocdbClusterSnapshot(r, ctyValue)
}

func DecodeDocdbClusterSnapshot(prev *DocdbClusterSnapshot, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDocdbClusterSnapshot_DbClusterSnapshotIdentifier(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterSnapshot_DbClusterIdentifier(&new.Spec.ForProvider, valMap)
	DecodeDocdbClusterSnapshot_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDocdbClusterSnapshot_EngineVersion(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterSnapshot_DbClusterSnapshotArn(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterSnapshot_Status(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterSnapshot_Port(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterSnapshot_SourceDbClusterSnapshotArn(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterSnapshot_StorageEncrypted(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterSnapshot_VpcId(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterSnapshot_KmsKeyId(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterSnapshot_SnapshotType(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterSnapshot_AvailabilityZones(&new.Status.AtProvider, valMap)
	DecodeDocdbClusterSnapshot_Engine(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_DbClusterSnapshotIdentifier(p *DocdbClusterSnapshotParameters, vals map[string]cty.Value) {
	p.DbClusterSnapshotIdentifier = ctwhy.ValueAsString(vals["db_cluster_snapshot_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_DbClusterIdentifier(p *DocdbClusterSnapshotParameters, vals map[string]cty.Value) {
	p.DbClusterIdentifier = ctwhy.ValueAsString(vals["db_cluster_identifier"])
}

//containerTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDocdbClusterSnapshot_Timeouts_Create(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_EngineVersion(p *DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.EngineVersion = ctwhy.ValueAsString(vals["engine_version"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_DbClusterSnapshotArn(p *DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.DbClusterSnapshotArn = ctwhy.ValueAsString(vals["db_cluster_snapshot_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_Status(p *DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.Status = ctwhy.ValueAsString(vals["status"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_Port(p *DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.Port = ctwhy.ValueAsInt64(vals["port"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_SourceDbClusterSnapshotArn(p *DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.SourceDbClusterSnapshotArn = ctwhy.ValueAsString(vals["source_db_cluster_snapshot_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_StorageEncrypted(p *DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.StorageEncrypted = ctwhy.ValueAsBool(vals["storage_encrypted"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_VpcId(p *DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.VpcId = ctwhy.ValueAsString(vals["vpc_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_KmsKeyId(p *DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_SnapshotType(p *DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.SnapshotType = ctwhy.ValueAsString(vals["snapshot_type"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_AvailabilityZones(p *DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["availability_zones"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.AvailabilityZones = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDocdbClusterSnapshot_Engine(p *DocdbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.Engine = ctwhy.ValueAsString(vals["engine"])
}