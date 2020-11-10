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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true

// IotTopicRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IotTopicRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IotTopicRuleSpec   `json:"spec"`
	Status IotTopicRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IotTopicRule contains a list of IotTopicRuleList
type IotTopicRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IotTopicRule `json:"items"`
}

// A IotTopicRuleSpec defines the desired state of a IotTopicRule
type IotTopicRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IotTopicRuleParameters `json:",inline"`
}

// A IotTopicRuleParameters defines the desired state of a IotTopicRule
type IotTopicRuleParameters struct {
	Name             string             `json:"name"`
	Sql              string             `json:"sql"`
	SqlVersion       string             `json:"sql_version"`
	Tags             map[string]string  `json:"tags"`
	Description      string             `json:"description"`
	Enabled          bool               `json:"enabled"`
	Sqs              []Sqs              `json:"sqs"`
	StepFunctions    []StepFunctions    `json:"step_functions"`
	CloudwatchAlarm  []CloudwatchAlarm  `json:"cloudwatch_alarm"`
	CloudwatchMetric []CloudwatchMetric `json:"cloudwatch_metric"`
	Dynamodbv2       []Dynamodbv2       `json:"dynamodbv2"`
	Elasticsearch    []Elasticsearch    `json:"elasticsearch"`
	IotAnalytics     []IotAnalytics     `json:"iot_analytics"`
	Republish        []Republish        `json:"republish"`
	Sns              []Sns              `json:"sns"`
	Dynamodb         []Dynamodb         `json:"dynamodb"`
	Firehose         []Firehose         `json:"firehose"`
	IotEvents        []IotEvents        `json:"iot_events"`
	Kinesis          []Kinesis          `json:"kinesis"`
	Lambda           []Lambda           `json:"lambda"`
	S3               []S3               `json:"s3"`
	ErrorAction      ErrorAction        `json:"error_action"`
}

type Sqs struct {
	UseBase64 bool   `json:"use_base64"`
	QueueUrl  string `json:"queue_url"`
	RoleArn   string `json:"role_arn"`
}

type StepFunctions struct {
	ExecutionNamePrefix string `json:"execution_name_prefix"`
	RoleArn             string `json:"role_arn"`
	StateMachineName    string `json:"state_machine_name"`
}

type CloudwatchAlarm struct {
	StateValue  string `json:"state_value"`
	AlarmName   string `json:"alarm_name"`
	RoleArn     string `json:"role_arn"`
	StateReason string `json:"state_reason"`
}

type CloudwatchMetric struct {
	MetricName      string `json:"metric_name"`
	MetricNamespace string `json:"metric_namespace"`
	MetricTimestamp string `json:"metric_timestamp"`
	MetricUnit      string `json:"metric_unit"`
	MetricValue     string `json:"metric_value"`
	RoleArn         string `json:"role_arn"`
}

type Dynamodbv2 struct {
	RoleArn string  `json:"role_arn"`
	PutItem PutItem `json:"put_item"`
}

type PutItem struct {
	TableName string `json:"table_name"`
}

type Elasticsearch struct {
	Endpoint string `json:"endpoint"`
	Id       string `json:"id"`
	Index    string `json:"index"`
	RoleArn  string `json:"role_arn"`
	Type     string `json:"type"`
}

type IotAnalytics struct {
	ChannelName string `json:"channel_name"`
	RoleArn     string `json:"role_arn"`
}

type Republish struct {
	Topic   string `json:"topic"`
	Qos     int    `json:"qos"`
	RoleArn string `json:"role_arn"`
}

type Sns struct {
	MessageFormat string `json:"message_format"`
	RoleArn       string `json:"role_arn"`
	TargetArn     string `json:"target_arn"`
}

type Dynamodb struct {
	HashKeyValue  string `json:"hash_key_value"`
	Operation     string `json:"operation"`
	TableName     string `json:"table_name"`
	HashKeyType   string `json:"hash_key_type"`
	PayloadField  string `json:"payload_field"`
	RangeKeyField string `json:"range_key_field"`
	RangeKeyType  string `json:"range_key_type"`
	RangeKeyValue string `json:"range_key_value"`
	RoleArn       string `json:"role_arn"`
	HashKeyField  string `json:"hash_key_field"`
}

type Firehose struct {
	DeliveryStreamName string `json:"delivery_stream_name"`
	RoleArn            string `json:"role_arn"`
	Separator          string `json:"separator"`
}

type IotEvents struct {
	MessageId string `json:"message_id"`
	RoleArn   string `json:"role_arn"`
	InputName string `json:"input_name"`
}

type Kinesis struct {
	PartitionKey string `json:"partition_key"`
	RoleArn      string `json:"role_arn"`
	StreamName   string `json:"stream_name"`
}

