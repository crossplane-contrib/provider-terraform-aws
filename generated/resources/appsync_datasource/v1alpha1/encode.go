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
	r, ok := mr.(*AppsyncDatasource)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a AppsyncDatasource.")
	}
	return EncodeAppsyncDatasource(*r), nil
}

func EncodeAppsyncDatasource(r AppsyncDatasource) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncDatasource_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncDatasource_Name(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncDatasource_ServiceRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncDatasource_Type(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncDatasource_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncDatasource_Description(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncDatasource_DynamodbConfig(r.Spec.ForProvider.DynamodbConfig, ctyVal)
	EncodeAppsyncDatasource_ElasticsearchConfig(r.Spec.ForProvider.ElasticsearchConfig, ctyVal)
	EncodeAppsyncDatasource_HttpConfig(r.Spec.ForProvider.HttpConfig, ctyVal)
	EncodeAppsyncDatasource_LambdaConfig(r.Spec.ForProvider.LambdaConfig, ctyVal)
	EncodeAppsyncDatasource_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeAppsyncDatasource_Id(p AppsyncDatasourceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppsyncDatasource_Name(p AppsyncDatasourceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppsyncDatasource_ServiceRoleArn(p AppsyncDatasourceParameters, vals map[string]cty.Value) {
	vals["service_role_arn"] = cty.StringVal(p.ServiceRoleArn)
}

func EncodeAppsyncDatasource_Type(p AppsyncDatasourceParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeAppsyncDatasource_ApiId(p AppsyncDatasourceParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeAppsyncDatasource_Description(p AppsyncDatasourceParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeAppsyncDatasource_DynamodbConfig(p DynamodbConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncDatasource_DynamodbConfig_Region(p, ctyVal)
	EncodeAppsyncDatasource_DynamodbConfig_TableName(p, ctyVal)
	EncodeAppsyncDatasource_DynamodbConfig_UseCallerCredentials(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["dynamodb_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncDatasource_DynamodbConfig_Region(p DynamodbConfig, vals map[string]cty.Value) {
	vals["region"] = cty.StringVal(p.Region)
}

func EncodeAppsyncDatasource_DynamodbConfig_TableName(p DynamodbConfig, vals map[string]cty.Value) {
	vals["table_name"] = cty.StringVal(p.TableName)
}

func EncodeAppsyncDatasource_DynamodbConfig_UseCallerCredentials(p DynamodbConfig, vals map[string]cty.Value) {
	vals["use_caller_credentials"] = cty.BoolVal(p.UseCallerCredentials)
}

func EncodeAppsyncDatasource_ElasticsearchConfig(p ElasticsearchConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncDatasource_ElasticsearchConfig_Endpoint(p, ctyVal)
	EncodeAppsyncDatasource_ElasticsearchConfig_Region(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["elasticsearch_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncDatasource_ElasticsearchConfig_Endpoint(p ElasticsearchConfig, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeAppsyncDatasource_ElasticsearchConfig_Region(p ElasticsearchConfig, vals map[string]cty.Value) {
	vals["region"] = cty.StringVal(p.Region)
}

func EncodeAppsyncDatasource_HttpConfig(p HttpConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncDatasource_HttpConfig_Endpoint(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["http_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncDatasource_HttpConfig_Endpoint(p HttpConfig, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeAppsyncDatasource_LambdaConfig(p LambdaConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncDatasource_LambdaConfig_FunctionArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["lambda_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncDatasource_LambdaConfig_FunctionArn(p LambdaConfig, vals map[string]cty.Value) {
	vals["function_arn"] = cty.StringVal(p.FunctionArn)
}

func EncodeAppsyncDatasource_Arn(p AppsyncDatasourceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}