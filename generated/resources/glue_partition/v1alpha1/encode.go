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

func EncodeGluePartition(r GluePartition) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGluePartition_TableName(r.Spec.ForProvider, ctyVal)
	EncodeGluePartition_Parameters(r.Spec.ForProvider, ctyVal)
	EncodeGluePartition_PartitionValues(r.Spec.ForProvider, ctyVal)
	EncodeGluePartition_CatalogId(r.Spec.ForProvider, ctyVal)
	EncodeGluePartition_DatabaseName(r.Spec.ForProvider, ctyVal)
	EncodeGluePartition_Id(r.Spec.ForProvider, ctyVal)
	EncodeGluePartition_StorageDescriptor(r.Spec.ForProvider.StorageDescriptor, ctyVal)
	EncodeGluePartition_CreationTime(r.Status.AtProvider, ctyVal)
	EncodeGluePartition_LastAccessedTime(r.Status.AtProvider, ctyVal)
	EncodeGluePartition_LastAnalyzedTime(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGluePartition_TableName(p GluePartitionParameters, vals map[string]cty.Value) {
	vals["table_name"] = cty.StringVal(p.TableName)
}

func EncodeGluePartition_Parameters(p GluePartitionParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Parameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["parameters"] = cty.MapVal(mVals)
}

func EncodeGluePartition_PartitionValues(p GluePartitionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.PartitionValues {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["partition_values"] = cty.SetVal(colVals)
}

func EncodeGluePartition_CatalogId(p GluePartitionParameters, vals map[string]cty.Value) {
	vals["catalog_id"] = cty.StringVal(p.CatalogId)
}

func EncodeGluePartition_DatabaseName(p GluePartitionParameters, vals map[string]cty.Value) {
	vals["database_name"] = cty.StringVal(p.DatabaseName)
}

func EncodeGluePartition_Id(p GluePartitionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGluePartition_StorageDescriptor(p StorageDescriptor, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGluePartition_StorageDescriptor_InputFormat(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_Location(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_NumberOfBuckets(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_OutputFormat(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_Parameters(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_StoredAsSubDirectories(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_BucketColumns(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_Compressed(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_SerDeInfo(p.SerDeInfo, ctyVal)
	EncodeGluePartition_StorageDescriptor_SkewedInfo(p.SkewedInfo, ctyVal)
	EncodeGluePartition_StorageDescriptor_SortColumns(p.SortColumns, ctyVal)
	EncodeGluePartition_StorageDescriptor_Columns(p.Columns, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["storage_descriptor"] = cty.ListVal(valsForCollection)
}

func EncodeGluePartition_StorageDescriptor_InputFormat(p StorageDescriptor, vals map[string]cty.Value) {
	vals["input_format"] = cty.StringVal(p.InputFormat)
}

func EncodeGluePartition_StorageDescriptor_Location(p StorageDescriptor, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeGluePartition_StorageDescriptor_NumberOfBuckets(p StorageDescriptor, vals map[string]cty.Value) {
	vals["number_of_buckets"] = cty.NumberIntVal(p.NumberOfBuckets)
}

func EncodeGluePartition_StorageDescriptor_OutputFormat(p StorageDescriptor, vals map[string]cty.Value) {
	vals["output_format"] = cty.StringVal(p.OutputFormat)
}

func EncodeGluePartition_StorageDescriptor_Parameters(p StorageDescriptor, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Parameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["parameters"] = cty.MapVal(mVals)
}

func EncodeGluePartition_StorageDescriptor_StoredAsSubDirectories(p StorageDescriptor, vals map[string]cty.Value) {
	vals["stored_as_sub_directories"] = cty.BoolVal(p.StoredAsSubDirectories)
}

func EncodeGluePartition_StorageDescriptor_BucketColumns(p StorageDescriptor, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.BucketColumns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["bucket_columns"] = cty.ListVal(colVals)
}

func EncodeGluePartition_StorageDescriptor_Compressed(p StorageDescriptor, vals map[string]cty.Value) {
	vals["compressed"] = cty.BoolVal(p.Compressed)
}

func EncodeGluePartition_StorageDescriptor_SerDeInfo(p SerDeInfo, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGluePartition_StorageDescriptor_SerDeInfo_Name(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_SerDeInfo_Parameters(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_SerDeInfo_SerializationLibrary(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ser_de_info"] = cty.ListVal(valsForCollection)
}

func EncodeGluePartition_StorageDescriptor_SerDeInfo_Name(p SerDeInfo, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGluePartition_StorageDescriptor_SerDeInfo_Parameters(p SerDeInfo, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Parameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["parameters"] = cty.MapVal(mVals)
}

func EncodeGluePartition_StorageDescriptor_SerDeInfo_SerializationLibrary(p SerDeInfo, vals map[string]cty.Value) {
	vals["serialization_library"] = cty.StringVal(p.SerializationLibrary)
}

func EncodeGluePartition_StorageDescriptor_SkewedInfo(p SkewedInfo, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGluePartition_StorageDescriptor_SkewedInfo_SkewedColumnNames(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_SkewedInfo_SkewedColumnValueLocationMaps(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_SkewedInfo_SkewedColumnValues(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["skewed_info"] = cty.ListVal(valsForCollection)
}

func EncodeGluePartition_StorageDescriptor_SkewedInfo_SkewedColumnNames(p SkewedInfo, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SkewedColumnNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["skewed_column_names"] = cty.ListVal(colVals)
}

func EncodeGluePartition_StorageDescriptor_SkewedInfo_SkewedColumnValueLocationMaps(p SkewedInfo, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.SkewedColumnValueLocationMaps {
		mVals[key] = cty.StringVal(value)
	}
	vals["skewed_column_value_location_maps"] = cty.MapVal(mVals)
}

func EncodeGluePartition_StorageDescriptor_SkewedInfo_SkewedColumnValues(p SkewedInfo, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SkewedColumnValues {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["skewed_column_values"] = cty.ListVal(colVals)
}

func EncodeGluePartition_StorageDescriptor_SortColumns(p SortColumns, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGluePartition_StorageDescriptor_SortColumns_Column(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_SortColumns_SortOrder(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["sort_columns"] = cty.ListVal(valsForCollection)
}

func EncodeGluePartition_StorageDescriptor_SortColumns_Column(p SortColumns, vals map[string]cty.Value) {
	vals["column"] = cty.StringVal(p.Column)
}

func EncodeGluePartition_StorageDescriptor_SortColumns_SortOrder(p SortColumns, vals map[string]cty.Value) {
	vals["sort_order"] = cty.NumberIntVal(p.SortOrder)
}

func EncodeGluePartition_StorageDescriptor_Columns(p Columns, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGluePartition_StorageDescriptor_Columns_Comment(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_Columns_Name(p, ctyVal)
	EncodeGluePartition_StorageDescriptor_Columns_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["columns"] = cty.ListVal(valsForCollection)
}

func EncodeGluePartition_StorageDescriptor_Columns_Comment(p Columns, vals map[string]cty.Value) {
	vals["comment"] = cty.StringVal(p.Comment)
}

func EncodeGluePartition_StorageDescriptor_Columns_Name(p Columns, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGluePartition_StorageDescriptor_Columns_Type(p Columns, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeGluePartition_CreationTime(p GluePartitionObservation, vals map[string]cty.Value) {
	vals["creation_time"] = cty.StringVal(p.CreationTime)
}

func EncodeGluePartition_LastAccessedTime(p GluePartitionObservation, vals map[string]cty.Value) {
	vals["last_accessed_time"] = cty.StringVal(p.LastAccessedTime)
}

func EncodeGluePartition_LastAnalyzedTime(p GluePartitionObservation, vals map[string]cty.Value) {
	vals["last_analyzed_time"] = cty.StringVal(p.LastAnalyzedTime)
}