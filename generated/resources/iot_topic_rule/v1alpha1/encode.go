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

func EncodeIotTopicRule(r IotTopicRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_SqlVersion(r.Spec.ForProvider, ctyVal)
	EncodeIotTopicRule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeIotTopicRule_Description(r.Spec.ForProvider, ctyVal)
	EncodeIotTopicRule_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeIotTopicRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeIotTopicRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeIotTopicRule_Sql(r.Spec.ForProvider, ctyVal)
	EncodeIotTopicRule_Lambda(r.Spec.ForProvider.Lambda, ctyVal)
	EncodeIotTopicRule_CloudwatchMetric(r.Spec.ForProvider.CloudwatchMetric, ctyVal)
	EncodeIotTopicRule_Dynamodb(r.Spec.ForProvider.Dynamodb, ctyVal)
	EncodeIotTopicRule_Elasticsearch(r.Spec.ForProvider.Elasticsearch, ctyVal)
	EncodeIotTopicRule_IotAnalytics(r.Spec.ForProvider.IotAnalytics, ctyVal)
	EncodeIotTopicRule_CloudwatchAlarm(r.Spec.ForProvider.CloudwatchAlarm, ctyVal)
	EncodeIotTopicRule_Republish(r.Spec.ForProvider.Republish, ctyVal)
	EncodeIotTopicRule_Sns(r.Spec.ForProvider.Sns, ctyVal)
	EncodeIotTopicRule_StepFunctions(r.Spec.ForProvider.StepFunctions, ctyVal)
	EncodeIotTopicRule_Dynamodbv2(r.Spec.ForProvider.Dynamodbv2, ctyVal)
	EncodeIotTopicRule_Firehose(r.Spec.ForProvider.Firehose, ctyVal)
	EncodeIotTopicRule_IotEvents(r.Spec.ForProvider.IotEvents, ctyVal)
	EncodeIotTopicRule_Kinesis(r.Spec.ForProvider.Kinesis, ctyVal)
	EncodeIotTopicRule_ErrorAction(r.Spec.ForProvider.ErrorAction, ctyVal)
	EncodeIotTopicRule_S3(r.Spec.ForProvider.S3, ctyVal)
	EncodeIotTopicRule_Sqs(r.Spec.ForProvider.Sqs, ctyVal)
	EncodeIotTopicRule_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeIotTopicRule_SqlVersion(p IotTopicRuleParameters, vals map[string]cty.Value) {
	vals["sql_version"] = cty.StringVal(p.SqlVersion)
}

func EncodeIotTopicRule_Tags(p IotTopicRuleParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeIotTopicRule_Description(p IotTopicRuleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeIotTopicRule_Enabled(p IotTopicRuleParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeIotTopicRule_Id(p IotTopicRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIotTopicRule_Name(p IotTopicRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIotTopicRule_Sql(p IotTopicRuleParameters, vals map[string]cty.Value) {
	vals["sql"] = cty.StringVal(p.Sql)
}

func EncodeIotTopicRule_Lambda(p Lambda, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_Lambda_FunctionArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["lambda"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_Lambda_FunctionArn(p Lambda, vals map[string]cty.Value) {
	vals["function_arn"] = cty.StringVal(p.FunctionArn)
}

func EncodeIotTopicRule_CloudwatchMetric(p CloudwatchMetric, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_CloudwatchMetric_MetricName(p, ctyVal)
	EncodeIotTopicRule_CloudwatchMetric_MetricNamespace(p, ctyVal)
	EncodeIotTopicRule_CloudwatchMetric_MetricTimestamp(p, ctyVal)
	EncodeIotTopicRule_CloudwatchMetric_MetricUnit(p, ctyVal)
	EncodeIotTopicRule_CloudwatchMetric_MetricValue(p, ctyVal)
	EncodeIotTopicRule_CloudwatchMetric_RoleArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["cloudwatch_metric"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_CloudwatchMetric_MetricName(p CloudwatchMetric, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeIotTopicRule_CloudwatchMetric_MetricNamespace(p CloudwatchMetric, vals map[string]cty.Value) {
	vals["metric_namespace"] = cty.StringVal(p.MetricNamespace)
}

func EncodeIotTopicRule_CloudwatchMetric_MetricTimestamp(p CloudwatchMetric, vals map[string]cty.Value) {
	vals["metric_timestamp"] = cty.StringVal(p.MetricTimestamp)
}

func EncodeIotTopicRule_CloudwatchMetric_MetricUnit(p CloudwatchMetric, vals map[string]cty.Value) {
	vals["metric_unit"] = cty.StringVal(p.MetricUnit)
}

func EncodeIotTopicRule_CloudwatchMetric_MetricValue(p CloudwatchMetric, vals map[string]cty.Value) {
	vals["metric_value"] = cty.StringVal(p.MetricValue)
}

func EncodeIotTopicRule_CloudwatchMetric_RoleArn(p CloudwatchMetric, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_Dynamodb(p Dynamodb, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_Dynamodb_HashKeyType(p, ctyVal)
	EncodeIotTopicRule_Dynamodb_HashKeyValue(p, ctyVal)
	EncodeIotTopicRule_Dynamodb_Operation(p, ctyVal)
	EncodeIotTopicRule_Dynamodb_RangeKeyField(p, ctyVal)
	EncodeIotTopicRule_Dynamodb_RangeKeyType(p, ctyVal)
	EncodeIotTopicRule_Dynamodb_TableName(p, ctyVal)
	EncodeIotTopicRule_Dynamodb_HashKeyField(p, ctyVal)
	EncodeIotTopicRule_Dynamodb_RangeKeyValue(p, ctyVal)
	EncodeIotTopicRule_Dynamodb_RoleArn(p, ctyVal)
	EncodeIotTopicRule_Dynamodb_PayloadField(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["dynamodb"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_Dynamodb_HashKeyType(p Dynamodb, vals map[string]cty.Value) {
	vals["hash_key_type"] = cty.StringVal(p.HashKeyType)
}

func EncodeIotTopicRule_Dynamodb_HashKeyValue(p Dynamodb, vals map[string]cty.Value) {
	vals["hash_key_value"] = cty.StringVal(p.HashKeyValue)
}

func EncodeIotTopicRule_Dynamodb_Operation(p Dynamodb, vals map[string]cty.Value) {
	vals["operation"] = cty.StringVal(p.Operation)
}

func EncodeIotTopicRule_Dynamodb_RangeKeyField(p Dynamodb, vals map[string]cty.Value) {
	vals["range_key_field"] = cty.StringVal(p.RangeKeyField)
}

func EncodeIotTopicRule_Dynamodb_RangeKeyType(p Dynamodb, vals map[string]cty.Value) {
	vals["range_key_type"] = cty.StringVal(p.RangeKeyType)
}

func EncodeIotTopicRule_Dynamodb_TableName(p Dynamodb, vals map[string]cty.Value) {
	vals["table_name"] = cty.StringVal(p.TableName)
}

func EncodeIotTopicRule_Dynamodb_HashKeyField(p Dynamodb, vals map[string]cty.Value) {
	vals["hash_key_field"] = cty.StringVal(p.HashKeyField)
}

func EncodeIotTopicRule_Dynamodb_RangeKeyValue(p Dynamodb, vals map[string]cty.Value) {
	vals["range_key_value"] = cty.StringVal(p.RangeKeyValue)
}

func EncodeIotTopicRule_Dynamodb_RoleArn(p Dynamodb, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_Dynamodb_PayloadField(p Dynamodb, vals map[string]cty.Value) {
	vals["payload_field"] = cty.StringVal(p.PayloadField)
}

func EncodeIotTopicRule_Elasticsearch(p Elasticsearch, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_Elasticsearch_Index(p, ctyVal)
	EncodeIotTopicRule_Elasticsearch_RoleArn(p, ctyVal)
	EncodeIotTopicRule_Elasticsearch_Type(p, ctyVal)
	EncodeIotTopicRule_Elasticsearch_Endpoint(p, ctyVal)
	EncodeIotTopicRule_Elasticsearch_Id(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["elasticsearch"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_Elasticsearch_Index(p Elasticsearch, vals map[string]cty.Value) {
	vals["index"] = cty.StringVal(p.Index)
}

func EncodeIotTopicRule_Elasticsearch_RoleArn(p Elasticsearch, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_Elasticsearch_Type(p Elasticsearch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeIotTopicRule_Elasticsearch_Endpoint(p Elasticsearch, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeIotTopicRule_Elasticsearch_Id(p Elasticsearch, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIotTopicRule_IotAnalytics(p IotAnalytics, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_IotAnalytics_ChannelName(p, ctyVal)
	EncodeIotTopicRule_IotAnalytics_RoleArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["iot_analytics"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_IotAnalytics_ChannelName(p IotAnalytics, vals map[string]cty.Value) {
	vals["channel_name"] = cty.StringVal(p.ChannelName)
}

func EncodeIotTopicRule_IotAnalytics_RoleArn(p IotAnalytics, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_CloudwatchAlarm(p CloudwatchAlarm, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_CloudwatchAlarm_AlarmName(p, ctyVal)
	EncodeIotTopicRule_CloudwatchAlarm_RoleArn(p, ctyVal)
	EncodeIotTopicRule_CloudwatchAlarm_StateReason(p, ctyVal)
	EncodeIotTopicRule_CloudwatchAlarm_StateValue(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["cloudwatch_alarm"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_CloudwatchAlarm_AlarmName(p CloudwatchAlarm, vals map[string]cty.Value) {
	vals["alarm_name"] = cty.StringVal(p.AlarmName)
}

func EncodeIotTopicRule_CloudwatchAlarm_RoleArn(p CloudwatchAlarm, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_CloudwatchAlarm_StateReason(p CloudwatchAlarm, vals map[string]cty.Value) {
	vals["state_reason"] = cty.StringVal(p.StateReason)
}

func EncodeIotTopicRule_CloudwatchAlarm_StateValue(p CloudwatchAlarm, vals map[string]cty.Value) {
	vals["state_value"] = cty.StringVal(p.StateValue)
}

func EncodeIotTopicRule_Republish(p Republish, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_Republish_Qos(p, ctyVal)
	EncodeIotTopicRule_Republish_RoleArn(p, ctyVal)
	EncodeIotTopicRule_Republish_Topic(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["republish"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_Republish_Qos(p Republish, vals map[string]cty.Value) {
	vals["qos"] = cty.NumberIntVal(p.Qos)
}

func EncodeIotTopicRule_Republish_RoleArn(p Republish, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_Republish_Topic(p Republish, vals map[string]cty.Value) {
	vals["topic"] = cty.StringVal(p.Topic)
}

func EncodeIotTopicRule_Sns(p Sns, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_Sns_MessageFormat(p, ctyVal)
	EncodeIotTopicRule_Sns_RoleArn(p, ctyVal)
	EncodeIotTopicRule_Sns_TargetArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["sns"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_Sns_MessageFormat(p Sns, vals map[string]cty.Value) {
	vals["message_format"] = cty.StringVal(p.MessageFormat)
}

func EncodeIotTopicRule_Sns_RoleArn(p Sns, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_Sns_TargetArn(p Sns, vals map[string]cty.Value) {
	vals["target_arn"] = cty.StringVal(p.TargetArn)
}

func EncodeIotTopicRule_StepFunctions(p StepFunctions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_StepFunctions_ExecutionNamePrefix(p, ctyVal)
	EncodeIotTopicRule_StepFunctions_RoleArn(p, ctyVal)
	EncodeIotTopicRule_StepFunctions_StateMachineName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["step_functions"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_StepFunctions_ExecutionNamePrefix(p StepFunctions, vals map[string]cty.Value) {
	vals["execution_name_prefix"] = cty.StringVal(p.ExecutionNamePrefix)
}

func EncodeIotTopicRule_StepFunctions_RoleArn(p StepFunctions, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_StepFunctions_StateMachineName(p StepFunctions, vals map[string]cty.Value) {
	vals["state_machine_name"] = cty.StringVal(p.StateMachineName)
}

func EncodeIotTopicRule_Dynamodbv2(p Dynamodbv2, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_Dynamodbv2_RoleArn(p, ctyVal)
	EncodeIotTopicRule_Dynamodbv2_PutItem(p.PutItem, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["dynamodbv2"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_Dynamodbv2_RoleArn(p Dynamodbv2, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_Dynamodbv2_PutItem(p PutItem, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_Dynamodbv2_PutItem_TableName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["put_item"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_Dynamodbv2_PutItem_TableName(p PutItem, vals map[string]cty.Value) {
	vals["table_name"] = cty.StringVal(p.TableName)
}

func EncodeIotTopicRule_Firehose(p Firehose, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_Firehose_DeliveryStreamName(p, ctyVal)
	EncodeIotTopicRule_Firehose_RoleArn(p, ctyVal)
	EncodeIotTopicRule_Firehose_Separator(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["firehose"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_Firehose_DeliveryStreamName(p Firehose, vals map[string]cty.Value) {
	vals["delivery_stream_name"] = cty.StringVal(p.DeliveryStreamName)
}

func EncodeIotTopicRule_Firehose_RoleArn(p Firehose, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_Firehose_Separator(p Firehose, vals map[string]cty.Value) {
	vals["separator"] = cty.StringVal(p.Separator)
}

func EncodeIotTopicRule_IotEvents(p IotEvents, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_IotEvents_RoleArn(p, ctyVal)
	EncodeIotTopicRule_IotEvents_InputName(p, ctyVal)
	EncodeIotTopicRule_IotEvents_MessageId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["iot_events"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_IotEvents_RoleArn(p IotEvents, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_IotEvents_InputName(p IotEvents, vals map[string]cty.Value) {
	vals["input_name"] = cty.StringVal(p.InputName)
}

func EncodeIotTopicRule_IotEvents_MessageId(p IotEvents, vals map[string]cty.Value) {
	vals["message_id"] = cty.StringVal(p.MessageId)
}

func EncodeIotTopicRule_Kinesis(p Kinesis, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_Kinesis_PartitionKey(p, ctyVal)
	EncodeIotTopicRule_Kinesis_RoleArn(p, ctyVal)
	EncodeIotTopicRule_Kinesis_StreamName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["kinesis"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_Kinesis_PartitionKey(p Kinesis, vals map[string]cty.Value) {
	vals["partition_key"] = cty.StringVal(p.PartitionKey)
}

func EncodeIotTopicRule_Kinesis_RoleArn(p Kinesis, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_Kinesis_StreamName(p Kinesis, vals map[string]cty.Value) {
	vals["stream_name"] = cty.StringVal(p.StreamName)
}

func EncodeIotTopicRule_ErrorAction(p ErrorAction, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_Firehose(p.Firehose, ctyVal)
	EncodeIotTopicRule_ErrorAction_IotEvents(p.IotEvents, ctyVal)
	EncodeIotTopicRule_ErrorAction_CloudwatchAlarm(p.CloudwatchAlarm, ctyVal)
	EncodeIotTopicRule_ErrorAction_CloudwatchMetric(p.CloudwatchMetric, ctyVal)
	EncodeIotTopicRule_ErrorAction_Dynamodb(p.Dynamodb, ctyVal)
	EncodeIotTopicRule_ErrorAction_Republish(p.Republish, ctyVal)
	EncodeIotTopicRule_ErrorAction_S3(p.S3, ctyVal)
	EncodeIotTopicRule_ErrorAction_Dynamodbv2(p.Dynamodbv2, ctyVal)
	EncodeIotTopicRule_ErrorAction_Elasticsearch(p.Elasticsearch, ctyVal)
	EncodeIotTopicRule_ErrorAction_Lambda(p.Lambda, ctyVal)
	EncodeIotTopicRule_ErrorAction_IotAnalytics(p.IotAnalytics, ctyVal)
	EncodeIotTopicRule_ErrorAction_Sns(p.Sns, ctyVal)
	EncodeIotTopicRule_ErrorAction_Sqs(p.Sqs, ctyVal)
	EncodeIotTopicRule_ErrorAction_Kinesis(p.Kinesis, ctyVal)
	EncodeIotTopicRule_ErrorAction_StepFunctions(p.StepFunctions, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["error_action"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_Firehose(p Firehose, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_Firehose_DeliveryStreamName(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Firehose_RoleArn(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Firehose_Separator(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["firehose"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_Firehose_DeliveryStreamName(p Firehose, vals map[string]cty.Value) {
	vals["delivery_stream_name"] = cty.StringVal(p.DeliveryStreamName)
}

func EncodeIotTopicRule_ErrorAction_Firehose_RoleArn(p Firehose, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_Firehose_Separator(p Firehose, vals map[string]cty.Value) {
	vals["separator"] = cty.StringVal(p.Separator)
}

func EncodeIotTopicRule_ErrorAction_IotEvents(p IotEvents, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_IotEvents_InputName(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_IotEvents_MessageId(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_IotEvents_RoleArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["iot_events"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_IotEvents_InputName(p IotEvents, vals map[string]cty.Value) {
	vals["input_name"] = cty.StringVal(p.InputName)
}

func EncodeIotTopicRule_ErrorAction_IotEvents_MessageId(p IotEvents, vals map[string]cty.Value) {
	vals["message_id"] = cty.StringVal(p.MessageId)
}

func EncodeIotTopicRule_ErrorAction_IotEvents_RoleArn(p IotEvents, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_CloudwatchAlarm(p CloudwatchAlarm, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_CloudwatchAlarm_StateValue(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_CloudwatchAlarm_AlarmName(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_CloudwatchAlarm_RoleArn(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_CloudwatchAlarm_StateReason(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["cloudwatch_alarm"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_CloudwatchAlarm_StateValue(p CloudwatchAlarm, vals map[string]cty.Value) {
	vals["state_value"] = cty.StringVal(p.StateValue)
}

func EncodeIotTopicRule_ErrorAction_CloudwatchAlarm_AlarmName(p CloudwatchAlarm, vals map[string]cty.Value) {
	vals["alarm_name"] = cty.StringVal(p.AlarmName)
}

func EncodeIotTopicRule_ErrorAction_CloudwatchAlarm_RoleArn(p CloudwatchAlarm, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_CloudwatchAlarm_StateReason(p CloudwatchAlarm, vals map[string]cty.Value) {
	vals["state_reason"] = cty.StringVal(p.StateReason)
}

func EncodeIotTopicRule_ErrorAction_CloudwatchMetric(p CloudwatchMetric, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_CloudwatchMetric_MetricName(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_CloudwatchMetric_MetricNamespace(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_CloudwatchMetric_MetricTimestamp(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_CloudwatchMetric_MetricUnit(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_CloudwatchMetric_MetricValue(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_CloudwatchMetric_RoleArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["cloudwatch_metric"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_CloudwatchMetric_MetricName(p CloudwatchMetric, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeIotTopicRule_ErrorAction_CloudwatchMetric_MetricNamespace(p CloudwatchMetric, vals map[string]cty.Value) {
	vals["metric_namespace"] = cty.StringVal(p.MetricNamespace)
}

func EncodeIotTopicRule_ErrorAction_CloudwatchMetric_MetricTimestamp(p CloudwatchMetric, vals map[string]cty.Value) {
	vals["metric_timestamp"] = cty.StringVal(p.MetricTimestamp)
}

func EncodeIotTopicRule_ErrorAction_CloudwatchMetric_MetricUnit(p CloudwatchMetric, vals map[string]cty.Value) {
	vals["metric_unit"] = cty.StringVal(p.MetricUnit)
}

func EncodeIotTopicRule_ErrorAction_CloudwatchMetric_MetricValue(p CloudwatchMetric, vals map[string]cty.Value) {
	vals["metric_value"] = cty.StringVal(p.MetricValue)
}

func EncodeIotTopicRule_ErrorAction_CloudwatchMetric_RoleArn(p CloudwatchMetric, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_Dynamodb(p Dynamodb, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_Dynamodb_RoleArn(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Dynamodb_HashKeyType(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Dynamodb_HashKeyValue(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Dynamodb_Operation(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Dynamodb_PayloadField(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Dynamodb_RangeKeyField(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Dynamodb_RangeKeyValue(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Dynamodb_HashKeyField(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Dynamodb_RangeKeyType(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Dynamodb_TableName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["dynamodb"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_Dynamodb_RoleArn(p Dynamodb, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_Dynamodb_HashKeyType(p Dynamodb, vals map[string]cty.Value) {
	vals["hash_key_type"] = cty.StringVal(p.HashKeyType)
}

func EncodeIotTopicRule_ErrorAction_Dynamodb_HashKeyValue(p Dynamodb, vals map[string]cty.Value) {
	vals["hash_key_value"] = cty.StringVal(p.HashKeyValue)
}

func EncodeIotTopicRule_ErrorAction_Dynamodb_Operation(p Dynamodb, vals map[string]cty.Value) {
	vals["operation"] = cty.StringVal(p.Operation)
}

func EncodeIotTopicRule_ErrorAction_Dynamodb_PayloadField(p Dynamodb, vals map[string]cty.Value) {
	vals["payload_field"] = cty.StringVal(p.PayloadField)
}

func EncodeIotTopicRule_ErrorAction_Dynamodb_RangeKeyField(p Dynamodb, vals map[string]cty.Value) {
	vals["range_key_field"] = cty.StringVal(p.RangeKeyField)
}

func EncodeIotTopicRule_ErrorAction_Dynamodb_RangeKeyValue(p Dynamodb, vals map[string]cty.Value) {
	vals["range_key_value"] = cty.StringVal(p.RangeKeyValue)
}

func EncodeIotTopicRule_ErrorAction_Dynamodb_HashKeyField(p Dynamodb, vals map[string]cty.Value) {
	vals["hash_key_field"] = cty.StringVal(p.HashKeyField)
}

func EncodeIotTopicRule_ErrorAction_Dynamodb_RangeKeyType(p Dynamodb, vals map[string]cty.Value) {
	vals["range_key_type"] = cty.StringVal(p.RangeKeyType)
}

func EncodeIotTopicRule_ErrorAction_Dynamodb_TableName(p Dynamodb, vals map[string]cty.Value) {
	vals["table_name"] = cty.StringVal(p.TableName)
}

func EncodeIotTopicRule_ErrorAction_Republish(p Republish, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_Republish_RoleArn(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Republish_Topic(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Republish_Qos(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["republish"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_Republish_RoleArn(p Republish, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_Republish_Topic(p Republish, vals map[string]cty.Value) {
	vals["topic"] = cty.StringVal(p.Topic)
}

func EncodeIotTopicRule_ErrorAction_Republish_Qos(p Republish, vals map[string]cty.Value) {
	vals["qos"] = cty.NumberIntVal(p.Qos)
}

func EncodeIotTopicRule_ErrorAction_S3(p S3, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_S3_BucketName(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_S3_Key(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_S3_RoleArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["s3"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_S3_BucketName(p S3, vals map[string]cty.Value) {
	vals["bucket_name"] = cty.StringVal(p.BucketName)
}

func EncodeIotTopicRule_ErrorAction_S3_Key(p S3, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeIotTopicRule_ErrorAction_S3_RoleArn(p S3, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_Dynamodbv2(p Dynamodbv2, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_Dynamodbv2_RoleArn(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Dynamodbv2_PutItem(p.PutItem, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["dynamodbv2"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_Dynamodbv2_RoleArn(p Dynamodbv2, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_Dynamodbv2_PutItem(p PutItem, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_Dynamodbv2_PutItem_TableName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["put_item"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_Dynamodbv2_PutItem_TableName(p PutItem, vals map[string]cty.Value) {
	vals["table_name"] = cty.StringVal(p.TableName)
}

func EncodeIotTopicRule_ErrorAction_Elasticsearch(p Elasticsearch, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_Elasticsearch_Endpoint(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Elasticsearch_Id(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Elasticsearch_Index(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Elasticsearch_RoleArn(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Elasticsearch_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["elasticsearch"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_Elasticsearch_Endpoint(p Elasticsearch, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeIotTopicRule_ErrorAction_Elasticsearch_Id(p Elasticsearch, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIotTopicRule_ErrorAction_Elasticsearch_Index(p Elasticsearch, vals map[string]cty.Value) {
	vals["index"] = cty.StringVal(p.Index)
}

func EncodeIotTopicRule_ErrorAction_Elasticsearch_RoleArn(p Elasticsearch, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_Elasticsearch_Type(p Elasticsearch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeIotTopicRule_ErrorAction_Lambda(p Lambda, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_Lambda_FunctionArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["lambda"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_Lambda_FunctionArn(p Lambda, vals map[string]cty.Value) {
	vals["function_arn"] = cty.StringVal(p.FunctionArn)
}

func EncodeIotTopicRule_ErrorAction_IotAnalytics(p IotAnalytics, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_IotAnalytics_ChannelName(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_IotAnalytics_RoleArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["iot_analytics"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_IotAnalytics_ChannelName(p IotAnalytics, vals map[string]cty.Value) {
	vals["channel_name"] = cty.StringVal(p.ChannelName)
}

func EncodeIotTopicRule_ErrorAction_IotAnalytics_RoleArn(p IotAnalytics, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_Sns(p Sns, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_Sns_TargetArn(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Sns_MessageFormat(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Sns_RoleArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["sns"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_Sns_TargetArn(p Sns, vals map[string]cty.Value) {
	vals["target_arn"] = cty.StringVal(p.TargetArn)
}

func EncodeIotTopicRule_ErrorAction_Sns_MessageFormat(p Sns, vals map[string]cty.Value) {
	vals["message_format"] = cty.StringVal(p.MessageFormat)
}

func EncodeIotTopicRule_ErrorAction_Sns_RoleArn(p Sns, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_Sqs(p Sqs, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_Sqs_QueueUrl(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Sqs_RoleArn(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Sqs_UseBase64(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["sqs"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_Sqs_QueueUrl(p Sqs, vals map[string]cty.Value) {
	vals["queue_url"] = cty.StringVal(p.QueueUrl)
}

func EncodeIotTopicRule_ErrorAction_Sqs_RoleArn(p Sqs, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_Sqs_UseBase64(p Sqs, vals map[string]cty.Value) {
	vals["use_base64"] = cty.BoolVal(p.UseBase64)
}

func EncodeIotTopicRule_ErrorAction_Kinesis(p Kinesis, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_Kinesis_PartitionKey(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Kinesis_RoleArn(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_Kinesis_StreamName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["kinesis"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_Kinesis_PartitionKey(p Kinesis, vals map[string]cty.Value) {
	vals["partition_key"] = cty.StringVal(p.PartitionKey)
}

func EncodeIotTopicRule_ErrorAction_Kinesis_RoleArn(p Kinesis, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_Kinesis_StreamName(p Kinesis, vals map[string]cty.Value) {
	vals["stream_name"] = cty.StringVal(p.StreamName)
}

func EncodeIotTopicRule_ErrorAction_StepFunctions(p StepFunctions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_ErrorAction_StepFunctions_ExecutionNamePrefix(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_StepFunctions_RoleArn(p, ctyVal)
	EncodeIotTopicRule_ErrorAction_StepFunctions_StateMachineName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["step_functions"] = cty.ListVal(valsForCollection)
}

func EncodeIotTopicRule_ErrorAction_StepFunctions_ExecutionNamePrefix(p StepFunctions, vals map[string]cty.Value) {
	vals["execution_name_prefix"] = cty.StringVal(p.ExecutionNamePrefix)
}

func EncodeIotTopicRule_ErrorAction_StepFunctions_RoleArn(p StepFunctions, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_ErrorAction_StepFunctions_StateMachineName(p StepFunctions, vals map[string]cty.Value) {
	vals["state_machine_name"] = cty.StringVal(p.StateMachineName)
}

func EncodeIotTopicRule_S3(p S3, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_S3_BucketName(p, ctyVal)
	EncodeIotTopicRule_S3_Key(p, ctyVal)
	EncodeIotTopicRule_S3_RoleArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["s3"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_S3_BucketName(p S3, vals map[string]cty.Value) {
	vals["bucket_name"] = cty.StringVal(p.BucketName)
}

func EncodeIotTopicRule_S3_Key(p S3, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeIotTopicRule_S3_RoleArn(p S3, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_Sqs(p Sqs, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeIotTopicRule_Sqs_QueueUrl(p, ctyVal)
	EncodeIotTopicRule_Sqs_RoleArn(p, ctyVal)
	EncodeIotTopicRule_Sqs_UseBase64(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["sqs"] = cty.SetVal(valsForCollection)
}

func EncodeIotTopicRule_Sqs_QueueUrl(p Sqs, vals map[string]cty.Value) {
	vals["queue_url"] = cty.StringVal(p.QueueUrl)
}

func EncodeIotTopicRule_Sqs_RoleArn(p Sqs, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeIotTopicRule_Sqs_UseBase64(p Sqs, vals map[string]cty.Value) {
	vals["use_base64"] = cty.BoolVal(p.UseBase64)
}

func EncodeIotTopicRule_Arn(p IotTopicRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}