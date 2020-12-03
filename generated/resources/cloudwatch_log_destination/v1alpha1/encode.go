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

package v1alpha1func EncodeCloudwatchLogDestination(r CloudwatchLogDestination) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCloudwatchLogDestination_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogDestination_TargetArn(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogDestination_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogDestination_Name(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogDestination_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeCloudwatchLogDestination_RoleArn(p *CloudwatchLogDestinationParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeCloudwatchLogDestination_TargetArn(p *CloudwatchLogDestinationParameters, vals map[string]cty.Value) {
	vals["target_arn"] = cty.StringVal(p.TargetArn)
}

func EncodeCloudwatchLogDestination_Id(p *CloudwatchLogDestinationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudwatchLogDestination_Name(p *CloudwatchLogDestinationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloudwatchLogDestination_Arn(p *CloudwatchLogDestinationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}