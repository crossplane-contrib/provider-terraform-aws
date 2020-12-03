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

package v1alpha1func EncodeApigatewayv2VpcLink(r Apigatewayv2VpcLink) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeApigatewayv2VpcLink_Id(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2VpcLink_Name(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2VpcLink_SecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2VpcLink_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2VpcLink_Tags(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2VpcLink_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeApigatewayv2VpcLink_Id(p *Apigatewayv2VpcLinkParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApigatewayv2VpcLink_Name(p *Apigatewayv2VpcLinkParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeApigatewayv2VpcLink_SecurityGroupIds(p *Apigatewayv2VpcLinkParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeApigatewayv2VpcLink_SubnetIds(p *Apigatewayv2VpcLinkParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeApigatewayv2VpcLink_Tags(p *Apigatewayv2VpcLinkParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeApigatewayv2VpcLink_Arn(p *Apigatewayv2VpcLinkObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}