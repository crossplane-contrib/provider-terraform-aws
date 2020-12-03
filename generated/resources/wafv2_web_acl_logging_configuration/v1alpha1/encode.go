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

package v1alpha1func EncodeWafv2WebAclLoggingConfiguration(r Wafv2WebAclLoggingConfiguration) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeWafv2WebAclLoggingConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafv2WebAclLoggingConfiguration_LogDestinationConfigs(r.Spec.ForProvider, ctyVal)
	EncodeWafv2WebAclLoggingConfiguration_ResourceArn(r.Spec.ForProvider, ctyVal)
	EncodeWafv2WebAclLoggingConfiguration_RedactedFields(r.Spec.ForProvider.RedactedFields, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeWafv2WebAclLoggingConfiguration_Id(p *Wafv2WebAclLoggingConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafv2WebAclLoggingConfiguration_LogDestinationConfigs(p *Wafv2WebAclLoggingConfigurationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.LogDestinationConfigs {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["log_destination_configs"] = cty.SetVal(colVals)
}

func EncodeWafv2WebAclLoggingConfiguration_ResourceArn(p *Wafv2WebAclLoggingConfigurationParameters, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields(p *RedactedFields, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RedactedFields {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_UriPath(v.UriPath, ctyVal)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_Body(v.Body, ctyVal)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_Method(v.Method, ctyVal)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_QueryString(v.QueryString, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["redacted_fields"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}