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

package v1alpha1func EncodeLambdaAlias(r LambdaAlias) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeLambdaAlias_FunctionVersion(r.Spec.ForProvider, ctyVal)
	EncodeLambdaAlias_Id(r.Spec.ForProvider, ctyVal)
	EncodeLambdaAlias_Name(r.Spec.ForProvider, ctyVal)
	EncodeLambdaAlias_Description(r.Spec.ForProvider, ctyVal)
	EncodeLambdaAlias_FunctionName(r.Spec.ForProvider, ctyVal)
	EncodeLambdaAlias_RoutingConfig(r.Spec.ForProvider.RoutingConfig, ctyVal)
	EncodeLambdaAlias_InvokeArn(r.Status.AtProvider, ctyVal)
	EncodeLambdaAlias_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeLambdaAlias_FunctionVersion(p *LambdaAliasParameters, vals map[string]cty.Value) {
	vals["function_version"] = cty.StringVal(p.FunctionVersion)
}

func EncodeLambdaAlias_Id(p *LambdaAliasParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLambdaAlias_Name(p *LambdaAliasParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLambdaAlias_Description(p *LambdaAliasParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeLambdaAlias_FunctionName(p *LambdaAliasParameters, vals map[string]cty.Value) {
	vals["function_name"] = cty.StringVal(p.FunctionName)
}

func EncodeLambdaAlias_RoutingConfig(p *RoutingConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RoutingConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeLambdaAlias_RoutingConfig_AdditionalVersionWeights(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["routing_config"] = cty.ListVal(valsForCollection)
}

func EncodeLambdaAlias_RoutingConfig_AdditionalVersionWeights(p *RoutingConfig, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.AdditionalVersionWeights {
		mVals[key] = cty.IntVal(value)
	}
	vals["additional_version_weights"] = cty.MapVal(mVals)
}

func EncodeLambdaAlias_InvokeArn(p *LambdaAliasObservation, vals map[string]cty.Value) {
	vals["invoke_arn"] = cty.StringVal(p.InvokeArn)
}

func EncodeLambdaAlias_Arn(p *LambdaAliasObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}