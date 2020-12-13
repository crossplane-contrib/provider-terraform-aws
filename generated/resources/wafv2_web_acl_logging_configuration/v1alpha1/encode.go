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
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*Wafv2WebAclLoggingConfiguration)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Wafv2WebAclLoggingConfiguration.")
	}
	return EncodeWafv2WebAclLoggingConfiguration(*r), nil
}

func EncodeWafv2WebAclLoggingConfiguration(r Wafv2WebAclLoggingConfiguration) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafv2WebAclLoggingConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafv2WebAclLoggingConfiguration_LogDestinationConfigs(r.Spec.ForProvider, ctyVal)
	EncodeWafv2WebAclLoggingConfiguration_ResourceArn(r.Spec.ForProvider, ctyVal)
	EncodeWafv2WebAclLoggingConfiguration_RedactedFields(r.Spec.ForProvider.RedactedFields, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeWafv2WebAclLoggingConfiguration_Id(p Wafv2WebAclLoggingConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafv2WebAclLoggingConfiguration_LogDestinationConfigs(p Wafv2WebAclLoggingConfigurationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.LogDestinationConfigs {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["log_destination_configs"] = cty.SetVal(colVals)
}

func EncodeWafv2WebAclLoggingConfiguration_ResourceArn(p Wafv2WebAclLoggingConfigurationParameters, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields(p []RedactedFields, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_UriPath(v.UriPath, ctyVal)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_Body(v.Body, ctyVal)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_Method(v.Method, ctyVal)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_QueryString(v.QueryString, ctyVal)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["redacted_fields"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_UriPath(p UriPath, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)

	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_AllQueryArguments(p AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)

	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_Body(p Body, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)

	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_Method(p Method, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)

	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_QueryString(p QueryString, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)

	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleHeader(p SingleHeader, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleHeader_Name(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleHeader_Name(p SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleQueryArgument(p SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleQueryArgument_Name(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2WebAclLoggingConfiguration_RedactedFields_SingleQueryArgument_Name(p SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}