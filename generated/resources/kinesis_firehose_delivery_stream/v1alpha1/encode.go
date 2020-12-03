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

package v1alpha1func EncodeKinesisFirehoseDeliveryStream(r KinesisFirehoseDeliveryStream) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeKinesisFirehoseDeliveryStream_Destination(r.Spec.ForProvider, ctyVal)
	EncodeKinesisFirehoseDeliveryStream_DestinationId(r.Spec.ForProvider, ctyVal)
	EncodeKinesisFirehoseDeliveryStream_Id(r.Spec.ForProvider, ctyVal)
	EncodeKinesisFirehoseDeliveryStream_Name(r.Spec.ForProvider, ctyVal)
	EncodeKinesisFirehoseDeliveryStream_Tags(r.Spec.ForProvider, ctyVal)
	EncodeKinesisFirehoseDeliveryStream_VersionId(r.Spec.ForProvider, ctyVal)
	EncodeKinesisFirehoseDeliveryStream_Arn(r.Spec.ForProvider, ctyVal)
	EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration(r.Spec.ForProvider.ElasticsearchConfiguration, ctyVal)
	EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration(r.Spec.ForProvider.ExtendedS3Configuration, ctyVal)
	EncodeKinesisFirehoseDeliveryStream_KinesisSourceConfiguration(r.Spec.ForProvider.KinesisSourceConfiguration, ctyVal)
	EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration(r.Spec.ForProvider.RedshiftConfiguration, ctyVal)
	EncodeKinesisFirehoseDeliveryStream_S3Configuration(r.Spec.ForProvider.S3Configuration, ctyVal)
	EncodeKinesisFirehoseDeliveryStream_ServerSideEncryption(r.Spec.ForProvider.ServerSideEncryption, ctyVal)
	EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration(r.Spec.ForProvider.SplunkConfiguration, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeKinesisFirehoseDeliveryStream_Destination(p *KinesisFirehoseDeliveryStreamParameters, vals map[string]cty.Value) {
	vals["destination"] = cty.StringVal(p.Destination)
}

func EncodeKinesisFirehoseDeliveryStream_DestinationId(p *KinesisFirehoseDeliveryStreamParameters, vals map[string]cty.Value) {
	vals["destination_id"] = cty.StringVal(p.DestinationId)
}

func EncodeKinesisFirehoseDeliveryStream_Id(p *KinesisFirehoseDeliveryStreamParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKinesisFirehoseDeliveryStream_Name(p *KinesisFirehoseDeliveryStreamParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeKinesisFirehoseDeliveryStream_Tags(p *KinesisFirehoseDeliveryStreamParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeKinesisFirehoseDeliveryStream_VersionId(p *KinesisFirehoseDeliveryStreamParameters, vals map[string]cty.Value) {
	vals["version_id"] = cty.StringVal(p.VersionId)
}

func EncodeKinesisFirehoseDeliveryStream_Arn(p *KinesisFirehoseDeliveryStreamParameters, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration(p *ElasticsearchConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ElasticsearchConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_RetryDuration(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_RoleArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_BufferingSize(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ClusterEndpoint(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_DomainArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_IndexName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_BufferingInterval(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_IndexRotationPeriod(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_S3BackupMode(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_TypeName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_CloudwatchLoggingOptions(v.CloudwatchLoggingOptions, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration(v.ProcessingConfiguration, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_VpcConfig(v.VpcConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["elasticsearch_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_RetryDuration(p *ElasticsearchConfiguration, vals map[string]cty.Value) {
	vals["retry_duration"] = cty.IntVal(p.RetryDuration)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_RoleArn(p *ElasticsearchConfiguration, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_BufferingSize(p *ElasticsearchConfiguration, vals map[string]cty.Value) {
	vals["buffering_size"] = cty.IntVal(p.BufferingSize)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ClusterEndpoint(p *ElasticsearchConfiguration, vals map[string]cty.Value) {
	vals["cluster_endpoint"] = cty.StringVal(p.ClusterEndpoint)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_DomainArn(p *ElasticsearchConfiguration, vals map[string]cty.Value) {
	vals["domain_arn"] = cty.StringVal(p.DomainArn)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_IndexName(p *ElasticsearchConfiguration, vals map[string]cty.Value) {
	vals["index_name"] = cty.StringVal(p.IndexName)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_BufferingInterval(p *ElasticsearchConfiguration, vals map[string]cty.Value) {
	vals["buffering_interval"] = cty.IntVal(p.BufferingInterval)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_IndexRotationPeriod(p *ElasticsearchConfiguration, vals map[string]cty.Value) {
	vals["index_rotation_period"] = cty.StringVal(p.IndexRotationPeriod)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_S3BackupMode(p *ElasticsearchConfiguration, vals map[string]cty.Value) {
	vals["s3_backup_mode"] = cty.StringVal(p.S3BackupMode)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_TypeName(p *ElasticsearchConfiguration, vals map[string]cty.Value) {
	vals["type_name"] = cty.StringVal(p.TypeName)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_CloudwatchLoggingOptions(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CloudwatchLoggingOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_CloudwatchLoggingOptions_Enabled(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_CloudwatchLoggingOptions_LogGroupName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_CloudwatchLoggingOptions_LogStreamName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cloudwatch_logging_options"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_CloudwatchLoggingOptions_Enabled(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_CloudwatchLoggingOptions_LogGroupName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_group_name"] = cty.StringVal(p.LogGroupName)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_CloudwatchLoggingOptions_LogStreamName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_stream_name"] = cty.StringVal(p.LogStreamName)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration(p *ProcessingConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ProcessingConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration_Enabled(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration_Processors(v.Processors, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["processing_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration_Enabled(p *ProcessingConfiguration, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration_Processors(p *Processors, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Processors {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration_Processors_Type(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration_Processors_Parameters(v.Parameters, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["processors"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration_Processors_Type(p *Processors, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration_Processors_Parameters(p *Parameters, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Parameters {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration_Processors_Parameters_ParameterName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration_Processors_Parameters_ParameterValue(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["parameters"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration_Processors_Parameters_ParameterName(p *Parameters, vals map[string]cty.Value) {
	vals["parameter_name"] = cty.StringVal(p.ParameterName)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_ProcessingConfiguration_Processors_Parameters_ParameterValue(p *Parameters, vals map[string]cty.Value) {
	vals["parameter_value"] = cty.StringVal(p.ParameterValue)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_VpcConfig(p *VpcConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.VpcConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_VpcConfig_RoleArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_VpcConfig_SecurityGroupIds(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_VpcConfig_SubnetIds(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_VpcConfig_VpcId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["vpc_config"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_VpcConfig_RoleArn(p *VpcConfig, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_VpcConfig_SecurityGroupIds(p *VpcConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_VpcConfig_SubnetIds(p *VpcConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeKinesisFirehoseDeliveryStream_ElasticsearchConfiguration_VpcConfig_VpcId(p *VpcConfig, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration(p *ExtendedS3Configuration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ExtendedS3Configuration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_BucketArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ErrorOutputPrefix(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_KmsKeyArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_RoleArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_BufferInterval(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_BufferSize(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_CompressionFormat(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_Prefix(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupMode(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration(v.DataFormatConversionConfiguration, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration(v.ProcessingConfiguration, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration(v.S3BackupConfiguration, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_CloudwatchLoggingOptions(v.CloudwatchLoggingOptions, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["extended_s3_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_BucketArn(p *ExtendedS3Configuration, vals map[string]cty.Value) {
	vals["bucket_arn"] = cty.StringVal(p.BucketArn)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ErrorOutputPrefix(p *ExtendedS3Configuration, vals map[string]cty.Value) {
	vals["error_output_prefix"] = cty.StringVal(p.ErrorOutputPrefix)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_KmsKeyArn(p *ExtendedS3Configuration, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_RoleArn(p *ExtendedS3Configuration, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_BufferInterval(p *ExtendedS3Configuration, vals map[string]cty.Value) {
	vals["buffer_interval"] = cty.IntVal(p.BufferInterval)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_BufferSize(p *ExtendedS3Configuration, vals map[string]cty.Value) {
	vals["buffer_size"] = cty.IntVal(p.BufferSize)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_CompressionFormat(p *ExtendedS3Configuration, vals map[string]cty.Value) {
	vals["compression_format"] = cty.StringVal(p.CompressionFormat)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_Prefix(p *ExtendedS3Configuration, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupMode(p *ExtendedS3Configuration, vals map[string]cty.Value) {
	vals["s3_backup_mode"] = cty.StringVal(p.S3BackupMode)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration(p *DataFormatConversionConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DataFormatConversionConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_Enabled(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration(v.SchemaConfiguration, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration(v.InputFormatConfiguration, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration(v.OutputFormatConfiguration, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["data_format_conversion_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_Enabled(p *DataFormatConversionConfiguration, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration(p *SchemaConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SchemaConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration_DatabaseName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration_Region(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration_RoleArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration_TableName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration_VersionId(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration_CatalogId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["schema_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration_DatabaseName(p *SchemaConfiguration, vals map[string]cty.Value) {
	vals["database_name"] = cty.StringVal(p.DatabaseName)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration_Region(p *SchemaConfiguration, vals map[string]cty.Value) {
	vals["region"] = cty.StringVal(p.Region)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration_RoleArn(p *SchemaConfiguration, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration_TableName(p *SchemaConfiguration, vals map[string]cty.Value) {
	vals["table_name"] = cty.StringVal(p.TableName)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration_VersionId(p *SchemaConfiguration, vals map[string]cty.Value) {
	vals["version_id"] = cty.StringVal(p.VersionId)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_SchemaConfiguration_CatalogId(p *SchemaConfiguration, vals map[string]cty.Value) {
	vals["catalog_id"] = cty.StringVal(p.CatalogId)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration(p *InputFormatConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.InputFormatConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer(v.Deserializer, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["input_format_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer(p *Deserializer, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Deserializer {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer_HiveJsonSerDe(v.HiveJsonSerDe, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer_OpenXJsonSerDe(v.OpenXJsonSerDe, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["deserializer"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer_HiveJsonSerDe(p *HiveJsonSerDe, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.HiveJsonSerDe {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer_HiveJsonSerDe_TimestampFormats(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["hive_json_ser_de"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer_HiveJsonSerDe_TimestampFormats(p *HiveJsonSerDe, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.TimestampFormats {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["timestamp_formats"] = cty.ListVal(colVals)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer_OpenXJsonSerDe(p *OpenXJsonSerDe, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OpenXJsonSerDe {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer_OpenXJsonSerDe_CaseInsensitive(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer_OpenXJsonSerDe_ColumnToJsonKeyMappings(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer_OpenXJsonSerDe_ConvertDotsInJsonKeysToUnderscores(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["open_x_json_ser_de"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer_OpenXJsonSerDe_CaseInsensitive(p *OpenXJsonSerDe, vals map[string]cty.Value) {
	vals["case_insensitive"] = cty.BoolVal(p.CaseInsensitive)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer_OpenXJsonSerDe_ColumnToJsonKeyMappings(p *OpenXJsonSerDe, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.ColumnToJsonKeyMappings {
		mVals[key] = cty.StringVal(value)
	}
	vals["column_to_json_key_mappings"] = cty.MapVal(mVals)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_InputFormatConfiguration_Deserializer_OpenXJsonSerDe_ConvertDotsInJsonKeysToUnderscores(p *OpenXJsonSerDe, vals map[string]cty.Value) {
	vals["convert_dots_in_json_keys_to_underscores"] = cty.BoolVal(p.ConvertDotsInJsonKeysToUnderscores)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration(p *OutputFormatConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OutputFormatConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer(v.Serializer, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["output_format_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer(p *Serializer, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Serializer {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe(v.OrcSerDe, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe(v.ParquetSerDe, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["serializer"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe(p *OrcSerDe, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OrcSerDe {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_BloomFilterFalsePositiveProbability(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_RowIndexStride(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_EnablePadding(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_FormatVersion(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_PaddingTolerance(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_StripeSizeBytes(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_BlockSizeBytes(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_BloomFilterColumns(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_Compression(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_DictionaryKeyThreshold(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["orc_ser_de"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_BloomFilterFalsePositiveProbability(p *OrcSerDe, vals map[string]cty.Value) {
	vals["bloom_filter_false_positive_probability"] = cty.IntVal(p.BloomFilterFalsePositiveProbability)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_RowIndexStride(p *OrcSerDe, vals map[string]cty.Value) {
	vals["row_index_stride"] = cty.IntVal(p.RowIndexStride)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_EnablePadding(p *OrcSerDe, vals map[string]cty.Value) {
	vals["enable_padding"] = cty.BoolVal(p.EnablePadding)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_FormatVersion(p *OrcSerDe, vals map[string]cty.Value) {
	vals["format_version"] = cty.StringVal(p.FormatVersion)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_PaddingTolerance(p *OrcSerDe, vals map[string]cty.Value) {
	vals["padding_tolerance"] = cty.IntVal(p.PaddingTolerance)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_StripeSizeBytes(p *OrcSerDe, vals map[string]cty.Value) {
	vals["stripe_size_bytes"] = cty.IntVal(p.StripeSizeBytes)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_BlockSizeBytes(p *OrcSerDe, vals map[string]cty.Value) {
	vals["block_size_bytes"] = cty.IntVal(p.BlockSizeBytes)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_BloomFilterColumns(p *OrcSerDe, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.BloomFilterColumns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["bloom_filter_columns"] = cty.ListVal(colVals)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_Compression(p *OrcSerDe, vals map[string]cty.Value) {
	vals["compression"] = cty.StringVal(p.Compression)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_OrcSerDe_DictionaryKeyThreshold(p *OrcSerDe, vals map[string]cty.Value) {
	vals["dictionary_key_threshold"] = cty.IntVal(p.DictionaryKeyThreshold)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe(p *ParquetSerDe, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ParquetSerDe {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe_BlockSizeBytes(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe_Compression(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe_EnableDictionaryCompression(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe_MaxPaddingBytes(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe_PageSizeBytes(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe_WriterVersion(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["parquet_ser_de"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe_BlockSizeBytes(p *ParquetSerDe, vals map[string]cty.Value) {
	vals["block_size_bytes"] = cty.IntVal(p.BlockSizeBytes)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe_Compression(p *ParquetSerDe, vals map[string]cty.Value) {
	vals["compression"] = cty.StringVal(p.Compression)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe_EnableDictionaryCompression(p *ParquetSerDe, vals map[string]cty.Value) {
	vals["enable_dictionary_compression"] = cty.BoolVal(p.EnableDictionaryCompression)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe_MaxPaddingBytes(p *ParquetSerDe, vals map[string]cty.Value) {
	vals["max_padding_bytes"] = cty.IntVal(p.MaxPaddingBytes)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe_PageSizeBytes(p *ParquetSerDe, vals map[string]cty.Value) {
	vals["page_size_bytes"] = cty.IntVal(p.PageSizeBytes)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_DataFormatConversionConfiguration_OutputFormatConfiguration_Serializer_ParquetSerDe_WriterVersion(p *ParquetSerDe, vals map[string]cty.Value) {
	vals["writer_version"] = cty.StringVal(p.WriterVersion)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration(p *ProcessingConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ProcessingConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration_Enabled(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration_Processors(v.Processors, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["processing_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration_Enabled(p *ProcessingConfiguration, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration_Processors(p *Processors, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Processors {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration_Processors_Type(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration_Processors_Parameters(v.Parameters, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["processors"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration_Processors_Type(p *Processors, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration_Processors_Parameters(p *Parameters, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Parameters {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration_Processors_Parameters_ParameterName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration_Processors_Parameters_ParameterValue(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["parameters"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration_Processors_Parameters_ParameterName(p *Parameters, vals map[string]cty.Value) {
	vals["parameter_name"] = cty.StringVal(p.ParameterName)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_ProcessingConfiguration_Processors_Parameters_ParameterValue(p *Parameters, vals map[string]cty.Value) {
	vals["parameter_value"] = cty.StringVal(p.ParameterValue)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration(p *S3BackupConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.S3BackupConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_RoleArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_BucketArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_BufferInterval(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_BufferSize(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_CompressionFormat(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_KmsKeyArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_Prefix(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_CloudwatchLoggingOptions(v.CloudwatchLoggingOptions, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["s3_backup_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_RoleArn(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_BucketArn(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["bucket_arn"] = cty.StringVal(p.BucketArn)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_BufferInterval(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["buffer_interval"] = cty.IntVal(p.BufferInterval)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_BufferSize(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["buffer_size"] = cty.IntVal(p.BufferSize)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_CompressionFormat(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["compression_format"] = cty.StringVal(p.CompressionFormat)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_KmsKeyArn(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_Prefix(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_CloudwatchLoggingOptions(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CloudwatchLoggingOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_CloudwatchLoggingOptions_Enabled(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_CloudwatchLoggingOptions_LogGroupName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_CloudwatchLoggingOptions_LogStreamName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cloudwatch_logging_options"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_CloudwatchLoggingOptions_Enabled(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_CloudwatchLoggingOptions_LogGroupName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_group_name"] = cty.StringVal(p.LogGroupName)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_S3BackupConfiguration_CloudwatchLoggingOptions_LogStreamName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_stream_name"] = cty.StringVal(p.LogStreamName)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_CloudwatchLoggingOptions(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CloudwatchLoggingOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_CloudwatchLoggingOptions_Enabled(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_CloudwatchLoggingOptions_LogGroupName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_CloudwatchLoggingOptions_LogStreamName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cloudwatch_logging_options"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_CloudwatchLoggingOptions_Enabled(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_CloudwatchLoggingOptions_LogGroupName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_group_name"] = cty.StringVal(p.LogGroupName)
}

func EncodeKinesisFirehoseDeliveryStream_ExtendedS3Configuration_CloudwatchLoggingOptions_LogStreamName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_stream_name"] = cty.StringVal(p.LogStreamName)
}

func EncodeKinesisFirehoseDeliveryStream_KinesisSourceConfiguration(p *KinesisSourceConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.KinesisSourceConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_KinesisSourceConfiguration_RoleArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_KinesisSourceConfiguration_KinesisStreamArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["kinesis_source_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_KinesisSourceConfiguration_RoleArn(p *KinesisSourceConfiguration, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisFirehoseDeliveryStream_KinesisSourceConfiguration_KinesisStreamArn(p *KinesisSourceConfiguration, vals map[string]cty.Value) {
	vals["kinesis_stream_arn"] = cty.StringVal(p.KinesisStreamArn)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration(p *RedshiftConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RedshiftConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_RetryDuration(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_RoleArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_Username(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_CopyOptions(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_DataTableColumns(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_Password(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupMode(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ClusterJdbcurl(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_DataTableName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration(v.S3BackupConfiguration, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_CloudwatchLoggingOptions(v.CloudwatchLoggingOptions, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration(v.ProcessingConfiguration, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["redshift_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_RetryDuration(p *RedshiftConfiguration, vals map[string]cty.Value) {
	vals["retry_duration"] = cty.IntVal(p.RetryDuration)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_RoleArn(p *RedshiftConfiguration, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_Username(p *RedshiftConfiguration, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_CopyOptions(p *RedshiftConfiguration, vals map[string]cty.Value) {
	vals["copy_options"] = cty.StringVal(p.CopyOptions)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_DataTableColumns(p *RedshiftConfiguration, vals map[string]cty.Value) {
	vals["data_table_columns"] = cty.StringVal(p.DataTableColumns)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_Password(p *RedshiftConfiguration, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupMode(p *RedshiftConfiguration, vals map[string]cty.Value) {
	vals["s3_backup_mode"] = cty.StringVal(p.S3BackupMode)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ClusterJdbcurl(p *RedshiftConfiguration, vals map[string]cty.Value) {
	vals["cluster_jdbcurl"] = cty.StringVal(p.ClusterJdbcurl)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_DataTableName(p *RedshiftConfiguration, vals map[string]cty.Value) {
	vals["data_table_name"] = cty.StringVal(p.DataTableName)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration(p *S3BackupConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.S3BackupConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_BufferInterval(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_BufferSize(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_CompressionFormat(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_KmsKeyArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_Prefix(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_RoleArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_BucketArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_CloudwatchLoggingOptions(v.CloudwatchLoggingOptions, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["s3_backup_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_BufferInterval(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["buffer_interval"] = cty.IntVal(p.BufferInterval)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_BufferSize(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["buffer_size"] = cty.IntVal(p.BufferSize)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_CompressionFormat(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["compression_format"] = cty.StringVal(p.CompressionFormat)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_KmsKeyArn(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_Prefix(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_RoleArn(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_BucketArn(p *S3BackupConfiguration, vals map[string]cty.Value) {
	vals["bucket_arn"] = cty.StringVal(p.BucketArn)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_CloudwatchLoggingOptions(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CloudwatchLoggingOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_CloudwatchLoggingOptions_Enabled(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_CloudwatchLoggingOptions_LogGroupName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_CloudwatchLoggingOptions_LogStreamName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cloudwatch_logging_options"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_CloudwatchLoggingOptions_Enabled(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_CloudwatchLoggingOptions_LogGroupName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_group_name"] = cty.StringVal(p.LogGroupName)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_S3BackupConfiguration_CloudwatchLoggingOptions_LogStreamName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_stream_name"] = cty.StringVal(p.LogStreamName)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_CloudwatchLoggingOptions(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CloudwatchLoggingOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_CloudwatchLoggingOptions_Enabled(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_CloudwatchLoggingOptions_LogGroupName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_CloudwatchLoggingOptions_LogStreamName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cloudwatch_logging_options"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_CloudwatchLoggingOptions_Enabled(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_CloudwatchLoggingOptions_LogGroupName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_group_name"] = cty.StringVal(p.LogGroupName)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_CloudwatchLoggingOptions_LogStreamName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_stream_name"] = cty.StringVal(p.LogStreamName)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration(p *ProcessingConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ProcessingConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration_Enabled(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration_Processors(v.Processors, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["processing_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration_Enabled(p *ProcessingConfiguration, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration_Processors(p *Processors, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Processors {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration_Processors_Type(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration_Processors_Parameters(v.Parameters, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["processors"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration_Processors_Type(p *Processors, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration_Processors_Parameters(p *Parameters, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Parameters {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration_Processors_Parameters_ParameterName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration_Processors_Parameters_ParameterValue(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["parameters"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration_Processors_Parameters_ParameterName(p *Parameters, vals map[string]cty.Value) {
	vals["parameter_name"] = cty.StringVal(p.ParameterName)
}

func EncodeKinesisFirehoseDeliveryStream_RedshiftConfiguration_ProcessingConfiguration_Processors_Parameters_ParameterValue(p *Parameters, vals map[string]cty.Value) {
	vals["parameter_value"] = cty.StringVal(p.ParameterValue)
}

func EncodeKinesisFirehoseDeliveryStream_S3Configuration(p *S3Configuration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.S3Configuration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_S3Configuration_RoleArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_S3Configuration_BucketArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_S3Configuration_BufferInterval(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_S3Configuration_BufferSize(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_S3Configuration_CompressionFormat(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_S3Configuration_KmsKeyArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_S3Configuration_Prefix(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_S3Configuration_CloudwatchLoggingOptions(v.CloudwatchLoggingOptions, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["s3_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_S3Configuration_RoleArn(p *S3Configuration, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisFirehoseDeliveryStream_S3Configuration_BucketArn(p *S3Configuration, vals map[string]cty.Value) {
	vals["bucket_arn"] = cty.StringVal(p.BucketArn)
}

func EncodeKinesisFirehoseDeliveryStream_S3Configuration_BufferInterval(p *S3Configuration, vals map[string]cty.Value) {
	vals["buffer_interval"] = cty.IntVal(p.BufferInterval)
}

func EncodeKinesisFirehoseDeliveryStream_S3Configuration_BufferSize(p *S3Configuration, vals map[string]cty.Value) {
	vals["buffer_size"] = cty.IntVal(p.BufferSize)
}

func EncodeKinesisFirehoseDeliveryStream_S3Configuration_CompressionFormat(p *S3Configuration, vals map[string]cty.Value) {
	vals["compression_format"] = cty.StringVal(p.CompressionFormat)
}

func EncodeKinesisFirehoseDeliveryStream_S3Configuration_KmsKeyArn(p *S3Configuration, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeKinesisFirehoseDeliveryStream_S3Configuration_Prefix(p *S3Configuration, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeKinesisFirehoseDeliveryStream_S3Configuration_CloudwatchLoggingOptions(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CloudwatchLoggingOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_S3Configuration_CloudwatchLoggingOptions_Enabled(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_S3Configuration_CloudwatchLoggingOptions_LogGroupName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_S3Configuration_CloudwatchLoggingOptions_LogStreamName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cloudwatch_logging_options"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_S3Configuration_CloudwatchLoggingOptions_Enabled(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKinesisFirehoseDeliveryStream_S3Configuration_CloudwatchLoggingOptions_LogGroupName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_group_name"] = cty.StringVal(p.LogGroupName)
}

func EncodeKinesisFirehoseDeliveryStream_S3Configuration_CloudwatchLoggingOptions_LogStreamName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_stream_name"] = cty.StringVal(p.LogStreamName)
}

func EncodeKinesisFirehoseDeliveryStream_ServerSideEncryption(p *ServerSideEncryption, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ServerSideEncryption {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_ServerSideEncryption_Enabled(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ServerSideEncryption_KeyArn(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_ServerSideEncryption_KeyType(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["server_side_encryption"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_ServerSideEncryption_Enabled(p *ServerSideEncryption, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKinesisFirehoseDeliveryStream_ServerSideEncryption_KeyArn(p *ServerSideEncryption, vals map[string]cty.Value) {
	vals["key_arn"] = cty.StringVal(p.KeyArn)
}

func EncodeKinesisFirehoseDeliveryStream_ServerSideEncryption_KeyType(p *ServerSideEncryption, vals map[string]cty.Value) {
	vals["key_type"] = cty.StringVal(p.KeyType)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration(p *SplunkConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SplunkConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_HecEndpoint(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_HecEndpointType(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_HecToken(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_RetryDuration(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_S3BackupMode(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_HecAcknowledgmentTimeout(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_CloudwatchLoggingOptions(v.CloudwatchLoggingOptions, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration(v.ProcessingConfiguration, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["splunk_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_HecEndpoint(p *SplunkConfiguration, vals map[string]cty.Value) {
	vals["hec_endpoint"] = cty.StringVal(p.HecEndpoint)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_HecEndpointType(p *SplunkConfiguration, vals map[string]cty.Value) {
	vals["hec_endpoint_type"] = cty.StringVal(p.HecEndpointType)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_HecToken(p *SplunkConfiguration, vals map[string]cty.Value) {
	vals["hec_token"] = cty.StringVal(p.HecToken)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_RetryDuration(p *SplunkConfiguration, vals map[string]cty.Value) {
	vals["retry_duration"] = cty.IntVal(p.RetryDuration)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_S3BackupMode(p *SplunkConfiguration, vals map[string]cty.Value) {
	vals["s3_backup_mode"] = cty.StringVal(p.S3BackupMode)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_HecAcknowledgmentTimeout(p *SplunkConfiguration, vals map[string]cty.Value) {
	vals["hec_acknowledgment_timeout"] = cty.IntVal(p.HecAcknowledgmentTimeout)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_CloudwatchLoggingOptions(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CloudwatchLoggingOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_CloudwatchLoggingOptions_LogStreamName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_CloudwatchLoggingOptions_Enabled(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_CloudwatchLoggingOptions_LogGroupName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cloudwatch_logging_options"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_CloudwatchLoggingOptions_LogStreamName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_stream_name"] = cty.StringVal(p.LogStreamName)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_CloudwatchLoggingOptions_Enabled(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_CloudwatchLoggingOptions_LogGroupName(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_group_name"] = cty.StringVal(p.LogGroupName)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration(p *ProcessingConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ProcessingConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration_Enabled(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration_Processors(v.Processors, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["processing_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration_Enabled(p *ProcessingConfiguration, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration_Processors(p *Processors, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Processors {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration_Processors_Type(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration_Processors_Parameters(v.Parameters, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["processors"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration_Processors_Type(p *Processors, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration_Processors_Parameters(p *Parameters, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Parameters {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration_Processors_Parameters_ParameterName(v, ctyVal)
		EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration_Processors_Parameters_ParameterValue(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["parameters"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration_Processors_Parameters_ParameterName(p *Parameters, vals map[string]cty.Value) {
	vals["parameter_name"] = cty.StringVal(p.ParameterName)
}

func EncodeKinesisFirehoseDeliveryStream_SplunkConfiguration_ProcessingConfiguration_Processors_Parameters_ParameterValue(p *Parameters, vals map[string]cty.Value) {
	vals["parameter_value"] = cty.StringVal(p.ParameterValue)
}