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

func EncodeMqConfiguration(r MqConfiguration) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeMqConfiguration_Data(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_EngineType(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_EngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_Tags(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_Description(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_Name(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_Arn(r.Status.AtProvider, ctyVal)
	EncodeMqConfiguration_LatestRevision(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeMqConfiguration_Data(p MqConfigurationParameters, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeMqConfiguration_EngineType(p MqConfigurationParameters, vals map[string]cty.Value) {
	vals["engine_type"] = cty.StringVal(p.EngineType)
}

func EncodeMqConfiguration_EngineVersion(p MqConfigurationParameters, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeMqConfiguration_Tags(p MqConfigurationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeMqConfiguration_Description(p MqConfigurationParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeMqConfiguration_Id(p MqConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeMqConfiguration_Name(p MqConfigurationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeMqConfiguration_Arn(p MqConfigurationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeMqConfiguration_LatestRevision(p MqConfigurationObservation, vals map[string]cty.Value) {
	vals["latest_revision"] = cty.NumberIntVal(p.LatestRevision)
}