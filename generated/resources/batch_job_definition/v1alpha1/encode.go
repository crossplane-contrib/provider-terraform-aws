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

package v1alpha1func EncodeBatchJobDefinition(r BatchJobDefinition) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeBatchJobDefinition_ContainerProperties(r.Spec.ForProvider, ctyVal)
	EncodeBatchJobDefinition_Id(r.Spec.ForProvider, ctyVal)
	EncodeBatchJobDefinition_Name(r.Spec.ForProvider, ctyVal)
	EncodeBatchJobDefinition_Parameters(r.Spec.ForProvider, ctyVal)
	EncodeBatchJobDefinition_Type(r.Spec.ForProvider, ctyVal)
	EncodeBatchJobDefinition_RetryStrategy(r.Spec.ForProvider.RetryStrategy, ctyVal)
	EncodeBatchJobDefinition_Timeout(r.Spec.ForProvider.Timeout, ctyVal)
	EncodeBatchJobDefinition_Arn(r.Status.AtProvider, ctyVal)
	EncodeBatchJobDefinition_Revision(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeBatchJobDefinition_ContainerProperties(p *BatchJobDefinitionParameters, vals map[string]cty.Value) {
	vals["container_properties"] = cty.StringVal(p.ContainerProperties)
}

func EncodeBatchJobDefinition_Id(p *BatchJobDefinitionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeBatchJobDefinition_Name(p *BatchJobDefinitionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeBatchJobDefinition_Parameters(p *BatchJobDefinitionParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Parameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["parameters"] = cty.MapVal(mVals)
}

func EncodeBatchJobDefinition_Type(p *BatchJobDefinitionParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeBatchJobDefinition_RetryStrategy(p *RetryStrategy, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RetryStrategy {
		ctyVal = make(map[string]cty.Value)
		EncodeBatchJobDefinition_RetryStrategy_Attempts(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["retry_strategy"] = cty.ListVal(valsForCollection)
}

func EncodeBatchJobDefinition_RetryStrategy_Attempts(p *RetryStrategy, vals map[string]cty.Value) {
	vals["attempts"] = cty.IntVal(p.Attempts)
}

func EncodeBatchJobDefinition_Timeout(p *Timeout, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Timeout {
		ctyVal = make(map[string]cty.Value)
		EncodeBatchJobDefinition_Timeout_AttemptDurationSeconds(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["timeout"] = cty.ListVal(valsForCollection)
}

func EncodeBatchJobDefinition_Timeout_AttemptDurationSeconds(p *Timeout, vals map[string]cty.Value) {
	vals["attempt_duration_seconds"] = cty.IntVal(p.AttemptDurationSeconds)
}

func EncodeBatchJobDefinition_Arn(p *BatchJobDefinitionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeBatchJobDefinition_Revision(p *BatchJobDefinitionObservation, vals map[string]cty.Value) {
	vals["revision"] = cty.IntVal(p.Revision)
}