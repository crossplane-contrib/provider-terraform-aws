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

package v1alpha1func EncodeElastictranscoderPipeline(r ElastictranscoderPipeline) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeElastictranscoderPipeline_Name(r.Spec.ForProvider, ctyVal)
	EncodeElastictranscoderPipeline_OutputBucket(r.Spec.ForProvider, ctyVal)
	EncodeElastictranscoderPipeline_Role(r.Spec.ForProvider, ctyVal)
	EncodeElastictranscoderPipeline_AwsKmsKeyArn(r.Spec.ForProvider, ctyVal)
	EncodeElastictranscoderPipeline_Id(r.Spec.ForProvider, ctyVal)
	EncodeElastictranscoderPipeline_InputBucket(r.Spec.ForProvider, ctyVal)
	EncodeElastictranscoderPipeline_ContentConfig(r.Spec.ForProvider.ContentConfig, ctyVal)
	EncodeElastictranscoderPipeline_ContentConfigPermissions(r.Spec.ForProvider.ContentConfigPermissions, ctyVal)
	EncodeElastictranscoderPipeline_Notifications(r.Spec.ForProvider.Notifications, ctyVal)
	EncodeElastictranscoderPipeline_ThumbnailConfig(r.Spec.ForProvider.ThumbnailConfig, ctyVal)
	EncodeElastictranscoderPipeline_ThumbnailConfigPermissions(r.Spec.ForProvider.ThumbnailConfigPermissions, ctyVal)
	EncodeElastictranscoderPipeline_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeElastictranscoderPipeline_Name(p *ElastictranscoderPipelineParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElastictranscoderPipeline_OutputBucket(p *ElastictranscoderPipelineParameters, vals map[string]cty.Value) {
	vals["output_bucket"] = cty.StringVal(p.OutputBucket)
}

func EncodeElastictranscoderPipeline_Role(p *ElastictranscoderPipelineParameters, vals map[string]cty.Value) {
	vals["role"] = cty.StringVal(p.Role)
}

func EncodeElastictranscoderPipeline_AwsKmsKeyArn(p *ElastictranscoderPipelineParameters, vals map[string]cty.Value) {
	vals["aws_kms_key_arn"] = cty.StringVal(p.AwsKmsKeyArn)
}

func EncodeElastictranscoderPipeline_Id(p *ElastictranscoderPipelineParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElastictranscoderPipeline_InputBucket(p *ElastictranscoderPipelineParameters, vals map[string]cty.Value) {
	vals["input_bucket"] = cty.StringVal(p.InputBucket)
}

func EncodeElastictranscoderPipeline_ContentConfig(p *ContentConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ContentConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeElastictranscoderPipeline_ContentConfig_Bucket(v, ctyVal)
		EncodeElastictranscoderPipeline_ContentConfig_StorageClass(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["content_config"] = cty.ListVal(valsForCollection)
}

func EncodeElastictranscoderPipeline_ContentConfig_Bucket(p *ContentConfig, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeElastictranscoderPipeline_ContentConfig_StorageClass(p *ContentConfig, vals map[string]cty.Value) {
	vals["storage_class"] = cty.StringVal(p.StorageClass)
}

func EncodeElastictranscoderPipeline_ContentConfigPermissions(p *ContentConfigPermissions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ContentConfigPermissions {
		ctyVal = make(map[string]cty.Value)
		EncodeElastictranscoderPipeline_ContentConfigPermissions_Access(v, ctyVal)
		EncodeElastictranscoderPipeline_ContentConfigPermissions_Grantee(v, ctyVal)
		EncodeElastictranscoderPipeline_ContentConfigPermissions_GranteeType(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["content_config_permissions"] = cty.SetVal(valsForCollection)
}

func EncodeElastictranscoderPipeline_ContentConfigPermissions_Access(p *ContentConfigPermissions, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Access {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["access"] = cty.ListVal(colVals)
}

func EncodeElastictranscoderPipeline_ContentConfigPermissions_Grantee(p *ContentConfigPermissions, vals map[string]cty.Value) {
	vals["grantee"] = cty.StringVal(p.Grantee)
}

func EncodeElastictranscoderPipeline_ContentConfigPermissions_GranteeType(p *ContentConfigPermissions, vals map[string]cty.Value) {
	vals["grantee_type"] = cty.StringVal(p.GranteeType)
}

func EncodeElastictranscoderPipeline_Notifications(p *Notifications, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Notifications {
		ctyVal = make(map[string]cty.Value)
		EncodeElastictranscoderPipeline_Notifications_Completed(v, ctyVal)
		EncodeElastictranscoderPipeline_Notifications_Error(v, ctyVal)
		EncodeElastictranscoderPipeline_Notifications_Progressing(v, ctyVal)
		EncodeElastictranscoderPipeline_Notifications_Warning(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["notifications"] = cty.ListVal(valsForCollection)
}

func EncodeElastictranscoderPipeline_Notifications_Completed(p *Notifications, vals map[string]cty.Value) {
	vals["completed"] = cty.StringVal(p.Completed)
}

func EncodeElastictranscoderPipeline_Notifications_Error(p *Notifications, vals map[string]cty.Value) {
	vals["error"] = cty.StringVal(p.Error)
}

func EncodeElastictranscoderPipeline_Notifications_Progressing(p *Notifications, vals map[string]cty.Value) {
	vals["progressing"] = cty.StringVal(p.Progressing)
}

func EncodeElastictranscoderPipeline_Notifications_Warning(p *Notifications, vals map[string]cty.Value) {
	vals["warning"] = cty.StringVal(p.Warning)
}

func EncodeElastictranscoderPipeline_ThumbnailConfig(p *ThumbnailConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ThumbnailConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeElastictranscoderPipeline_ThumbnailConfig_Bucket(v, ctyVal)
		EncodeElastictranscoderPipeline_ThumbnailConfig_StorageClass(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["thumbnail_config"] = cty.ListVal(valsForCollection)
}

func EncodeElastictranscoderPipeline_ThumbnailConfig_Bucket(p *ThumbnailConfig, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeElastictranscoderPipeline_ThumbnailConfig_StorageClass(p *ThumbnailConfig, vals map[string]cty.Value) {
	vals["storage_class"] = cty.StringVal(p.StorageClass)
}

func EncodeElastictranscoderPipeline_ThumbnailConfigPermissions(p *ThumbnailConfigPermissions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ThumbnailConfigPermissions {
		ctyVal = make(map[string]cty.Value)
		EncodeElastictranscoderPipeline_ThumbnailConfigPermissions_Access(v, ctyVal)
		EncodeElastictranscoderPipeline_ThumbnailConfigPermissions_Grantee(v, ctyVal)
		EncodeElastictranscoderPipeline_ThumbnailConfigPermissions_GranteeType(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["thumbnail_config_permissions"] = cty.SetVal(valsForCollection)
}

func EncodeElastictranscoderPipeline_ThumbnailConfigPermissions_Access(p *ThumbnailConfigPermissions, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Access {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["access"] = cty.ListVal(colVals)
}

func EncodeElastictranscoderPipeline_ThumbnailConfigPermissions_Grantee(p *ThumbnailConfigPermissions, vals map[string]cty.Value) {
	vals["grantee"] = cty.StringVal(p.Grantee)
}

func EncodeElastictranscoderPipeline_ThumbnailConfigPermissions_GranteeType(p *ThumbnailConfigPermissions, vals map[string]cty.Value) {
	vals["grantee_type"] = cty.StringVal(p.GranteeType)
}

func EncodeElastictranscoderPipeline_Arn(p *ElastictranscoderPipelineObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}