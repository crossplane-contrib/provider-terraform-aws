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
	r, ok := mr.(*KinesisStream)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a KinesisStream.")
	}
	return EncodeKinesisStream(*r), nil
}

func EncodeKinesisStream(r KinesisStream) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeKinesisStream_Id(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_Name(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_Tags(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_EncryptionType(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_EnforceConsumerDeletion(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_RetentionPeriod(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_ShardCount(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_ShardLevelMetrics(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_Arn(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeKinesisStream_Id(p KinesisStreamParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKinesisStream_KmsKeyId(p KinesisStreamParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeKinesisStream_Name(p KinesisStreamParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeKinesisStream_Tags(p KinesisStreamParameters, vals map[string]cty.Value) {
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

func EncodeKinesisStream_EncryptionType(p KinesisStreamParameters, vals map[string]cty.Value) {
	vals["encryption_type"] = cty.StringVal(p.EncryptionType)
}

func EncodeKinesisStream_EnforceConsumerDeletion(p KinesisStreamParameters, vals map[string]cty.Value) {
	vals["enforce_consumer_deletion"] = cty.BoolVal(p.EnforceConsumerDeletion)
}

func EncodeKinesisStream_RetentionPeriod(p KinesisStreamParameters, vals map[string]cty.Value) {
	vals["retention_period"] = cty.NumberIntVal(p.RetentionPeriod)
}

func EncodeKinesisStream_ShardCount(p KinesisStreamParameters, vals map[string]cty.Value) {
	vals["shard_count"] = cty.NumberIntVal(p.ShardCount)
}

func EncodeKinesisStream_ShardLevelMetrics(p KinesisStreamParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ShardLevelMetrics {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["shard_level_metrics"] = cty.SetVal(colVals)
}

func EncodeKinesisStream_Arn(p KinesisStreamParameters, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeKinesisStream_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeKinesisStream_Timeouts_Create(p, ctyVal)
	EncodeKinesisStream_Timeouts_Delete(p, ctyVal)
	EncodeKinesisStream_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeKinesisStream_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeKinesisStream_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeKinesisStream_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}