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
	r, ok := mr.(*DmsReplicationTask)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DmsReplicationTask.")
	}
	return EncodeDmsReplicationTask(*r), nil
}

func EncodeDmsReplicationTask(r DmsReplicationTask) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDmsReplicationTask_ReplicationTaskSettings(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_SourceEndpointArn(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_TableMappings(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_CdcStartTime(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_MigrationType(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_ReplicationInstanceArn(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_ReplicationTaskId(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_TargetEndpointArn(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_ReplicationTaskArn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeDmsReplicationTask_ReplicationTaskSettings(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["replication_task_settings"] = cty.StringVal(p.ReplicationTaskSettings)
}

func EncodeDmsReplicationTask_SourceEndpointArn(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["source_endpoint_arn"] = cty.StringVal(p.SourceEndpointArn)
}

func EncodeDmsReplicationTask_TableMappings(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["table_mappings"] = cty.StringVal(p.TableMappings)
}

func EncodeDmsReplicationTask_Tags(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
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

func EncodeDmsReplicationTask_CdcStartTime(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["cdc_start_time"] = cty.StringVal(p.CdcStartTime)
}

func EncodeDmsReplicationTask_MigrationType(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["migration_type"] = cty.StringVal(p.MigrationType)
}

func EncodeDmsReplicationTask_ReplicationInstanceArn(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["replication_instance_arn"] = cty.StringVal(p.ReplicationInstanceArn)
}

func EncodeDmsReplicationTask_ReplicationTaskId(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["replication_task_id"] = cty.StringVal(p.ReplicationTaskId)
}

func EncodeDmsReplicationTask_TargetEndpointArn(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["target_endpoint_arn"] = cty.StringVal(p.TargetEndpointArn)
}

func EncodeDmsReplicationTask_ReplicationTaskArn(p DmsReplicationTaskObservation, vals map[string]cty.Value) {
	vals["replication_task_arn"] = cty.StringVal(p.ReplicationTaskArn)
}