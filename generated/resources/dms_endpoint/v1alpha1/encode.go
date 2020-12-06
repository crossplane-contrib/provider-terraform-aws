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

func EncodeDmsEndpoint(r DmsEndpoint) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDmsEndpoint_CertificateArn(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_EndpointType(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_DatabaseName(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_EndpointId(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_ExtraConnectionAttributes(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_Id(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_Password(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_EngineName(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_KmsKeyArn(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_Port(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_ServerName(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_SslMode(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_Username(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_ServiceAccessRole(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDmsEndpoint_S3Settings(r.Spec.ForProvider.S3Settings, ctyVal)
	EncodeDmsEndpoint_ElasticsearchSettings(r.Spec.ForProvider.ElasticsearchSettings, ctyVal)
	EncodeDmsEndpoint_KafkaSettings(r.Spec.ForProvider.KafkaSettings, ctyVal)
	EncodeDmsEndpoint_KinesisSettings(r.Spec.ForProvider.KinesisSettings, ctyVal)
	EncodeDmsEndpoint_MongodbSettings(r.Spec.ForProvider.MongodbSettings, ctyVal)
	EncodeDmsEndpoint_EndpointArn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDmsEndpoint_CertificateArn(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["certificate_arn"] = cty.StringVal(p.CertificateArn)
}

func EncodeDmsEndpoint_EndpointType(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["endpoint_type"] = cty.StringVal(p.EndpointType)
}

func EncodeDmsEndpoint_DatabaseName(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["database_name"] = cty.StringVal(p.DatabaseName)
}

func EncodeDmsEndpoint_EndpointId(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["endpoint_id"] = cty.StringVal(p.EndpointId)
}

func EncodeDmsEndpoint_ExtraConnectionAttributes(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["extra_connection_attributes"] = cty.StringVal(p.ExtraConnectionAttributes)
}

func EncodeDmsEndpoint_Id(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDmsEndpoint_Password(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeDmsEndpoint_EngineName(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["engine_name"] = cty.StringVal(p.EngineName)
}

func EncodeDmsEndpoint_KmsKeyArn(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeDmsEndpoint_Port(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeDmsEndpoint_ServerName(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["server_name"] = cty.StringVal(p.ServerName)
}

func EncodeDmsEndpoint_SslMode(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["ssl_mode"] = cty.StringVal(p.SslMode)
}

func EncodeDmsEndpoint_Username(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeDmsEndpoint_ServiceAccessRole(p DmsEndpointParameters, vals map[string]cty.Value) {
	vals["service_access_role"] = cty.StringVal(p.ServiceAccessRole)
}

func EncodeDmsEndpoint_Tags(p DmsEndpointParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDmsEndpoint_S3Settings(p S3Settings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDmsEndpoint_S3Settings_BucketFolder(p, ctyVal)
	EncodeDmsEndpoint_S3Settings_BucketName(p, ctyVal)
	EncodeDmsEndpoint_S3Settings_CompressionType(p, ctyVal)
	EncodeDmsEndpoint_S3Settings_CsvDelimiter(p, ctyVal)
	EncodeDmsEndpoint_S3Settings_CsvRowDelimiter(p, ctyVal)
	EncodeDmsEndpoint_S3Settings_ExternalTableDefinition(p, ctyVal)
	EncodeDmsEndpoint_S3Settings_ServiceAccessRoleArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["s3_settings"] = cty.ListVal(valsForCollection)
}

func EncodeDmsEndpoint_S3Settings_BucketFolder(p S3Settings, vals map[string]cty.Value) {
	vals["bucket_folder"] = cty.StringVal(p.BucketFolder)
}

func EncodeDmsEndpoint_S3Settings_BucketName(p S3Settings, vals map[string]cty.Value) {
	vals["bucket_name"] = cty.StringVal(p.BucketName)
}

func EncodeDmsEndpoint_S3Settings_CompressionType(p S3Settings, vals map[string]cty.Value) {
	vals["compression_type"] = cty.StringVal(p.CompressionType)
}

func EncodeDmsEndpoint_S3Settings_CsvDelimiter(p S3Settings, vals map[string]cty.Value) {
	vals["csv_delimiter"] = cty.StringVal(p.CsvDelimiter)
}

func EncodeDmsEndpoint_S3Settings_CsvRowDelimiter(p S3Settings, vals map[string]cty.Value) {
	vals["csv_row_delimiter"] = cty.StringVal(p.CsvRowDelimiter)
}

func EncodeDmsEndpoint_S3Settings_ExternalTableDefinition(p S3Settings, vals map[string]cty.Value) {
	vals["external_table_definition"] = cty.StringVal(p.ExternalTableDefinition)
}

func EncodeDmsEndpoint_S3Settings_ServiceAccessRoleArn(p S3Settings, vals map[string]cty.Value) {
	vals["service_access_role_arn"] = cty.StringVal(p.ServiceAccessRoleArn)
}

func EncodeDmsEndpoint_ElasticsearchSettings(p ElasticsearchSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDmsEndpoint_ElasticsearchSettings_EndpointUri(p, ctyVal)
	EncodeDmsEndpoint_ElasticsearchSettings_ErrorRetryDuration(p, ctyVal)
	EncodeDmsEndpoint_ElasticsearchSettings_FullLoadErrorPercentage(p, ctyVal)
	EncodeDmsEndpoint_ElasticsearchSettings_ServiceAccessRoleArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["elasticsearch_settings"] = cty.ListVal(valsForCollection)
}

func EncodeDmsEndpoint_ElasticsearchSettings_EndpointUri(p ElasticsearchSettings, vals map[string]cty.Value) {
	vals["endpoint_uri"] = cty.StringVal(p.EndpointUri)
}

func EncodeDmsEndpoint_ElasticsearchSettings_ErrorRetryDuration(p ElasticsearchSettings, vals map[string]cty.Value) {
	vals["error_retry_duration"] = cty.NumberIntVal(p.ErrorRetryDuration)
}

func EncodeDmsEndpoint_ElasticsearchSettings_FullLoadErrorPercentage(p ElasticsearchSettings, vals map[string]cty.Value) {
	vals["full_load_error_percentage"] = cty.NumberIntVal(p.FullLoadErrorPercentage)
}

func EncodeDmsEndpoint_ElasticsearchSettings_ServiceAccessRoleArn(p ElasticsearchSettings, vals map[string]cty.Value) {
	vals["service_access_role_arn"] = cty.StringVal(p.ServiceAccessRoleArn)
}

func EncodeDmsEndpoint_KafkaSettings(p KafkaSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDmsEndpoint_KafkaSettings_Broker(p, ctyVal)
	EncodeDmsEndpoint_KafkaSettings_Topic(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["kafka_settings"] = cty.ListVal(valsForCollection)
}

func EncodeDmsEndpoint_KafkaSettings_Broker(p KafkaSettings, vals map[string]cty.Value) {
	vals["broker"] = cty.StringVal(p.Broker)
}

func EncodeDmsEndpoint_KafkaSettings_Topic(p KafkaSettings, vals map[string]cty.Value) {
	vals["topic"] = cty.StringVal(p.Topic)
}

func EncodeDmsEndpoint_KinesisSettings(p KinesisSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDmsEndpoint_KinesisSettings_StreamArn(p, ctyVal)
	EncodeDmsEndpoint_KinesisSettings_MessageFormat(p, ctyVal)
	EncodeDmsEndpoint_KinesisSettings_ServiceAccessRoleArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["kinesis_settings"] = cty.ListVal(valsForCollection)
}

func EncodeDmsEndpoint_KinesisSettings_StreamArn(p KinesisSettings, vals map[string]cty.Value) {
	vals["stream_arn"] = cty.StringVal(p.StreamArn)
}

func EncodeDmsEndpoint_KinesisSettings_MessageFormat(p KinesisSettings, vals map[string]cty.Value) {
	vals["message_format"] = cty.StringVal(p.MessageFormat)
}

func EncodeDmsEndpoint_KinesisSettings_ServiceAccessRoleArn(p KinesisSettings, vals map[string]cty.Value) {
	vals["service_access_role_arn"] = cty.StringVal(p.ServiceAccessRoleArn)
}

func EncodeDmsEndpoint_MongodbSettings(p MongodbSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDmsEndpoint_MongodbSettings_NestingLevel(p, ctyVal)
	EncodeDmsEndpoint_MongodbSettings_AuthMechanism(p, ctyVal)
	EncodeDmsEndpoint_MongodbSettings_AuthSource(p, ctyVal)
	EncodeDmsEndpoint_MongodbSettings_AuthType(p, ctyVal)
	EncodeDmsEndpoint_MongodbSettings_DocsToInvestigate(p, ctyVal)
	EncodeDmsEndpoint_MongodbSettings_ExtractDocId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["mongodb_settings"] = cty.ListVal(valsForCollection)
}

func EncodeDmsEndpoint_MongodbSettings_NestingLevel(p MongodbSettings, vals map[string]cty.Value) {
	vals["nesting_level"] = cty.StringVal(p.NestingLevel)
}

func EncodeDmsEndpoint_MongodbSettings_AuthMechanism(p MongodbSettings, vals map[string]cty.Value) {
	vals["auth_mechanism"] = cty.StringVal(p.AuthMechanism)
}

func EncodeDmsEndpoint_MongodbSettings_AuthSource(p MongodbSettings, vals map[string]cty.Value) {
	vals["auth_source"] = cty.StringVal(p.AuthSource)
}

func EncodeDmsEndpoint_MongodbSettings_AuthType(p MongodbSettings, vals map[string]cty.Value) {
	vals["auth_type"] = cty.StringVal(p.AuthType)
}

func EncodeDmsEndpoint_MongodbSettings_DocsToInvestigate(p MongodbSettings, vals map[string]cty.Value) {
	vals["docs_to_investigate"] = cty.StringVal(p.DocsToInvestigate)
}

func EncodeDmsEndpoint_MongodbSettings_ExtractDocId(p MongodbSettings, vals map[string]cty.Value) {
	vals["extract_doc_id"] = cty.StringVal(p.ExtractDocId)
}

func EncodeDmsEndpoint_EndpointArn(p DmsEndpointObservation, vals map[string]cty.Value) {
	vals["endpoint_arn"] = cty.StringVal(p.EndpointArn)
}