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

package v1alpha1func EncodeKinesisAnalyticsApplication(r KinesisAnalyticsApplication) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeKinesisAnalyticsApplication_Code(r.Spec.ForProvider, ctyVal)
	EncodeKinesisAnalyticsApplication_Id(r.Spec.ForProvider, ctyVal)
	EncodeKinesisAnalyticsApplication_Name(r.Spec.ForProvider, ctyVal)
	EncodeKinesisAnalyticsApplication_Description(r.Spec.ForProvider, ctyVal)
	EncodeKinesisAnalyticsApplication_Tags(r.Spec.ForProvider, ctyVal)
	EncodeKinesisAnalyticsApplication_Inputs(r.Spec.ForProvider.Inputs, ctyVal)
	EncodeKinesisAnalyticsApplication_Outputs(r.Spec.ForProvider.Outputs, ctyVal)
	EncodeKinesisAnalyticsApplication_ReferenceDataSources(r.Spec.ForProvider.ReferenceDataSources, ctyVal)
	EncodeKinesisAnalyticsApplication_CloudwatchLoggingOptions(r.Spec.ForProvider.CloudwatchLoggingOptions, ctyVal)
	EncodeKinesisAnalyticsApplication_Version(r.Status.AtProvider, ctyVal)
	EncodeKinesisAnalyticsApplication_CreateTimestamp(r.Status.AtProvider, ctyVal)
	EncodeKinesisAnalyticsApplication_LastUpdateTimestamp(r.Status.AtProvider, ctyVal)
	EncodeKinesisAnalyticsApplication_Status(r.Status.AtProvider, ctyVal)
	EncodeKinesisAnalyticsApplication_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeKinesisAnalyticsApplication_Code(p *KinesisAnalyticsApplicationParameters, vals map[string]cty.Value) {
	vals["code"] = cty.StringVal(p.Code)
}

