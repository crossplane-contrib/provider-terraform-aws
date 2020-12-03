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

package v1alpha1func EncodeApiGatewayDocumentationVersion(r ApiGatewayDocumentationVersion) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeApiGatewayDocumentationVersion_Version(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDocumentationVersion_Description(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDocumentationVersion_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDocumentationVersion_RestApiId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeApiGatewayDocumentationVersion_Version(p *ApiGatewayDocumentationVersionParameters, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeApiGatewayDocumentationVersion_Description(p *ApiGatewayDocumentationVersionParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeApiGatewayDocumentationVersion_Id(p *ApiGatewayDocumentationVersionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayDocumentationVersion_RestApiId(p *ApiGatewayDocumentationVersionParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}