type Lambda struct {
	FunctionArn string `json:"function_arn"`
}

type S3 struct {
	BucketName string `json:"bucket_name"`
	Key        string `json:"key"`
	RoleArn    string `json:"role_arn"`
}

type ErrorAction struct {
	Lambda           Lambda           `json:"lambda"`
	Sns              Sns              `json:"sns"`
	Dynamodb         Dynamodb         `json:"dynamodb"`
	Elasticsearch    Elasticsearch    `json:"elasticsearch"`
	Kinesis          Kinesis          `json:"kinesis"`
	Republish        Republish        `json:"republish"`
	StepFunctions    StepFunctions    `json:"step_functions"`
	S3               S3               `json:"s3"`
	CloudwatchAlarm  CloudwatchAlarm  `json:"cloudwatch_alarm"`
	CloudwatchMetric CloudwatchMetric `json:"cloudwatch_metric"`
	Dynamodbv2       Dynamodbv2       `json:"dynamodbv2"`
	Firehose         Firehose         `json:"firehose"`
	IotAnalytics     IotAnalytics     `json:"iot_analytics"`
	IotEvents        IotEvents        `json:"iot_events"`
	Sqs              Sqs              `json:"sqs"`
}

type Lambda struct {
	FunctionArn string `json:"function_arn"`
}

type Sns struct {
	TargetArn     string `json:"target_arn"`
	MessageFormat string `json:"message_format"`
	RoleArn       string `json:"role_arn"`
}

type Dynamodb struct {
	TableName     string `json:"table_name"`
	HashKeyValue  string `json:"hash_key_value"`
	Operation     string `json:"operation"`
	PayloadField  string `json:"payload_field"`
	RangeKeyField string `json:"range_key_field"`
	RangeKeyType  string `json:"range_key_type"`
	RangeKeyValue string `json:"range_key_value"`
	HashKeyField  string `json:"hash_key_field"`
	HashKeyType   string `json:"hash_key_type"`
	RoleArn       string `json:"role_arn"`
}

type Elasticsearch struct {
	Endpoint string `json:"endpoint"`
	Id       string `json:"id"`
	Index    string `json:"index"`
	RoleArn  string `json:"role_arn"`
	Type     string `json:"type"`
}

type Kinesis struct {
	PartitionKey string `json:"partition_key"`
	RoleArn      string `json:"role_arn"`
	StreamName   string `json:"stream_name"`
}

type Republish struct {
	Qos     int    `json:"qos"`
	RoleArn string `json:"role_arn"`
	Topic   string `json:"topic"`
}

type StepFunctions struct {
	StateMachineName    string `json:"state_machine_name"`
	ExecutionNamePrefix string `json:"execution_name_prefix"`
	RoleArn             string `json:"role_arn"`
}

type S3 struct {
	BucketName string `json:"bucket_name"`
	Key        string `json:"key"`
	RoleArn    string `json:"role_arn"`
}

type CloudwatchAlarm struct {
	AlarmName   string `json:"alarm_name"`
	RoleArn     string `json:"role_arn"`
	StateReason string `json:"state_reason"`
	StateValue  string `json:"state_value"`
}

type CloudwatchMetric struct {
	MetricNamespace string `json:"metric_namespace"`
	MetricTimestamp string `json:"metric_timestamp"`
	MetricUnit      string `json:"metric_unit"`
	MetricValue     string `json:"metric_value"`
	RoleArn         string `json:"role_arn"`
	MetricName      string `json:"metric_name"`
}

type Dynamodbv2 struct {
	RoleArn string  `json:"role_arn"`
	PutItem PutItem `json:"put_item"`
}

type PutItem struct {
	TableName string `json:"table_name"`
}

type Firehose struct {
	DeliveryStreamName string `json:"delivery_stream_name"`
	RoleArn            string `json:"role_arn"`
	Separator          string `json:"separator"`
}

type IotAnalytics struct {
	ChannelName string `json:"channel_name"`
	RoleArn     string `json:"role_arn"`
}

type IotEvents struct {
	InputName string `json:"input_name"`
	MessageId string `json:"message_id"`
	RoleArn   string `json:"role_arn"`
}

type Sqs struct {
	UseBase64 bool   `json:"use_base64"`
	QueueUrl  string `json:"queue_url"`
	RoleArn   string `json:"role_arn"`
}

// A IotTopicRuleStatus defines the observed state of a IotTopicRule
type IotTopicRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IotTopicRuleObservation `json:",inline"`
}

// A IotTopicRuleObservation records the observed state of a IotTopicRule
type IotTopicRuleObservation struct {
	Id  string `json:"id"`
	Arn string `json:"arn"`
}