func EncodeKinesisAnalyticsApplication_Id(p *KinesisAnalyticsApplicationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKinesisAnalyticsApplication_Name(p *KinesisAnalyticsApplicationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeKinesisAnalyticsApplication_Description(p *KinesisAnalyticsApplicationParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeKinesisAnalyticsApplication_Tags(p *KinesisAnalyticsApplicationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeKinesisAnalyticsApplication_Inputs(p *Inputs, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Inputs {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Inputs_Id(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_NamePrefix(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_StartingPositionConfiguration(v.StartingPositionConfiguration, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_StreamNames(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_KinesisFirehose(v.KinesisFirehose, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_KinesisStream(v.KinesisStream, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_Parallelism(v.Parallelism, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_ProcessingConfiguration(v.ProcessingConfiguration, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_Schema(v.Schema, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["inputs"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Inputs_Id(p *Inputs, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKinesisAnalyticsApplication_Inputs_NamePrefix(p *Inputs, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeKinesisAnalyticsApplication_Inputs_StartingPositionConfiguration(p *StartingPositionConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.StartingPositionConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Inputs_StartingPositionConfiguration_StartingPosition(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["starting_position_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Inputs_StartingPositionConfiguration_StartingPosition(p *StartingPositionConfiguration, vals map[string]cty.Value) {
	vals["starting_position"] = cty.StringVal(p.StartingPosition)
}

func EncodeKinesisAnalyticsApplication_Inputs_StreamNames(p *Inputs, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.StreamNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["stream_names"] = cty.SetVal(colVals)
}

func EncodeKinesisAnalyticsApplication_Inputs_KinesisFirehose(p *KinesisFirehose, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.KinesisFirehose {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Inputs_KinesisFirehose_ResourceArn(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_KinesisFirehose_RoleArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["kinesis_firehose"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Inputs_KinesisFirehose_ResourceArn(p *KinesisFirehose, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}

func EncodeKinesisAnalyticsApplication_Inputs_KinesisFirehose_RoleArn(p *KinesisFirehose, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisAnalyticsApplication_Inputs_KinesisStream(p *KinesisStream, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.KinesisStream {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Inputs_KinesisStream_ResourceArn(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_KinesisStream_RoleArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["kinesis_stream"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Inputs_KinesisStream_ResourceArn(p *KinesisStream, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}

func EncodeKinesisAnalyticsApplication_Inputs_KinesisStream_RoleArn(p *KinesisStream, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisAnalyticsApplication_Inputs_Parallelism(p *Parallelism, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Parallelism {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Inputs_Parallelism_Count(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["parallelism"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Inputs_Parallelism_Count(p *Parallelism, vals map[string]cty.Value) {
	vals["count"] = cty.IntVal(p.Count)
}

func EncodeKinesisAnalyticsApplication_Inputs_ProcessingConfiguration(p *ProcessingConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ProcessingConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Inputs_ProcessingConfiguration_Lambda(v.Lambda, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["processing_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Inputs_ProcessingConfiguration_Lambda(p *Lambda, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Lambda {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Inputs_ProcessingConfiguration_Lambda_ResourceArn(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_ProcessingConfiguration_Lambda_RoleArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["lambda"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Inputs_ProcessingConfiguration_Lambda_ResourceArn(p *Lambda, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}

func EncodeKinesisAnalyticsApplication_Inputs_ProcessingConfiguration_Lambda_RoleArn(p *Lambda, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema(p *Schema, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Schema {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordEncoding(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordColumns(v.RecordColumns, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat(v.RecordFormat, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["schema"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordEncoding(p *Schema, vals map[string]cty.Value) {
	vals["record_encoding"] = cty.StringVal(p.RecordEncoding)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordColumns(p *RecordColumns, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RecordColumns {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordColumns_SqlType(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordColumns_Mapping(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordColumns_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["record_columns"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordColumns_SqlType(p *RecordColumns, vals map[string]cty.Value) {
	vals["sql_type"] = cty.StringVal(p.SqlType)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordColumns_Mapping(p *RecordColumns, vals map[string]cty.Value) {
	vals["mapping"] = cty.StringVal(p.Mapping)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordColumns_Name(p *RecordColumns, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat(p *RecordFormat, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RecordFormat {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_RecordFormatType(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_MappingParameters(v.MappingParameters, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["record_format"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_RecordFormatType(p *RecordFormat, vals map[string]cty.Value) {
	vals["record_format_type"] = cty.StringVal(p.RecordFormatType)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_MappingParameters(p *MappingParameters, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.MappingParameters {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_MappingParameters_Csv(v.Csv, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_MappingParameters_Json(v.Json, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["mapping_parameters"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_MappingParameters_Csv(p *Csv, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Csv {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_MappingParameters_Csv_RecordColumnDelimiter(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_MappingParameters_Csv_RecordRowDelimiter(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["csv"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_MappingParameters_Csv_RecordColumnDelimiter(p *Csv, vals map[string]cty.Value) {
	vals["record_column_delimiter"] = cty.StringVal(p.RecordColumnDelimiter)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_MappingParameters_Csv_RecordRowDelimiter(p *Csv, vals map[string]cty.Value) {
	vals["record_row_delimiter"] = cty.StringVal(p.RecordRowDelimiter)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_MappingParameters_Json(p *Json, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Json {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_MappingParameters_Json_RecordRowPath(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["json"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Inputs_Schema_RecordFormat_MappingParameters_Json_RecordRowPath(p *Json, vals map[string]cty.Value) {
	vals["record_row_path"] = cty.StringVal(p.RecordRowPath)
}

func EncodeKinesisAnalyticsApplication_Outputs(p *Outputs, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Outputs {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Outputs_Id(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Outputs_Name(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Outputs_KinesisFirehose(v.KinesisFirehose, ctyVal)
		EncodeKinesisAnalyticsApplication_Outputs_KinesisStream(v.KinesisStream, ctyVal)
		EncodeKinesisAnalyticsApplication_Outputs_Lambda(v.Lambda, ctyVal)
		EncodeKinesisAnalyticsApplication_Outputs_Schema(v.Schema, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["outputs"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Outputs_Id(p *Outputs, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKinesisAnalyticsApplication_Outputs_Name(p *Outputs, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeKinesisAnalyticsApplication_Outputs_KinesisFirehose(p *KinesisFirehose, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.KinesisFirehose {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Outputs_KinesisFirehose_RoleArn(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Outputs_KinesisFirehose_ResourceArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["kinesis_firehose"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Outputs_KinesisFirehose_RoleArn(p *KinesisFirehose, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisAnalyticsApplication_Outputs_KinesisFirehose_ResourceArn(p *KinesisFirehose, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}

func EncodeKinesisAnalyticsApplication_Outputs_KinesisStream(p *KinesisStream, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.KinesisStream {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Outputs_KinesisStream_ResourceArn(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Outputs_KinesisStream_RoleArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["kinesis_stream"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Outputs_KinesisStream_ResourceArn(p *KinesisStream, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}

func EncodeKinesisAnalyticsApplication_Outputs_KinesisStream_RoleArn(p *KinesisStream, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisAnalyticsApplication_Outputs_Lambda(p *Lambda, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Lambda {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Outputs_Lambda_ResourceArn(v, ctyVal)
		EncodeKinesisAnalyticsApplication_Outputs_Lambda_RoleArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["lambda"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Outputs_Lambda_ResourceArn(p *Lambda, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}

func EncodeKinesisAnalyticsApplication_Outputs_Lambda_RoleArn(p *Lambda, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisAnalyticsApplication_Outputs_Schema(p *Schema, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Schema {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_Outputs_Schema_RecordFormatType(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["schema"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_Outputs_Schema_RecordFormatType(p *Schema, vals map[string]cty.Value) {
	vals["record_format_type"] = cty.StringVal(p.RecordFormatType)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources(p *ReferenceDataSources, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ReferenceDataSources {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Id(v, ctyVal)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_TableName(v, ctyVal)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_S3(v.S3, ctyVal)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema(v.Schema, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["reference_data_sources"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Id(p *ReferenceDataSources, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_TableName(p *ReferenceDataSources, vals map[string]cty.Value) {
	vals["table_name"] = cty.StringVal(p.TableName)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_S3(p *S3, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.S3 {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_S3_BucketArn(v, ctyVal)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_S3_FileKey(v, ctyVal)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_S3_RoleArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["s3"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_S3_BucketArn(p *S3, vals map[string]cty.Value) {
	vals["bucket_arn"] = cty.StringVal(p.BucketArn)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_S3_FileKey(p *S3, vals map[string]cty.Value) {
	vals["file_key"] = cty.StringVal(p.FileKey)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_S3_RoleArn(p *S3, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema(p *Schema, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Schema {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordEncoding(v, ctyVal)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat(v.RecordFormat, ctyVal)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordColumns(v.RecordColumns, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["schema"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordEncoding(p *Schema, vals map[string]cty.Value) {
	vals["record_encoding"] = cty.StringVal(p.RecordEncoding)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat(p *RecordFormat, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RecordFormat {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_RecordFormatType(v, ctyVal)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_MappingParameters(v.MappingParameters, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["record_format"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_RecordFormatType(p *RecordFormat, vals map[string]cty.Value) {
	vals["record_format_type"] = cty.StringVal(p.RecordFormatType)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_MappingParameters(p *MappingParameters, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.MappingParameters {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_MappingParameters_Csv(v.Csv, ctyVal)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_MappingParameters_Json(v.Json, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["mapping_parameters"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_MappingParameters_Csv(p *Csv, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Csv {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_MappingParameters_Csv_RecordColumnDelimiter(v, ctyVal)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_MappingParameters_Csv_RecordRowDelimiter(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["csv"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_MappingParameters_Csv_RecordColumnDelimiter(p *Csv, vals map[string]cty.Value) {
	vals["record_column_delimiter"] = cty.StringVal(p.RecordColumnDelimiter)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_MappingParameters_Csv_RecordRowDelimiter(p *Csv, vals map[string]cty.Value) {
	vals["record_row_delimiter"] = cty.StringVal(p.RecordRowDelimiter)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_MappingParameters_Json(p *Json, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Json {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_MappingParameters_Json_RecordRowPath(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["json"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordFormat_MappingParameters_Json_RecordRowPath(p *Json, vals map[string]cty.Value) {
	vals["record_row_path"] = cty.StringVal(p.RecordRowPath)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordColumns(p *RecordColumns, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RecordColumns {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordColumns_Mapping(v, ctyVal)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordColumns_Name(v, ctyVal)
		EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordColumns_SqlType(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["record_columns"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordColumns_Mapping(p *RecordColumns, vals map[string]cty.Value) {
	vals["mapping"] = cty.StringVal(p.Mapping)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordColumns_Name(p *RecordColumns, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeKinesisAnalyticsApplication_ReferenceDataSources_Schema_RecordColumns_SqlType(p *RecordColumns, vals map[string]cty.Value) {
	vals["sql_type"] = cty.StringVal(p.SqlType)
}

func EncodeKinesisAnalyticsApplication_CloudwatchLoggingOptions(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CloudwatchLoggingOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeKinesisAnalyticsApplication_CloudwatchLoggingOptions_LogStreamArn(v, ctyVal)
		EncodeKinesisAnalyticsApplication_CloudwatchLoggingOptions_RoleArn(v, ctyVal)
		EncodeKinesisAnalyticsApplication_CloudwatchLoggingOptions_Id(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cloudwatch_logging_options"] = cty.ListVal(valsForCollection)
}

func EncodeKinesisAnalyticsApplication_CloudwatchLoggingOptions_LogStreamArn(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["log_stream_arn"] = cty.StringVal(p.LogStreamArn)
}

func EncodeKinesisAnalyticsApplication_CloudwatchLoggingOptions_RoleArn(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeKinesisAnalyticsApplication_CloudwatchLoggingOptions_Id(p *CloudwatchLoggingOptions, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKinesisAnalyticsApplication_Version(p *KinesisAnalyticsApplicationObservation, vals map[string]cty.Value) {
	vals["version"] = cty.IntVal(p.Version)
}

func EncodeKinesisAnalyticsApplication_CreateTimestamp(p *KinesisAnalyticsApplicationObservation, vals map[string]cty.Value) {
	vals["create_timestamp"] = cty.StringVal(p.CreateTimestamp)
}

func EncodeKinesisAnalyticsApplication_LastUpdateTimestamp(p *KinesisAnalyticsApplicationObservation, vals map[string]cty.Value) {
	vals["last_update_timestamp"] = cty.StringVal(p.LastUpdateTimestamp)
}

func EncodeKinesisAnalyticsApplication_Status(p *KinesisAnalyticsApplicationObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeKinesisAnalyticsApplication_Arn(p *KinesisAnalyticsApplicationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}