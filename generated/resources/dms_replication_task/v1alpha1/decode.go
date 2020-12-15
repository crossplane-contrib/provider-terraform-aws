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
	r, ok := mr.(*DmsReplicationTask)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDmsReplicationTask(r, ctyValue)
}

func DecodeDmsReplicationTask(prev *DmsReplicationTask, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDmsReplicationTask_Id(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationTask_MigrationType(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationTask_ReplicationInstanceArn(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationTask_ReplicationTaskSettings(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationTask_TableMappings(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationTask_Tags(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationTask_CdcStartTime(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationTask_ReplicationTaskId(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationTask_SourceEndpointArn(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationTask_TargetEndpointArn(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationTask_ReplicationTaskArn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDmsReplicationTask_Id(p *DmsReplicationTaskParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsReplicationTask_MigrationType(p *DmsReplicationTaskParameters, vals map[string]cty.Value) {
	p.MigrationType = ctwhy.ValueAsString(vals["migration_type"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsReplicationTask_ReplicationInstanceArn(p *DmsReplicationTaskParameters, vals map[string]cty.Value) {
	p.ReplicationInstanceArn = ctwhy.ValueAsString(vals["replication_instance_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsReplicationTask_ReplicationTaskSettings(p *DmsReplicationTaskParameters, vals map[string]cty.Value) {
	p.ReplicationTaskSettings = ctwhy.ValueAsString(vals["replication_task_settings"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsReplicationTask_TableMappings(p *DmsReplicationTaskParameters, vals map[string]cty.Value) {
	p.TableMappings = ctwhy.ValueAsString(vals["table_mappings"])
}

//primitiveMapTypeDecodeTemplate
func DecodeDmsReplicationTask_Tags(p *DmsReplicationTaskParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeDmsReplicationTask_CdcStartTime(p *DmsReplicationTaskParameters, vals map[string]cty.Value) {
	p.CdcStartTime = ctwhy.ValueAsString(vals["cdc_start_time"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsReplicationTask_ReplicationTaskId(p *DmsReplicationTaskParameters, vals map[string]cty.Value) {
	p.ReplicationTaskId = ctwhy.ValueAsString(vals["replication_task_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsReplicationTask_SourceEndpointArn(p *DmsReplicationTaskParameters, vals map[string]cty.Value) {
	p.SourceEndpointArn = ctwhy.ValueAsString(vals["source_endpoint_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsReplicationTask_TargetEndpointArn(p *DmsReplicationTaskParameters, vals map[string]cty.Value) {
	p.TargetEndpointArn = ctwhy.ValueAsString(vals["target_endpoint_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsReplicationTask_ReplicationTaskArn(p *DmsReplicationTaskObservation, vals map[string]cty.Value) {
	p.ReplicationTaskArn = ctwhy.ValueAsString(vals["replication_task_arn"])
}