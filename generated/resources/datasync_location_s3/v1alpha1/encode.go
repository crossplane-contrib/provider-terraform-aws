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

func EncodeDatasyncLocationS3(r DatasyncLocationS3) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncLocationS3_Subdirectory(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationS3_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationS3_Id(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationS3_S3BucketArn(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationS3_S3Config(r.Spec.ForProvider.S3Config, ctyVal)
	EncodeDatasyncLocationS3_Uri(r.Status.AtProvider, ctyVal)
	EncodeDatasyncLocationS3_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDatasyncLocationS3_Subdirectory(p DatasyncLocationS3Parameters, vals map[string]cty.Value) {
	vals["subdirectory"] = cty.StringVal(p.Subdirectory)
}

func EncodeDatasyncLocationS3_Tags(p DatasyncLocationS3Parameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDatasyncLocationS3_Id(p DatasyncLocationS3Parameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDatasyncLocationS3_S3BucketArn(p DatasyncLocationS3Parameters, vals map[string]cty.Value) {
	vals["s3_bucket_arn"] = cty.StringVal(p.S3BucketArn)
}

func EncodeDatasyncLocationS3_S3Config(p S3Config, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncLocationS3_S3Config_BucketAccessRoleArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["s3_config"] = cty.ListVal(valsForCollection)
}

func EncodeDatasyncLocationS3_S3Config_BucketAccessRoleArn(p S3Config, vals map[string]cty.Value) {
	vals["bucket_access_role_arn"] = cty.StringVal(p.BucketAccessRoleArn)
}

func EncodeDatasyncLocationS3_Uri(p DatasyncLocationS3Observation, vals map[string]cty.Value) {
	vals["uri"] = cty.StringVal(p.Uri)
}

func EncodeDatasyncLocationS3_Arn(p DatasyncLocationS3Observation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}