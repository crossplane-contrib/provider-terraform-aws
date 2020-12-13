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
	r, ok := mr.(*S3BucketNotification)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a S3BucketNotification.")
	}
	return EncodeS3BucketNotification(*r), nil
}

func EncodeS3BucketNotification(r S3BucketNotification) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketNotification_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketNotification_Id(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketNotification_LambdaFunction(r.Spec.ForProvider.LambdaFunction, ctyVal)
	EncodeS3BucketNotification_Queue(r.Spec.ForProvider.Queue, ctyVal)
	EncodeS3BucketNotification_Topic(r.Spec.ForProvider.Topic, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeS3BucketNotification_Bucket(p S3BucketNotificationParameters, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeS3BucketNotification_Id(p S3BucketNotificationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3BucketNotification_LambdaFunction(p LambdaFunction, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketNotification_LambdaFunction_Id(p, ctyVal)
	EncodeS3BucketNotification_LambdaFunction_LambdaFunctionArn(p, ctyVal)
	EncodeS3BucketNotification_LambdaFunction_Events(p, ctyVal)
	EncodeS3BucketNotification_LambdaFunction_FilterPrefix(p, ctyVal)
	EncodeS3BucketNotification_LambdaFunction_FilterSuffix(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["lambda_function"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketNotification_LambdaFunction_Id(p LambdaFunction, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3BucketNotification_LambdaFunction_LambdaFunctionArn(p LambdaFunction, vals map[string]cty.Value) {
	vals["lambda_function_arn"] = cty.StringVal(p.LambdaFunctionArn)
}

func EncodeS3BucketNotification_LambdaFunction_Events(p LambdaFunction, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Events {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["events"] = cty.SetVal(colVals)
}

func EncodeS3BucketNotification_LambdaFunction_FilterPrefix(p LambdaFunction, vals map[string]cty.Value) {
	vals["filter_prefix"] = cty.StringVal(p.FilterPrefix)
}

func EncodeS3BucketNotification_LambdaFunction_FilterSuffix(p LambdaFunction, vals map[string]cty.Value) {
	vals["filter_suffix"] = cty.StringVal(p.FilterSuffix)
}

func EncodeS3BucketNotification_Queue(p Queue, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketNotification_Queue_FilterPrefix(p, ctyVal)
	EncodeS3BucketNotification_Queue_FilterSuffix(p, ctyVal)
	EncodeS3BucketNotification_Queue_Id(p, ctyVal)
	EncodeS3BucketNotification_Queue_QueueArn(p, ctyVal)
	EncodeS3BucketNotification_Queue_Events(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["queue"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketNotification_Queue_FilterPrefix(p Queue, vals map[string]cty.Value) {
	vals["filter_prefix"] = cty.StringVal(p.FilterPrefix)
}

func EncodeS3BucketNotification_Queue_FilterSuffix(p Queue, vals map[string]cty.Value) {
	vals["filter_suffix"] = cty.StringVal(p.FilterSuffix)
}

func EncodeS3BucketNotification_Queue_Id(p Queue, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3BucketNotification_Queue_QueueArn(p Queue, vals map[string]cty.Value) {
	vals["queue_arn"] = cty.StringVal(p.QueueArn)
}

func EncodeS3BucketNotification_Queue_Events(p Queue, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Events {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["events"] = cty.SetVal(colVals)
}

func EncodeS3BucketNotification_Topic(p Topic, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketNotification_Topic_Events(p, ctyVal)
	EncodeS3BucketNotification_Topic_FilterPrefix(p, ctyVal)
	EncodeS3BucketNotification_Topic_FilterSuffix(p, ctyVal)
	EncodeS3BucketNotification_Topic_Id(p, ctyVal)
	EncodeS3BucketNotification_Topic_TopicArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["topic"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketNotification_Topic_Events(p Topic, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Events {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["events"] = cty.SetVal(colVals)
}

func EncodeS3BucketNotification_Topic_FilterPrefix(p Topic, vals map[string]cty.Value) {
	vals["filter_prefix"] = cty.StringVal(p.FilterPrefix)
}

func EncodeS3BucketNotification_Topic_FilterSuffix(p Topic, vals map[string]cty.Value) {
	vals["filter_suffix"] = cty.StringVal(p.FilterSuffix)
}

func EncodeS3BucketNotification_Topic_Id(p Topic, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3BucketNotification_Topic_TopicArn(p Topic, vals map[string]cty.Value) {
	vals["topic_arn"] = cty.StringVal(p.TopicArn)
}