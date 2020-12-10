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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*ApiGatewayRestApi)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ApiGatewayRestApi.")
	}
	return EncodeApiGatewayRestApi(*r), nil
}

func EncodeApiGatewayRestApi(r ApiGatewayRestApi) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayRestApi_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayRestApi_MinimumCompressionSize(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayRestApi_Policy(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayRestApi_Body(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayRestApi_Name(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayRestApi_Tags(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayRestApi_ApiKeySource(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayRestApi_BinaryMediaTypes(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayRestApi_Description(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayRestApi_EndpointConfiguration(r.Spec.ForProvider.EndpointConfiguration, ctyVal)
	EncodeApiGatewayRestApi_ExecutionArn(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayRestApi_RootResourceId(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayRestApi_Arn(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayRestApi_CreatedDate(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayRestApi_Id(p ApiGatewayRestApiParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayRestApi_MinimumCompressionSize(p ApiGatewayRestApiParameters, vals map[string]cty.Value) {
	vals["minimum_compression_size"] = cty.NumberIntVal(p.MinimumCompressionSize)
}

func EncodeApiGatewayRestApi_Policy(p ApiGatewayRestApiParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeApiGatewayRestApi_Body(p ApiGatewayRestApiParameters, vals map[string]cty.Value) {
	vals["body"] = cty.StringVal(p.Body)
}

func EncodeApiGatewayRestApi_Name(p ApiGatewayRestApiParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeApiGatewayRestApi_Tags(p ApiGatewayRestApiParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeApiGatewayRestApi_ApiKeySource(p ApiGatewayRestApiParameters, vals map[string]cty.Value) {
	vals["api_key_source"] = cty.StringVal(p.ApiKeySource)
}

func EncodeApiGatewayRestApi_BinaryMediaTypes(p ApiGatewayRestApiParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.BinaryMediaTypes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["binary_media_types"] = cty.ListVal(colVals)
}

func EncodeApiGatewayRestApi_Description(p ApiGatewayRestApiParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeApiGatewayRestApi_EndpointConfiguration(p EndpointConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayRestApi_EndpointConfiguration_VpcEndpointIds(p, ctyVal)
	EncodeApiGatewayRestApi_EndpointConfiguration_Types(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["endpoint_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeApiGatewayRestApi_EndpointConfiguration_VpcEndpointIds(p EndpointConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcEndpointIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_endpoint_ids"] = cty.SetVal(colVals)
}

func EncodeApiGatewayRestApi_EndpointConfiguration_Types(p EndpointConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Types {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["types"] = cty.ListVal(colVals)
}

func EncodeApiGatewayRestApi_ExecutionArn(p ApiGatewayRestApiObservation, vals map[string]cty.Value) {
	vals["execution_arn"] = cty.StringVal(p.ExecutionArn)
}

func EncodeApiGatewayRestApi_RootResourceId(p ApiGatewayRestApiObservation, vals map[string]cty.Value) {
	vals["root_resource_id"] = cty.StringVal(p.RootResourceId)
}

func EncodeApiGatewayRestApi_Arn(p ApiGatewayRestApiObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeApiGatewayRestApi_CreatedDate(p ApiGatewayRestApiObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}