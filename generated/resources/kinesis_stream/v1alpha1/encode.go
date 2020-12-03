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

package v1alpha1func EncodeKinesisStream(r KinesisStream) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeKinesisStream_Arn(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_EncryptionType(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_EnforceConsumerDeletion(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_Id(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_Name(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_RetentionPeriod(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_ShardCount(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_ShardLevelMetrics(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_Tags(r.Spec.ForProvider, ctyVal)
	EncodeKinesisStream_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeKinesisStream_Arn(p *KinesisStreamParameters, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeKinesisStream_EncryptionType(p *KinesisStreamParameters, vals map[string]cty.Value) {
	vals["encryption_type"] = cty.StringVal(p.EncryptionType)
}

func EncodeKinesisStream_EnforceConsumerDeletion(p *KinesisStreamParameters, vals map[string]cty.Value) {
	vals["enforce_consumer_deletion"] = cty.BoolVal(p.EnforceConsumerDeletion)
}

func EncodeKinesisStream_Id(p *KinesisStreamParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKinesisStream_KmsKeyId(p *KinesisStreamParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeKinesisStream_Name(p *KinesisStreamParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeKinesisStream_RetentionPeriod(p *KinesisStreamParameters, vals map[string]cty.Value) {
	vals["retention_period"] = cty.IntVal(p.RetentionPeriod)
}

func EncodeKinesisStream_ShardCount(p *KinesisStreamParameters, vals map[string]cty.Value) {
	vals["shard_count"] = cty.IntVal(p.ShardCount)
}

func EncodeKinesisStream_ShardLevelMetrics(p *KinesisStreamParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ShardLevelMetrics {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["shard_level_metrics"] = cty.SetVal(colVals)
}

func EncodeKinesisStream_Tags(p *KinesisStreamParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeKinesisStream_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeKinesisStream_Timeouts_Delete(p, ctyVal)
	EncodeKinesisStream_Timeouts_Update(p, ctyVal)
	EncodeKinesisStream_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeKinesisStream_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeKinesisStream_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeKinesisStream_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}