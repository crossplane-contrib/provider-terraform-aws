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
	"github.com/zclconf/go-cty/cty"
)

func EncodeDmsReplicationTask(r DmsReplicationTask) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDmsReplicationTask_ReplicationTaskSettings(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_TableMappings(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_TargetEndpointArn(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_ReplicationInstanceArn(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_ReplicationTaskId(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_SourceEndpointArn(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_CdcStartTime(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_Id(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_MigrationType(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationTask_ReplicationTaskArn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDmsReplicationTask_ReplicationTaskSettings(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["replication_task_settings"] = cty.StringVal(p.ReplicationTaskSettings)
}

func EncodeDmsReplicationTask_TableMappings(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["table_mappings"] = cty.StringVal(p.TableMappings)
}

func EncodeDmsReplicationTask_Tags(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDmsReplicationTask_TargetEndpointArn(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["target_endpoint_arn"] = cty.StringVal(p.TargetEndpointArn)
}

func EncodeDmsReplicationTask_ReplicationInstanceArn(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["replication_instance_arn"] = cty.StringVal(p.ReplicationInstanceArn)
}

func EncodeDmsReplicationTask_ReplicationTaskId(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["replication_task_id"] = cty.StringVal(p.ReplicationTaskId)
}

func EncodeDmsReplicationTask_SourceEndpointArn(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["source_endpoint_arn"] = cty.StringVal(p.SourceEndpointArn)
}

func EncodeDmsReplicationTask_CdcStartTime(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["cdc_start_time"] = cty.StringVal(p.CdcStartTime)
}

func EncodeDmsReplicationTask_Id(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDmsReplicationTask_MigrationType(p DmsReplicationTaskParameters, vals map[string]cty.Value) {
	vals["migration_type"] = cty.StringVal(p.MigrationType)
}

func EncodeDmsReplicationTask_ReplicationTaskArn(p DmsReplicationTaskObservation, vals map[string]cty.Value) {
	vals["replication_task_arn"] = cty.StringVal(p.ReplicationTaskArn)
}