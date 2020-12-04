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

func EncodeCloudwatchLogMetricFilter(r CloudwatchLogMetricFilter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchLogMetricFilter_Name(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogMetricFilter_Pattern(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogMetricFilter_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogMetricFilter_LogGroupName(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogMetricFilter_MetricTransformation(r.Spec.ForProvider.MetricTransformation, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeCloudwatchLogMetricFilter_Name(p CloudwatchLogMetricFilterParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloudwatchLogMetricFilter_Pattern(p CloudwatchLogMetricFilterParameters, vals map[string]cty.Value) {
	vals["pattern"] = cty.StringVal(p.Pattern)
}

func EncodeCloudwatchLogMetricFilter_Id(p CloudwatchLogMetricFilterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudwatchLogMetricFilter_LogGroupName(p CloudwatchLogMetricFilterParameters, vals map[string]cty.Value) {
	vals["log_group_name"] = cty.StringVal(p.LogGroupName)
}

func EncodeCloudwatchLogMetricFilter_MetricTransformation(p MetricTransformation, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchLogMetricFilter_MetricTransformation_DefaultValue(p, ctyVal)
	EncodeCloudwatchLogMetricFilter_MetricTransformation_Name(p, ctyVal)
	EncodeCloudwatchLogMetricFilter_MetricTransformation_Namespace(p, ctyVal)
	EncodeCloudwatchLogMetricFilter_MetricTransformation_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["metric_transformation"] = cty.ListVal(valsForCollection)
}

func EncodeCloudwatchLogMetricFilter_MetricTransformation_DefaultValue(p MetricTransformation, vals map[string]cty.Value) {
	vals["default_value"] = cty.StringVal(p.DefaultValue)
}

func EncodeCloudwatchLogMetricFilter_MetricTransformation_Name(p MetricTransformation, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloudwatchLogMetricFilter_MetricTransformation_Namespace(p MetricTransformation, vals map[string]cty.Value) {
	vals["namespace"] = cty.StringVal(p.Namespace)
}

func EncodeCloudwatchLogMetricFilter_MetricTransformation_Value(p MetricTransformation, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}