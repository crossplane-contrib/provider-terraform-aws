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
	r, ok := mr.(*DbClusterSnapshot)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDbClusterSnapshot(r, ctyValue)
}

func DecodeDbClusterSnapshot(prev *DbClusterSnapshot, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDbClusterSnapshot_DbClusterSnapshotIdentifier(&new.Spec.ForProvider, valMap)
	DecodeDbClusterSnapshot_Tags(&new.Spec.ForProvider, valMap)
	DecodeDbClusterSnapshot_Id(&new.Spec.ForProvider, valMap)
	DecodeDbClusterSnapshot_DbClusterIdentifier(&new.Spec.ForProvider, valMap)
	DecodeDbClusterSnapshot_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDbClusterSnapshot_AvailabilityZones(&new.Status.AtProvider, valMap)
	DecodeDbClusterSnapshot_VpcId(&new.Status.AtProvider, valMap)
	DecodeDbClusterSnapshot_Engine(&new.Status.AtProvider, valMap)
	DecodeDbClusterSnapshot_KmsKeyId(&new.Status.AtProvider, valMap)
	DecodeDbClusterSnapshot_Status(&new.Status.AtProvider, valMap)
	DecodeDbClusterSnapshot_DbClusterSnapshotArn(&new.Status.AtProvider, valMap)
	DecodeDbClusterSnapshot_EngineVersion(&new.Status.AtProvider, valMap)
	DecodeDbClusterSnapshot_SnapshotType(&new.Status.AtProvider, valMap)
	DecodeDbClusterSnapshot_SourceDbClusterSnapshotArn(&new.Status.AtProvider, valMap)
	DecodeDbClusterSnapshot_StorageEncrypted(&new.Status.AtProvider, valMap)
	DecodeDbClusterSnapshot_AllocatedStorage(&new.Status.AtProvider, valMap)
	DecodeDbClusterSnapshot_LicenseModel(&new.Status.AtProvider, valMap)
	DecodeDbClusterSnapshot_Port(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_DbClusterSnapshotIdentifier(p *DbClusterSnapshotParameters, vals map[string]cty.Value) {
	p.DbClusterSnapshotIdentifier = ctwhy.ValueAsString(vals["db_cluster_snapshot_identifier"])
}

//primitiveMapTypeDecodeTemplate
func DecodeDbClusterSnapshot_Tags(p *DbClusterSnapshotParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_Id(p *DbClusterSnapshotParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_DbClusterIdentifier(p *DbClusterSnapshotParameters, vals map[string]cty.Value) {
	p.DbClusterIdentifier = ctwhy.ValueAsString(vals["db_cluster_identifier"])
}

//containerTypeDecodeTemplate
func DecodeDbClusterSnapshot_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDbClusterSnapshot_Timeouts_Create(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDbClusterSnapshot_AvailabilityZones(p *DbClusterSnapshotObservation, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["availability_zones"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.AvailabilityZones = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_VpcId(p *DbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.VpcId = ctwhy.ValueAsString(vals["vpc_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_Engine(p *DbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.Engine = ctwhy.ValueAsString(vals["engine"])
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_KmsKeyId(p *DbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_Status(p *DbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.Status = ctwhy.ValueAsString(vals["status"])
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_DbClusterSnapshotArn(p *DbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.DbClusterSnapshotArn = ctwhy.ValueAsString(vals["db_cluster_snapshot_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_EngineVersion(p *DbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.EngineVersion = ctwhy.ValueAsString(vals["engine_version"])
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_SnapshotType(p *DbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.SnapshotType = ctwhy.ValueAsString(vals["snapshot_type"])
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_SourceDbClusterSnapshotArn(p *DbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.SourceDbClusterSnapshotArn = ctwhy.ValueAsString(vals["source_db_cluster_snapshot_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_StorageEncrypted(p *DbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.StorageEncrypted = ctwhy.ValueAsBool(vals["storage_encrypted"])
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_AllocatedStorage(p *DbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.AllocatedStorage = ctwhy.ValueAsInt64(vals["allocated_storage"])
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_LicenseModel(p *DbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.LicenseModel = ctwhy.ValueAsString(vals["license_model"])
}

//primitiveTypeDecodeTemplate
func DecodeDbClusterSnapshot_Port(p *DbClusterSnapshotObservation, vals map[string]cty.Value) {
	p.Port = ctwhy.ValueAsInt64(vals["port"])
}