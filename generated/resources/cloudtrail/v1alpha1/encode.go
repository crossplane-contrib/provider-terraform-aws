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

package v1alpha1func EncodeCloudtrail(r Cloudtrail) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCloudtrail_S3BucketName(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_IsMultiRegionTrail(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_IsOrganizationTrail(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_Name(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_SnsTopicName(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_CloudWatchLogsGroupArn(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_CloudWatchLogsRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_IncludeGlobalServiceEvents(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_S3KeyPrefix(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_EnableLogFileValidation(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_EnableLogging(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeCloudtrail_EventSelector(r.Spec.ForProvider.EventSelector, ctyVal)
	EncodeCloudtrail_InsightSelector(r.Spec.ForProvider.InsightSelector, ctyVal)
	EncodeCloudtrail_Arn(r.Status.AtProvider, ctyVal)
	EncodeCloudtrail_HomeRegion(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeCloudtrail_S3BucketName(p *CloudtrailParameters, vals map[string]cty.Value) {
	vals["s3_bucket_name"] = cty.StringVal(p.S3BucketName)
}

func EncodeCloudtrail_Tags(p *CloudtrailParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCloudtrail_IsMultiRegionTrail(p *CloudtrailParameters, vals map[string]cty.Value) {
	vals["is_multi_region_trail"] = cty.BoolVal(p.IsMultiRegionTrail)
}

func EncodeCloudtrail_IsOrganizationTrail(p *CloudtrailParameters, vals map[string]cty.Value) {
	vals["is_organization_trail"] = cty.BoolVal(p.IsOrganizationTrail)
}

func EncodeCloudtrail_Name(p *CloudtrailParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloudtrail_SnsTopicName(p *CloudtrailParameters, vals map[string]cty.Value) {
	vals["sns_topic_name"] = cty.StringVal(p.SnsTopicName)
}

func EncodeCloudtrail_CloudWatchLogsGroupArn(p *CloudtrailParameters, vals map[string]cty.Value) {
	vals["cloud_watch_logs_group_arn"] = cty.StringVal(p.CloudWatchLogsGroupArn)
}

func EncodeCloudtrail_CloudWatchLogsRoleArn(p *CloudtrailParameters, vals map[string]cty.Value) {
	vals["cloud_watch_logs_role_arn"] = cty.StringVal(p.CloudWatchLogsRoleArn)
}

func EncodeCloudtrail_IncludeGlobalServiceEvents(p *CloudtrailParameters, vals map[string]cty.Value) {
	vals["include_global_service_events"] = cty.BoolVal(p.IncludeGlobalServiceEvents)
}

func EncodeCloudtrail_S3KeyPrefix(p *CloudtrailParameters, vals map[string]cty.Value) {
	vals["s3_key_prefix"] = cty.StringVal(p.S3KeyPrefix)
}

func EncodeCloudtrail_EnableLogFileValidation(p *CloudtrailParameters, vals map[string]cty.Value) {
	vals["enable_log_file_validation"] = cty.BoolVal(p.EnableLogFileValidation)
}

func EncodeCloudtrail_EnableLogging(p *CloudtrailParameters, vals map[string]cty.Value) {
	vals["enable_logging"] = cty.BoolVal(p.EnableLogging)
}

func EncodeCloudtrail_Id(p *CloudtrailParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudtrail_KmsKeyId(p *CloudtrailParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeCloudtrail_EventSelector(p *EventSelector, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EventSelector {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudtrail_EventSelector_IncludeManagementEvents(v, ctyVal)
		EncodeCloudtrail_EventSelector_ReadWriteType(v, ctyVal)
		EncodeCloudtrail_EventSelector_DataResource(v.DataResource, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["event_selector"] = cty.ListVal(valsForCollection)
}

func EncodeCloudtrail_EventSelector_IncludeManagementEvents(p *EventSelector, vals map[string]cty.Value) {
	vals["include_management_events"] = cty.BoolVal(p.IncludeManagementEvents)
}

func EncodeCloudtrail_EventSelector_ReadWriteType(p *EventSelector, vals map[string]cty.Value) {
	vals["read_write_type"] = cty.StringVal(p.ReadWriteType)
}

func EncodeCloudtrail_EventSelector_DataResource(p *DataResource, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DataResource {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudtrail_EventSelector_DataResource_Type(v, ctyVal)
		EncodeCloudtrail_EventSelector_DataResource_Values(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["data_resource"] = cty.ListVal(valsForCollection)
}

func EncodeCloudtrail_EventSelector_DataResource_Type(p *DataResource, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCloudtrail_EventSelector_DataResource_Values(p *DataResource, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.ListVal(colVals)
}

func EncodeCloudtrail_InsightSelector(p *InsightSelector, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.InsightSelector {
		ctyVal = make(map[string]cty.Value)
		EncodeCloudtrail_InsightSelector_InsightType(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["insight_selector"] = cty.ListVal(valsForCollection)
}

func EncodeCloudtrail_InsightSelector_InsightType(p *InsightSelector, vals map[string]cty.Value) {
	vals["insight_type"] = cty.StringVal(p.InsightType)
}

func EncodeCloudtrail_Arn(p *CloudtrailObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeCloudtrail_HomeRegion(p *CloudtrailObservation, vals map[string]cty.Value) {
	vals["home_region"] = cty.StringVal(p.HomeRegion)
}