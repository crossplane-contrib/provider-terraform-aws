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

package v1alpha1func EncodeSesEventDestination(r SesEventDestination) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSesEventDestination_ConfigurationSetName(r.Spec.ForProvider, ctyVal)
	EncodeSesEventDestination_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeSesEventDestination_Id(r.Spec.ForProvider, ctyVal)
	EncodeSesEventDestination_MatchingTypes(r.Spec.ForProvider, ctyVal)
	EncodeSesEventDestination_Name(r.Spec.ForProvider, ctyVal)
	EncodeSesEventDestination_CloudwatchDestination(r.Spec.ForProvider.CloudwatchDestination, ctyVal)
	EncodeSesEventDestination_KinesisDestination(r.Spec.ForProvider.KinesisDestination, ctyVal)
	EncodeSesEventDestination_SnsDestination(r.Spec.ForProvider.SnsDestination, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeSesEventDestination_ConfigurationSetName(p *SesEventDestinationParameters, vals map[string]cty.Value) {
	vals["configuration_set_name"] = cty.StringVal(p.ConfigurationSetName)
}

func EncodeSesEventDestination_Enabled(p *SesEventDestinationParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeSesEventDestination_Id(p *SesEventDestinationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSesEventDestination_MatchingTypes(p *SesEventDestinationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.MatchingTypes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["matching_types"] = cty.SetVal(colVals)
}

func EncodeSesEventDestination_Name(p *SesEventDestinationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSesEventDestination_CloudwatchDestination(p *CloudwatchDestination, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CloudwatchDestination {
		ctyVal = make(map[string]cty.Value)
		EncodeSesEventDestination_CloudwatchDestination_DefaultValue(v, ctyVal)
		EncodeSesEventDestination_CloudwatchDestination_DimensionName(v, ctyVal)
		EncodeSesEventDestination_CloudwatchDestination_ValueSource(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cloudwatch_destination"] = cty.SetVal(valsForCollection)
}

func EncodeSesEventDestination_CloudwatchDestination_DefaultValue(p *CloudwatchDestination, vals map[string]cty.Value) {
	vals["default_value"] = cty.StringVal(p.DefaultValue)
}

func EncodeSesEventDestination_CloudwatchDestination_DimensionName(p *CloudwatchDestination, vals map[string]cty.Value) {
	vals["dimension_name"] = cty.StringVal(p.DimensionName)
}

func EncodeSesEventDestination_CloudwatchDestination_ValueSource(p *CloudwatchDestination, vals map[string]cty.Value) {
	vals["value_source"] = cty.StringVal(p.ValueSource)
}

func EncodeSesEventDestination_KinesisDestination(p *KinesisDestination, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.KinesisDestination {
		ctyVal = make(map[string]cty.Value)
		EncodeSesEventDestination_KinesisDestination_RoleArn(v, ctyVal)
		EncodeSesEventDestination_KinesisDestination_StreamArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["kinesis_destination"] = cty.ListVal(valsForCollection)
}

func EncodeSesEventDestination_KinesisDestination_RoleArn(p *KinesisDestination, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeSesEventDestination_KinesisDestination_StreamArn(p *KinesisDestination, vals map[string]cty.Value) {
	vals["stream_arn"] = cty.StringVal(p.StreamArn)
}

func EncodeSesEventDestination_SnsDestination(p *SnsDestination, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SnsDestination {
		ctyVal = make(map[string]cty.Value)
		EncodeSesEventDestination_SnsDestination_TopicArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sns_destination"] = cty.ListVal(valsForCollection)
}

func EncodeSesEventDestination_SnsDestination_TopicArn(p *SnsDestination, vals map[string]cty.Value) {
	vals["topic_arn"] = cty.StringVal(p.TopicArn)
}