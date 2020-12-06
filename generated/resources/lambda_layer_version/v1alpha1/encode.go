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

func EncodeLambdaLayerVersion(r LambdaLayerVersion) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaLayerVersion_CompatibleRuntimes(r.Spec.ForProvider, ctyVal)
	EncodeLambdaLayerVersion_Id(r.Spec.ForProvider, ctyVal)
	EncodeLambdaLayerVersion_LicenseInfo(r.Spec.ForProvider, ctyVal)
	EncodeLambdaLayerVersion_Description(r.Spec.ForProvider, ctyVal)
	EncodeLambdaLayerVersion_S3Bucket(r.Spec.ForProvider, ctyVal)
	EncodeLambdaLayerVersion_S3Key(r.Spec.ForProvider, ctyVal)
	EncodeLambdaLayerVersion_LayerName(r.Spec.ForProvider, ctyVal)
	EncodeLambdaLayerVersion_Filename(r.Spec.ForProvider, ctyVal)
	EncodeLambdaLayerVersion_S3ObjectVersion(r.Spec.ForProvider, ctyVal)
	EncodeLambdaLayerVersion_SourceCodeHash(r.Spec.ForProvider, ctyVal)
	EncodeLambdaLayerVersion_Arn(r.Status.AtProvider, ctyVal)
	EncodeLambdaLayerVersion_CreatedDate(r.Status.AtProvider, ctyVal)
	EncodeLambdaLayerVersion_LayerArn(r.Status.AtProvider, ctyVal)
	EncodeLambdaLayerVersion_SourceCodeSize(r.Status.AtProvider, ctyVal)
	EncodeLambdaLayerVersion_Version(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeLambdaLayerVersion_CompatibleRuntimes(p LambdaLayerVersionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CompatibleRuntimes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["compatible_runtimes"] = cty.SetVal(colVals)
}

func EncodeLambdaLayerVersion_Id(p LambdaLayerVersionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLambdaLayerVersion_LicenseInfo(p LambdaLayerVersionParameters, vals map[string]cty.Value) {
	vals["license_info"] = cty.StringVal(p.LicenseInfo)
}

func EncodeLambdaLayerVersion_Description(p LambdaLayerVersionParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeLambdaLayerVersion_S3Bucket(p LambdaLayerVersionParameters, vals map[string]cty.Value) {
	vals["s3_bucket"] = cty.StringVal(p.S3Bucket)
}

func EncodeLambdaLayerVersion_S3Key(p LambdaLayerVersionParameters, vals map[string]cty.Value) {
	vals["s3_key"] = cty.StringVal(p.S3Key)
}

func EncodeLambdaLayerVersion_LayerName(p LambdaLayerVersionParameters, vals map[string]cty.Value) {
	vals["layer_name"] = cty.StringVal(p.LayerName)
}

func EncodeLambdaLayerVersion_Filename(p LambdaLayerVersionParameters, vals map[string]cty.Value) {
	vals["filename"] = cty.StringVal(p.Filename)
}

func EncodeLambdaLayerVersion_S3ObjectVersion(p LambdaLayerVersionParameters, vals map[string]cty.Value) {
	vals["s3_object_version"] = cty.StringVal(p.S3ObjectVersion)
}

func EncodeLambdaLayerVersion_SourceCodeHash(p LambdaLayerVersionParameters, vals map[string]cty.Value) {
	vals["source_code_hash"] = cty.StringVal(p.SourceCodeHash)
}

func EncodeLambdaLayerVersion_Arn(p LambdaLayerVersionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeLambdaLayerVersion_CreatedDate(p LambdaLayerVersionObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}

func EncodeLambdaLayerVersion_LayerArn(p LambdaLayerVersionObservation, vals map[string]cty.Value) {
	vals["layer_arn"] = cty.StringVal(p.LayerArn)
}

func EncodeLambdaLayerVersion_SourceCodeSize(p LambdaLayerVersionObservation, vals map[string]cty.Value) {
	vals["source_code_size"] = cty.NumberIntVal(p.SourceCodeSize)
}

func EncodeLambdaLayerVersion_Version(p LambdaLayerVersionObservation, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}