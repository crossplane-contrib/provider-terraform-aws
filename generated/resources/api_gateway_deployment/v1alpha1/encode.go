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

package v1alpha1func EncodeApiGatewayDeployment(r ApiGatewayDeployment) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeApiGatewayDeployment_Description(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDeployment_Variables(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDeployment_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDeployment_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDeployment_StageDescription(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDeployment_StageName(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDeployment_Triggers(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDeployment_ExecutionArn(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayDeployment_InvokeUrl(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayDeployment_CreatedDate(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeApiGatewayDeployment_Description(p *ApiGatewayDeploymentParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeApiGatewayDeployment_Variables(p *ApiGatewayDeploymentParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Variables {
		mVals[key] = cty.StringVal(value)
	}
	vals["variables"] = cty.MapVal(mVals)
}

func EncodeApiGatewayDeployment_Id(p *ApiGatewayDeploymentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayDeployment_RestApiId(p *ApiGatewayDeploymentParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayDeployment_StageDescription(p *ApiGatewayDeploymentParameters, vals map[string]cty.Value) {
	vals["stage_description"] = cty.StringVal(p.StageDescription)
}

func EncodeApiGatewayDeployment_StageName(p *ApiGatewayDeploymentParameters, vals map[string]cty.Value) {
	vals["stage_name"] = cty.StringVal(p.StageName)
}

func EncodeApiGatewayDeployment_Triggers(p *ApiGatewayDeploymentParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Triggers {
		mVals[key] = cty.StringVal(value)
	}
	vals["triggers"] = cty.MapVal(mVals)
}

func EncodeApiGatewayDeployment_ExecutionArn(p *ApiGatewayDeploymentObservation, vals map[string]cty.Value) {
	vals["execution_arn"] = cty.StringVal(p.ExecutionArn)
}

func EncodeApiGatewayDeployment_InvokeUrl(p *ApiGatewayDeploymentObservation, vals map[string]cty.Value) {
	vals["invoke_url"] = cty.StringVal(p.InvokeUrl)
}

func EncodeApiGatewayDeployment_CreatedDate(p *ApiGatewayDeploymentObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}