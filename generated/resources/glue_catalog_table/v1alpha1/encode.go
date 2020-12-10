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
	r, ok := mr.(*GlueCatalogTable)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GlueCatalogTable.")
	}
	return EncodeGlueCatalogTable(*r), nil
}

func EncodeGlueCatalogTable(r GlueCatalogTable) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCatalogTable_CatalogId(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogTable_Description(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogTable_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogTable_Owner(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogTable_TableType(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogTable_DatabaseName(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogTable_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogTable_Parameters(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogTable_Retention(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogTable_ViewExpandedText(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogTable_ViewOriginalText(r.Spec.ForProvider, ctyVal)
	EncodeGlueCatalogTable_PartitionKeys(r.Spec.ForProvider.PartitionKeys, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor(r.Spec.ForProvider.StorageDescriptor, ctyVal)
	EncodeGlueCatalogTable_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGlueCatalogTable_CatalogId(p GlueCatalogTableParameters, vals map[string]cty.Value) {
	vals["catalog_id"] = cty.StringVal(p.CatalogId)
}

func EncodeGlueCatalogTable_Description(p GlueCatalogTableParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeGlueCatalogTable_Id(p GlueCatalogTableParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlueCatalogTable_Owner(p GlueCatalogTableParameters, vals map[string]cty.Value) {
	vals["owner"] = cty.StringVal(p.Owner)
}

func EncodeGlueCatalogTable_TableType(p GlueCatalogTableParameters, vals map[string]cty.Value) {
	vals["table_type"] = cty.StringVal(p.TableType)
}

func EncodeGlueCatalogTable_DatabaseName(p GlueCatalogTableParameters, vals map[string]cty.Value) {
	vals["database_name"] = cty.StringVal(p.DatabaseName)
}

func EncodeGlueCatalogTable_Name(p GlueCatalogTableParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueCatalogTable_Parameters(p GlueCatalogTableParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Parameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["parameters"] = cty.MapVal(mVals)
}

func EncodeGlueCatalogTable_Retention(p GlueCatalogTableParameters, vals map[string]cty.Value) {
	vals["retention"] = cty.NumberIntVal(p.Retention)
}

func EncodeGlueCatalogTable_ViewExpandedText(p GlueCatalogTableParameters, vals map[string]cty.Value) {
	vals["view_expanded_text"] = cty.StringVal(p.ViewExpandedText)
}

func EncodeGlueCatalogTable_ViewOriginalText(p GlueCatalogTableParameters, vals map[string]cty.Value) {
	vals["view_original_text"] = cty.StringVal(p.ViewOriginalText)
}

func EncodeGlueCatalogTable_PartitionKeys(p PartitionKeys, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCatalogTable_PartitionKeys_Type(p, ctyVal)
	EncodeGlueCatalogTable_PartitionKeys_Comment(p, ctyVal)
	EncodeGlueCatalogTable_PartitionKeys_Name(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["partition_keys"] = cty.ListVal(valsForCollection)
}

func EncodeGlueCatalogTable_PartitionKeys_Type(p PartitionKeys, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeGlueCatalogTable_PartitionKeys_Comment(p PartitionKeys, vals map[string]cty.Value) {
	vals["comment"] = cty.StringVal(p.Comment)
}

func EncodeGlueCatalogTable_PartitionKeys_Name(p PartitionKeys, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueCatalogTable_StorageDescriptor(p StorageDescriptor, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCatalogTable_StorageDescriptor_Parameters(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_StoredAsSubDirectories(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_BucketColumns(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_Compressed(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_InputFormat(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_Location(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_NumberOfBuckets(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_OutputFormat(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_Columns(p.Columns, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_SerDeInfo(p.SerDeInfo, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_SkewedInfo(p.SkewedInfo, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_SortColumns(p.SortColumns, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["storage_descriptor"] = cty.ListVal(valsForCollection)
}

func EncodeGlueCatalogTable_StorageDescriptor_Parameters(p StorageDescriptor, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Parameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["parameters"] = cty.MapVal(mVals)
}

func EncodeGlueCatalogTable_StorageDescriptor_StoredAsSubDirectories(p StorageDescriptor, vals map[string]cty.Value) {
	vals["stored_as_sub_directories"] = cty.BoolVal(p.StoredAsSubDirectories)
}

func EncodeGlueCatalogTable_StorageDescriptor_BucketColumns(p StorageDescriptor, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.BucketColumns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["bucket_columns"] = cty.ListVal(colVals)
}

func EncodeGlueCatalogTable_StorageDescriptor_Compressed(p StorageDescriptor, vals map[string]cty.Value) {
	vals["compressed"] = cty.BoolVal(p.Compressed)
}

func EncodeGlueCatalogTable_StorageDescriptor_InputFormat(p StorageDescriptor, vals map[string]cty.Value) {
	vals["input_format"] = cty.StringVal(p.InputFormat)
}

func EncodeGlueCatalogTable_StorageDescriptor_Location(p StorageDescriptor, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeGlueCatalogTable_StorageDescriptor_NumberOfBuckets(p StorageDescriptor, vals map[string]cty.Value) {
	vals["number_of_buckets"] = cty.NumberIntVal(p.NumberOfBuckets)
}

func EncodeGlueCatalogTable_StorageDescriptor_OutputFormat(p StorageDescriptor, vals map[string]cty.Value) {
	vals["output_format"] = cty.StringVal(p.OutputFormat)
}

func EncodeGlueCatalogTable_StorageDescriptor_Columns(p Columns, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCatalogTable_StorageDescriptor_Columns_Comment(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_Columns_Name(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_Columns_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["columns"] = cty.ListVal(valsForCollection)
}

func EncodeGlueCatalogTable_StorageDescriptor_Columns_Comment(p Columns, vals map[string]cty.Value) {
	vals["comment"] = cty.StringVal(p.Comment)
}

func EncodeGlueCatalogTable_StorageDescriptor_Columns_Name(p Columns, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueCatalogTable_StorageDescriptor_Columns_Type(p Columns, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeGlueCatalogTable_StorageDescriptor_SerDeInfo(p SerDeInfo, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCatalogTable_StorageDescriptor_SerDeInfo_Name(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_SerDeInfo_Parameters(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_SerDeInfo_SerializationLibrary(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ser_de_info"] = cty.ListVal(valsForCollection)
}

func EncodeGlueCatalogTable_StorageDescriptor_SerDeInfo_Name(p SerDeInfo, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueCatalogTable_StorageDescriptor_SerDeInfo_Parameters(p SerDeInfo, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Parameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["parameters"] = cty.MapVal(mVals)
}

func EncodeGlueCatalogTable_StorageDescriptor_SerDeInfo_SerializationLibrary(p SerDeInfo, vals map[string]cty.Value) {
	vals["serialization_library"] = cty.StringVal(p.SerializationLibrary)
}

func EncodeGlueCatalogTable_StorageDescriptor_SkewedInfo(p SkewedInfo, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCatalogTable_StorageDescriptor_SkewedInfo_SkewedColumnNames(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_SkewedInfo_SkewedColumnValueLocationMaps(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_SkewedInfo_SkewedColumnValues(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["skewed_info"] = cty.ListVal(valsForCollection)
}

func EncodeGlueCatalogTable_StorageDescriptor_SkewedInfo_SkewedColumnNames(p SkewedInfo, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SkewedColumnNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["skewed_column_names"] = cty.ListVal(colVals)
}

func EncodeGlueCatalogTable_StorageDescriptor_SkewedInfo_SkewedColumnValueLocationMaps(p SkewedInfo, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.SkewedColumnValueLocationMaps {
		mVals[key] = cty.StringVal(value)
	}
	vals["skewed_column_value_location_maps"] = cty.MapVal(mVals)
}

func EncodeGlueCatalogTable_StorageDescriptor_SkewedInfo_SkewedColumnValues(p SkewedInfo, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SkewedColumnValues {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["skewed_column_values"] = cty.ListVal(colVals)
}

func EncodeGlueCatalogTable_StorageDescriptor_SortColumns(p SortColumns, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueCatalogTable_StorageDescriptor_SortColumns_SortOrder(p, ctyVal)
	EncodeGlueCatalogTable_StorageDescriptor_SortColumns_Column(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["sort_columns"] = cty.ListVal(valsForCollection)
}

func EncodeGlueCatalogTable_StorageDescriptor_SortColumns_SortOrder(p SortColumns, vals map[string]cty.Value) {
	vals["sort_order"] = cty.NumberIntVal(p.SortOrder)
}

func EncodeGlueCatalogTable_StorageDescriptor_SortColumns_Column(p SortColumns, vals map[string]cty.Value) {
	vals["column"] = cty.StringVal(p.Column)
}

func EncodeGlueCatalogTable_Arn(p GlueCatalogTableObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}