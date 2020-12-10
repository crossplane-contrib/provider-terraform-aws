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
	r, ok := mr.(*GlueMlTransform)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GlueMlTransform.")
	}
	return EncodeGlueMlTransform(*r), nil
}

func EncodeGlueMlTransform(r GlueMlTransform) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlueMlTransform_Description(r.Spec.ForProvider, ctyVal)
	EncodeGlueMlTransform_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlueMlTransform_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeGlueMlTransform_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlueMlTransform_NumberOfWorkers(r.Spec.ForProvider, ctyVal)
	EncodeGlueMlTransform_WorkerType(r.Spec.ForProvider, ctyVal)
	EncodeGlueMlTransform_GlueVersion(r.Spec.ForProvider, ctyVal)
	EncodeGlueMlTransform_MaxCapacity(r.Spec.ForProvider, ctyVal)
	EncodeGlueMlTransform_MaxRetries(r.Spec.ForProvider, ctyVal)
	EncodeGlueMlTransform_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGlueMlTransform_Timeout(r.Spec.ForProvider, ctyVal)
	EncodeGlueMlTransform_InputRecordTables(r.Spec.ForProvider.InputRecordTables, ctyVal)
	EncodeGlueMlTransform_Parameters(r.Spec.ForProvider.Parameters, ctyVal)
	EncodeGlueMlTransform_Arn(r.Status.AtProvider, ctyVal)
	EncodeGlueMlTransform_LabelCount(r.Status.AtProvider, ctyVal)
	EncodeGlueMlTransform_Schema(r.Status.AtProvider.Schema, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGlueMlTransform_Description(p GlueMlTransformParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeGlueMlTransform_Id(p GlueMlTransformParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlueMlTransform_RoleArn(p GlueMlTransformParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeGlueMlTransform_Name(p GlueMlTransformParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueMlTransform_NumberOfWorkers(p GlueMlTransformParameters, vals map[string]cty.Value) {
	vals["number_of_workers"] = cty.NumberIntVal(p.NumberOfWorkers)
}

func EncodeGlueMlTransform_WorkerType(p GlueMlTransformParameters, vals map[string]cty.Value) {
	vals["worker_type"] = cty.StringVal(p.WorkerType)
}

func EncodeGlueMlTransform_GlueVersion(p GlueMlTransformParameters, vals map[string]cty.Value) {
	vals["glue_version"] = cty.StringVal(p.GlueVersion)
}

func EncodeGlueMlTransform_MaxCapacity(p GlueMlTransformParameters, vals map[string]cty.Value) {
	vals["max_capacity"] = cty.NumberIntVal(p.MaxCapacity)
}

func EncodeGlueMlTransform_MaxRetries(p GlueMlTransformParameters, vals map[string]cty.Value) {
	vals["max_retries"] = cty.NumberIntVal(p.MaxRetries)
}

func EncodeGlueMlTransform_Tags(p GlueMlTransformParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGlueMlTransform_Timeout(p GlueMlTransformParameters, vals map[string]cty.Value) {
	vals["timeout"] = cty.NumberIntVal(p.Timeout)
}

func EncodeGlueMlTransform_InputRecordTables(p []InputRecordTables, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeGlueMlTransform_InputRecordTables_CatalogId(v, ctyVal)
		EncodeGlueMlTransform_InputRecordTables_ConnectionName(v, ctyVal)
		EncodeGlueMlTransform_InputRecordTables_DatabaseName(v, ctyVal)
		EncodeGlueMlTransform_InputRecordTables_TableName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["input_record_tables"] = cty.ListVal(valsForCollection)
}

func EncodeGlueMlTransform_InputRecordTables_CatalogId(p InputRecordTables, vals map[string]cty.Value) {
	vals["catalog_id"] = cty.StringVal(p.CatalogId)
}

func EncodeGlueMlTransform_InputRecordTables_ConnectionName(p InputRecordTables, vals map[string]cty.Value) {
	vals["connection_name"] = cty.StringVal(p.ConnectionName)
}

func EncodeGlueMlTransform_InputRecordTables_DatabaseName(p InputRecordTables, vals map[string]cty.Value) {
	vals["database_name"] = cty.StringVal(p.DatabaseName)
}

func EncodeGlueMlTransform_InputRecordTables_TableName(p InputRecordTables, vals map[string]cty.Value) {
	vals["table_name"] = cty.StringVal(p.TableName)
}

func EncodeGlueMlTransform_Parameters(p Parameters, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueMlTransform_Parameters_TransformType(p, ctyVal)
	EncodeGlueMlTransform_Parameters_FindMatchesParameters(p.FindMatchesParameters, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["parameters"] = cty.ListVal(valsForCollection)
}

func EncodeGlueMlTransform_Parameters_TransformType(p Parameters, vals map[string]cty.Value) {
	vals["transform_type"] = cty.StringVal(p.TransformType)
}

func EncodeGlueMlTransform_Parameters_FindMatchesParameters(p FindMatchesParameters, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueMlTransform_Parameters_FindMatchesParameters_AccuracyCostTradeOff(p, ctyVal)
	EncodeGlueMlTransform_Parameters_FindMatchesParameters_EnforceProvidedLabels(p, ctyVal)
	EncodeGlueMlTransform_Parameters_FindMatchesParameters_PrecisionRecallTradeOff(p, ctyVal)
	EncodeGlueMlTransform_Parameters_FindMatchesParameters_PrimaryKeyColumnName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["find_matches_parameters"] = cty.ListVal(valsForCollection)
}

func EncodeGlueMlTransform_Parameters_FindMatchesParameters_AccuracyCostTradeOff(p FindMatchesParameters, vals map[string]cty.Value) {
	vals["accuracy_cost_trade_off"] = cty.NumberIntVal(p.AccuracyCostTradeOff)
}

func EncodeGlueMlTransform_Parameters_FindMatchesParameters_EnforceProvidedLabels(p FindMatchesParameters, vals map[string]cty.Value) {
	vals["enforce_provided_labels"] = cty.BoolVal(p.EnforceProvidedLabels)
}

func EncodeGlueMlTransform_Parameters_FindMatchesParameters_PrecisionRecallTradeOff(p FindMatchesParameters, vals map[string]cty.Value) {
	vals["precision_recall_trade_off"] = cty.NumberIntVal(p.PrecisionRecallTradeOff)
}

func EncodeGlueMlTransform_Parameters_FindMatchesParameters_PrimaryKeyColumnName(p FindMatchesParameters, vals map[string]cty.Value) {
	vals["primary_key_column_name"] = cty.StringVal(p.PrimaryKeyColumnName)
}

func EncodeGlueMlTransform_Arn(p GlueMlTransformObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeGlueMlTransform_LabelCount(p GlueMlTransformObservation, vals map[string]cty.Value) {
	vals["label_count"] = cty.NumberIntVal(p.LabelCount)
}

func EncodeGlueMlTransform_Schema(p []Schema, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeGlueMlTransform_Schema_DataType(v, ctyVal)
		EncodeGlueMlTransform_Schema_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["schema"] = cty.ListVal(valsForCollection)
}

func EncodeGlueMlTransform_Schema_DataType(p Schema, vals map[string]cty.Value) {
	vals["data_type"] = cty.StringVal(p.DataType)
}

func EncodeGlueMlTransform_Schema_Name(p Schema, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}