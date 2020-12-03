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

package v1alpha1func EncodeS3BucketNotification(r S3BucketNotification) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeS3BucketNotification_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketNotification_Id(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketNotification_LambdaFunction(r.Spec.ForProvider.LambdaFunction, ctyVal)
	EncodeS3BucketNotification_Queue(r.Spec.ForProvider.Queue, ctyVal)
	EncodeS3BucketNotification_Topic(r.Spec.ForProvider.Topic, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeS3BucketNotification_Bucket(p *S3BucketNotificationParameters, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeS3BucketNotification_Id(p *S3BucketNotificationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3BucketNotification_LambdaFunction(p *LambdaFunction, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LambdaFunction {
		ctyVal = make(map[string]cty.Value)
		EncodeS3BucketNotification_LambdaFunction_FilterSuffix(v, ctyVal)
		EncodeS3BucketNotification_LambdaFunction_Id(v, ctyVal)
		EncodeS3BucketNotification_LambdaFunction_LambdaFunctionArn(v, ctyVal)
		EncodeS3BucketNotification_LambdaFunction_Events(v, ctyVal)
		EncodeS3BucketNotification_LambdaFunction_FilterPrefix(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["lambda_function"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketNotification_LambdaFunction_FilterSuffix(p *LambdaFunction, vals map[string]cty.Value) {
	vals["filter_suffix"] = cty.StringVal(p.FilterSuffix)
}

func EncodeS3BucketNotification_LambdaFunction_Id(p *LambdaFunction, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3BucketNotification_LambdaFunction_LambdaFunctionArn(p *LambdaFunction, vals map[string]cty.Value) {
	vals["lambda_function_arn"] = cty.StringVal(p.LambdaFunctionArn)
}

func EncodeS3BucketNotification_LambdaFunction_Events(p *LambdaFunction, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Events {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["events"] = cty.SetVal(colVals)
}

func EncodeS3BucketNotification_LambdaFunction_FilterPrefix(p *LambdaFunction, vals map[string]cty.Value) {
	vals["filter_prefix"] = cty.StringVal(p.FilterPrefix)
}

func EncodeS3BucketNotification_Queue(p *Queue, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Queue {
		ctyVal = make(map[string]cty.Value)
		EncodeS3BucketNotification_Queue_Id(v, ctyVal)
		EncodeS3BucketNotification_Queue_QueueArn(v, ctyVal)
		EncodeS3BucketNotification_Queue_Events(v, ctyVal)
		EncodeS3BucketNotification_Queue_FilterPrefix(v, ctyVal)
		EncodeS3BucketNotification_Queue_FilterSuffix(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["queue"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketNotification_Queue_Id(p *Queue, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3BucketNotification_Queue_QueueArn(p *Queue, vals map[string]cty.Value) {
	vals["queue_arn"] = cty.StringVal(p.QueueArn)
}

func EncodeS3BucketNotification_Queue_Events(p *Queue, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Events {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["events"] = cty.SetVal(colVals)
}

func EncodeS3BucketNotification_Queue_FilterPrefix(p *Queue, vals map[string]cty.Value) {
	vals["filter_prefix"] = cty.StringVal(p.FilterPrefix)
}

func EncodeS3BucketNotification_Queue_FilterSuffix(p *Queue, vals map[string]cty.Value) {
	vals["filter_suffix"] = cty.StringVal(p.FilterSuffix)
}

func EncodeS3BucketNotification_Topic(p *Topic, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Topic {
		ctyVal = make(map[string]cty.Value)
		EncodeS3BucketNotification_Topic_Events(v, ctyVal)
		EncodeS3BucketNotification_Topic_FilterPrefix(v, ctyVal)
		EncodeS3BucketNotification_Topic_FilterSuffix(v, ctyVal)
		EncodeS3BucketNotification_Topic_Id(v, ctyVal)
		EncodeS3BucketNotification_Topic_TopicArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["topic"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketNotification_Topic_Events(p *Topic, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Events {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["events"] = cty.SetVal(colVals)
}

func EncodeS3BucketNotification_Topic_FilterPrefix(p *Topic, vals map[string]cty.Value) {
	vals["filter_prefix"] = cty.StringVal(p.FilterPrefix)
}

func EncodeS3BucketNotification_Topic_FilterSuffix(p *Topic, vals map[string]cty.Value) {
	vals["filter_suffix"] = cty.StringVal(p.FilterSuffix)
}

func EncodeS3BucketNotification_Topic_Id(p *Topic, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3BucketNotification_Topic_TopicArn(p *Topic, vals map[string]cty.Value) {
	vals["topic_arn"] = cty.StringVal(p.TopicArn)
}