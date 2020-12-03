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

package v1alpha1func EncodeAccessanalyzerAnalyzer(r AccessanalyzerAnalyzer) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeAccessanalyzerAnalyzer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAccessanalyzerAnalyzer_Type(r.Spec.ForProvider, ctyVal)
	EncodeAccessanalyzerAnalyzer_AnalyzerName(r.Spec.ForProvider, ctyVal)
	EncodeAccessanalyzerAnalyzer_Id(r.Spec.ForProvider, ctyVal)
	EncodeAccessanalyzerAnalyzer_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeAccessanalyzerAnalyzer_Tags(p *AccessanalyzerAnalyzerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAccessanalyzerAnalyzer_Type(p *AccessanalyzerAnalyzerParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeAccessanalyzerAnalyzer_AnalyzerName(p *AccessanalyzerAnalyzerParameters, vals map[string]cty.Value) {
	vals["analyzer_name"] = cty.StringVal(p.AnalyzerName)
}

func EncodeAccessanalyzerAnalyzer_Id(p *AccessanalyzerAnalyzerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAccessanalyzerAnalyzer_Arn(p *AccessanalyzerAnalyzerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}