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

func EncodeApigatewayv2DomainName(r Apigatewayv2DomainName) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2DomainName_DomainName(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2DomainName_Id(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2DomainName_Tags(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2DomainName_DomainNameConfiguration(r.Spec.ForProvider.DomainNameConfiguration, ctyVal)
	EncodeApigatewayv2DomainName_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeApigatewayv2DomainName_ApiMappingSelectionExpression(r.Status.AtProvider, ctyVal)
	EncodeApigatewayv2DomainName_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeApigatewayv2DomainName_DomainName(p Apigatewayv2DomainNameParameters, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeApigatewayv2DomainName_Id(p Apigatewayv2DomainNameParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApigatewayv2DomainName_Tags(p Apigatewayv2DomainNameParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeApigatewayv2DomainName_DomainNameConfiguration(p DomainNameConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2DomainName_DomainNameConfiguration_CertificateArn(p, ctyVal)
	EncodeApigatewayv2DomainName_DomainNameConfiguration_EndpointType(p, ctyVal)
	EncodeApigatewayv2DomainName_DomainNameConfiguration_HostedZoneId(p, ctyVal)
	EncodeApigatewayv2DomainName_DomainNameConfiguration_SecurityPolicy(p, ctyVal)
	EncodeApigatewayv2DomainName_DomainNameConfiguration_TargetDomainName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["domain_name_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeApigatewayv2DomainName_DomainNameConfiguration_CertificateArn(p DomainNameConfiguration, vals map[string]cty.Value) {
	vals["certificate_arn"] = cty.StringVal(p.CertificateArn)
}

func EncodeApigatewayv2DomainName_DomainNameConfiguration_EndpointType(p DomainNameConfiguration, vals map[string]cty.Value) {
	vals["endpoint_type"] = cty.StringVal(p.EndpointType)
}

func EncodeApigatewayv2DomainName_DomainNameConfiguration_HostedZoneId(p DomainNameConfiguration, vals map[string]cty.Value) {
	vals["hosted_zone_id"] = cty.StringVal(p.HostedZoneId)
}

func EncodeApigatewayv2DomainName_DomainNameConfiguration_SecurityPolicy(p DomainNameConfiguration, vals map[string]cty.Value) {
	vals["security_policy"] = cty.StringVal(p.SecurityPolicy)
}

func EncodeApigatewayv2DomainName_DomainNameConfiguration_TargetDomainName(p DomainNameConfiguration, vals map[string]cty.Value) {
	vals["target_domain_name"] = cty.StringVal(p.TargetDomainName)
}

func EncodeApigatewayv2DomainName_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2DomainName_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeApigatewayv2DomainName_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeApigatewayv2DomainName_ApiMappingSelectionExpression(p Apigatewayv2DomainNameObservation, vals map[string]cty.Value) {
	vals["api_mapping_selection_expression"] = cty.StringVal(p.ApiMappingSelectionExpression)
}

func EncodeApigatewayv2DomainName_Arn(p Apigatewayv2DomainNameObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}