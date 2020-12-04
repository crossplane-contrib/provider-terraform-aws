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

func EncodeApigatewayv2Deployment(r Apigatewayv2Deployment) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2Deployment_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Deployment_Description(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Deployment_Id(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Deployment_Triggers(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Deployment_AutoDeployed(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeApigatewayv2Deployment_ApiId(p Apigatewayv2DeploymentParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeApigatewayv2Deployment_Description(p Apigatewayv2DeploymentParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeApigatewayv2Deployment_Id(p Apigatewayv2DeploymentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApigatewayv2Deployment_Triggers(p Apigatewayv2DeploymentParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Triggers {
		mVals[key] = cty.StringVal(value)
	}
	vals["triggers"] = cty.MapVal(mVals)
}

func EncodeApigatewayv2Deployment_AutoDeployed(p Apigatewayv2DeploymentObservation, vals map[string]cty.Value) {
	vals["auto_deployed"] = cty.BoolVal(p.AutoDeployed)
}