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

package v1alpha1func EncodeApiGatewayResource(r ApiGatewayResource) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeApiGatewayResource_PathPart(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayResource_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayResource_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayResource_ParentId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayResource_Path(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeApiGatewayResource_PathPart(p *ApiGatewayResourceParameters, vals map[string]cty.Value) {
	vals["path_part"] = cty.StringVal(p.PathPart)
}

func EncodeApiGatewayResource_RestApiId(p *ApiGatewayResourceParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayResource_Id(p *ApiGatewayResourceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayResource_ParentId(p *ApiGatewayResourceParameters, vals map[string]cty.Value) {
	vals["parent_id"] = cty.StringVal(p.ParentId)
}

func EncodeApiGatewayResource_Path(p *ApiGatewayResourceObservation, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}