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
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*DatasyncLocationS3)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DatasyncLocationS3.")
	}
	return EncodeDatasyncLocationS3(*r), nil
}

func EncodeDatasyncLocationS3(r DatasyncLocationS3) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncLocationS3_Id(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationS3_S3BucketArn(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationS3_Subdirectory(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationS3_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationS3_S3Config(r.Spec.ForProvider.S3Config, ctyVal)
	EncodeDatasyncLocationS3_Uri(r.Status.AtProvider, ctyVal)
	EncodeDatasyncLocationS3_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDatasyncLocationS3_Id(p DatasyncLocationS3Parameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDatasyncLocationS3_S3BucketArn(p DatasyncLocationS3Parameters, vals map[string]cty.Value) {
	vals["s3_bucket_arn"] = cty.StringVal(p.S3BucketArn)
}

func EncodeDatasyncLocationS3_Subdirectory(p DatasyncLocationS3Parameters, vals map[string]cty.Value) {
	vals["subdirectory"] = cty.StringVal(p.Subdirectory)
}

func EncodeDatasyncLocationS3_Tags(p DatasyncLocationS3Parameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
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