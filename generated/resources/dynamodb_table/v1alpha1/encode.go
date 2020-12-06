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

func EncodeDynamodbTable(r DynamodbTable) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDynamodbTable_HashKey(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTable_Id(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTable_RangeKey(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTable_StreamEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTable_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTable_BillingMode(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTable_Name(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTable_ReadCapacity(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTable_StreamViewType(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTable_WriteCapacity(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTable_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDynamodbTable_Ttl(r.Spec.ForProvider.Ttl, ctyVal)
	EncodeDynamodbTable_Attribute(r.Spec.ForProvider.Attribute, ctyVal)
	EncodeDynamodbTable_GlobalSecondaryIndex(r.Spec.ForProvider.GlobalSecondaryIndex, ctyVal)
	EncodeDynamodbTable_LocalSecondaryIndex(r.Spec.ForProvider.LocalSecondaryIndex, ctyVal)
	EncodeDynamodbTable_PointInTimeRecovery(r.Spec.ForProvider.PointInTimeRecovery, ctyVal)
	EncodeDynamodbTable_Replica(r.Spec.ForProvider.Replica, ctyVal)
	EncodeDynamodbTable_ServerSideEncryption(r.Spec.ForProvider.ServerSideEncryption, ctyVal)
	EncodeDynamodbTable_StreamArn(r.Status.AtProvider, ctyVal)
	EncodeDynamodbTable_Arn(r.Status.AtProvider, ctyVal)
	EncodeDynamodbTable_StreamLabel(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDynamodbTable_HashKey(p DynamodbTableParameters, vals map[string]cty.Value) {
	vals["hash_key"] = cty.StringVal(p.HashKey)
}

func EncodeDynamodbTable_Id(p DynamodbTableParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDynamodbTable_RangeKey(p DynamodbTableParameters, vals map[string]cty.Value) {
	vals["range_key"] = cty.StringVal(p.RangeKey)
}

func EncodeDynamodbTable_StreamEnabled(p DynamodbTableParameters, vals map[string]cty.Value) {
	vals["stream_enabled"] = cty.BoolVal(p.StreamEnabled)
}

func EncodeDynamodbTable_Tags(p DynamodbTableParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDynamodbTable_BillingMode(p DynamodbTableParameters, vals map[string]cty.Value) {
	vals["billing_mode"] = cty.StringVal(p.BillingMode)
}

func EncodeDynamodbTable_Name(p DynamodbTableParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDynamodbTable_ReadCapacity(p DynamodbTableParameters, vals map[string]cty.Value) {
	vals["read_capacity"] = cty.NumberIntVal(p.ReadCapacity)
}

func EncodeDynamodbTable_StreamViewType(p DynamodbTableParameters, vals map[string]cty.Value) {
	vals["stream_view_type"] = cty.StringVal(p.StreamViewType)
}

func EncodeDynamodbTable_WriteCapacity(p DynamodbTableParameters, vals map[string]cty.Value) {
	vals["write_capacity"] = cty.NumberIntVal(p.WriteCapacity)
}

func EncodeDynamodbTable_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDynamodbTable_Timeouts_Update(p, ctyVal)
	EncodeDynamodbTable_Timeouts_Create(p, ctyVal)
	EncodeDynamodbTable_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDynamodbTable_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDynamodbTable_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDynamodbTable_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDynamodbTable_Ttl(p Ttl, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDynamodbTable_Ttl_AttributeName(p, ctyVal)
	EncodeDynamodbTable_Ttl_Enabled(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ttl"] = cty.ListVal(valsForCollection)
}

func EncodeDynamodbTable_Ttl_AttributeName(p Ttl, vals map[string]cty.Value) {
	vals["attribute_name"] = cty.StringVal(p.AttributeName)
}

func EncodeDynamodbTable_Ttl_Enabled(p Ttl, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeDynamodbTable_Attribute(p []Attribute, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeDynamodbTable_Attribute_Name(v, ctyVal)
		EncodeDynamodbTable_Attribute_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["attribute"] = cty.SetVal(valsForCollection)
}

func EncodeDynamodbTable_Attribute_Name(p Attribute, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDynamodbTable_Attribute_Type(p Attribute, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeDynamodbTable_GlobalSecondaryIndex(p GlobalSecondaryIndex, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDynamodbTable_GlobalSecondaryIndex_NonKeyAttributes(p, ctyVal)
	EncodeDynamodbTable_GlobalSecondaryIndex_ProjectionType(p, ctyVal)
	EncodeDynamodbTable_GlobalSecondaryIndex_RangeKey(p, ctyVal)
	EncodeDynamodbTable_GlobalSecondaryIndex_ReadCapacity(p, ctyVal)
	EncodeDynamodbTable_GlobalSecondaryIndex_WriteCapacity(p, ctyVal)
	EncodeDynamodbTable_GlobalSecondaryIndex_HashKey(p, ctyVal)
	EncodeDynamodbTable_GlobalSecondaryIndex_Name(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["global_secondary_index"] = cty.SetVal(valsForCollection)
}

func EncodeDynamodbTable_GlobalSecondaryIndex_NonKeyAttributes(p GlobalSecondaryIndex, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NonKeyAttributes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["non_key_attributes"] = cty.SetVal(colVals)
}

func EncodeDynamodbTable_GlobalSecondaryIndex_ProjectionType(p GlobalSecondaryIndex, vals map[string]cty.Value) {
	vals["projection_type"] = cty.StringVal(p.ProjectionType)
}

func EncodeDynamodbTable_GlobalSecondaryIndex_RangeKey(p GlobalSecondaryIndex, vals map[string]cty.Value) {
	vals["range_key"] = cty.StringVal(p.RangeKey)
}

func EncodeDynamodbTable_GlobalSecondaryIndex_ReadCapacity(p GlobalSecondaryIndex, vals map[string]cty.Value) {
	vals["read_capacity"] = cty.NumberIntVal(p.ReadCapacity)
}

func EncodeDynamodbTable_GlobalSecondaryIndex_WriteCapacity(p GlobalSecondaryIndex, vals map[string]cty.Value) {
	vals["write_capacity"] = cty.NumberIntVal(p.WriteCapacity)
}

func EncodeDynamodbTable_GlobalSecondaryIndex_HashKey(p GlobalSecondaryIndex, vals map[string]cty.Value) {
	vals["hash_key"] = cty.StringVal(p.HashKey)
}

func EncodeDynamodbTable_GlobalSecondaryIndex_Name(p GlobalSecondaryIndex, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDynamodbTable_LocalSecondaryIndex(p LocalSecondaryIndex, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDynamodbTable_LocalSecondaryIndex_Name(p, ctyVal)
	EncodeDynamodbTable_LocalSecondaryIndex_NonKeyAttributes(p, ctyVal)
	EncodeDynamodbTable_LocalSecondaryIndex_ProjectionType(p, ctyVal)
	EncodeDynamodbTable_LocalSecondaryIndex_RangeKey(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["local_secondary_index"] = cty.SetVal(valsForCollection)
}

func EncodeDynamodbTable_LocalSecondaryIndex_Name(p LocalSecondaryIndex, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDynamodbTable_LocalSecondaryIndex_NonKeyAttributes(p LocalSecondaryIndex, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NonKeyAttributes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["non_key_attributes"] = cty.ListVal(colVals)
}

func EncodeDynamodbTable_LocalSecondaryIndex_ProjectionType(p LocalSecondaryIndex, vals map[string]cty.Value) {
	vals["projection_type"] = cty.StringVal(p.ProjectionType)
}

func EncodeDynamodbTable_LocalSecondaryIndex_RangeKey(p LocalSecondaryIndex, vals map[string]cty.Value) {
	vals["range_key"] = cty.StringVal(p.RangeKey)
}

func EncodeDynamodbTable_PointInTimeRecovery(p PointInTimeRecovery, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDynamodbTable_PointInTimeRecovery_Enabled(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["point_in_time_recovery"] = cty.ListVal(valsForCollection)
}

func EncodeDynamodbTable_PointInTimeRecovery_Enabled(p PointInTimeRecovery, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeDynamodbTable_Replica(p Replica, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDynamodbTable_Replica_RegionName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["replica"] = cty.SetVal(valsForCollection)
}

func EncodeDynamodbTable_Replica_RegionName(p Replica, vals map[string]cty.Value) {
	vals["region_name"] = cty.StringVal(p.RegionName)
}

func EncodeDynamodbTable_ServerSideEncryption(p ServerSideEncryption, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDynamodbTable_ServerSideEncryption_Enabled(p, ctyVal)
	EncodeDynamodbTable_ServerSideEncryption_KmsKeyArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["server_side_encryption"] = cty.ListVal(valsForCollection)
}

func EncodeDynamodbTable_ServerSideEncryption_Enabled(p ServerSideEncryption, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeDynamodbTable_ServerSideEncryption_KmsKeyArn(p ServerSideEncryption, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeDynamodbTable_StreamArn(p DynamodbTableObservation, vals map[string]cty.Value) {
	vals["stream_arn"] = cty.StringVal(p.StreamArn)
}

func EncodeDynamodbTable_Arn(p DynamodbTableObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDynamodbTable_StreamLabel(p DynamodbTableObservation, vals map[string]cty.Value) {
	vals["stream_label"] = cty.StringVal(p.StreamLabel)
}