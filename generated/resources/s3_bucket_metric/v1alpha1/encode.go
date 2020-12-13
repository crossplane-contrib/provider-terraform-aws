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
	r, ok := mr.(*S3BucketMetric)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a S3BucketMetric.")
	}
	return EncodeS3BucketMetric(*r), nil
}

func EncodeS3BucketMetric(r S3BucketMetric) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketMetric_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketMetric_Id(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketMetric_Name(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketMetric_Filter(r.Spec.ForProvider.Filter, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeS3BucketMetric_Bucket(p S3BucketMetricParameters, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeS3BucketMetric_Id(p S3BucketMetricParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3BucketMetric_Name(p S3BucketMetricParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeS3BucketMetric_Filter(p Filter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketMetric_Filter_Tags(p, ctyVal)
	EncodeS3BucketMetric_Filter_Prefix(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["filter"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketMetric_Filter_Tags(p Filter, vals map[string]cty.Value) {
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

func EncodeS3BucketMetric_Filter_Prefix(p Filter, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}