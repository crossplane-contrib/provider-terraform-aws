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
	r, ok := mr.(*NeptuneClusterSnapshot)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeNeptuneClusterSnapshot(r, ctyValue)
}

func DecodeNeptuneClusterSnapshot(prev *NeptuneClusterSnapshot, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeNeptuneClusterSnapshot_DbClusterIdentifier(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterSnapshot_DbClusterSnapshotIdentifier(&new.Spec.ForProvider, valMap)
	DecodeNeptuneClusterSnapshot_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeNeptuneClusterSnapshot_VpcId(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterSnapshot_AllocatedStorage(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterSnapshot_LicenseModel(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterSnapshot_StorageEncrypted(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterSnapshot_Port(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterSnapshot_Status(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterSnapshot_DbClusterSnapshotArn(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterSnapshot_EngineVersion(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterSnapshot_KmsKeyId(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterSnapshot_Engine(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterSnapshot_AvailabilityZones(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterSnapshot_SnapshotType(&new.Status.AtProvider, valMap)
	DecodeNeptuneClusterSnapshot_SourceDbClusterSnapshotArn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_DbClusterIdentifier(p *NeptuneClusterSnapshotParameters, vals map[string]cty.Value) {
	p.DbClusterIdentifier = ctwhy.ValueAsString(vals["db_cluster_identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_DbClusterSnapshotIdentifier(p *NeptuneClusterSnapshotParameters, vals map[string]cty.Value) {
	p.DbClusterSnapshotIdentifier = ctwhy.ValueAsString(vals["db_cluster_snapshot_identifier"])
}

//containerTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeNeptuneClusterSnapshot_Timeouts_Create(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_VpcId(p *NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	p.VpcId = ctwhy.ValueAsString(vals["vpc_id"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_AllocatedStorage(p *NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	p.AllocatedStorage = ctwhy.ValueAsInt64(vals["allocated_storage"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_LicenseModel(p *NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	p.LicenseModel = ctwhy.ValueAsString(vals["license_model"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_StorageEncrypted(p *NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	p.StorageEncrypted = ctwhy.ValueAsBool(vals["storage_encrypted"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_Port(p *NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	p.Port = ctwhy.ValueAsInt64(vals["port"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_Status(p *NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	p.Status = ctwhy.ValueAsString(vals["status"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_DbClusterSnapshotArn(p *NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	p.DbClusterSnapshotArn = ctwhy.ValueAsString(vals["db_cluster_snapshot_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_EngineVersion(p *NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	p.EngineVersion = ctwhy.ValueAsString(vals["engine_version"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_KmsKeyId(p *NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_Engine(p *NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	p.Engine = ctwhy.ValueAsString(vals["engine"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_AvailabilityZones(p *NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["availability_zones"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.AvailabilityZones = goVals
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_SnapshotType(p *NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	p.SnapshotType = ctwhy.ValueAsString(vals["snapshot_type"])
}

//primitiveTypeDecodeTemplate
func DecodeNeptuneClusterSnapshot_SourceDbClusterSnapshotArn(p *NeptuneClusterSnapshotObservation, vals map[string]cty.Value) {
	p.SourceDbClusterSnapshotArn = ctwhy.ValueAsString(vals["source_db_cluster_snapshot_arn"])
}