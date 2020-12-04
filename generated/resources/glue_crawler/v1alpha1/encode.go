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

func EncodeGlueCrawler(r GlueCrawler) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCrawler_Configuration(r.Spec.ForProvider, ctyVal)
	EncodeGlueCrawler_DatabaseName(r.Spec.ForProvider, ctyVal)
	EncodeGlueCrawler_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlueCrawler_SecurityConfiguration(r.Spec.ForProvider, ctyVal)
	EncodeGlueCrawler_TablePrefix(r.Spec.ForProvider, ctyVal)
	EncodeGlueCrawler_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGlueCrawler_Classifiers(r.Spec.ForProvider, ctyVal)
	EncodeGlueCrawler_Role(r.Spec.ForProvider, ctyVal)
	EncodeGlueCrawler_Schedule(r.Spec.ForProvider, ctyVal)
	EncodeGlueCrawler_Description(r.Spec.ForProvider, ctyVal)
	EncodeGlueCrawler_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlueCrawler_CatalogTarget(r.Spec.ForProvider.CatalogTarget, ctyVal)
	EncodeGlueCrawler_DynamodbTarget(r.Spec.ForProvider.DynamodbTarget, ctyVal)
	EncodeGlueCrawler_JdbcTarget(r.Spec.ForProvider.JdbcTarget, ctyVal)
	EncodeGlueCrawler_S3Target(r.Spec.ForProvider.S3Target, ctyVal)
	EncodeGlueCrawler_SchemaChangePolicy(r.Spec.ForProvider.SchemaChangePolicy, ctyVal)
	EncodeGlueCrawler_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGlueCrawler_Configuration(p GlueCrawlerParameters, vals map[string]cty.Value) {
	vals["configuration"] = cty.StringVal(p.Configuration)
}

func EncodeGlueCrawler_DatabaseName(p GlueCrawlerParameters, vals map[string]cty.Value) {
	vals["database_name"] = cty.StringVal(p.DatabaseName)
}

func EncodeGlueCrawler_Id(p GlueCrawlerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlueCrawler_SecurityConfiguration(p GlueCrawlerParameters, vals map[string]cty.Value) {
	vals["security_configuration"] = cty.StringVal(p.SecurityConfiguration)
}

func EncodeGlueCrawler_TablePrefix(p GlueCrawlerParameters, vals map[string]cty.Value) {
	vals["table_prefix"] = cty.StringVal(p.TablePrefix)
}

func EncodeGlueCrawler_Tags(p GlueCrawlerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGlueCrawler_Classifiers(p GlueCrawlerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Classifiers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["classifiers"] = cty.ListVal(colVals)
}

func EncodeGlueCrawler_Role(p GlueCrawlerParameters, vals map[string]cty.Value) {
	vals["role"] = cty.StringVal(p.Role)
}

func EncodeGlueCrawler_Schedule(p GlueCrawlerParameters, vals map[string]cty.Value) {
	vals["schedule"] = cty.StringVal(p.Schedule)
}

func EncodeGlueCrawler_Description(p GlueCrawlerParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeGlueCrawler_Name(p GlueCrawlerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueCrawler_CatalogTarget(p CatalogTarget, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCrawler_CatalogTarget_DatabaseName(p, ctyVal)
	EncodeGlueCrawler_CatalogTarget_Tables(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["catalog_target"] = cty.ListVal(valsForCollection)
}

func EncodeGlueCrawler_CatalogTarget_DatabaseName(p CatalogTarget, vals map[string]cty.Value) {
	vals["database_name"] = cty.StringVal(p.DatabaseName)
}

func EncodeGlueCrawler_CatalogTarget_Tables(p CatalogTarget, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Tables {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["tables"] = cty.ListVal(colVals)
}

func EncodeGlueCrawler_DynamodbTarget(p DynamodbTarget, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCrawler_DynamodbTarget_ScanRate(p, ctyVal)
	EncodeGlueCrawler_DynamodbTarget_Path(p, ctyVal)
	EncodeGlueCrawler_DynamodbTarget_ScanAll(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["dynamodb_target"] = cty.ListVal(valsForCollection)
}

func EncodeGlueCrawler_DynamodbTarget_ScanRate(p DynamodbTarget, vals map[string]cty.Value) {
	vals["scan_rate"] = cty.NumberIntVal(p.ScanRate)
}

func EncodeGlueCrawler_DynamodbTarget_Path(p DynamodbTarget, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeGlueCrawler_DynamodbTarget_ScanAll(p DynamodbTarget, vals map[string]cty.Value) {
	vals["scan_all"] = cty.BoolVal(p.ScanAll)
}

func EncodeGlueCrawler_JdbcTarget(p JdbcTarget, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCrawler_JdbcTarget_Exclusions(p, ctyVal)
	EncodeGlueCrawler_JdbcTarget_Path(p, ctyVal)
	EncodeGlueCrawler_JdbcTarget_ConnectionName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["jdbc_target"] = cty.ListVal(valsForCollection)
}

func EncodeGlueCrawler_JdbcTarget_Exclusions(p JdbcTarget, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Exclusions {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["exclusions"] = cty.ListVal(colVals)
}

func EncodeGlueCrawler_JdbcTarget_Path(p JdbcTarget, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeGlueCrawler_JdbcTarget_ConnectionName(p JdbcTarget, vals map[string]cty.Value) {
	vals["connection_name"] = cty.StringVal(p.ConnectionName)
}

func EncodeGlueCrawler_S3Target(p S3Target, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCrawler_S3Target_ConnectionName(p, ctyVal)
	EncodeGlueCrawler_S3Target_Exclusions(p, ctyVal)
	EncodeGlueCrawler_S3Target_Path(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["s3_target"] = cty.ListVal(valsForCollection)
}

func EncodeGlueCrawler_S3Target_ConnectionName(p S3Target, vals map[string]cty.Value) {
	vals["connection_name"] = cty.StringVal(p.ConnectionName)
}

func EncodeGlueCrawler_S3Target_Exclusions(p S3Target, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Exclusions {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["exclusions"] = cty.ListVal(colVals)
}

func EncodeGlueCrawler_S3Target_Path(p S3Target, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeGlueCrawler_SchemaChangePolicy(p SchemaChangePolicy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCrawler_SchemaChangePolicy_DeleteBehavior(p, ctyVal)
	EncodeGlueCrawler_SchemaChangePolicy_UpdateBehavior(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["schema_change_policy"] = cty.ListVal(valsForCollection)
}

func EncodeGlueCrawler_SchemaChangePolicy_DeleteBehavior(p SchemaChangePolicy, vals map[string]cty.Value) {
	vals["delete_behavior"] = cty.StringVal(p.DeleteBehavior)
}

func EncodeGlueCrawler_SchemaChangePolicy_UpdateBehavior(p SchemaChangePolicy, vals map[string]cty.Value) {
	vals["update_behavior"] = cty.StringVal(p.UpdateBehavior)
}

func EncodeGlueCrawler_Arn(p GlueCrawlerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}