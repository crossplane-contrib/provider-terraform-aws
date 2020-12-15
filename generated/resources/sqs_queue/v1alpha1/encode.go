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
	r, ok := mr.(*SqsQueue)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SqsQueue.")
	}
	return EncodeSqsQueue(*r), nil
}

func EncodeSqsQueue(r SqsQueue) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSqsQueue_MessageRetentionSeconds(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_Name(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_ReceiveWaitTimeSeconds(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_DelaySeconds(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_MaxMessageSize(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_VisibilityTimeoutSeconds(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_FifoQueue(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_RedrivePolicy(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_KmsDataKeyReusePeriodSeconds(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_KmsMasterKeyId(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_Policy(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_ContentBasedDeduplication(r.Spec.ForProvider, ctyVal)
	EncodeSqsQueue_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeSqsQueue_MessageRetentionSeconds(p SqsQueueParameters, vals map[string]cty.Value) {
	vals["message_retention_seconds"] = cty.NumberIntVal(p.MessageRetentionSeconds)
}

func EncodeSqsQueue_Name(p SqsQueueParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSqsQueue_NamePrefix(p SqsQueueParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeSqsQueue_ReceiveWaitTimeSeconds(p SqsQueueParameters, vals map[string]cty.Value) {
	vals["receive_wait_time_seconds"] = cty.NumberIntVal(p.ReceiveWaitTimeSeconds)
}

func EncodeSqsQueue_DelaySeconds(p SqsQueueParameters, vals map[string]cty.Value) {
	vals["delay_seconds"] = cty.NumberIntVal(p.DelaySeconds)
}

func EncodeSqsQueue_MaxMessageSize(p SqsQueueParameters, vals map[string]cty.Value) {
	vals["max_message_size"] = cty.NumberIntVal(p.MaxMessageSize)
}

func EncodeSqsQueue_VisibilityTimeoutSeconds(p SqsQueueParameters, vals map[string]cty.Value) {
	vals["visibility_timeout_seconds"] = cty.NumberIntVal(p.VisibilityTimeoutSeconds)
}

func EncodeSqsQueue_FifoQueue(p SqsQueueParameters, vals map[string]cty.Value) {
	vals["fifo_queue"] = cty.BoolVal(p.FifoQueue)
}

func EncodeSqsQueue_RedrivePolicy(p SqsQueueParameters, vals map[string]cty.Value) {
	vals["redrive_policy"] = cty.StringVal(p.RedrivePolicy)
}

func EncodeSqsQueue_KmsDataKeyReusePeriodSeconds(p SqsQueueParameters, vals map[string]cty.Value) {
	vals["kms_data_key_reuse_period_seconds"] = cty.NumberIntVal(p.KmsDataKeyReusePeriodSeconds)
}

func EncodeSqsQueue_KmsMasterKeyId(p SqsQueueParameters, vals map[string]cty.Value) {
	vals["kms_master_key_id"] = cty.StringVal(p.KmsMasterKeyId)
}

func EncodeSqsQueue_Policy(p SqsQueueParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeSqsQueue_Tags(p SqsQueueParameters, vals map[string]cty.Value) {
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

func EncodeSqsQueue_ContentBasedDeduplication(p SqsQueueParameters, vals map[string]cty.Value) {
	vals["content_based_deduplication"] = cty.BoolVal(p.ContentBasedDeduplication)
}

func EncodeSqsQueue_Arn(p SqsQueueObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}