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

func EncodeS3BucketMetric(r S3BucketMetric) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketMetric_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketMetric_Id(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketMetric_Name(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketMetric_Filter(r.Spec.ForProvider.Filter, ctyVal)

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
	EncodeS3BucketMetric_Filter_Prefix(p, ctyVal)
	EncodeS3BucketMetric_Filter_Tags(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["filter"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketMetric_Filter_Prefix(p Filter, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeS3BucketMetric_Filter_Tags(p Filter, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}