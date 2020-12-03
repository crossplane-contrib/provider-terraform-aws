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

package v1alpha1func EncodeDynamodbTableItem(r DynamodbTableItem) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeDynamodbTableItem_TableName(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTableItem_HashKey(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTableItem_Id(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTableItem_Item(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTableItem_RangeKey(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeDynamodbTableItem_TableName(p *DynamodbTableItemParameters, vals map[string]cty.Value) {
	vals["table_name"] = cty.StringVal(p.TableName)
}

func EncodeDynamodbTableItem_HashKey(p *DynamodbTableItemParameters, vals map[string]cty.Value) {
	vals["hash_key"] = cty.StringVal(p.HashKey)
}

func EncodeDynamodbTableItem_Id(p *DynamodbTableItemParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDynamodbTableItem_Item(p *DynamodbTableItemParameters, vals map[string]cty.Value) {
	vals["item"] = cty.StringVal(p.Item)
}

func EncodeDynamodbTableItem_RangeKey(p *DynamodbTableItemParameters, vals map[string]cty.Value) {
	vals["range_key"] = cty.StringVal(p.RangeKey)
}