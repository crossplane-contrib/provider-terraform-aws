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

func EncodeCloudwatchLogStream(r CloudwatchLogStream) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchLogStream_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogStream_LogGroupName(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogStream_Name(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogStream_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudwatchLogStream_Id(p CloudwatchLogStreamParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudwatchLogStream_LogGroupName(p CloudwatchLogStreamParameters, vals map[string]cty.Value) {
	vals["log_group_name"] = cty.StringVal(p.LogGroupName)
}

func EncodeCloudwatchLogStream_Name(p CloudwatchLogStreamParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloudwatchLogStream_Arn(p CloudwatchLogStreamObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}