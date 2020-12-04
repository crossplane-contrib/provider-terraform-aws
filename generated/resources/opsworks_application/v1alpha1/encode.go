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

func EncodeOpsworksApplication(r OpsworksApplication) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksApplication_DataSourceType(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_Description(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_StackId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_AutoBundleOnDeploy(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_AwsFlowRubySettings(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_DataSourceDatabaseName(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_ShortName(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_DataSourceArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_DocumentRoot(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_EnableSsl(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_Type(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_Domains(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_Name(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_RailsEnv(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksApplication_AppSource(r.Spec.ForProvider.AppSource, ctyVal)
	EncodeOpsworksApplication_Environment(r.Spec.ForProvider.Environment, ctyVal)
	EncodeOpsworksApplication_SslConfiguration(r.Spec.ForProvider.SslConfiguration, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksApplication_DataSourceType(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["data_source_type"] = cty.StringVal(p.DataSourceType)
}

func EncodeOpsworksApplication_Description(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeOpsworksApplication_StackId(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}

func EncodeOpsworksApplication_AutoBundleOnDeploy(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["auto_bundle_on_deploy"] = cty.StringVal(p.AutoBundleOnDeploy)
}

func EncodeOpsworksApplication_AwsFlowRubySettings(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["aws_flow_ruby_settings"] = cty.StringVal(p.AwsFlowRubySettings)
}

func EncodeOpsworksApplication_DataSourceDatabaseName(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["data_source_database_name"] = cty.StringVal(p.DataSourceDatabaseName)
}

func EncodeOpsworksApplication_ShortName(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["short_name"] = cty.StringVal(p.ShortName)
}

func EncodeOpsworksApplication_DataSourceArn(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["data_source_arn"] = cty.StringVal(p.DataSourceArn)
}

func EncodeOpsworksApplication_DocumentRoot(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["document_root"] = cty.StringVal(p.DocumentRoot)
}

func EncodeOpsworksApplication_EnableSsl(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["enable_ssl"] = cty.BoolVal(p.EnableSsl)
}

func EncodeOpsworksApplication_Type(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOpsworksApplication_Domains(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Domains {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["domains"] = cty.ListVal(colVals)
}

func EncodeOpsworksApplication_Id(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksApplication_Name(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOpsworksApplication_RailsEnv(p OpsworksApplicationParameters, vals map[string]cty.Value) {
	vals["rails_env"] = cty.StringVal(p.RailsEnv)
}

func EncodeOpsworksApplication_AppSource(p AppSource, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksApplication_AppSource_Type(p, ctyVal)
	EncodeOpsworksApplication_AppSource_Url(p, ctyVal)
	EncodeOpsworksApplication_AppSource_Username(p, ctyVal)
	EncodeOpsworksApplication_AppSource_Password(p, ctyVal)
	EncodeOpsworksApplication_AppSource_Revision(p, ctyVal)
	EncodeOpsworksApplication_AppSource_SshKey(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["app_source"] = cty.ListVal(valsForCollection)
}

func EncodeOpsworksApplication_AppSource_Type(p AppSource, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOpsworksApplication_AppSource_Url(p AppSource, vals map[string]cty.Value) {
	vals["url"] = cty.StringVal(p.Url)
}

func EncodeOpsworksApplication_AppSource_Username(p AppSource, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeOpsworksApplication_AppSource_Password(p AppSource, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeOpsworksApplication_AppSource_Revision(p AppSource, vals map[string]cty.Value) {
	vals["revision"] = cty.StringVal(p.Revision)
}

func EncodeOpsworksApplication_AppSource_SshKey(p AppSource, vals map[string]cty.Value) {
	vals["ssh_key"] = cty.StringVal(p.SshKey)
}

func EncodeOpsworksApplication_Environment(p Environment, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksApplication_Environment_Key(p, ctyVal)
	EncodeOpsworksApplication_Environment_Secure(p, ctyVal)
	EncodeOpsworksApplication_Environment_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["environment"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksApplication_Environment_Key(p Environment, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeOpsworksApplication_Environment_Secure(p Environment, vals map[string]cty.Value) {
	vals["secure"] = cty.BoolVal(p.Secure)
}

func EncodeOpsworksApplication_Environment_Value(p Environment, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeOpsworksApplication_SslConfiguration(p SslConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksApplication_SslConfiguration_Chain(p, ctyVal)
	EncodeOpsworksApplication_SslConfiguration_PrivateKey(p, ctyVal)
	EncodeOpsworksApplication_SslConfiguration_Certificate(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ssl_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeOpsworksApplication_SslConfiguration_Chain(p SslConfiguration, vals map[string]cty.Value) {
	vals["chain"] = cty.StringVal(p.Chain)
}

func EncodeOpsworksApplication_SslConfiguration_PrivateKey(p SslConfiguration, vals map[string]cty.Value) {
	vals["private_key"] = cty.StringVal(p.PrivateKey)
}

func EncodeOpsworksApplication_SslConfiguration_Certificate(p SslConfiguration, vals map[string]cty.Value) {
	vals["certificate"] = cty.StringVal(p.Certificate)
}